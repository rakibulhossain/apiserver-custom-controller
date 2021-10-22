// Code generated by lister-gen. DO NOT EDIT.

package v1

import (
	v1 "github.com/Rakibul-Hossain/apiserver-custom-controller/pkg/apis/customcontroller/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// CustomLister helps list Customs.
// All objects returned here must be treated as read-only.
type CustomLister interface {
	// List lists all Customs in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.Custom, err error)
	// Customs returns an object that can list and get Customs.
	Customs(namespace string) CustomNamespaceLister
	CustomListerExpansion
}

// customLister implements the CustomLister interface.
type customLister struct {
	indexer cache.Indexer
}

// NewCustomLister returns a new CustomLister.
func NewCustomLister(indexer cache.Indexer) CustomLister {
	return &customLister{indexer: indexer}
}

// List lists all Customs in the indexer.
func (s *customLister) List(selector labels.Selector) (ret []*v1.Custom, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.Custom))
	})
	return ret, err
}

// Customs returns an object that can list and get Customs.
func (s *customLister) Customs(namespace string) CustomNamespaceLister {
	return customNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// CustomNamespaceLister helps list and get Customs.
// All objects returned here must be treated as read-only.
type CustomNamespaceLister interface {
	// List lists all Customs in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.Custom, err error)
	// Get retrieves the Custom from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1.Custom, error)
	CustomNamespaceListerExpansion
}

// customNamespaceLister implements the CustomNamespaceLister
// interface.
type customNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all Customs in the indexer for a given namespace.
func (s customNamespaceLister) List(selector labels.Selector) (ret []*v1.Custom, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.Custom))
	})
	return ret, err
}

// Get retrieves the Custom from the indexer for a given namespace and name.
func (s customNamespaceLister) Get(name string) (*v1.Custom, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("custom"), name)
	}
	return obj.(*v1.Custom), nil
}