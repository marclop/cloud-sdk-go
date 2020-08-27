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
)

// NewGetProxiesFilteredGroupParams creates a new GetProxiesFilteredGroupParams object
// with the default values initialized.
func NewGetProxiesFilteredGroupParams() *GetProxiesFilteredGroupParams {
	var ()
	return &GetProxiesFilteredGroupParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetProxiesFilteredGroupParamsWithTimeout creates a new GetProxiesFilteredGroupParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetProxiesFilteredGroupParamsWithTimeout(timeout time.Duration) *GetProxiesFilteredGroupParams {
	var ()
	return &GetProxiesFilteredGroupParams{

		timeout: timeout,
	}
}

// NewGetProxiesFilteredGroupParamsWithContext creates a new GetProxiesFilteredGroupParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetProxiesFilteredGroupParamsWithContext(ctx context.Context) *GetProxiesFilteredGroupParams {
	var ()
	return &GetProxiesFilteredGroupParams{

		Context: ctx,
	}
}

// NewGetProxiesFilteredGroupParamsWithHTTPClient creates a new GetProxiesFilteredGroupParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetProxiesFilteredGroupParamsWithHTTPClient(client *http.Client) *GetProxiesFilteredGroupParams {
	var ()
	return &GetProxiesFilteredGroupParams{
		HTTPClient: client,
	}
}

/*GetProxiesFilteredGroupParams contains all the parameters to send to the API endpoint
for the get proxies filtered group operation typically these are written to a http.Request
*/
type GetProxiesFilteredGroupParams struct {

	/*ProxiesFilteredGroupID
	  "The identifier for the filtered group of proxies

	*/
	ProxiesFilteredGroupID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get proxies filtered group params
func (o *GetProxiesFilteredGroupParams) WithTimeout(timeout time.Duration) *GetProxiesFilteredGroupParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get proxies filtered group params
func (o *GetProxiesFilteredGroupParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get proxies filtered group params
func (o *GetProxiesFilteredGroupParams) WithContext(ctx context.Context) *GetProxiesFilteredGroupParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get proxies filtered group params
func (o *GetProxiesFilteredGroupParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get proxies filtered group params
func (o *GetProxiesFilteredGroupParams) WithHTTPClient(client *http.Client) *GetProxiesFilteredGroupParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get proxies filtered group params
func (o *GetProxiesFilteredGroupParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithProxiesFilteredGroupID adds the proxiesFilteredGroupID to the get proxies filtered group params
func (o *GetProxiesFilteredGroupParams) WithProxiesFilteredGroupID(proxiesFilteredGroupID string) *GetProxiesFilteredGroupParams {
	o.SetProxiesFilteredGroupID(proxiesFilteredGroupID)
	return o
}

// SetProxiesFilteredGroupID adds the proxiesFilteredGroupId to the get proxies filtered group params
func (o *GetProxiesFilteredGroupParams) SetProxiesFilteredGroupID(proxiesFilteredGroupID string) {
	o.ProxiesFilteredGroupID = proxiesFilteredGroupID
}

// WriteToRequest writes these params to a swagger request
func (o *GetProxiesFilteredGroupParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param proxies_filtered_group_id
	if err := r.SetPathParam("proxies_filtered_group_id", o.ProxiesFilteredGroupID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
