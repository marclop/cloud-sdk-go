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

// ElasticsearchClusterBlockingIssues Issues that prevent the Elasticsearch cluster or index from correctly operating.
//
// swagger:model ElasticsearchClusterBlockingIssues
type ElasticsearchClusterBlockingIssues struct {

	// A list of issues that affect availability of entire cluster
	// Required: true
	ClusterLevel []*ElasticsearchClusterBlockingIssueElement `json:"cluster_level"`

	// Whether the cluster has issues (false) or not (true)
	// Required: true
	Healthy *bool `json:"healthy"`

	// A list of issues that affect availability of the cluster's indices
	// Required: true
	IndexLevel []*ElasticsearchClusterBlockingIssueElement `json:"index_level"`
}

// Validate validates this elasticsearch cluster blocking issues
func (m *ElasticsearchClusterBlockingIssues) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateClusterLevel(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHealthy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIndexLevel(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ElasticsearchClusterBlockingIssues) validateClusterLevel(formats strfmt.Registry) error {

	if err := validate.Required("cluster_level", "body", m.ClusterLevel); err != nil {
		return err
	}

	for i := 0; i < len(m.ClusterLevel); i++ {
		if swag.IsZero(m.ClusterLevel[i]) { // not required
			continue
		}

		if m.ClusterLevel[i] != nil {
			if err := m.ClusterLevel[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("cluster_level" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ElasticsearchClusterBlockingIssues) validateHealthy(formats strfmt.Registry) error {

	if err := validate.Required("healthy", "body", m.Healthy); err != nil {
		return err
	}

	return nil
}

func (m *ElasticsearchClusterBlockingIssues) validateIndexLevel(formats strfmt.Registry) error {

	if err := validate.Required("index_level", "body", m.IndexLevel); err != nil {
		return err
	}

	for i := 0; i < len(m.IndexLevel); i++ {
		if swag.IsZero(m.IndexLevel[i]) { // not required
			continue
		}

		if m.IndexLevel[i] != nil {
			if err := m.IndexLevel[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("index_level" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ElasticsearchClusterBlockingIssues) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ElasticsearchClusterBlockingIssues) UnmarshalBinary(b []byte) error {
	var res ElasticsearchClusterBlockingIssues
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
