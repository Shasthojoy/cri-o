/*
Copyright 2017 The Kubernetes Authors.

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

package v1alpha1

import (
	fmt "fmt"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	serializer "k8s.io/apimachinery/pkg/runtime/serializer"
	rest "k8s.io/client-go/rest"
	api "k8s.io/kubernetes/pkg/api"
)

type ApiregistrationV1alpha1Interface interface {
	RESTClient() rest.Interface
	APIServicesGetter
}

// ApiregistrationV1alpha1Client is used to interact with features provided by the apiregistration.k8s.io group.
type ApiregistrationV1alpha1Client struct {
	restClient rest.Interface
}

func (c *ApiregistrationV1alpha1Client) APIServices() APIServiceInterface {
	return newAPIServices(c)
}

// NewForConfig creates a new ApiregistrationV1alpha1Client for the given config.
func NewForConfig(c *rest.Config) (*ApiregistrationV1alpha1Client, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	client, err := rest.RESTClientFor(&config)
	if err != nil {
		return nil, err
	}
	return &ApiregistrationV1alpha1Client{client}, nil
}

// NewForConfigOrDie creates a new ApiregistrationV1alpha1Client for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *ApiregistrationV1alpha1Client {
	client, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return client
}

// New creates a new ApiregistrationV1alpha1Client for the given RESTClient.
func New(c rest.Interface) *ApiregistrationV1alpha1Client {
	return &ApiregistrationV1alpha1Client{c}
}

func setConfigDefaults(config *rest.Config) error {
	gv, err := schema.ParseGroupVersion("apiregistration.k8s.io/v1alpha1")
	if err != nil {
		return err
	}
	// if apiregistration.k8s.io/v1alpha1 is not enabled, return an error
	if !api.Registry.IsEnabledVersion(gv) {
		return fmt.Errorf("apiregistration.k8s.io/v1alpha1 is not enabled")
	}
	config.APIPath = "/apis"
	if config.UserAgent == "" {
		config.UserAgent = rest.DefaultKubernetesUserAgent()
	}
	copyGroupVersion := gv
	config.GroupVersion = &copyGroupVersion

	config.NegotiatedSerializer = serializer.DirectCodecFactory{CodecFactory: api.Codecs}

	return nil
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *ApiregistrationV1alpha1Client) RESTClient() rest.Interface {
	if c == nil {
		return nil
	}
	return c.restClient
}