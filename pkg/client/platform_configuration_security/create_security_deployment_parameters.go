// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

// Code generated by go-swagger; DO NOT EDIT.

package platform_configuration_security

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/elastic/cloud-sdk-go/pkg/models"
)

// NewCreateSecurityDeploymentParams creates a new CreateSecurityDeploymentParams object
// with the default values initialized.
func NewCreateSecurityDeploymentParams() *CreateSecurityDeploymentParams {
	var ()
	return &CreateSecurityDeploymentParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateSecurityDeploymentParamsWithTimeout creates a new CreateSecurityDeploymentParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateSecurityDeploymentParamsWithTimeout(timeout time.Duration) *CreateSecurityDeploymentParams {
	var ()
	return &CreateSecurityDeploymentParams{

		timeout: timeout,
	}
}

// NewCreateSecurityDeploymentParamsWithContext creates a new CreateSecurityDeploymentParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateSecurityDeploymentParamsWithContext(ctx context.Context) *CreateSecurityDeploymentParams {
	var ()
	return &CreateSecurityDeploymentParams{

		Context: ctx,
	}
}

// NewCreateSecurityDeploymentParamsWithHTTPClient creates a new CreateSecurityDeploymentParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateSecurityDeploymentParamsWithHTTPClient(client *http.Client) *CreateSecurityDeploymentParams {
	var ()
	return &CreateSecurityDeploymentParams{
		HTTPClient: client,
	}
}

/*CreateSecurityDeploymentParams contains all the parameters to send to the API endpoint
for the create security deployment operation typically these are written to a http.Request
*/
type CreateSecurityDeploymentParams struct {

	/*Body
	  The deployment request

	*/
	Body *models.SecurityDeploymentCreateRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create security deployment params
func (o *CreateSecurityDeploymentParams) WithTimeout(timeout time.Duration) *CreateSecurityDeploymentParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create security deployment params
func (o *CreateSecurityDeploymentParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create security deployment params
func (o *CreateSecurityDeploymentParams) WithContext(ctx context.Context) *CreateSecurityDeploymentParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create security deployment params
func (o *CreateSecurityDeploymentParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create security deployment params
func (o *CreateSecurityDeploymentParams) WithHTTPClient(client *http.Client) *CreateSecurityDeploymentParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create security deployment params
func (o *CreateSecurityDeploymentParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create security deployment params
func (o *CreateSecurityDeploymentParams) WithBody(body *models.SecurityDeploymentCreateRequest) *CreateSecurityDeploymentParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create security deployment params
func (o *CreateSecurityDeploymentParams) SetBody(body *models.SecurityDeploymentCreateRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *CreateSecurityDeploymentParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
