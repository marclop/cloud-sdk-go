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

// SnapshotStatusInfo Information about the snapshot status for the Elasticsearch cluster. For example, the health status.
// swagger:model SnapshotStatusInfo
type SnapshotStatusInfo struct {

	// Number of snapshots stored for this cluster
	// Required: true
	Count *int32 `json:"count"`

	// Health status of snapshots for this cluster
	// Required: true
	Healthy *bool `json:"healthy"`

	// The end time of the most recently attempted snapshot
	// Format: date-time
	LatestEndTime strfmt.DateTime `json:"latest_end_time,omitempty"`

	// Status of the latest snapshot attempt, if any exist.
	LatestStatus string `json:"latest_status,omitempty"`

	// Latest snapshot status
	LatestSuccessful *bool `json:"latest_successful,omitempty"`

	// The end time of the most recently successful snapshot
	// Format: date-time
	LatestSuccessfulEndTime strfmt.DateTime `json:"latest_successful_end_time,omitempty"`

	// Indicates whether the cluster has a relatively recent successful snapshot.
	RecentSuccess *bool `json:"recent_success,omitempty"`

	// Scheduled time of next snapshot attempt
	// Format: date-time
	ScheduledTime strfmt.DateTime `json:"scheduled_time,omitempty"`
}

// Validate validates this snapshot status info
func (m *SnapshotStatusInfo) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHealthy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLatestEndTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLatestSuccessfulEndTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateScheduledTime(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SnapshotStatusInfo) validateCount(formats strfmt.Registry) error {

	if err := validate.Required("count", "body", m.Count); err != nil {
		return err
	}

	return nil
}

func (m *SnapshotStatusInfo) validateHealthy(formats strfmt.Registry) error {

	if err := validate.Required("healthy", "body", m.Healthy); err != nil {
		return err
	}

	return nil
}

func (m *SnapshotStatusInfo) validateLatestEndTime(formats strfmt.Registry) error {

	if swag.IsZero(m.LatestEndTime) { // not required
		return nil
	}

	if err := validate.FormatOf("latest_end_time", "body", "date-time", m.LatestEndTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *SnapshotStatusInfo) validateLatestSuccessfulEndTime(formats strfmt.Registry) error {

	if swag.IsZero(m.LatestSuccessfulEndTime) { // not required
		return nil
	}

	if err := validate.FormatOf("latest_successful_end_time", "body", "date-time", m.LatestSuccessfulEndTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *SnapshotStatusInfo) validateScheduledTime(formats strfmt.Registry) error {

	if swag.IsZero(m.ScheduledTime) { // not required
		return nil
	}

	if err := validate.FormatOf("scheduled_time", "body", "date-time", m.ScheduledTime.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SnapshotStatusInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SnapshotStatusInfo) UnmarshalBinary(b []byte) error {
	var res SnapshotStatusInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
