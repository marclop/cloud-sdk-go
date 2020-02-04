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

// NewMoveClustersParams creates a new MoveClustersParams object
// with the default values initialized.
func NewMoveClustersParams() *MoveClustersParams {
	var (
		forceUpdateDefault  = bool(false)
		moveOnlyDefault     = bool(true)
		validateOnlyDefault = bool(false)
	)
	return &MoveClustersParams{
		ForceUpdate:  &forceUpdateDefault,
		MoveOnly:     &moveOnlyDefault,
		ValidateOnly: &validateOnlyDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewMoveClustersParamsWithTimeout creates a new MoveClustersParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewMoveClustersParamsWithTimeout(timeout time.Duration) *MoveClustersParams {
	var (
		forceUpdateDefault  = bool(false)
		moveOnlyDefault     = bool(true)
		validateOnlyDefault = bool(false)
	)
	return &MoveClustersParams{
		ForceUpdate:  &forceUpdateDefault,
		MoveOnly:     &moveOnlyDefault,
		ValidateOnly: &validateOnlyDefault,

		timeout: timeout,
	}
}

// NewMoveClustersParamsWithContext creates a new MoveClustersParams object
// with the default values initialized, and the ability to set a context for a request
func NewMoveClustersParamsWithContext(ctx context.Context) *MoveClustersParams {
	var (
		forceUpdateDefault  = bool(false)
		moveOnlyDefault     = bool(true)
		validateOnlyDefault = bool(false)
	)
	return &MoveClustersParams{
		ForceUpdate:  &forceUpdateDefault,
		MoveOnly:     &moveOnlyDefault,
		ValidateOnly: &validateOnlyDefault,

		Context: ctx,
	}
}

// NewMoveClustersParamsWithHTTPClient creates a new MoveClustersParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewMoveClustersParamsWithHTTPClient(client *http.Client) *MoveClustersParams {
	var (
		forceUpdateDefault  = bool(false)
		moveOnlyDefault     = bool(true)
		validateOnlyDefault = bool(false)
	)
	return &MoveClustersParams{
		ForceUpdate:  &forceUpdateDefault,
		MoveOnly:     &moveOnlyDefault,
		ValidateOnly: &validateOnlyDefault,
		HTTPClient:   client,
	}
}

/*MoveClustersParams contains all the parameters to send to the API endpoint
for the move clusters operation typically these are written to a http.Request
*/
type MoveClustersParams struct {

	/*AllocatorDown
	  When `true`, considers all instances on the allocator as permanently shut down when deciding how to migrate data to new nodes.When left blank, the system automatically decides. NOTE: The default treats the allocator as up.

	*/
	AllocatorDown *bool
	/*AllocatorID
	  The allocator identifier.

	*/
	AllocatorID string
	/*Body
	  Overrides defaults for the move of each cluster

	*/
	Body *models.MoveClustersRequest
	/*ForceUpdate
	  When `true`, cancels and overwrites the pending plans, or treats the instance as an error.

	*/
	ForceUpdate *bool
	/*MoveOnly
	  When `true`, moves the specified instances and ignores the changes for the cluster state.

	*/
	MoveOnly *bool
	/*ValidateOnly
	  When `true`, validates the plan overrides, then returns the plan without performing the move.

	*/
	ValidateOnly *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the move clusters params
func (o *MoveClustersParams) WithTimeout(timeout time.Duration) *MoveClustersParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the move clusters params
func (o *MoveClustersParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the move clusters params
func (o *MoveClustersParams) WithContext(ctx context.Context) *MoveClustersParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the move clusters params
func (o *MoveClustersParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the move clusters params
func (o *MoveClustersParams) WithHTTPClient(client *http.Client) *MoveClustersParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the move clusters params
func (o *MoveClustersParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAllocatorDown adds the allocatorDown to the move clusters params
func (o *MoveClustersParams) WithAllocatorDown(allocatorDown *bool) *MoveClustersParams {
	o.SetAllocatorDown(allocatorDown)
	return o
}

// SetAllocatorDown adds the allocatorDown to the move clusters params
func (o *MoveClustersParams) SetAllocatorDown(allocatorDown *bool) {
	o.AllocatorDown = allocatorDown
}

// WithAllocatorID adds the allocatorID to the move clusters params
func (o *MoveClustersParams) WithAllocatorID(allocatorID string) *MoveClustersParams {
	o.SetAllocatorID(allocatorID)
	return o
}

// SetAllocatorID adds the allocatorId to the move clusters params
func (o *MoveClustersParams) SetAllocatorID(allocatorID string) {
	o.AllocatorID = allocatorID
}

// WithBody adds the body to the move clusters params
func (o *MoveClustersParams) WithBody(body *models.MoveClustersRequest) *MoveClustersParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the move clusters params
func (o *MoveClustersParams) SetBody(body *models.MoveClustersRequest) {
	o.Body = body
}

// WithForceUpdate adds the forceUpdate to the move clusters params
func (o *MoveClustersParams) WithForceUpdate(forceUpdate *bool) *MoveClustersParams {
	o.SetForceUpdate(forceUpdate)
	return o
}

// SetForceUpdate adds the forceUpdate to the move clusters params
func (o *MoveClustersParams) SetForceUpdate(forceUpdate *bool) {
	o.ForceUpdate = forceUpdate
}

// WithMoveOnly adds the moveOnly to the move clusters params
func (o *MoveClustersParams) WithMoveOnly(moveOnly *bool) *MoveClustersParams {
	o.SetMoveOnly(moveOnly)
	return o
}

// SetMoveOnly adds the moveOnly to the move clusters params
func (o *MoveClustersParams) SetMoveOnly(moveOnly *bool) {
	o.MoveOnly = moveOnly
}

// WithValidateOnly adds the validateOnly to the move clusters params
func (o *MoveClustersParams) WithValidateOnly(validateOnly *bool) *MoveClustersParams {
	o.SetValidateOnly(validateOnly)
	return o
}

// SetValidateOnly adds the validateOnly to the move clusters params
func (o *MoveClustersParams) SetValidateOnly(validateOnly *bool) {
	o.ValidateOnly = validateOnly
}

// WriteToRequest writes these params to a swagger request
func (o *MoveClustersParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.AllocatorDown != nil {

		// query param allocator_down
		var qrAllocatorDown bool
		if o.AllocatorDown != nil {
			qrAllocatorDown = *o.AllocatorDown
		}
		qAllocatorDown := swag.FormatBool(qrAllocatorDown)
		if qAllocatorDown != "" {
			if err := r.SetQueryParam("allocator_down", qAllocatorDown); err != nil {
				return err
			}
		}

	}

	// path param allocator_id
	if err := r.SetPathParam("allocator_id", o.AllocatorID); err != nil {
		return err
	}

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if o.ForceUpdate != nil {

		// query param force_update
		var qrForceUpdate bool
		if o.ForceUpdate != nil {
			qrForceUpdate = *o.ForceUpdate
		}
		qForceUpdate := swag.FormatBool(qrForceUpdate)
		if qForceUpdate != "" {
			if err := r.SetQueryParam("force_update", qForceUpdate); err != nil {
				return err
			}
		}

	}

	if o.MoveOnly != nil {

		// query param move_only
		var qrMoveOnly bool
		if o.MoveOnly != nil {
			qrMoveOnly = *o.MoveOnly
		}
		qMoveOnly := swag.FormatBool(qrMoveOnly)
		if qMoveOnly != "" {
			if err := r.SetQueryParam("move_only", qMoveOnly); err != nil {
				return err
			}
		}

	}

	if o.ValidateOnly != nil {

		// query param validate_only
		var qrValidateOnly bool
		if o.ValidateOnly != nil {
			qrValidateOnly = *o.ValidateOnly
		}
		qValidateOnly := swag.FormatBool(qrValidateOnly)
		if qValidateOnly != "" {
			if err := r.SetQueryParam("validate_only", qValidateOnly); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
