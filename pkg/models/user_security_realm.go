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
	"encoding/json"

	"github.com/go-openapi/errors"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// UserSecurityRealm The security realm the user belongs to
// swagger:model UserSecurityRealm
type UserSecurityRealm struct {

	// The identifier for the security realm
	// Required: true
	ID *string `json:"id"`

	// The type of the security realm
	// Required: true
	// Enum: [native ldap saml active_directory]
	Type *string `json:"type"`
}

// Validate validates this user security realm
func (m *UserSecurityRealm) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UserSecurityRealm) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

var userSecurityRealmTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["native","ldap","saml","active_directory"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		userSecurityRealmTypeTypePropEnum = append(userSecurityRealmTypeTypePropEnum, v)
	}
}

const (

	// UserSecurityRealmTypeNative captures enum value "native"
	UserSecurityRealmTypeNative string = "native"

	// UserSecurityRealmTypeLdap captures enum value "ldap"
	UserSecurityRealmTypeLdap string = "ldap"

	// UserSecurityRealmTypeSaml captures enum value "saml"
	UserSecurityRealmTypeSaml string = "saml"

	// UserSecurityRealmTypeActiveDirectory captures enum value "active_directory"
	UserSecurityRealmTypeActiveDirectory string = "active_directory"
)

// prop value enum
func (m *UserSecurityRealm) validateTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, userSecurityRealmTypeTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *UserSecurityRealm) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", *m.Type); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *UserSecurityRealm) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UserSecurityRealm) UnmarshalBinary(b []byte) error {
	var res UserSecurityRealm
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
