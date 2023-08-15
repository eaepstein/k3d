//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright The KubeStellar Authors.

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

// Code generated by kcp code-generator. DO NOT EDIT.

package v1alpha1

import (
	"context"
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/tools/cache"

	kcpcache "github.com/kcp-dev/apimachinery/v2/pkg/cache"
	kcpinformers "github.com/kcp-dev/apimachinery/v2/third_party/informers"
	"github.com/kcp-dev/logicalcluster/v3"

	spacev1alpha1 "github.com/kubestellar/kubestellar/pkg/apis/space/v1alpha1"
	scopedclientset "github.com/kubestellar/kubestellar/pkg/client/clientset/versioned"
	clientset "github.com/kubestellar/kubestellar/pkg/client/clientset/versioned/cluster"
	"github.com/kubestellar/kubestellar/pkg/client/informers/externalversions/internalinterfaces"
	spacev1alpha1listers "github.com/kubestellar/kubestellar/pkg/client/listers/space/v1alpha1"
)

// SpaceClusterInformer provides access to a shared informer and lister for
// Spaces.
type SpaceClusterInformer interface {
	Cluster(logicalcluster.Name) SpaceInformer
	Informer() kcpcache.ScopeableSharedIndexInformer
	Lister() spacev1alpha1listers.SpaceClusterLister
}

type spaceClusterInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewSpaceClusterInformer constructs a new informer for Space type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewSpaceClusterInformer(client clientset.ClusterInterface, resyncPeriod time.Duration, indexers cache.Indexers) kcpcache.ScopeableSharedIndexInformer {
	return NewFilteredSpaceClusterInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredSpaceClusterInformer constructs a new informer for Space type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredSpaceClusterInformer(client clientset.ClusterInterface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) kcpcache.ScopeableSharedIndexInformer {
	return kcpinformers.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.SpaceV1alpha1().Spaces().List(context.TODO(), options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.SpaceV1alpha1().Spaces().Watch(context.TODO(), options)
			},
		},
		&spacev1alpha1.Space{},
		resyncPeriod,
		indexers,
	)
}

func (f *spaceClusterInformer) defaultInformer(client clientset.ClusterInterface, resyncPeriod time.Duration) kcpcache.ScopeableSharedIndexInformer {
	return NewFilteredSpaceClusterInformer(client, resyncPeriod, cache.Indexers{
		kcpcache.ClusterIndexName:             kcpcache.ClusterIndexFunc,
		kcpcache.ClusterAndNamespaceIndexName: kcpcache.ClusterAndNamespaceIndexFunc},
		f.tweakListOptions,
	)
}

func (f *spaceClusterInformer) Informer() kcpcache.ScopeableSharedIndexInformer {
	return f.factory.InformerFor(&spacev1alpha1.Space{}, f.defaultInformer)
}

func (f *spaceClusterInformer) Lister() spacev1alpha1listers.SpaceClusterLister {
	return spacev1alpha1listers.NewSpaceClusterLister(f.Informer().GetIndexer())
}

// SpaceInformer provides access to a shared informer and lister for
// Spaces.
type SpaceInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() spacev1alpha1listers.SpaceLister
}

func (f *spaceClusterInformer) Cluster(clusterName logicalcluster.Name) SpaceInformer {
	return &spaceInformer{
		informer: f.Informer().Cluster(clusterName),
		lister:   f.Lister().Cluster(clusterName),
	}
}

type spaceInformer struct {
	informer cache.SharedIndexInformer
	lister   spacev1alpha1listers.SpaceLister
}

func (f *spaceInformer) Informer() cache.SharedIndexInformer {
	return f.informer
}

func (f *spaceInformer) Lister() spacev1alpha1listers.SpaceLister {
	return f.lister
}

type spaceScopedInformer struct {
	factory          internalinterfaces.SharedScopedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

func (f *spaceScopedInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&spacev1alpha1.Space{}, f.defaultInformer)
}

func (f *spaceScopedInformer) Lister() spacev1alpha1listers.SpaceLister {
	return spacev1alpha1listers.NewSpaceLister(f.Informer().GetIndexer())
}

// NewSpaceInformer constructs a new informer for Space type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewSpaceInformer(client scopedclientset.Interface, resyncPeriod time.Duration, namespace string, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredSpaceInformer(client, resyncPeriod, namespace, indexers, nil)
}

// NewFilteredSpaceInformer constructs a new informer for Space type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredSpaceInformer(client scopedclientset.Interface, resyncPeriod time.Duration, namespace string, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.SpaceV1alpha1().Spaces(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.SpaceV1alpha1().Spaces(namespace).Watch(context.TODO(), options)
			},
		},
		&spacev1alpha1.Space{},
		resyncPeriod,
		indexers,
	)
}

func (f *spaceScopedInformer) defaultInformer(client scopedclientset.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredSpaceInformer(client, resyncPeriod, f.namespace, cache.Indexers{
		cache.NamespaceIndex: cache.MetaNamespaceIndexFunc,
	}, f.tweakListOptions)
}