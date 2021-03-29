package registry

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/registry"
	jsoniter "github.com/json-iterator/go"
	"github.com/realotz/whole/pkg/kube/utils"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

type DiscoveryOption func(r *kubeDiscovery)

func LoadDiscoveryClient(client *kubernetes.Clientset) DiscoveryOption {
	return func(o *kubeDiscovery) {
		o.client = client
	}
}

// KubeConfig with kube config.
func KubeDiscoveryConfig(config string) DiscoveryOption {
	return func(o *kubeDiscovery) {
		o.kubeConfig = config
	}
}

// Namespace with kube config.
func DiscoveryNamespace(config string) DiscoveryOption {
	return func(o *kubeDiscovery) {
		if utils.GetNamespace() == "" {
			o.dev = true
		}
		o.namespace = config
	}
}

// Master with kube master.
func DiscoveryMaster(master string) DiscoveryOption {
	return func(o *kubeDiscovery) {
		o.master = master
	}
}

type kubeDiscovery struct {
	client *kubernetes.Clientset
	// kube namespace
	namespace string
	// kube master
	master string
	// set KubeConfig out-of-cluster Use outside cluster
	kubeConfig string
	dev        bool
}

func (k *kubeDiscovery) ToService(service *v1.Service) ([]*registry.ServiceInstance, error) {
	var err error
	ser := &registry.ServiceInstance{}
	if v, ok := service.Annotations[DevelopsKey]; ok && k.dev {
		err = jsoniter.UnmarshalFromString(v, ser)
		if err != nil {
			return nil, err
		}
		return []*registry.ServiceInstance{ser}, nil
	}
	if v, ok := service.Annotations[MetaData]; ok {
		_ = jsoniter.UnmarshalFromString(v, ser)
	}
	if ser.ID == "" {
		ser.ID = string(service.UID)
	}
	if ser.Name == "" {
		ser.ID = service.Name
	}
	ser.Endpoints = []string{
		fmt.Sprintf("grpc://%s.%s.svc.cluster.local:9001", service.Name, service.Namespace),
		fmt.Sprintf("http://%s.%s.svc.cluster.local:9002", service.Name, service.Namespace),
	}
	fmt.Println(ser)
	return []*registry.ServiceInstance{ser}, err
}

func (k *kubeDiscovery) GetService(ctx context.Context, serviceName string) ([]*registry.ServiceInstance, error) {
	serviceSet := k.client.CoreV1().Services(k.namespace)
	service, err := serviceSet.Get(ctx, serviceName, metav1.GetOptions{})
	if err != nil {
		return nil, err
	}
	return k.ToService(service)
}

func (k *kubeDiscovery) Watch(ctx context.Context, serviceName string) (registry.Watcher, error) {
	return newWatcher(k, serviceName)
}

func NewKubeDiscovery(opts ...DiscoveryOption) registry.Discovery {
	m := &kubeDiscovery{}
	for _, o := range opts {
		o(m)
	}
	return m
}
