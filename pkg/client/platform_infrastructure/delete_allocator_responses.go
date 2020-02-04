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

// DeleteAllocatorReader is a Reader for the DeleteAllocator structure.
type DeleteAllocatorReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteAllocatorReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteAllocatorOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteAllocatorBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteAllocatorNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 449:
		result := NewDeleteAllocatorRetryWith()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteAllocatorOK creates a DeleteAllocatorOK with default headers values
func NewDeleteAllocatorOK() *DeleteAllocatorOK {
	return &DeleteAllocatorOK{}
}

/*DeleteAllocatorOK handles this case with default header values.

The allocator specified by {allocator_id} was successfully deleted
*/
type DeleteAllocatorOK struct {
	Payload models.EmptyResponse
}

func (o *DeleteAllocatorOK) Error() string {
	return fmt.Sprintf("[DELETE /platform/infrastructure/allocators/{allocator_id}][%d] deleteAllocatorOK  %+v", 200, o.Payload)
}

func (o *DeleteAllocatorOK) GetPayload() models.EmptyResponse {
	return o.Payload
}

func (o *DeleteAllocatorOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteAllocatorBadRequest creates a DeleteAllocatorBadRequest with default headers values
func NewDeleteAllocatorBadRequest() *DeleteAllocatorBadRequest {
	return &DeleteAllocatorBadRequest{}
}

/*DeleteAllocatorBadRequest handles this case with default header values.

* The allocator specified by {allocator_id} could not be deleted. (code: `allocators.delete_connected_allocator_attempt`)
* The allocator specified by {allocator_id} could not be deleted. (code: `allocators.delete_allocator_with_instances_attempt`)
 */
type DeleteAllocatorBadRequest struct {
	/*The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

func (o *DeleteAllocatorBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /platform/infrastructure/allocators/{allocator_id}][%d] deleteAllocatorBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteAllocatorBadRequest) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *DeleteAllocatorBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header x-cloud-error-codes
	o.XCloudErrorCodes = response.GetHeader("x-cloud-error-codes")

	o.Payload = new(models.BasicFailedReply)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteAllocatorNotFound creates a DeleteAllocatorNotFound with default headers values
func NewDeleteAllocatorNotFound() *DeleteAllocatorNotFound {
	return &DeleteAllocatorNotFound{}
}

/*DeleteAllocatorNotFound handles this case with default header values.

The allocator specified by {allocator_id} cannot be found. (code: `allocators.allocator_not_found`)
*/
type DeleteAllocatorNotFound struct {
	/*The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

func (o *DeleteAllocatorNotFound) Error() string {
	return fmt.Sprintf("[DELETE /platform/infrastructure/allocators/{allocator_id}][%d] deleteAllocatorNotFound  %+v", 404, o.Payload)
}

func (o *DeleteAllocatorNotFound) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *DeleteAllocatorNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header x-cloud-error-codes
	o.XCloudErrorCodes = response.GetHeader("x-cloud-error-codes")

	o.Payload = new(models.BasicFailedReply)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteAllocatorRetryWith creates a DeleteAllocatorRetryWith with default headers values
func NewDeleteAllocatorRetryWith() *DeleteAllocatorRetryWith {
	return &DeleteAllocatorRetryWith{}
}

/*DeleteAllocatorRetryWith handles this case with default header values.

Elevated permissions are required. (code: `root.unauthorized.rbac.elevated_permissions_required`)
*/
type DeleteAllocatorRetryWith struct {
	/*The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

func (o *DeleteAllocatorRetryWith) Error() string {
	return fmt.Sprintf("[DELETE /platform/infrastructure/allocators/{allocator_id}][%d] deleteAllocatorRetryWith  %+v", 449, o.Payload)
}

func (o *DeleteAllocatorRetryWith) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *DeleteAllocatorRetryWith) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header x-cloud-error-codes
	o.XCloudErrorCodes = response.GetHeader("x-cloud-error-codes")

	o.Payload = new(models.BasicFailedReply)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
