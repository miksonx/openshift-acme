// Code generated by informer-gen. DO NOT EDIT.

package v1

import (
	time "time"

	operator_v1 "github.com/openshift/api/operator/v1"
	versioned "github.com/openshift/client-go/operator/clientset/versioned"
	internalinterfaces "github.com/openshift/client-go/operator/informers/externalversions/internalinterfaces"
	v1 "github.com/openshift/client-go/operator/listers/operator/v1"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// IngressControllerInformer provides access to a shared informer and lister for
// IngressControllers.
type IngressControllerInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1.IngressControllerLister
}

type ingressControllerInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewIngressControllerInformer constructs a new informer for IngressController type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewIngressControllerInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredIngressControllerInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredIngressControllerInformer constructs a new informer for IngressController type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredIngressControllerInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options meta_v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.OperatorV1().IngressControllers(namespace).List(options)
			},
			WatchFunc: func(options meta_v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.OperatorV1().IngressControllers(namespace).Watch(options)
			},
		},
		&operator_v1.IngressController{},
		resyncPeriod,
		indexers,
	)
}

func (f *ingressControllerInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredIngressControllerInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *ingressControllerInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&operator_v1.IngressController{}, f.defaultInformer)
}

func (f *ingressControllerInformer) Lister() v1.IngressControllerLister {
	return v1.NewIngressControllerLister(f.Informer().GetIndexer())
}
