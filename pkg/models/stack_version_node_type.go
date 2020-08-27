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

// StackVersionNodeType The configuration for an Elastic Stack node type.
//
// swagger:model StackVersionNodeType
type StackVersionNodeType struct {

	// capacity constraints
	CapacityConstraints *StackVersionInstanceCapacityConstraint `json:"capacity_constraints,omitempty"`

	// Node types that are compatible with this one
	CompatibleNodeTypes []string `json:"compatible_node_types"`

	// Description of the node type
	// Required: true
	Description *string `json:"description"`

	// Flag to specify a node type is mandatory in a cluster's plan
	Mandatory *bool `json:"mandatory,omitempty"`

	// Name of the node type
	// Required: true
	Name *string `json:"name"`

	// Type of the node (master, data, ...)
	// Required: true
	NodeType *string `json:"node_type"`

	// Settings that are applied to all nodes of this type
	Settings interface{} `json:"settings,omitempty"`
}

// Validate validates this stack version node type
func (m *StackVersionNodeType) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCapacityConstraints(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNodeType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *StackVersionNodeType) validateCapacityConstraints(formats strfmt.Registry) error {

	if swag.IsZero(m.CapacityConstraints) { // not required
		return nil
	}

	if m.CapacityConstraints != nil {
		if err := m.CapacityConstraints.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("capacity_constraints")
			}
			return err
		}
	}

	return nil
}

func (m *StackVersionNodeType) validateDescription(formats strfmt.Registry) error {

	if err := validate.Required("description", "body", m.Description); err != nil {
		return err
	}

	return nil
}

func (m *StackVersionNodeType) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *StackVersionNodeType) validateNodeType(formats strfmt.Registry) error {

	if err := validate.Required("node_type", "body", m.NodeType); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *StackVersionNodeType) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StackVersionNodeType) UnmarshalBinary(b []byte) error {
	var res StackVersionNodeType
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
