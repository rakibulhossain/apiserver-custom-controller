// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1 "github.com/Rakibul-Hossain/apiserver-custom-controller/pkg/generated/clientset/versioned/typed/customcontroller/v1"
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeCustomcontrollerV1 struct {
	*testing.Fake
}

func (c *FakeCustomcontrollerV1) Customs(namespace string) v1.CustomInterface {
	return &FakeCustoms{c, namespace}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeCustomcontrollerV1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
