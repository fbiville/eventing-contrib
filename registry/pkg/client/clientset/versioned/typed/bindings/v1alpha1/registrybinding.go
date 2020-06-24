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

// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	"time"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
	v1alpha1 "knative.dev/eventing-contrib/registry/pkg/apis/bindings/v1alpha1"
	scheme "knative.dev/eventing-contrib/registry/pkg/client/clientset/versioned/scheme"
)

// RegistryBindingsGetter has a method to return a RegistryBindingInterface.
// A group's client should implement this interface.
type RegistryBindingsGetter interface {
	RegistryBindings(namespace string) RegistryBindingInterface
}

// RegistryBindingInterface has methods to work with RegistryBinding resources.
type RegistryBindingInterface interface {
	Create(*v1alpha1.RegistryBinding) (*v1alpha1.RegistryBinding, error)
	Update(*v1alpha1.RegistryBinding) (*v1alpha1.RegistryBinding, error)
	UpdateStatus(*v1alpha1.RegistryBinding) (*v1alpha1.RegistryBinding, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.RegistryBinding, error)
	List(opts v1.ListOptions) (*v1alpha1.RegistryBindingList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.RegistryBinding, err error)
	RegistryBindingExpansion
}

// registryBindings implements RegistryBindingInterface
type registryBindings struct {
	client rest.Interface
	ns     string
}

// newRegistryBindings returns a RegistryBindings
func newRegistryBindings(c *BindingsV1alpha1Client, namespace string) *registryBindings {
	return &registryBindings{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the registryBinding, and returns the corresponding registryBinding object, and an error if there is any.
func (c *registryBindings) Get(name string, options v1.GetOptions) (result *v1alpha1.RegistryBinding, err error) {
	result = &v1alpha1.RegistryBinding{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("registrybindings").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of RegistryBindings that match those selectors.
func (c *registryBindings) List(opts v1.ListOptions) (result *v1alpha1.RegistryBindingList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.RegistryBindingList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("registrybindings").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested registryBindings.
func (c *registryBindings) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("registrybindings").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a registryBinding and creates it.  Returns the server's representation of the registryBinding, and an error, if there is any.
func (c *registryBindings) Create(registryBinding *v1alpha1.RegistryBinding) (result *v1alpha1.RegistryBinding, err error) {
	result = &v1alpha1.RegistryBinding{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("registrybindings").
		Body(registryBinding).
		Do().
		Into(result)
	return
}

// Update takes the representation of a registryBinding and updates it. Returns the server's representation of the registryBinding, and an error, if there is any.
func (c *registryBindings) Update(registryBinding *v1alpha1.RegistryBinding) (result *v1alpha1.RegistryBinding, err error) {
	result = &v1alpha1.RegistryBinding{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("registrybindings").
		Name(registryBinding.Name).
		Body(registryBinding).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *registryBindings) UpdateStatus(registryBinding *v1alpha1.RegistryBinding) (result *v1alpha1.RegistryBinding, err error) {
	result = &v1alpha1.RegistryBinding{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("registrybindings").
		Name(registryBinding.Name).
		SubResource("status").
		Body(registryBinding).
		Do().
		Into(result)
	return
}

// Delete takes name of the registryBinding and deletes it. Returns an error if one occurs.
func (c *registryBindings) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("registrybindings").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *registryBindings) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("registrybindings").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched registryBinding.
func (c *registryBindings) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.RegistryBinding, err error) {
	result = &v1alpha1.RegistryBinding{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("registrybindings").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
