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

// Code generated by main. DO NOT EDIT.

package v1alpha2

import (
	v1alpha2 "github.com/jetstack/cert-manager/pkg/apis/certmanager/v1alpha2"
	clientset "github.com/rancher/wrangler-api/pkg/generated/clientset/versioned/typed/certmanager/v1alpha2"
	informers "github.com/rancher/wrangler-api/pkg/generated/informers/externalversions/certmanager/v1alpha2"
	"github.com/rancher/wrangler/pkg/generic"
)

type Interface interface {
	Certificate() CertificateController
	ClusterIssuer() ClusterIssuerController
	Issuer() IssuerController
}

func New(controllerManager *generic.ControllerManager, client clientset.CertmanagerV1alpha2Interface,
	informers informers.Interface) Interface {
	return &version{
		controllerManager: controllerManager,
		client:            client,
		informers:         informers,
	}
}

type version struct {
	controllerManager *generic.ControllerManager
	informers         informers.Interface
	client            clientset.CertmanagerV1alpha2Interface
}

func (c *version) Certificate() CertificateController {
	return NewCertificateController(v1alpha2.SchemeGroupVersion.WithKind("Certificate"), c.controllerManager, c.client, c.informers.Certificates())
}
func (c *version) ClusterIssuer() ClusterIssuerController {
	return NewClusterIssuerController(v1alpha2.SchemeGroupVersion.WithKind("ClusterIssuer"), c.controllerManager, c.client, c.informers.ClusterIssuers())
}
func (c *version) Issuer() IssuerController {
	return NewIssuerController(v1alpha2.SchemeGroupVersion.WithKind("Issuer"), c.controllerManager, c.client, c.informers.Issuers())
}
