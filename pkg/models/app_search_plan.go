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
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// AppSearchPlan app search plan
//
// swagger:model AppSearchPlan
type AppSearchPlan struct {

	// appsearch
	// Required: true
	Appsearch *AppSearchConfiguration `json:"appsearch"`

	// cluster topology
	ClusterTopology []*AppSearchTopologyElement `json:"cluster_topology"`

	// transient
	Transient *TransientAppSearchPlanConfiguration `json:"transient,omitempty"`
}

// Validate validates this app search plan
func (m *AppSearchPlan) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAppsearch(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateClusterTopology(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTransient(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AppSearchPlan) validateAppsearch(formats strfmt.Registry) error {

	if err := validate.Required("appsearch", "body", m.Appsearch); err != nil {
		return err
	}

	if m.Appsearch != nil {
		if err := m.Appsearch.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("appsearch")
			}
			return err
		}
	}

	return nil
}

func (m *AppSearchPlan) validateClusterTopology(formats strfmt.Registry) error {

	if swag.IsZero(m.ClusterTopology) { // not required
		return nil
	}

	for i := 0; i < len(m.ClusterTopology); i++ {
		if swag.IsZero(m.ClusterTopology[i]) { // not required
			continue
		}

		if m.ClusterTopology[i] != nil {
			if err := m.ClusterTopology[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("cluster_topology" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *AppSearchPlan) validateTransient(formats strfmt.Registry) error {

	if swag.IsZero(m.Transient) { // not required
		return nil
	}

	if m.Transient != nil {
		if err := m.Transient.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("transient")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AppSearchPlan) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AppSearchPlan) UnmarshalBinary(b []byte) error {
	var res AppSearchPlan
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
