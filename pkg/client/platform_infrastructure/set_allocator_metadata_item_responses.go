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

// SetAllocatorMetadataItemReader is a Reader for the SetAllocatorMetadataItem structure.
type SetAllocatorMetadataItemReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SetAllocatorMetadataItemReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSetAllocatorMetadataItemOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewSetAllocatorMetadataItemBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewSetAllocatorMetadataItemNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 449:
		result := NewSetAllocatorMetadataItemRetryWith()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewSetAllocatorMetadataItemOK creates a SetAllocatorMetadataItemOK with default headers values
func NewSetAllocatorMetadataItemOK() *SetAllocatorMetadataItemOK {
	return &SetAllocatorMetadataItemOK{}
}

/*SetAllocatorMetadataItemOK handles this case with default header values.

The allocator metadata was successfully changed (the updated JSON is returned)
*/
type SetAllocatorMetadataItemOK struct {
	Payload []*models.MetadataItem
}

func (o *SetAllocatorMetadataItemOK) Error() string {
	return fmt.Sprintf("[PUT /platform/infrastructure/allocators/{allocator_id}/metadata/{key}][%d] setAllocatorMetadataItemOK  %+v", 200, o.Payload)
}

func (o *SetAllocatorMetadataItemOK) GetPayload() []*models.MetadataItem {
	return o.Payload
}

func (o *SetAllocatorMetadataItemOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSetAllocatorMetadataItemBadRequest creates a SetAllocatorMetadataItemBadRequest with default headers values
func NewSetAllocatorMetadataItemBadRequest() *SetAllocatorMetadataItemBadRequest {
	return &SetAllocatorMetadataItemBadRequest{}
}

/*SetAllocatorMetadataItemBadRequest handles this case with default header values.

The value specified for the metadata tag is empty. (code: `allocators.invalid_empty_metadata_item`)
*/
type SetAllocatorMetadataItemBadRequest struct {
	/*The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

func (o *SetAllocatorMetadataItemBadRequest) Error() string {
	return fmt.Sprintf("[PUT /platform/infrastructure/allocators/{allocator_id}/metadata/{key}][%d] setAllocatorMetadataItemBadRequest  %+v", 400, o.Payload)
}

func (o *SetAllocatorMetadataItemBadRequest) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *SetAllocatorMetadataItemBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header x-cloud-error-codes
	o.XCloudErrorCodes = response.GetHeader("x-cloud-error-codes")

	o.Payload = new(models.BasicFailedReply)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSetAllocatorMetadataItemNotFound creates a SetAllocatorMetadataItemNotFound with default headers values
func NewSetAllocatorMetadataItemNotFound() *SetAllocatorMetadataItemNotFound {
	return &SetAllocatorMetadataItemNotFound{}
}

/*SetAllocatorMetadataItemNotFound handles this case with default header values.

* The allocator specified by {allocator_id} cannot be found. (code: `allocators.allocator_not_found`)
* The metadata item specified by {key} cannot be found. (code: `allocators.metadata_item_not_found`)
 */
type SetAllocatorMetadataItemNotFound struct {
	/*The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

func (o *SetAllocatorMetadataItemNotFound) Error() string {
	return fmt.Sprintf("[PUT /platform/infrastructure/allocators/{allocator_id}/metadata/{key}][%d] setAllocatorMetadataItemNotFound  %+v", 404, o.Payload)
}

func (o *SetAllocatorMetadataItemNotFound) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *SetAllocatorMetadataItemNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header x-cloud-error-codes
	o.XCloudErrorCodes = response.GetHeader("x-cloud-error-codes")

	o.Payload = new(models.BasicFailedReply)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSetAllocatorMetadataItemRetryWith creates a SetAllocatorMetadataItemRetryWith with default headers values
func NewSetAllocatorMetadataItemRetryWith() *SetAllocatorMetadataItemRetryWith {
	return &SetAllocatorMetadataItemRetryWith{}
}

/*SetAllocatorMetadataItemRetryWith handles this case with default header values.

Elevated permissions are required. (code: `root.unauthorized.rbac.elevated_permissions_required`)
*/
type SetAllocatorMetadataItemRetryWith struct {
	/*The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

func (o *SetAllocatorMetadataItemRetryWith) Error() string {
	return fmt.Sprintf("[PUT /platform/infrastructure/allocators/{allocator_id}/metadata/{key}][%d] setAllocatorMetadataItemRetryWith  %+v", 449, o.Payload)
}

func (o *SetAllocatorMetadataItemRetryWith) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *SetAllocatorMetadataItemRetryWith) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header x-cloud-error-codes
	o.XCloudErrorCodes = response.GetHeader("x-cloud-error-codes")

	o.Payload = new(models.BasicFailedReply)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
