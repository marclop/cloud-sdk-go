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
	"github.com/go-openapi/strfmt"

	"github.com/elastic/cloud-sdk-go/pkg/models"
)

// NewCreateLdapConfigurationParams creates a new CreateLdapConfigurationParams object
// with the default values initialized.
func NewCreateLdapConfigurationParams() *CreateLdapConfigurationParams {
	var ()
	return &CreateLdapConfigurationParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateLdapConfigurationParamsWithTimeout creates a new CreateLdapConfigurationParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateLdapConfigurationParamsWithTimeout(timeout time.Duration) *CreateLdapConfigurationParams {
	var ()
	return &CreateLdapConfigurationParams{

		timeout: timeout,
	}
}

// NewCreateLdapConfigurationParamsWithContext creates a new CreateLdapConfigurationParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateLdapConfigurationParamsWithContext(ctx context.Context) *CreateLdapConfigurationParams {
	var ()
	return &CreateLdapConfigurationParams{

		Context: ctx,
	}
}

// NewCreateLdapConfigurationParamsWithHTTPClient creates a new CreateLdapConfigurationParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateLdapConfigurationParamsWithHTTPClient(client *http.Client) *CreateLdapConfigurationParams {
	var ()
	return &CreateLdapConfigurationParams{
		HTTPClient: client,
	}
}

/*CreateLdapConfigurationParams contains all the parameters to send to the API endpoint
for the create ldap configuration operation typically these are written to a http.Request
*/
type CreateLdapConfigurationParams struct {

	/*Body
	  The LDAP configuration

	*/
	Body *models.LdapSettings

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create ldap configuration params
func (o *CreateLdapConfigurationParams) WithTimeout(timeout time.Duration) *CreateLdapConfigurationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create ldap configuration params
func (o *CreateLdapConfigurationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create ldap configuration params
func (o *CreateLdapConfigurationParams) WithContext(ctx context.Context) *CreateLdapConfigurationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create ldap configuration params
func (o *CreateLdapConfigurationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create ldap configuration params
func (o *CreateLdapConfigurationParams) WithHTTPClient(client *http.Client) *CreateLdapConfigurationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create ldap configuration params
func (o *CreateLdapConfigurationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create ldap configuration params
func (o *CreateLdapConfigurationParams) WithBody(body *models.LdapSettings) *CreateLdapConfigurationParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create ldap configuration params
func (o *CreateLdapConfigurationParams) SetBody(body *models.LdapSettings) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *CreateLdapConfigurationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
