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

// DeploymentUpdateRequest A request for updating a Deployment consisting of multiple resources
//
// swagger:model DeploymentUpdateRequest
type DeploymentUpdateRequest struct {

	// New information about the current deployment object, otherwise stays the same
	Metadata *DeploymentUpdateMetadata `json:"metadata,omitempty"`

	// A new name for the deployment, otherwise stays the same.
	Name string `json:"name,omitempty"`

	// Whether or not to prune orphan resources that are no longer mentioned in this request. Note that resourcesare tracked by ref_id, and if a resource's ref_id is changed, any previous running resources created with that previousref_id are considered to be orphaned as well.
	// Required: true
	PruneOrphans *bool `json:"prune_orphans"`

	// New information about the Resources that will have this Deployment, otherwise they stay the same
	Resources *DeploymentUpdateResources `json:"resources,omitempty"`
}

// Validate validates this deployment update request
func (m *DeploymentUpdateRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMetadata(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePruneOrphans(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResources(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DeploymentUpdateRequest) validateMetadata(formats strfmt.Registry) error {

	if swag.IsZero(m.Metadata) { // not required
		return nil
	}

	if m.Metadata != nil {
		if err := m.Metadata.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("metadata")
			}
			return err
		}
	}

	return nil
}

func (m *DeploymentUpdateRequest) validatePruneOrphans(formats strfmt.Registry) error {

	if err := validate.Required("prune_orphans", "body", m.PruneOrphans); err != nil {
		return err
	}

	return nil
}

func (m *DeploymentUpdateRequest) validateResources(formats strfmt.Registry) error {

	if swag.IsZero(m.Resources) { // not required
		return nil
	}

	if m.Resources != nil {
		if err := m.Resources.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("resources")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DeploymentUpdateRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DeploymentUpdateRequest) UnmarshalBinary(b []byte) error {
	var res DeploymentUpdateRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
