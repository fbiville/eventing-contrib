/*
Copyright 2020 The Knative Authors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	time "time"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
	sourcesv1alpha1 "knative.dev/eventing-contrib/registry/pkg/apis/sources/v1alpha1"
	versioned "knative.dev/eventing-contrib/registry/pkg/client/clientset/versioned"
	internalinterfaces "knative.dev/eventing-contrib/registry/pkg/client/informers/externalversions/internalinterfaces"
	v1alpha1 "knative.dev/eventing-contrib/registry/pkg/client/listers/sources/v1alpha1"
)

// RegistrySourceInformer provides access to a shared informer and lister for
// RegistrySources.
type RegistrySourceInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.RegistrySourceLister
}

type registrySourceInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewRegistrySourceInformer constructs a new informer for RegistrySource type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewRegistrySourceInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredRegistrySourceInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredRegistrySourceInformer constructs a new informer for RegistrySource type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredRegistrySourceInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.SourcesV1alpha1().RegistrySources(namespace).List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.SourcesV1alpha1().RegistrySources(namespace).Watch(options)
			},
		},
		&sourcesv1alpha1.RegistrySource{},
		resyncPeriod,
		indexers,
	)
}

func (f *registrySourceInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredRegistrySourceInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *registrySourceInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&sourcesv1alpha1.RegistrySource{}, f.defaultInformer)
}

func (f *registrySourceInformer) Lister() v1alpha1.RegistrySourceLister {
	return v1alpha1.NewRegistrySourceLister(f.Informer().GetIndexer())
}