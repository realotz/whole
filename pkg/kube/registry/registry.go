package registry

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/registry"
	jsoniter "github.com/json-iterator/go"
	"github.com/realotz/whole/pkg/kube/utils"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
	"k8s.io/client-go/kubernetes"
	"net"
	"net/url"
	"strings"
	"time"
)

type RegistryOption func(r *kubeRegistry)

const MetaData = "metadata"

var DevelopsKey = fmt.Sprintf("develops.%s", LocalMac())

func LoadClient(client *kubernetes.Clientset) RegistryOption {
	return func(o *kubeRegistry) {
		o.client = client
	}
}

// KubeConfig with kube config.
func KubeConfig(config string) RegistryOption {
	return func(o *kubeRegistry) {
		o.kubeConfig = config
	}
}

// Namespace with kube config.
func Namespace(config string) RegistryOption {
	return func(o *kubeRegistry) {
		if utils.GetNamespace() == "" {
			o.dev = true
		}
		o.namespace = config
	}
}

// Master with kube master.
func Master(master string) RegistryOption {
	return func(o *kubeRegistry) {
		o.master = master
	}
}

type kubeRegistry struct {
	client *kubernetes.Clientset
	// kube namespace
	namespace string
	// kube master
	master string
	// set KubeConfig out-of-cluster Use outside cluster
	kubeConfig string
	dev        bool
}

func NewKubeRegistry(opts ...RegistryOption) registry.Registrar {
	m := &kubeRegistry{}
	for _, o := range opts {
		o(m)
	}
	return m
}

// 获取本机mac地址 用于开发模式下匹配对应服务地址
func LocalMac() string {
	interfaces, err := net.Interfaces()
	if err != nil {
		panic("Poor soul, here is what you got: " + err.Error())
	}
	inter := interfaces[0]
	mac := inter.HardwareAddr.String() //获取本机MAC地址
	return strings.ReplaceAll(mac, ":", ".")
}

//注册服务
func (r *kubeRegistry) Register(ctx context.Context, service *registry.ServiceInstance) error {
	serviceSet := r.client.CoreV1().Services(r.namespace)
	kubeService, err := serviceSet.Get(ctx, service.Name, metav1.GetOptions{})
	label := map[string]string{
		"app":     service.Name,
		"version": service.Version,
	}
	var annotations map[string]string
	annotations = kubeService.Annotations
	if annotations == nil {
		annotations = make(map[string]string)
	}
	annotations[MetaData], _ = jsoniter.MarshalToString(service)
	//测试环境设置mac地址
	if r.dev {
		annotations[DevelopsKey], _ = jsoniter.MarshalToString(service)
	}
	var ports []v1.ServicePort
	for _, ent := range service.Endpoints {
		u, _ := url.Parse(ent)
		var port int32 = 9001
		if u.Scheme == "http" {
			port = 9002
		}
		ports = append(ports, v1.ServicePort{
			Name:       u.Scheme,
			Protocol:   "TCP",
			Port:       port,
			TargetPort: intstr.Parse(u.Port()),
		})
	}
	if err == nil {
		kubeService.Annotations = annotations
		kubeService.Spec.Ports = ports
		kubeService.Labels = label
		kubeService.Spec.Type = v1.ServiceTypeClusterIP
		kubeService.Spec.Selector = label
		kubeService, err = serviceSet.Update(ctx, kubeService, metav1.UpdateOptions{})
		if err != nil {
			return err
		}
		return nil
	}
	kubeService, err = serviceSet.Create(ctx, &v1.Service{
		TypeMeta: metav1.TypeMeta{
			Kind:       "Service",
			APIVersion: "v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:        service.Name,
			Namespace:   r.namespace,
			Labels:      label,
			Annotations: annotations,
		},
		Spec: v1.ServiceSpec{
			Ports:    ports,
			Selector: label,
			Type:     v1.ServiceTypeClusterIP,
		},
	}, metav1.CreateOptions{})
	if err != nil {
		return err
	}
	return nil
}

//注销服务
func (r *kubeRegistry) Deregister(ctx context.Context, service *registry.ServiceInstance) error {
	serviceSet := r.client.CoreV1().Services(r.namespace)
	kubeService, err := serviceSet.Get(ctx, service.Name, metav1.GetOptions{})
	if err != nil {
		return err
	}
	annotations := kubeService.Annotations
	if _, ok := annotations[DevelopsKey]; ok {
		// 延迟3秒后修改
		time.Sleep(2 * time.Second)
		kubeService, err = serviceSet.Get(ctx, service.Name, metav1.GetOptions{})
		if err != nil {
			return err
		}
		delete(annotations, DevelopsKey)
		kubeService.Annotations = annotations
		kubeService, err = serviceSet.Update(ctx, kubeService, metav1.UpdateOptions{})
		if err != nil {
			return err
		}
	}
	return nil
}
