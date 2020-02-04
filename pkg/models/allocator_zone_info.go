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
	"strconv"

	"github.com/go-openapi/errors"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// AllocatorZoneInfo The allocators for the specified zone.
// swagger:model AllocatorZoneInfo
type AllocatorZoneInfo struct {

	// allocators
	// Required: true
	Allocators []*AllocatorInfo `json:"allocators"`

	// Identifier of the zone
	// Required: true
	ZoneID *string `json:"zone_id"`
}

// Validate validates this allocator zone info
func (m *AllocatorZoneInfo) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAllocators(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateZoneID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AllocatorZoneInfo) validateAllocators(formats strfmt.Registry) error {

	if err := validate.Required("allocators", "body", m.Allocators); err != nil {
		return err
	}

	for i := 0; i < len(m.Allocators); i++ {
		if swag.IsZero(m.Allocators[i]) { // not required
			continue
		}

		if m.Allocators[i] != nil {
			if err := m.Allocators[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("allocators" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *AllocatorZoneInfo) validateZoneID(formats strfmt.Registry) error {

	if err := validate.Required("zone_id", "body", m.ZoneID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AllocatorZoneInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AllocatorZoneInfo) UnmarshalBinary(b []byte) error {
	var res AllocatorZoneInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
