// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/kyma-project/kyma/components/service-binding-usage-controller/pkg/apis/servicecatalog/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// ServiceBindingUsageLister helps list ServiceBindingUsages.
type ServiceBindingUsageLister interface {
	// List lists all ServiceBindingUsages in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.ServiceBindingUsage, err error)
	// ServiceBindingUsages returns an object that can list and get ServiceBindingUsages.
	ServiceBindingUsages(namespace string) ServiceBindingUsageNamespaceLister
	ServiceBindingUsageListerExpansion
}

// serviceBindingUsageLister implements the ServiceBindingUsageLister interface.
type serviceBindingUsageLister struct {
	indexer cache.Indexer
}

// NewServiceBindingUsageLister returns a new ServiceBindingUsageLister.
func NewServiceBindingUsageLister(indexer cache.Indexer) ServiceBindingUsageLister {
	return &serviceBindingUsageLister{indexer: indexer}
}

// List lists all ServiceBindingUsages in the indexer.
func (s *serviceBindingUsageLister) List(selector labels.Selector) (ret []*v1alpha1.ServiceBindingUsage, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ServiceBindingUsage))
	})
	return ret, err
}

// ServiceBindingUsages returns an object that can list and get ServiceBindingUsages.
func (s *serviceBindingUsageLister) ServiceBindingUsages(namespace string) ServiceBindingUsageNamespaceLister {
	return serviceBindingUsageNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ServiceBindingUsageNamespaceLister helps list and get ServiceBindingUsages.
type ServiceBindingUsageNamespaceLister interface {
	// List lists all ServiceBindingUsages in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.ServiceBindingUsage, err error)
	// Get retrieves the ServiceBindingUsage from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.ServiceBindingUsage, error)
	ServiceBindingUsageNamespaceListerExpansion
}

// serviceBindingUsageNamespaceLister implements the ServiceBindingUsageNamespaceLister
// interface.
type serviceBindingUsageNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all ServiceBindingUsages in the indexer for a given namespace.
func (s serviceBindingUsageNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.ServiceBindingUsage, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ServiceBindingUsage))
	})
	return ret, err
}

// Get retrieves the ServiceBindingUsage from the indexer for a given namespace and name.
func (s serviceBindingUsageNamespaceLister) Get(name string) (*v1alpha1.ServiceBindingUsage, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("servicebindingusage"), name)
	}
	return obj.(*v1alpha1.ServiceBindingUsage), nil
}