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

// GetAllocatorMetadataReader is a Reader for the GetAllocatorMetadata structure.
type GetAllocatorMetadataReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAllocatorMetadataReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAllocatorMetadataOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetAllocatorMetadataNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetAllocatorMetadataOK creates a GetAllocatorMetadataOK with default headers values
func NewGetAllocatorMetadataOK() *GetAllocatorMetadataOK {
	return &GetAllocatorMetadataOK{}
}

/*GetAllocatorMetadataOK handles this case with default header values.

The allocator metadata was successfully returned
*/
type GetAllocatorMetadataOK struct {
	/*The date-time when the resource was created (ISO format relative to UTC)
	 */
	XCloudResourceCreated string
	/*The date-time when the resource was last modified (ISO format relative to UTC)
	 */
	XCloudResourceLastModified string
	/*The resource version, which is used to avoid update conflicts with concurrent operations
	 */
	XCloudResourceVersion string

	Payload []*models.MetadataItem
}

func (o *GetAllocatorMetadataOK) Error() string {
	return fmt.Sprintf("[GET /platform/infrastructure/allocators/{allocator_id}/metadata][%d] getAllocatorMetadataOK  %+v", 200, o.Payload)
}

func (o *GetAllocatorMetadataOK) GetPayload() []*models.MetadataItem {
	return o.Payload
}

func (o *GetAllocatorMetadataOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header x-cloud-resource-created
	o.XCloudResourceCreated = response.GetHeader("x-cloud-resource-created")

	// response header x-cloud-resource-last-modified
	o.XCloudResourceLastModified = response.GetHeader("x-cloud-resource-last-modified")

	// response header x-cloud-resource-version
	o.XCloudResourceVersion = response.GetHeader("x-cloud-resource-version")

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAllocatorMetadataNotFound creates a GetAllocatorMetadataNotFound with default headers values
func NewGetAllocatorMetadataNotFound() *GetAllocatorMetadataNotFound {
	return &GetAllocatorMetadataNotFound{}
}

/*GetAllocatorMetadataNotFound handles this case with default header values.

The allocator specified by {allocator_id} cannot be found. (code: `allocators.allocator_not_found`)
*/
type GetAllocatorMetadataNotFound struct {
	/*The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

func (o *GetAllocatorMetadataNotFound) Error() string {
	return fmt.Sprintf("[GET /platform/infrastructure/allocators/{allocator_id}/metadata][%d] getAllocatorMetadataNotFound  %+v", 404, o.Payload)
}

func (o *GetAllocatorMetadataNotFound) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *GetAllocatorMetadataNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header x-cloud-error-codes
	o.XCloudErrorCodes = response.GetHeader("x-cloud-error-codes")

	o.Payload = new(models.BasicFailedReply)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
