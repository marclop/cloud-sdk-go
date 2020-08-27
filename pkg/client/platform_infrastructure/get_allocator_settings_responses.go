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
	"github.com/go-openapi/strfmt"

	"github.com/elastic/cloud-sdk-go/pkg/models"
)

// GetAllocatorSettingsReader is a Reader for the GetAllocatorSettings structure.
type GetAllocatorSettingsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAllocatorSettingsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAllocatorSettingsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetAllocatorSettingsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetAllocatorSettingsOK creates a GetAllocatorSettingsOK with default headers values
func NewGetAllocatorSettingsOK() *GetAllocatorSettingsOK {
	return &GetAllocatorSettingsOK{}
}

/*GetAllocatorSettingsOK handles this case with default header values.

Returns the settings for the specified Allocator
*/
type GetAllocatorSettingsOK struct {
	/*The date-time when the resource was created (ISO format relative to UTC)
	 */
	XCloudResourceCreated string
	/*The date-time when the resource was last modified (ISO format relative to UTC)
	 */
	XCloudResourceLastModified string
	/*The resource version, which is used to avoid update conflicts with concurrent operations
	 */
	XCloudResourceVersion string

	Payload *models.AllocatorSettings
}

func (o *GetAllocatorSettingsOK) Error() string {
	return fmt.Sprintf("[GET /platform/infrastructure/allocators/{allocator_id}/settings][%d] getAllocatorSettingsOK  %+v", 200, o.Payload)
}

func (o *GetAllocatorSettingsOK) GetPayload() *models.AllocatorSettings {
	return o.Payload
}

func (o *GetAllocatorSettingsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header x-cloud-resource-created
	o.XCloudResourceCreated = response.GetHeader("x-cloud-resource-created")

	// response header x-cloud-resource-last-modified
	o.XCloudResourceLastModified = response.GetHeader("x-cloud-resource-last-modified")

	// response header x-cloud-resource-version
	o.XCloudResourceVersion = response.GetHeader("x-cloud-resource-version")

	o.Payload = new(models.AllocatorSettings)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAllocatorSettingsNotFound creates a GetAllocatorSettingsNotFound with default headers values
func NewGetAllocatorSettingsNotFound() *GetAllocatorSettingsNotFound {
	return &GetAllocatorSettingsNotFound{}
}

/*GetAllocatorSettingsNotFound handles this case with default header values.

The allocator specified by {allocator_id} cannot be found. (code: `allocators.allocator_not_found`)
*/
type GetAllocatorSettingsNotFound struct {
	/*The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

func (o *GetAllocatorSettingsNotFound) Error() string {
	return fmt.Sprintf("[GET /platform/infrastructure/allocators/{allocator_id}/settings][%d] getAllocatorSettingsNotFound  %+v", 404, o.Payload)
}

func (o *GetAllocatorSettingsNotFound) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *GetAllocatorSettingsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header x-cloud-error-codes
	o.XCloudErrorCodes = response.GetHeader("x-cloud-error-codes")

	o.Payload = new(models.BasicFailedReply)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
