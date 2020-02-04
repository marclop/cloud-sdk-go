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
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/elastic/cloud-sdk-go/pkg/models"
)

// NewSetProxiesSettingsParams creates a new SetProxiesSettingsParams object
// with the default values initialized.
func NewSetProxiesSettingsParams() *SetProxiesSettingsParams {
	var ()
	return &SetProxiesSettingsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewSetProxiesSettingsParamsWithTimeout creates a new SetProxiesSettingsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewSetProxiesSettingsParamsWithTimeout(timeout time.Duration) *SetProxiesSettingsParams {
	var ()
	return &SetProxiesSettingsParams{

		timeout: timeout,
	}
}

// NewSetProxiesSettingsParamsWithContext creates a new SetProxiesSettingsParams object
// with the default values initialized, and the ability to set a context for a request
func NewSetProxiesSettingsParamsWithContext(ctx context.Context) *SetProxiesSettingsParams {
	var ()
	return &SetProxiesSettingsParams{

		Context: ctx,
	}
}

// NewSetProxiesSettingsParamsWithHTTPClient creates a new SetProxiesSettingsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewSetProxiesSettingsParamsWithHTTPClient(client *http.Client) *SetProxiesSettingsParams {
	var ()
	return &SetProxiesSettingsParams{
		HTTPClient: client,
	}
}

/*SetProxiesSettingsParams contains all the parameters to send to the API endpoint
for the set proxies settings operation typically these are written to a http.Request
*/
type SetProxiesSettingsParams struct {

	/*Body
	  The proxy settings to apply

	*/
	Body *models.ProxiesSettings
	/*Version
	  If specified, checks for conflicts against the version of the settings (returned in 'x-cloud-resource-version' of the GET request)

	*/
	Version *int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the set proxies settings params
func (o *SetProxiesSettingsParams) WithTimeout(timeout time.Duration) *SetProxiesSettingsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the set proxies settings params
func (o *SetProxiesSettingsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the set proxies settings params
func (o *SetProxiesSettingsParams) WithContext(ctx context.Context) *SetProxiesSettingsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the set proxies settings params
func (o *SetProxiesSettingsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the set proxies settings params
func (o *SetProxiesSettingsParams) WithHTTPClient(client *http.Client) *SetProxiesSettingsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the set proxies settings params
func (o *SetProxiesSettingsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the set proxies settings params
func (o *SetProxiesSettingsParams) WithBody(body *models.ProxiesSettings) *SetProxiesSettingsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the set proxies settings params
func (o *SetProxiesSettingsParams) SetBody(body *models.ProxiesSettings) {
	o.Body = body
}

// WithVersion adds the version to the set proxies settings params
func (o *SetProxiesSettingsParams) WithVersion(version *int64) *SetProxiesSettingsParams {
	o.SetVersion(version)
	return o
}

// SetVersion adds the version to the set proxies settings params
func (o *SetProxiesSettingsParams) SetVersion(version *int64) {
	o.Version = version
}

// WriteToRequest writes these params to a swagger request
func (o *SetProxiesSettingsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
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
