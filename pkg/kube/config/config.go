package config

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-kratos/kratos/v2/config"
	"github.com/realotz/whole/pkg/kube/utils"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

// Option is kube option.
type Option func(*kube)

// Namespace with kube namespace.
func Namespace(ns string) Option {
	return func(o *kube) {
		o.namespace = ns
	}
}

// LabelSelector with kube label selector.
func LabelSelector(label string) Option {
	return func(o *kube) {
		o.labelSelector = label
	}
}

// FieldSelector with kube field selector.
func FieldSelector(field string) Option {
	return func(o *kube) {
		o.fieldSelector = field
	}
}

// KubeConfig with kube config.
func KubeConfig(config string) Option {
	return func(o *kube) {
		o.kubeConfig = config
	}
}

// Master with kube master.
func Master(master string) Option {
	return func(o *kube) {
		o.master = master
	}
}

// LoadClient with kube master.
func LoadClient(client *kubernetes.Clientset) Option {
	return func(o *kube) {
		o.client = client
	}
}

type kube struct {
	client *kubernetes.Clientset
	// kube namespace
	namespace string
	// kube labelSelector example `app=test`
	labelSelector string
	// kube fieldSelector example `app=test`
	fieldSelector string
	// set KubeConfig out-of-cluster Use outside cluster
	kubeConfig string
	// set master url
	master string
}

// NewSource new a kube config source.
func NewSource(opts ...Option) config.Source {
	conf := &kube{
		namespace: utils.GetNamespace(),
	}
	for _, o := range opts {
		o(conf)
	}
	return conf
}

func (k *kube) init() (err error) {
	if k.client != nil {
		return nil
	}
	cfg, err := utils.LoadRustConfig(k.master, k.kubeConfig)
	if err != nil {
		return err
	}
	if k.client, err = kubernetes.NewForConfig(cfg); err != nil {
		return err
	}
	return nil
}

func (k *kube) load() (kvs []*config.KeyValue, err error) {
	cmList, err := k.client.
		CoreV1().
		ConfigMaps(k.namespace).
		List(context.Background(), metav1.ListOptions{
			LabelSelector: k.labelSelector,
			FieldSelector: k.fieldSelector,
		})
	if err != nil {
		fmt.Println(k.namespace)
		return nil, err
	}
	for _, cm := range cmList.Items {
		kvs = append(kvs, k.configMap(cm)...)
	}
	return kvs, nil
}

func (k *kube) configMap(cm v1.ConfigMap) (kvs []*config.KeyValue) {
	for name, val := range cm.Data {
		kvs = append(kvs, &config.KeyValue{
			Key:   fmt.Sprintf("%s/%s/%s", k.namespace, cm.Name, name),
			Value: []byte(val),
		})
	}
	return kvs
}

func (k *kube) Load() ([]*config.KeyValue, error) {
	if k.namespace == "" {
		return nil, errors.New("options namespace not full")
	}
	if err := k.init(); err != nil {
		return nil, err
	}
	return k.load()
}

func (k *kube) Watch() (config.Watcher, error) {
	return newWatcher(k)
}
