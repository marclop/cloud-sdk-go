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

// ProxiesSSOSettings The single sign-on settings for all proxies.
//
// swagger:model ProxiesSSOSettings
type ProxiesSSOSettings struct {

	// Name of the HTTP cookie used for single-sign-on
	// Required: true
	CookieName *string `json:"cookie_name"`

	// Default path where users are redirected after a successful single-sign-on
	// Required: true
	DefaultRedirectPath *string `json:"default_redirect_path"`

	// If true, don't log requests
	// Required: true
	DontLogRequests *bool `json:"dont_log_requests"`

	// Name of the cookie that bypasses maintenance
	// Required: true
	MaintenanceBypassCookieName *string `json:"maintenance_bypass_cookie_name"`

	// Maximum age of single-sign-on token in milliseconds
	// Required: true
	MaxAge *int64 `json:"max_age"`

	// Secret string for single-sign-on
	// Required: true
	SsoSecret *string `json:"sso_secret"`
}

// Validate validates this proxies s s o settings
func (m *ProxiesSSOSettings) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCookieName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDefaultRedirectPath(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDontLogRequests(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMaintenanceBypassCookieName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMaxAge(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSsoSecret(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ProxiesSSOSettings) validateCookieName(formats strfmt.Registry) error {

	if err := validate.Required("cookie_name", "body", m.CookieName); err != nil {
		return err
	}

	return nil
}

func (m *ProxiesSSOSettings) validateDefaultRedirectPath(formats strfmt.Registry) error {

	if err := validate.Required("default_redirect_path", "body", m.DefaultRedirectPath); err != nil {
		return err
	}

	return nil
}

func (m *ProxiesSSOSettings) validateDontLogRequests(formats strfmt.Registry) error {

	if err := validate.Required("dont_log_requests", "body", m.DontLogRequests); err != nil {
		return err
	}

	return nil
}

func (m *ProxiesSSOSettings) validateMaintenanceBypassCookieName(formats strfmt.Registry) error {

	if err := validate.Required("maintenance_bypass_cookie_name", "body", m.MaintenanceBypassCookieName); err != nil {
		return err
	}

	return nil
}

func (m *ProxiesSSOSettings) validateMaxAge(formats strfmt.Registry) error {

	if err := validate.Required("max_age", "body", m.MaxAge); err != nil {
		return err
	}

	return nil
}

func (m *ProxiesSSOSettings) validateSsoSecret(formats strfmt.Registry) error {

	if err := validate.Required("sso_secret", "body", m.SsoSecret); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ProxiesSSOSettings) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ProxiesSSOSettings) UnmarshalBinary(b []byte) error {
	var res ProxiesSSOSettings
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
