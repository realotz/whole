package registry

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/registry"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/watch"
)

type watcher struct {
	discovery   *kubeDiscovery
	watcher     watch.Interface
	serviceName string
}

func newWatcher(discovery *kubeDiscovery, serviceName string) (registry.Watcher, error) {
	w, err := discovery.client.CoreV1().Services(discovery.namespace).Watch(context.Background(), metav1.ListOptions{
		LabelSelector: fmt.Sprintf("app=%s", serviceName),
	})
	if err != nil {
		return nil, err
	}
	return &watcher{
		discovery:   discovery,
		watcher:     w,
		serviceName: serviceName,
	}, nil
}

func (w *watcher) Next() ([]*registry.ServiceInstance, error) {
ResultChan:
	ch := <-w.watcher.ResultChan()
	if ch.Object == nil {
		// 重新获取watcher
		k8sWatcher, err := w.discovery.client.CoreV1().Services(w.discovery.namespace).Watch(context.Background(), metav1.ListOptions{
			LabelSelector: fmt.Sprintf("app=%s", w.serviceName),
		})
		if err != nil {
			return nil, err
		}
		w.watcher = k8sWatcher
		goto ResultChan
	}
	cm, ok := ch.Object.(*v1.Service)
	if !ok {
		return nil, fmt.Errorf("kube Object not service")
	}
	if ch.Type == "DELETED" {
		return nil, fmt.Errorf("kube service delete %s", cm.Name)
	}
	return w.discovery.ToService(cm)
}

func (w *watcher) Stop() error {
	w.watcher.Stop()
	return nil
}
