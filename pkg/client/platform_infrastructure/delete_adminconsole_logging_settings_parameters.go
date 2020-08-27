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

// NewDeleteAdminconsoleLoggingSettingsParams creates a new DeleteAdminconsoleLoggingSettingsParams object
// with the default values initialized.
func NewDeleteAdminconsoleLoggingSettingsParams() *DeleteAdminconsoleLoggingSettingsParams {
	var ()
	return &DeleteAdminconsoleLoggingSettingsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteAdminconsoleLoggingSettingsParamsWithTimeout creates a new DeleteAdminconsoleLoggingSettingsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteAdminconsoleLoggingSettingsParamsWithTimeout(timeout time.Duration) *DeleteAdminconsoleLoggingSettingsParams {
	var ()
	return &DeleteAdminconsoleLoggingSettingsParams{

		timeout: timeout,
	}
}

// NewDeleteAdminconsoleLoggingSettingsParamsWithContext creates a new DeleteAdminconsoleLoggingSettingsParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteAdminconsoleLoggingSettingsParamsWithContext(ctx context.Context) *DeleteAdminconsoleLoggingSettingsParams {
	var ()
	return &DeleteAdminconsoleLoggingSettingsParams{

		Context: ctx,
	}
}

// NewDeleteAdminconsoleLoggingSettingsParamsWithHTTPClient creates a new DeleteAdminconsoleLoggingSettingsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteAdminconsoleLoggingSettingsParamsWithHTTPClient(client *http.Client) *DeleteAdminconsoleLoggingSettingsParams {
	var ()
	return &DeleteAdminconsoleLoggingSettingsParams{
		HTTPClient: client,
	}
}

/*DeleteAdminconsoleLoggingSettingsParams contains all the parameters to send to the API endpoint
for the delete adminconsole logging settings operation typically these are written to a http.Request
*/
type DeleteAdminconsoleLoggingSettingsParams struct {

	/*AdminconsoleID
	  The identifier for the adminconsole instance

	*/
	AdminconsoleID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete adminconsole logging settings params
func (o *DeleteAdminconsoleLoggingSettingsParams) WithTimeout(timeout time.Duration) *DeleteAdminconsoleLoggingSettingsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete adminconsole logging settings params
func (o *DeleteAdminconsoleLoggingSettingsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete adminconsole logging settings params
func (o *DeleteAdminconsoleLoggingSettingsParams) WithContext(ctx context.Context) *DeleteAdminconsoleLoggingSettingsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete adminconsole logging settings params
func (o *DeleteAdminconsoleLoggingSettingsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete adminconsole logging settings params
func (o *DeleteAdminconsoleLoggingSettingsParams) WithHTTPClient(client *http.Client) *DeleteAdminconsoleLoggingSettingsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete adminconsole logging settings params
func (o *DeleteAdminconsoleLoggingSettingsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAdminconsoleID adds the adminconsoleID to the delete adminconsole logging settings params
func (o *DeleteAdminconsoleLoggingSettingsParams) WithAdminconsoleID(adminconsoleID string) *DeleteAdminconsoleLoggingSettingsParams {
	o.SetAdminconsoleID(adminconsoleID)
	return o
}

// SetAdminconsoleID adds the adminconsoleId to the delete adminconsole logging settings params
func (o *DeleteAdminconsoleLoggingSettingsParams) SetAdminconsoleID(adminconsoleID string) {
	o.AdminconsoleID = adminconsoleID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteAdminconsoleLoggingSettingsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param adminconsole_id
	if err := r.SetPathParam("adminconsole_id", o.AdminconsoleID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
