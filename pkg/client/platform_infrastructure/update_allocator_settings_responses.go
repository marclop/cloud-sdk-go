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

package platform_infrastructure

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/elastic/cloud-sdk-go/pkg/models"
)

// UpdateAllocatorSettingsReader is a Reader for the UpdateAllocatorSettings structure.
type UpdateAllocatorSettingsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateAllocatorSettingsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateAllocatorSettingsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewUpdateAllocatorSettingsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 449:
		result := NewUpdateAllocatorSettingsRetryWith()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUpdateAllocatorSettingsOK creates a UpdateAllocatorSettingsOK with default headers values
func NewUpdateAllocatorSettingsOK() *UpdateAllocatorSettingsOK {
	return &UpdateAllocatorSettingsOK{}
}

/*UpdateAllocatorSettingsOK handles this case with default header values.

Returns the updated settings for the specified allocator
*/
type UpdateAllocatorSettingsOK struct {
	Payload *models.AllocatorSettings
}

func (o *UpdateAllocatorSettingsOK) Error() string {
	return fmt.Sprintf("[PATCH /platform/infrastructure/allocators/{allocator_id}/settings][%d] updateAllocatorSettingsOK  %+v", 200, o.Payload)
}

func (o *UpdateAllocatorSettingsOK) GetPayload() *models.AllocatorSettings {
	return o.Payload
}

func (o *UpdateAllocatorSettingsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AllocatorSettings)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateAllocatorSettingsNotFound creates a UpdateAllocatorSettingsNotFound with default headers values
func NewUpdateAllocatorSettingsNotFound() *UpdateAllocatorSettingsNotFound {
	return &UpdateAllocatorSettingsNotFound{}
}

/*UpdateAllocatorSettingsNotFound handles this case with default header values.

The allocator specified by {allocator_id} cannot be found. (code: `allocators.allocator_not_found`)
*/
type UpdateAllocatorSettingsNotFound struct {
	/*The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

func (o *UpdateAllocatorSettingsNotFound) Error() string {
	return fmt.Sprintf("[PATCH /platform/infrastructure/allocators/{allocator_id}/settings][%d] updateAllocatorSettingsNotFound  %+v", 404, o.Payload)
}

func (o *UpdateAllocatorSettingsNotFound) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *UpdateAllocatorSettingsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header x-cloud-error-codes
	o.XCloudErrorCodes = response.GetHeader("x-cloud-error-codes")

	o.Payload = new(models.BasicFailedReply)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateAllocatorSettingsRetryWith creates a UpdateAllocatorSettingsRetryWith with default headers values
func NewUpdateAllocatorSettingsRetryWith() *UpdateAllocatorSettingsRetryWith {
	return &UpdateAllocatorSettingsRetryWith{}
}

/*UpdateAllocatorSettingsRetryWith handles this case with default header values.

Elevated permissions are required. (code: `root.unauthorized.rbac.elevated_permissions_required`)
*/
type UpdateAllocatorSettingsRetryWith struct {
	/*The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

func (o *UpdateAllocatorSettingsRetryWith) Error() string {
	return fmt.Sprintf("[PATCH /platform/infrastructure/allocators/{allocator_id}/settings][%d] updateAllocatorSettingsRetryWith  %+v", 449, o.Payload)
}

func (o *UpdateAllocatorSettingsRetryWith) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *UpdateAllocatorSettingsRetryWith) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header x-cloud-error-codes
	o.XCloudErrorCodes = response.GetHeader("x-cloud-error-codes")

	o.Payload = new(models.BasicFailedReply)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
