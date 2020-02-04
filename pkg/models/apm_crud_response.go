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

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ApmCrudResponse The response to an APM CRUD (create/update-plan) request.
// swagger:model ApmCrudResponse
type ApmCrudResponse struct {

	// For an operation creating or updating an APM server, the Id of that server
	ApmID string `json:"apm_id,omitempty"`

	// If the endpoint is called with URL param 'validate_only=true', then this contains advanced debug info (the internal plan representation)
	Diagnostics interface{} `json:"diagnostics,omitempty"`

	// The secret token for accessing the server
	SecretToken string `json:"secret_token,omitempty"`
}

// Validate validates this apm crud response
func (m *ApmCrudResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ApmCrudResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ApmCrudResponse) UnmarshalBinary(b []byte) error {
	var res ApmCrudResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
