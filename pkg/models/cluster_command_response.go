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

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// ClusterCommandResponse The response to an Elasticsearch cluster or Kibana instance command.
// swagger:model ClusterCommandResponse
type ClusterCommandResponse struct {

	// If validating the command only, then the calculated Apm plan that would be applied.
	CalculatedApmPlan *TransientApmPlanConfiguration `json:"calculated_apm_plan,omitempty"`

	// If validating the command only, then the calculated App search plan that would be applied.
	CalculatedAppSearchPlan *TransientAppSearchPlanConfiguration `json:"calculated_app_search_plan,omitempty"`

	// If validating the command only, then the calculated Elasticsearch plan that would be applied.
	CalculatedElasticsearchPlan *TransientElasticsearchPlanConfiguration `json:"calculated_elasticsearch_plan,omitempty"`

	// If validating the command only, then the calculated Kibana plan that would be applied.
	CalculatedKibanaPlan *TransientKibanaPlanConfiguration `json:"calculated_kibana_plan,omitempty"`
}

// Validate validates this cluster command response
func (m *ClusterCommandResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCalculatedApmPlan(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCalculatedAppSearchPlan(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCalculatedElasticsearchPlan(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCalculatedKibanaPlan(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClusterCommandResponse) validateCalculatedApmPlan(formats strfmt.Registry) error {

	if swag.IsZero(m.CalculatedApmPlan) { // not required
		return nil
	}

	if m.CalculatedApmPlan != nil {
		if err := m.CalculatedApmPlan.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("calculated_apm_plan")
			}
			return err
		}
	}

	return nil
}

func (m *ClusterCommandResponse) validateCalculatedAppSearchPlan(formats strfmt.Registry) error {

	if swag.IsZero(m.CalculatedAppSearchPlan) { // not required
		return nil
	}

	if m.CalculatedAppSearchPlan != nil {
		if err := m.CalculatedAppSearchPlan.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("calculated_app_search_plan")
			}
			return err
		}
	}

	return nil
}

func (m *ClusterCommandResponse) validateCalculatedElasticsearchPlan(formats strfmt.Registry) error {

	if swag.IsZero(m.CalculatedElasticsearchPlan) { // not required
		return nil
	}

	if m.CalculatedElasticsearchPlan != nil {
		if err := m.CalculatedElasticsearchPlan.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("calculated_elasticsearch_plan")
			}
			return err
		}
	}

	return nil
}

func (m *ClusterCommandResponse) validateCalculatedKibanaPlan(formats strfmt.Registry) error {

	if swag.IsZero(m.CalculatedKibanaPlan) { // not required
		return nil
	}

	if m.CalculatedKibanaPlan != nil {
		if err := m.CalculatedKibanaPlan.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("calculated_kibana_plan")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ClusterCommandResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClusterCommandResponse) UnmarshalBinary(b []byte) error {
	var res ClusterCommandResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
