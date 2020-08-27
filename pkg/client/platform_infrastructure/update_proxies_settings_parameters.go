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
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewUpdateProxiesSettingsParams creates a new UpdateProxiesSettingsParams object
// with the default values initialized.
func NewUpdateProxiesSettingsParams() *UpdateProxiesSettingsParams {
	var ()
	return &UpdateProxiesSettingsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateProxiesSettingsParamsWithTimeout creates a new UpdateProxiesSettingsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateProxiesSettingsParamsWithTimeout(timeout time.Duration) *UpdateProxiesSettingsParams {
	var ()
	return &UpdateProxiesSettingsParams{

		timeout: timeout,
	}
}

// NewUpdateProxiesSettingsParamsWithContext creates a new UpdateProxiesSettingsParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateProxiesSettingsParamsWithContext(ctx context.Context) *UpdateProxiesSettingsParams {
	var ()
	return &UpdateProxiesSettingsParams{

		Context: ctx,
	}
}

// NewUpdateProxiesSettingsParamsWithHTTPClient creates a new UpdateProxiesSettingsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateProxiesSettingsParamsWithHTTPClient(client *http.Client) *UpdateProxiesSettingsParams {
	var ()
	return &UpdateProxiesSettingsParams{
		HTTPClient: client,
	}
}

/*UpdateProxiesSettingsParams contains all the parameters to send to the API endpoint
for the update proxies settings operation typically these are written to a http.Request
*/
type UpdateProxiesSettingsParams struct {

	/*Body
	  A JSON to merge with the existing settings

	*/
	Body string
	/*Version
	  If specified, checks for conflicts against the version of the repository configuration (returned in 'x-cloud-resource-version' of the GET request)

	*/
	Version *int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update proxies settings params
func (o *UpdateProxiesSettingsParams) WithTimeout(timeout time.Duration) *UpdateProxiesSettingsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update proxies settings params
func (o *UpdateProxiesSettingsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update proxies settings params
func (o *UpdateProxiesSettingsParams) WithContext(ctx context.Context) *UpdateProxiesSettingsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update proxies settings params
func (o *UpdateProxiesSettingsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update proxies settings params
func (o *UpdateProxiesSettingsParams) WithHTTPClient(client *http.Client) *UpdateProxiesSettingsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update proxies settings params
func (o *UpdateProxiesSettingsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update proxies settings params
func (o *UpdateProxiesSettingsParams) WithBody(body string) *UpdateProxiesSettingsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update proxies settings params
func (o *UpdateProxiesSettingsParams) SetBody(body string) {
	o.Body = body
}

// WithVersion adds the version to the update proxies settings params
func (o *UpdateProxiesSettingsParams) WithVersion(version *int64) *UpdateProxiesSettingsParams {
	o.SetVersion(version)
	return o
}

// SetVersion adds the version to the update proxies settings params
func (o *UpdateProxiesSettingsParams) SetVersion(version *int64) {
	o.Version = version
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateProxiesSettingsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	if o.Version != nil {

		// query param version
		var qrVersion int64
		if o.Version != nil {
			qrVersion = *o.Version
		}
		qVersion := swag.FormatInt64(qrVersion)
		if qVersion != "" {
			if err := r.SetQueryParam("version", qVersion); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
