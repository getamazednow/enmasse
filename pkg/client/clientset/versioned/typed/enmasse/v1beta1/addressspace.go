/*
 * Copyright 2018, EnMasse authors.
 * License: Apache License 2.0 (see the file LICENSE or http://apache.org/licenses/LICENSE-2.0.html).
 */

// Code generated by client-gen. DO NOT EDIT.

package v1beta1

import (
	v1beta1 "github.com/enmasseproject/enmasse/pkg/apis/enmasse/v1beta1"
	scheme "github.com/enmasseproject/enmasse/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// AddressSpacesGetter has a method to return a AddressSpaceInterface.
// A group's client should implement this interface.
type AddressSpacesGetter interface {
	AddressSpaces(namespace string) AddressSpaceInterface
}

// AddressSpaceInterface has methods to work with AddressSpace resources.
type AddressSpaceInterface interface {
	Create(*v1beta1.AddressSpace) (*v1beta1.AddressSpace, error)
	Update(*v1beta1.AddressSpace) (*v1beta1.AddressSpace, error)
	UpdateStatus(*v1beta1.AddressSpace) (*v1beta1.AddressSpace, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1beta1.AddressSpace, error)
	List(opts v1.ListOptions) (*v1beta1.AddressSpaceList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1beta1.AddressSpace, err error)
	AddressSpaceExpansion
}

// addressSpaces implements AddressSpaceInterface
type addressSpaces struct {
	client rest.Interface
	ns     string
}

// newAddressSpaces returns a AddressSpaces
func newAddressSpaces(c *EnmasseV1beta1Client, namespace string) *addressSpaces {
	return &addressSpaces{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the addressSpace, and returns the corresponding addressSpace object, and an error if there is any.
func (c *addressSpaces) Get(name string, options v1.GetOptions) (result *v1beta1.AddressSpace, err error) {
	result = &v1beta1.AddressSpace{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("addressspaces").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AddressSpaces that match those selectors.
func (c *addressSpaces) List(opts v1.ListOptions) (result *v1beta1.AddressSpaceList, err error) {
	result = &v1beta1.AddressSpaceList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("addressspaces").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested addressSpaces.
func (c *addressSpaces) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("addressspaces").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a addressSpace and creates it.  Returns the server's representation of the addressSpace, and an error, if there is any.
func (c *addressSpaces) Create(addressSpace *v1beta1.AddressSpace) (result *v1beta1.AddressSpace, err error) {
	result = &v1beta1.AddressSpace{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("addressspaces").
		Body(addressSpace).
		Do().
		Into(result)
	return
}

// Update takes the representation of a addressSpace and updates it. Returns the server's representation of the addressSpace, and an error, if there is any.
func (c *addressSpaces) Update(addressSpace *v1beta1.AddressSpace) (result *v1beta1.AddressSpace, err error) {
	result = &v1beta1.AddressSpace{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("addressspaces").
		Name(addressSpace.Name).
		Body(addressSpace).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *addressSpaces) UpdateStatus(addressSpace *v1beta1.AddressSpace) (result *v1beta1.AddressSpace, err error) {
	result = &v1beta1.AddressSpace{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("addressspaces").
		Name(addressSpace.Name).
		SubResource("status").
		Body(addressSpace).
		Do().
		Into(result)
	return
}

// Delete takes name of the addressSpace and deletes it. Returns an error if one occurs.
func (c *addressSpaces) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("addressspaces").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *addressSpaces) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("addressspaces").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched addressSpace.
func (c *addressSpaces) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1beta1.AddressSpace, err error) {
	result = &v1beta1.AddressSpace{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("addressspaces").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
