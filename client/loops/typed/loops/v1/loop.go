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

package v1

import (
	"context"
	"time"

	v1 "github.com/clever-telemetry/miniloops/apis/loops/v1"
	scheme "github.com/clever-telemetry/miniloops/client/loops/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// LoopsGetter has a method to return a LoopInterface.
// A group's client should implement this interface.
type LoopsGetter interface {
	Loops(namespace string) LoopInterface
}

// LoopInterface has methods to work with Loop resources.
type LoopInterface interface {
	Create(ctx context.Context, loop *v1.Loop, opts metav1.CreateOptions) (*v1.Loop, error)
	Update(ctx context.Context, loop *v1.Loop, opts metav1.UpdateOptions) (*v1.Loop, error)
	UpdateStatus(ctx context.Context, loop *v1.Loop, opts metav1.UpdateOptions) (*v1.Loop, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.Loop, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.LoopList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.Loop, err error)
	LoopExpansion
}

// loops implements LoopInterface
type loops struct {
	client rest.Interface
	ns     string
}

// newLoops returns a Loops
func newLoops(c *LoopsV1Client, namespace string) *loops {
	return &loops{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the loop, and returns the corresponding loop object, and an error if there is any.
func (c *loops) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.Loop, err error) {
	result = &v1.Loop{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("loops").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Loops that match those selectors.
func (c *loops) List(ctx context.Context, opts metav1.ListOptions) (result *v1.LoopList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.LoopList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("loops").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested loops.
func (c *loops) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("loops").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a loop and creates it.  Returns the server's representation of the loop, and an error, if there is any.
func (c *loops) Create(ctx context.Context, loop *v1.Loop, opts metav1.CreateOptions) (result *v1.Loop, err error) {
	result = &v1.Loop{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("loops").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(loop).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a loop and updates it. Returns the server's representation of the loop, and an error, if there is any.
func (c *loops) Update(ctx context.Context, loop *v1.Loop, opts metav1.UpdateOptions) (result *v1.Loop, err error) {
	result = &v1.Loop{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("loops").
		Name(loop.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(loop).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *loops) UpdateStatus(ctx context.Context, loop *v1.Loop, opts metav1.UpdateOptions) (result *v1.Loop, err error) {
	result = &v1.Loop{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("loops").
		Name(loop.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(loop).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the loop and deletes it. Returns an error if one occurs.
func (c *loops) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("loops").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *loops) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("loops").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched loop.
func (c *loops) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.Loop, err error) {
	result = &v1.Loop{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("loops").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
