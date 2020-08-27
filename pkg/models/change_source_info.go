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
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ChangeSourceInfo A container for information about the source of a change.
//
// swagger:model ChangeSourceInfo
type ChangeSourceInfo struct {

	// The type of plan change that was initiated
	// Required: true
	Action *string `json:"action"`

	// The admin user that requested the change
	AdminID string `json:"admin_id,omitempty"`

	// The time the change was initiated
	// Required: true
	// Format: date-time
	Date *strfmt.DateTime `json:"date"`

	// The service where the change originated from
	// Required: true
	Facilitator *string `json:"facilitator"`

	// The host addresses of the user that originated the change
	RemoteAddresses []string `json:"remote_addresses"`

	// The user that requested the change
	UserID string `json:"user_id,omitempty"`
}

// Validate validates this change source info
func (m *ChangeSourceInfo) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAction(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFacilitator(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ChangeSourceInfo) validateAction(formats strfmt.Registry) error {

	if err := validate.Required("action", "body", m.Action); err != nil {
		return err
	}

	return nil
}

func (m *ChangeSourceInfo) validateDate(formats strfmt.Registry) error {

	if err := validate.Required("date", "body", m.Date); err != nil {
		return err
	}

	if err := validate.FormatOf("date", "body", "date-time", m.Date.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ChangeSourceInfo) validateFacilitator(formats strfmt.Registry) error {

	if err := validate.Required("facilitator", "body", m.Facilitator); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ChangeSourceInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ChangeSourceInfo) UnmarshalBinary(b []byte) error {
	var res ChangeSourceInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
