package secrets

import (
	"context"
	"github.com/realotz/whole/pkg/kube/utils"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

type MapOption func(p *secretObject)

func TypeOption(t string) MapOption {
	return func(p *secretObject) {
		p.tp = t
	}
}

type SecretMap interface {
	GetSecret(name string, ops ...MapOption) Map
}

type Map interface {
	Get(ctx context.Context) (map[string]string, error)
	Reload(ctx context.Context) (map[string]string, error)
	Set(ctx context.Context, data map[string]string) error
}

// Option is kube option.
type SecretOption func(*secret)

func LoadClient(client *kubernetes.Clientset) SecretOption {
	return func(o *secret) {
		o.client = client
	}
}

// KubeConfig with kube config.
func KubeConfig(config string) SecretOption {
	return func(o *secret) {
		o.kubeConfig = config
	}
}

// Namespace with kube config.
func Namespace(config string) SecretOption {
	return func(o *secret) {
		o.namespace = config
	}
}

// Master with kube master.
func Master(master string) SecretOption {
	return func(o *secret) {
		o.master = master
	}
}

// 密文字典存储
type secret struct {
	client *kubernetes.Clientset
	// kube namespace
	namespace string
	// kube master
	master string
	// set KubeConfig out-of-cluster Use outside cluster
	kubeConfig string
}

func NewSecretSource(opts ...SecretOption) *secret {
	m := &secret{}
	for _, o := range opts {
		o(m)
	}
	return m
}

func (k *secret) init() (err error) {
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

type secretObject struct {
	name      string
	tp        string
	namespace string
	sts       *secret
	data      *v1.Secret
}

// 获取密钥对象
func (s *secret) GetSecret(name string, ops ...MapOption) Map {
	obj := &secretObject{
		name:      name,
		namespace: s.namespace,
		sts:       s,
	}
	for _, o := range ops {
		o(obj)
	}
	return obj
}

//Get 获取密文配置
func (s *secretObject) Get(ctx context.Context) (map[string]string, error) {
	var err error
	if s.data != nil {
		return s.data.StringData, err
	}
	return s.Reload(ctx)
}

//Get 获取密文配置
func (s *secretObject) Reload(ctx context.Context) (map[string]string, error) {
	var err error
	s.data, err = s.sts.client.CoreV1().Secrets(s.namespace).Get(ctx, s.name, metav1.GetOptions{})
	if err != nil {
		return nil, err
	}
	if s.data.StringData == nil && s.data.Data != nil {
		s.data.StringData = make(map[string]string)
		for k, v := range s.data.Data {
			s.data.StringData[k] = string(v)
		}
	}
	return s.data.StringData, nil
}

// 设置密文配置
func (s *secretObject) Set(ctx context.Context, data map[string]string) error {
	var err error
	dataByte := make(map[string][]byte)
	for k, v := range data {
		dataByte[k] = []byte(v)
	}
	s.data, err = s.sts.client.CoreV1().Secrets(s.namespace).Get(ctx, s.name, metav1.GetOptions{})
	if err != nil {
		s.data, err = s.sts.client.CoreV1().Secrets(s.namespace).Create(ctx, &v1.Secret{
			TypeMeta: metav1.TypeMeta{
				Kind:       "Secret",
				APIVersion: "v1",
			},
			ObjectMeta: metav1.ObjectMeta{
				Name:      s.name,
				Namespace: s.namespace,
			},
			StringData: data,
			Data:       dataByte,
			Type:       v1.SecretType(s.tp),
		}, metav1.CreateOptions{})
		return err
	}
	s.data.StringData = data
	s.data.Data = dataByte
	s.data, err = s.sts.client.CoreV1().Secrets(s.namespace).Update(ctx, s.data, metav1.UpdateOptions{})
	return err
}
