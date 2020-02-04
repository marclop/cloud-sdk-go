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
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// SamlAttributeSettings The mapping configuration for the Elasticsearch security SAML attribute.
// swagger:model SamlAttributeSettings
type SamlAttributeSettings struct {

	// The Name of the SAML attribute that contains the user's X.50 Distinguished Name
	Dn string `json:"dn,omitempty"`

	// The Name of the SAML attribute that contains the user's groups
	// Required: true
	Groups *string `json:"groups"`

	// The Name of the SAML attribute that contains the user's email address
	Mail string `json:"mail,omitempty"`

	// The Name of the SAML attribute that contains the user's full name
	Name string `json:"name,omitempty"`

	// The Name of the SAML attribute that contains the user's principal (username)
	// Required: true
	Principal *string `json:"principal"`
}

// Validate validates this saml attribute settings
func (m *SamlAttributeSettings) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateGroups(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePrincipal(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SamlAttributeSettings) validateGroups(formats strfmt.Registry) error {

	if err := validate.Required("groups", "body", m.Groups); err != nil {
		return err
	}

	return nil
}

func (m *SamlAttributeSettings) validatePrincipal(formats strfmt.Registry) error {

	if err := validate.Required("principal", "body", m.Principal); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SamlAttributeSettings) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SamlAttributeSettings) UnmarshalBinary(b []byte) error {
	var res SamlAttributeSettings
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
