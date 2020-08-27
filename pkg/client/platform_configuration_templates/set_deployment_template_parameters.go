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

package platform_configuration_templates

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

	"github.com/elastic/cloud-sdk-go/pkg/models"
)

// NewSetDeploymentTemplateParams creates a new SetDeploymentTemplateParams object
// with the default values initialized.
func NewSetDeploymentTemplateParams() *SetDeploymentTemplateParams {
	var (
		createOnlyDefault = bool(false)
	)
	return &SetDeploymentTemplateParams{
		CreateOnly: &createOnlyDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewSetDeploymentTemplateParamsWithTimeout creates a new SetDeploymentTemplateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewSetDeploymentTemplateParamsWithTimeout(timeout time.Duration) *SetDeploymentTemplateParams {
	var (
		createOnlyDefault = bool(false)
	)
	return &SetDeploymentTemplateParams{
		CreateOnly: &createOnlyDefault,

		timeout: timeout,
	}
}

// NewSetDeploymentTemplateParamsWithContext creates a new SetDeploymentTemplateParams object
// with the default values initialized, and the ability to set a context for a request
func NewSetDeploymentTemplateParamsWithContext(ctx context.Context) *SetDeploymentTemplateParams {
	var (
		createOnlyDefault = bool(false)
	)
	return &SetDeploymentTemplateParams{
		CreateOnly: &createOnlyDefault,

		Context: ctx,
	}
}

// NewSetDeploymentTemplateParamsWithHTTPClient creates a new SetDeploymentTemplateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewSetDeploymentTemplateParamsWithHTTPClient(client *http.Client) *SetDeploymentTemplateParams {
	var (
		createOnlyDefault = bool(false)
	)
	return &SetDeploymentTemplateParams{
		CreateOnly: &createOnlyDefault,
		HTTPClient: client,
	}
}

/*SetDeploymentTemplateParams contains all the parameters to send to the API endpoint
for the set deployment template operation typically these are written to a http.Request
*/
type SetDeploymentTemplateParams struct {

	/*Body
	  The deployment template definition.

	*/
	Body *models.DeploymentTemplateInfo
	/*CreateOnly
	  If true, will fail if the deployment template already exists at the given id

	*/
	CreateOnly *bool
	/*TemplateID
	  The identifier for the deployment template.

	*/
	TemplateID string
	/*Version
	  If specified, checks for conflicts against the version of the template (returned in 'x-cloud-resource-version' of the GET request)

	*/
	Version *int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the set deployment template params
func (o *SetDeploymentTemplateParams) WithTimeout(timeout time.Duration) *SetDeploymentTemplateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the set deployment template params
func (o *SetDeploymentTemplateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the set deployment template params
func (o *SetDeploymentTemplateParams) WithContext(ctx context.Context) *SetDeploymentTemplateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the set deployment template params
func (o *SetDeploymentTemplateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the set deployment template params
func (o *SetDeploymentTemplateParams) WithHTTPClient(client *http.Client) *SetDeploymentTemplateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the set deployment template params
func (o *SetDeploymentTemplateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the set deployment template params
func (o *SetDeploymentTemplateParams) WithBody(body *models.DeploymentTemplateInfo) *SetDeploymentTemplateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the set deployment template params
func (o *SetDeploymentTemplateParams) SetBody(body *models.DeploymentTemplateInfo) {
	o.Body = body
}

// WithCreateOnly adds the createOnly to the set deployment template params
func (o *SetDeploymentTemplateParams) WithCreateOnly(createOnly *bool) *SetDeploymentTemplateParams {
	o.SetCreateOnly(createOnly)
	return o
}

// SetCreateOnly adds the createOnly to the set deployment template params
func (o *SetDeploymentTemplateParams) SetCreateOnly(createOnly *bool) {
	o.CreateOnly = createOnly
}

// WithTemplateID adds the templateID to the set deployment template params
func (o *SetDeploymentTemplateParams) WithTemplateID(templateID string) *SetDeploymentTemplateParams {
	o.SetTemplateID(templateID)
	return o
}

// SetTemplateID adds the templateId to the set deployment template params
func (o *SetDeploymentTemplateParams) SetTemplateID(templateID string) {
	o.TemplateID = templateID
}

// WithVersion adds the version to the set deployment template params
func (o *SetDeploymentTemplateParams) WithVersion(version *int64) *SetDeploymentTemplateParams {
	o.SetVersion(version)
	return o
}

// SetVersion adds the version to the set deployment template params
func (o *SetDeploymentTemplateParams) SetVersion(version *int64) {
	o.Version = version
}

// WriteToRequest writes these params to a swagger request
func (o *SetDeploymentTemplateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if o.CreateOnly != nil {

		// query param create_only
		var qrCreateOnly bool
		if o.CreateOnly != nil {
			qrCreateOnly = *o.CreateOnly
		}
		qCreateOnly := swag.FormatBool(qrCreateOnly)
		if qCreateOnly != "" {
			if err := r.SetQueryParam("create_only", qCreateOnly); err != nil {
				return err
			}
		}

	}

	// path param template_id
	if err := r.SetPathParam("template_id", o.TemplateID); err != nil {
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
