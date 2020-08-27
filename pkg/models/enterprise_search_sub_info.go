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

// EnterpriseSearchSubInfo Information about the APM Servers associated with the Elasticsearch cluster.
//
// swagger:model EnterpriseSearchSubInfo
type EnterpriseSearchSubInfo struct {

	// Whether the associated Enterprise Search is currently available
	// Required: true
	Enabled *bool `json:"enabled"`

	// The Enterprise Search Id
	// Required: true
	EnterpriseSearchID *string `json:"enterprise_search_id"`

	// A map of application-specific operations (which map to 'operationId's in the Swagger API) to metadata about that operation
	Links map[string]Hyperlink `json:"links,omitempty"`
}

// Validate validates this enterprise search sub info
func (m *EnterpriseSearchSubInfo) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEnabled(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEnterpriseSearchID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EnterpriseSearchSubInfo) validateEnabled(formats strfmt.Registry) error {

	if err := validate.Required("enabled", "body", m.Enabled); err != nil {
		return err
	}

	return nil
}

func (m *EnterpriseSearchSubInfo) validateEnterpriseSearchID(formats strfmt.Registry) error {

	if err := validate.Required("enterprise_search_id", "body", m.EnterpriseSearchID); err != nil {
		return err
	}

	return nil
}

func (m *EnterpriseSearchSubInfo) validateLinks(formats strfmt.Registry) error {

	if swag.IsZero(m.Links) { // not required
		return nil
	}

	for k := range m.Links {

		if err := validate.Required("links"+"."+k, "body", m.Links[k]); err != nil {
			return err
		}
		if val, ok := m.Links[k]; ok {
			if err := val.Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *EnterpriseSearchSubInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EnterpriseSearchSubInfo) UnmarshalBinary(b []byte) error {
	var res EnterpriseSearchSubInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
