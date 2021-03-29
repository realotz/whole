package utils

import (
	"io/ioutil"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

// ServiceAccountNamespacePath defines the location of the namespace file
const ServiceAccountNamespacePath = "/var/run/secrets/kubernetes.io/serviceaccount/namespace"

var currentNamespace = LoadNamespace()

// GetNamespace is used to get the namespace of the Pod where the current container is located
func GetNamespace() string {
	return currentNamespace
}

// LoadNamespace is used to get the current namespace from the file
func LoadNamespace() string {
	data, err := ioutil.ReadFile(ServiceAccountNamespacePath)
	if err != nil {
		return ""
	}
	return string(data)
}

func LoadRustConfig(master, config string) (*rest.Config, error) {
	if config != "" {
		return clientcmd.BuildConfigFromFlags(master, config)
	} else {
		return rest.InClusterConfig()
	}
}

func LoadClient(master, config string) (*kubernetes.Clientset, error) {
	cfg, err := LoadRustConfig(master, config)
	if err != nil {
		return nil, err
	}
	return kubernetes.NewForConfig(cfg)
}
