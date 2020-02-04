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
)

// BoolQuery A query for documents that match boolean combinations of other queries.
// swagger:model BoolQuery
type BoolQuery struct {

	// filter
	Filter []*QueryContainer `json:"filter,omitempty"`

	// The minimum number of optional should clauses to match.
	MinimumShouldMatch int32 `json:"minimum_should_match,omitempty"`

	// must
	Must []*QueryContainer `json:"must,omitempty"`

	// must not
	MustNot []*QueryContainer `json:"must_not,omitempty"`

	// should
	Should []*QueryContainer `json:"should,omitempty"`
}

// Validate validates this bool query
func (m *BoolQuery) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFilter(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMust(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMustNot(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateShould(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BoolQuery) validateFilter(formats strfmt.Registry) error {

	if swag.IsZero(m.Filter) { // not required
		return nil
	}

	for i := 0; i < len(m.Filter); i++ {
		if swag.IsZero(m.Filter[i]) { // not required
			continue
		}

		if m.Filter[i] != nil {
			if err := m.Filter[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("filter" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *BoolQuery) validateMust(formats strfmt.Registry) error {

	if swag.IsZero(m.Must) { // not required
		return nil
	}

	for i := 0; i < len(m.Must); i++ {
		if swag.IsZero(m.Must[i]) { // not required
			continue
		}

		if m.Must[i] != nil {
			if err := m.Must[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("must" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *BoolQuery) validateMustNot(formats strfmt.Registry) error {

	if swag.IsZero(m.MustNot) { // not required
		return nil
	}

	for i := 0; i < len(m.MustNot); i++ {
		if swag.IsZero(m.MustNot[i]) { // not required
			continue
		}

		if m.MustNot[i] != nil {
			if err := m.MustNot[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("must_not" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *BoolQuery) validateShould(formats strfmt.Registry) error {

	if swag.IsZero(m.Should) { // not required
		return nil
	}

	for i := 0; i < len(m.Should); i++ {
		if swag.IsZero(m.Should[i]) { // not required
			continue
		}

		if m.Should[i] != nil {
			if err := m.Should[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("should" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *BoolQuery) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BoolQuery) UnmarshalBinary(b []byte) error {
	var res BoolQuery
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
