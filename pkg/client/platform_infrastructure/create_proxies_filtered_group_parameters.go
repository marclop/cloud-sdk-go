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

package platform_infrastructure

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

// NewCreateProxiesFilteredGroupParams creates a new CreateProxiesFilteredGroupParams object
// with the default values initialized.
func NewCreateProxiesFilteredGroupParams() *CreateProxiesFilteredGroupParams {
	var ()
	return &CreateProxiesFilteredGroupParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateProxiesFilteredGroupParamsWithTimeout creates a new CreateProxiesFilteredGroupParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateProxiesFilteredGroupParamsWithTimeout(timeout time.Duration) *CreateProxiesFilteredGroupParams {
	var ()
	return &CreateProxiesFilteredGroupParams{

		timeout: timeout,
	}
}

// NewCreateProxiesFilteredGroupParamsWithContext creates a new CreateProxiesFilteredGroupParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateProxiesFilteredGroupParamsWithContext(ctx context.Context) *CreateProxiesFilteredGroupParams {
	var ()
	return &CreateProxiesFilteredGroupParams{

		Context: ctx,
	}
}

// NewCreateProxiesFilteredGroupParamsWithHTTPClient creates a new CreateProxiesFilteredGroupParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateProxiesFilteredGroupParamsWithHTTPClient(client *http.Client) *CreateProxiesFilteredGroupParams {
	var ()
	return &CreateProxiesFilteredGroupParams{
		HTTPClient: client,
	}
}

/*CreateProxiesFilteredGroupParams contains all the parameters to send to the API endpoint
for the create proxies filtered group operation typically these are written to a http.Request
*/
type CreateProxiesFilteredGroupParams struct {

	/*Body
	  Data for the filtered group of proxies to create

	*/
	Body *models.ProxiesFilteredGroup

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create proxies filtered group params
func (o *CreateProxiesFilteredGroupParams) WithTimeout(timeout time.Duration) *CreateProxiesFilteredGroupParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create proxies filtered group params
func (o *CreateProxiesFilteredGroupParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create proxies filtered group params
func (o *CreateProxiesFilteredGroupParams) WithContext(ctx context.Context) *CreateProxiesFilteredGroupParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create proxies filtered group params
func (o *CreateProxiesFilteredGroupParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create proxies filtered group params
func (o *CreateProxiesFilteredGroupParams) WithHTTPClient(client *http.Client) *CreateProxiesFilteredGroupParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create proxies filtered group params
func (o *CreateProxiesFilteredGroupParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create proxies filtered group params
func (o *CreateProxiesFilteredGroupParams) WithBody(body *models.ProxiesFilteredGroup) *CreateProxiesFilteredGroupParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create proxies filtered group params
func (o *CreateProxiesFilteredGroupParams) SetBody(body *models.ProxiesFilteredGroup) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *CreateProxiesFilteredGroupParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
