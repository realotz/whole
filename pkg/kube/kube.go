package kube

import (
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/registry"
	config2 "github.com/realotz/whole/pkg/kube/config"
	registry2 "github.com/realotz/whole/pkg/kube/registry"
	"github.com/realotz/whole/pkg/kube/secrets"
	"github.com/realotz/whole/pkg/kube/utils"
	"k8s.io/client-go/kubernetes"
)

var (
	globalConfig    config.Source
	globalRegistrar registry.Registrar
	globalDiscovery registry.Discovery
)

// Option is kube option.
type Option func(*kube)

func ConfigOption(conf string) Option {
	return func(o *kube) {
		o.conf = conf
	}
}

func NamespaceDefOption(ns string) Option {
	return func(o *kube) {
		if o.namespace == "" {
			o.namespace = ns
		}
	}
}

type Kube interface {
	NewConfig(opts ...config2.Option) config.Source
	NewSecretsMap(opts ...secrets.SecretOption) secrets.SecretMap
	NewRegistry(opts ...registry2.RegistryOption) registry.Registrar
	NewDiscovery(opts ...registry2.DiscoveryOption) registry.Discovery
}

type kube struct {
	conf      string
	namespace string
	client    *kubernetes.Clientset
}

func New(opts ...Option) (Kube, error) {
	k := &kube{
		conf:      "",
		namespace: utils.GetNamespace(),
	}
	for _, o := range opts {
		o(k)
	}
	var err error
	// 初始化kube
	k.client, err = utils.LoadClient("", k.conf)
	if err != nil {
		return nil, err
	}
	return k, nil
}

//新建配置
func (k *kube) NewConfig(opts ...config2.Option) config.Source {
	opts = append(opts, config2.Namespace(k.namespace), config2.LoadClient(k.client))
	globalConfig = config2.NewSource(opts...)
	return globalConfig
}

//新建密文字典
func (k *kube) NewSecretsMap(opts ...secrets.SecretOption) secrets.SecretMap {
	opts = append(opts, secrets.Namespace(k.namespace), secrets.LoadClient(k.client))
	return secrets.NewSecretSource(opts...)
}

// 新建Registry
func (k *kube) NewRegistry(opts ...registry2.RegistryOption) registry.Registrar {
	opts = append(opts, registry2.Namespace(k.namespace), registry2.LoadClient(k.client))
	globalRegistrar = registry2.NewKubeRegistry(opts...)
	return globalRegistrar
}

// NewDiscovery
func (k *kube) NewDiscovery(opts ...registry2.DiscoveryOption) registry.Discovery {
	opts = append(opts, registry2.DiscoveryNamespace(k.namespace), registry2.LoadDiscoveryClient(k.client))
	globalDiscovery = registry2.NewKubeDiscovery(opts...)
	return globalDiscovery
}

func Discover() registry.Discovery {
	return globalDiscovery
}

func Registrar() registry.Registrar {
	return globalRegistrar
}
