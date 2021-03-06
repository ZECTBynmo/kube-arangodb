//
// DISCLAIMER
//
// Copyright 2018 ArangoDB GmbH, Cologne, Germany
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Copyright holder is ArangoDB GmbH, Cologne, Germany
//
package v1alpha

import (
	v1alpha "github.com/arangodb/kube-arangodb/pkg/apis/storage/v1alpha"
	scheme "github.com/arangodb/kube-arangodb/pkg/generated/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// ArangoLocalStoragesGetter has a method to return a ArangoLocalStorageInterface.
// A group's client should implement this interface.
type ArangoLocalStoragesGetter interface {
	ArangoLocalStorages() ArangoLocalStorageInterface
}

// ArangoLocalStorageInterface has methods to work with ArangoLocalStorage resources.
type ArangoLocalStorageInterface interface {
	Create(*v1alpha.ArangoLocalStorage) (*v1alpha.ArangoLocalStorage, error)
	Update(*v1alpha.ArangoLocalStorage) (*v1alpha.ArangoLocalStorage, error)
	UpdateStatus(*v1alpha.ArangoLocalStorage) (*v1alpha.ArangoLocalStorage, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha.ArangoLocalStorage, error)
	List(opts v1.ListOptions) (*v1alpha.ArangoLocalStorageList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha.ArangoLocalStorage, err error)
	ArangoLocalStorageExpansion
}

// arangoLocalStorages implements ArangoLocalStorageInterface
type arangoLocalStorages struct {
	client rest.Interface
}

// newArangoLocalStorages returns a ArangoLocalStorages
func newArangoLocalStorages(c *StorageV1alphaClient) *arangoLocalStorages {
	return &arangoLocalStorages{
		client: c.RESTClient(),
	}
}

// Get takes name of the arangoLocalStorage, and returns the corresponding arangoLocalStorage object, and an error if there is any.
func (c *arangoLocalStorages) Get(name string, options v1.GetOptions) (result *v1alpha.ArangoLocalStorage, err error) {
	result = &v1alpha.ArangoLocalStorage{}
	err = c.client.Get().
		Resource("arangolocalstorages").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ArangoLocalStorages that match those selectors.
func (c *arangoLocalStorages) List(opts v1.ListOptions) (result *v1alpha.ArangoLocalStorageList, err error) {
	result = &v1alpha.ArangoLocalStorageList{}
	err = c.client.Get().
		Resource("arangolocalstorages").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested arangoLocalStorages.
func (c *arangoLocalStorages) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Resource("arangolocalstorages").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a arangoLocalStorage and creates it.  Returns the server's representation of the arangoLocalStorage, and an error, if there is any.
func (c *arangoLocalStorages) Create(arangoLocalStorage *v1alpha.ArangoLocalStorage) (result *v1alpha.ArangoLocalStorage, err error) {
	result = &v1alpha.ArangoLocalStorage{}
	err = c.client.Post().
		Resource("arangolocalstorages").
		Body(arangoLocalStorage).
		Do().
		Into(result)
	return
}

// Update takes the representation of a arangoLocalStorage and updates it. Returns the server's representation of the arangoLocalStorage, and an error, if there is any.
func (c *arangoLocalStorages) Update(arangoLocalStorage *v1alpha.ArangoLocalStorage) (result *v1alpha.ArangoLocalStorage, err error) {
	result = &v1alpha.ArangoLocalStorage{}
	err = c.client.Put().
		Resource("arangolocalstorages").
		Name(arangoLocalStorage.Name).
		Body(arangoLocalStorage).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *arangoLocalStorages) UpdateStatus(arangoLocalStorage *v1alpha.ArangoLocalStorage) (result *v1alpha.ArangoLocalStorage, err error) {
	result = &v1alpha.ArangoLocalStorage{}
	err = c.client.Put().
		Resource("arangolocalstorages").
		Name(arangoLocalStorage.Name).
		SubResource("status").
		Body(arangoLocalStorage).
		Do().
		Into(result)
	return
}

// Delete takes name of the arangoLocalStorage and deletes it. Returns an error if one occurs.
func (c *arangoLocalStorages) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("arangolocalstorages").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *arangoLocalStorages) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Resource("arangolocalstorages").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched arangoLocalStorage.
func (c *arangoLocalStorages) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha.ArangoLocalStorage, err error) {
	result = &v1alpha.ArangoLocalStorage{}
	err = c.client.Patch(pt).
		Resource("arangolocalstorages").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
