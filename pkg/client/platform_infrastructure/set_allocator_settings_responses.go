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

// SetAllocatorSettingsReader is a Reader for the SetAllocatorSettings structure.
type SetAllocatorSettingsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SetAllocatorSettingsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSetAllocatorSettingsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewSetAllocatorSettingsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 449:
		result := NewSetAllocatorSettingsRetryWith()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewSetAllocatorSettingsOK creates a SetAllocatorSettingsOK with default headers values
func NewSetAllocatorSettingsOK() *SetAllocatorSettingsOK {
	return &SetAllocatorSettingsOK{}
}

/*SetAllocatorSettingsOK handles this case with default header values.

Returns the updated settings for the specified allocator
*/
type SetAllocatorSettingsOK struct {
	Payload *models.AllocatorSettings
}

func (o *SetAllocatorSettingsOK) Error() string {
	return fmt.Sprintf("[PUT /platform/infrastructure/allocators/{allocator_id}/settings][%d] setAllocatorSettingsOK  %+v", 200, o.Payload)
}

func (o *SetAllocatorSettingsOK) GetPayload() *models.AllocatorSettings {
	return o.Payload
}

func (o *SetAllocatorSettingsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AllocatorSettings)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSetAllocatorSettingsNotFound creates a SetAllocatorSettingsNotFound with default headers values
func NewSetAllocatorSettingsNotFound() *SetAllocatorSettingsNotFound {
	return &SetAllocatorSettingsNotFound{}
}

/*SetAllocatorSettingsNotFound handles this case with default header values.

The allocator specified by {allocator_id} cannot be found. (code: `allocators.allocator_not_found`)
*/
type SetAllocatorSettingsNotFound struct {
	/*The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

func (o *SetAllocatorSettingsNotFound) Error() string {
	return fmt.Sprintf("[PUT /platform/infrastructure/allocators/{allocator_id}/settings][%d] setAllocatorSettingsNotFound  %+v", 404, o.Payload)
}

func (o *SetAllocatorSettingsNotFound) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *SetAllocatorSettingsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header x-cloud-error-codes
	o.XCloudErrorCodes = response.GetHeader("x-cloud-error-codes")

	o.Payload = new(models.BasicFailedReply)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSetAllocatorSettingsRetryWith creates a SetAllocatorSettingsRetryWith with default headers values
func NewSetAllocatorSettingsRetryWith() *SetAllocatorSettingsRetryWith {
	return &SetAllocatorSettingsRetryWith{}
}

/*SetAllocatorSettingsRetryWith handles this case with default header values.

Elevated permissions are required. (code: `root.unauthorized.rbac.elevated_permissions_required`)
*/
type SetAllocatorSettingsRetryWith struct {
	/*The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

func (o *SetAllocatorSettingsRetryWith) Error() string {
	return fmt.Sprintf("[PUT /platform/infrastructure/allocators/{allocator_id}/settings][%d] setAllocatorSettingsRetryWith  %+v", 449, o.Payload)
}

func (o *SetAllocatorSettingsRetryWith) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *SetAllocatorSettingsRetryWith) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header x-cloud-error-codes
	o.XCloudErrorCodes = response.GetHeader("x-cloud-error-codes")

	o.Payload = new(models.BasicFailedReply)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
