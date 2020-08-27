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

// NewResyncConstructorsParams creates a new ResyncConstructorsParams object
// with the default values initialized.
func NewResyncConstructorsParams() *ResyncConstructorsParams {
	var (
		skipMatchingVersionDefault = bool(true)
	)
	return &ResyncConstructorsParams{
		SkipMatchingVersion: &skipMatchingVersionDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewResyncConstructorsParamsWithTimeout creates a new ResyncConstructorsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewResyncConstructorsParamsWithTimeout(timeout time.Duration) *ResyncConstructorsParams {
	var (
		skipMatchingVersionDefault = bool(true)
	)
	return &ResyncConstructorsParams{
		SkipMatchingVersion: &skipMatchingVersionDefault,

		timeout: timeout,
	}
}

// NewResyncConstructorsParamsWithContext creates a new ResyncConstructorsParams object
// with the default values initialized, and the ability to set a context for a request
func NewResyncConstructorsParamsWithContext(ctx context.Context) *ResyncConstructorsParams {
	var (
		skipMatchingVersionDefault = bool(true)
	)
	return &ResyncConstructorsParams{
		SkipMatchingVersion: &skipMatchingVersionDefault,

		Context: ctx,
	}
}

// NewResyncConstructorsParamsWithHTTPClient creates a new ResyncConstructorsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewResyncConstructorsParamsWithHTTPClient(client *http.Client) *ResyncConstructorsParams {
	var (
		skipMatchingVersionDefault = bool(true)
	)
	return &ResyncConstructorsParams{
		SkipMatchingVersion: &skipMatchingVersionDefault,
		HTTPClient:          client,
	}
}

/*ResyncConstructorsParams contains all the parameters to send to the API endpoint
for the resync constructors operation typically these are written to a http.Request
*/
type ResyncConstructorsParams struct {

	/*SkipMatchingVersion
	  When true, skips the document indexing when the version matches the in-memory copy.

	*/
	SkipMatchingVersion *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the resync constructors params
func (o *ResyncConstructorsParams) WithTimeout(timeout time.Duration) *ResyncConstructorsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the resync constructors params
func (o *ResyncConstructorsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the resync constructors params
func (o *ResyncConstructorsParams) WithContext(ctx context.Context) *ResyncConstructorsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the resync constructors params
func (o *ResyncConstructorsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the resync constructors params
func (o *ResyncConstructorsParams) WithHTTPClient(client *http.Client) *ResyncConstructorsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the resync constructors params
func (o *ResyncConstructorsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSkipMatchingVersion adds the skipMatchingVersion to the resync constructors params
func (o *ResyncConstructorsParams) WithSkipMatchingVersion(skipMatchingVersion *bool) *ResyncConstructorsParams {
	o.SetSkipMatchingVersion(skipMatchingVersion)
	return o
}

// SetSkipMatchingVersion adds the skipMatchingVersion to the resync constructors params
func (o *ResyncConstructorsParams) SetSkipMatchingVersion(skipMatchingVersion *bool) {
	o.SkipMatchingVersion = skipMatchingVersion
}

// WriteToRequest writes these params to a swagger request
func (o *ResyncConstructorsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.SkipMatchingVersion != nil {

		// query param skip_matching_version
		var qrSkipMatchingVersion bool
		if o.SkipMatchingVersion != nil {
			qrSkipMatchingVersion = *o.SkipMatchingVersion
		}
		qSkipMatchingVersion := swag.FormatBool(qrSkipMatchingVersion)
		if qSkipMatchingVersion != "" {
			if err := r.SetQueryParam("skip_matching_version", qSkipMatchingVersion); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
