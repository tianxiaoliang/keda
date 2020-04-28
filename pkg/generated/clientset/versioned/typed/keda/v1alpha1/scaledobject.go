/*
Copyright 2020 The KEDA Authors

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

	v1alpha1 "github.com/kedacore/keda/pkg/apis/keda/v1alpha1"
	scheme "github.com/kedacore/keda/pkg/generated/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// ScaledObjectsGetter has a method to return a ScaledObjectInterface.
// A group's client should implement this interface.
type ScaledObjectsGetter interface {
	ScaledObjects(namespace string) ScaledObjectInterface
}

// ScaledObjectInterface has methods to work with ScaledObject resources.
type ScaledObjectInterface interface {
	Create(*v1alpha1.ScaledObject) (*v1alpha1.ScaledObject, error)
	Update(*v1alpha1.ScaledObject) (*v1alpha1.ScaledObject, error)
	UpdateStatus(*v1alpha1.ScaledObject) (*v1alpha1.ScaledObject, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.ScaledObject, error)
	List(opts v1.ListOptions) (*v1alpha1.ScaledObjectList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ScaledObject, err error)
	ScaledObjectExpansion
}

// scaledObjects implements ScaledObjectInterface
type scaledObjects struct {
	client rest.Interface
	ns     string
}

// newScaledObjects returns a ScaledObjects
func newScaledObjects(c *KedaV1alpha1Client, namespace string) *scaledObjects {
	return &scaledObjects{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the scaledObject, and returns the corresponding scaledObject object, and an error if there is any.
func (c *scaledObjects) Get(name string, options v1.GetOptions) (result *v1alpha1.ScaledObject, err error) {
	result = &v1alpha1.ScaledObject{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("scaledobjects").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ScaledObjects that match those selectors.
func (c *scaledObjects) List(opts v1.ListOptions) (result *v1alpha1.ScaledObjectList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.ScaledObjectList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("scaledobjects").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested scaledObjects.
func (c *scaledObjects) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("scaledobjects").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a scaledObject and creates it.  Returns the server's representation of the scaledObject, and an error, if there is any.
func (c *scaledObjects) Create(scaledObject *v1alpha1.ScaledObject) (result *v1alpha1.ScaledObject, err error) {
	result = &v1alpha1.ScaledObject{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("scaledobjects").
		Body(scaledObject).
		Do().
		Into(result)
	return
}

// Update takes the representation of a scaledObject and updates it. Returns the server's representation of the scaledObject, and an error, if there is any.
func (c *scaledObjects) Update(scaledObject *v1alpha1.ScaledObject) (result *v1alpha1.ScaledObject, err error) {
	result = &v1alpha1.ScaledObject{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("scaledobjects").
		Name(scaledObject.Name).
		Body(scaledObject).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *scaledObjects) UpdateStatus(scaledObject *v1alpha1.ScaledObject) (result *v1alpha1.ScaledObject, err error) {
	result = &v1alpha1.ScaledObject{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("scaledobjects").
		Name(scaledObject.Name).
		SubResource("status").
		Body(scaledObject).
		Do().
		Into(result)
	return
}

// Delete takes name of the scaledObject and deletes it. Returns an error if one occurs.
func (c *scaledObjects) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("scaledobjects").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *scaledObjects) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("scaledobjects").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched scaledObject.
func (c *scaledObjects) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ScaledObject, err error) {
	result = &v1alpha1.ScaledObject{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("scaledobjects").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
