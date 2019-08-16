/*
Copyright The Kubernetes Authors.

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

	v1alpha1 "github.com/cloud-ark/kubeplus/platform-operator/pkg/apis/platformstackcontroller/v1alpha1"
	scheme "github.com/cloud-ark/kubeplus/platform-operator/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// PlatformStacksGetter has a method to return a PlatformStackInterface.
// A group's client should implement this interface.
type PlatformStacksGetter interface {
	PlatformStacks(namespace string) PlatformStackInterface
}

// PlatformStackInterface has methods to work with PlatformStack resources.
type PlatformStackInterface interface {
	Create(*v1alpha1.PlatformStack) (*v1alpha1.PlatformStack, error)
	Update(*v1alpha1.PlatformStack) (*v1alpha1.PlatformStack, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.PlatformStack, error)
	List(opts v1.ListOptions) (*v1alpha1.PlatformStackList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.PlatformStack, err error)
	PlatformStackExpansion
}

// platformStacks implements PlatformStackInterface
type platformStacks struct {
	client rest.Interface
	ns     string
}

// newPlatformStacks returns a PlatformStacks
func newPlatformStacks(c *PlatformstackV1alpha1Client, namespace string) *platformStacks {
	return &platformStacks{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the platformStack, and returns the corresponding platformStack object, and an error if there is any.
func (c *platformStacks) Get(name string, options v1.GetOptions) (result *v1alpha1.PlatformStack, err error) {
	result = &v1alpha1.PlatformStack{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("platformstacks").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of PlatformStacks that match those selectors.
func (c *platformStacks) List(opts v1.ListOptions) (result *v1alpha1.PlatformStackList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.PlatformStackList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("platformstacks").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested platformStacks.
func (c *platformStacks) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("platformstacks").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a platformStack and creates it.  Returns the server's representation of the platformStack, and an error, if there is any.
func (c *platformStacks) Create(platformStack *v1alpha1.PlatformStack) (result *v1alpha1.PlatformStack, err error) {
	result = &v1alpha1.PlatformStack{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("platformstacks").
		Body(platformStack).
		Do().
		Into(result)
	return
}

// Update takes the representation of a platformStack and updates it. Returns the server's representation of the platformStack, and an error, if there is any.
func (c *platformStacks) Update(platformStack *v1alpha1.PlatformStack) (result *v1alpha1.PlatformStack, err error) {
	result = &v1alpha1.PlatformStack{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("platformstacks").
		Name(platformStack.Name).
		Body(platformStack).
		Do().
		Into(result)
	return
}

// Delete takes name of the platformStack and deletes it. Returns an error if one occurs.
func (c *platformStacks) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("platformstacks").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *platformStacks) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("platformstacks").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched platformStack.
func (c *platformStacks) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.PlatformStack, err error) {
	result = &v1alpha1.PlatformStack{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("platformstacks").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}