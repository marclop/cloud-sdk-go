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

package comments

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/elastic/cloud-sdk-go/pkg/models"
)

// UpdateCommentReader is a Reader for the UpdateComment structure.
type UpdateCommentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateCommentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateCommentOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewUpdateCommentUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateCommentNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewUpdateCommentConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateCommentOK creates a UpdateCommentOK with default headers values
func NewUpdateCommentOK() *UpdateCommentOK {
	return &UpdateCommentOK{}
}

/*UpdateCommentOK handles this case with default header values.

Comment updated successfully.
*/
type UpdateCommentOK struct {
	/*The date-time when the resource was created (ISO format relative to UTC)
	 */
	XCloudResourceCreated string
	/*The date-time when the resource was last modified (ISO format relative to UTC)
	 */
	XCloudResourceLastModified string
	/*The resource version, which is used to avoid update conflicts with concurrent operations
	 */
	XCloudResourceVersion string

	Payload *models.Comment
}

func (o *UpdateCommentOK) Error() string {
	return fmt.Sprintf("[PUT /comments/{resource_type}/{resource_id}/{comment_id}][%d] updateCommentOK  %+v", 200, o.Payload)
}

func (o *UpdateCommentOK) GetPayload() *models.Comment {
	return o.Payload
}

func (o *UpdateCommentOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header x-cloud-resource-created
	o.XCloudResourceCreated = response.GetHeader("x-cloud-resource-created")

	// response header x-cloud-resource-last-modified
	o.XCloudResourceLastModified = response.GetHeader("x-cloud-resource-last-modified")

	// response header x-cloud-resource-version
	o.XCloudResourceVersion = response.GetHeader("x-cloud-resource-version")

	o.Payload = new(models.Comment)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateCommentUnauthorized creates a UpdateCommentUnauthorized with default headers values
func NewUpdateCommentUnauthorized() *UpdateCommentUnauthorized {
	return &UpdateCommentUnauthorized{}
}

/*UpdateCommentUnauthorized handles this case with default header values.

* The Comment does not belong to you. (code: `comments.unauthorised`)
* Your current session does not have a user id associated with it. (code: `comments.no_user_id`)
 */
type UpdateCommentUnauthorized struct {
	/*The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

func (o *UpdateCommentUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /comments/{resource_type}/{resource_id}/{comment_id}][%d] updateCommentUnauthorized  %+v", 401, o.Payload)
}

func (o *UpdateCommentUnauthorized) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *UpdateCommentUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header x-cloud-error-codes
	o.XCloudErrorCodes = response.GetHeader("x-cloud-error-codes")

	o.Payload = new(models.BasicFailedReply)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateCommentNotFound creates a UpdateCommentNotFound with default headers values
func NewUpdateCommentNotFound() *UpdateCommentNotFound {
	return &UpdateCommentNotFound{}
}

/*UpdateCommentNotFound handles this case with default header values.

The Comment you want does not exist. (code: `comments.comment_does_not_exist`)
*/
type UpdateCommentNotFound struct {
	/*The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

func (o *UpdateCommentNotFound) Error() string {
	return fmt.Sprintf("[PUT /comments/{resource_type}/{resource_id}/{comment_id}][%d] updateCommentNotFound  %+v", 404, o.Payload)
}

func (o *UpdateCommentNotFound) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *UpdateCommentNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header x-cloud-error-codes
	o.XCloudErrorCodes = response.GetHeader("x-cloud-error-codes")

	o.Payload = new(models.BasicFailedReply)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateCommentConflict creates a UpdateCommentConflict with default headers values
func NewUpdateCommentConflict() *UpdateCommentConflict {
	return &UpdateCommentConflict{}
}

/*UpdateCommentConflict handles this case with default header values.

The version you sent does not match the persisted version. (code: `comments.version_conflict`)
*/
type UpdateCommentConflict struct {
	/*The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

func (o *UpdateCommentConflict) Error() string {
	return fmt.Sprintf("[PUT /comments/{resource_type}/{resource_id}/{comment_id}][%d] updateCommentConflict  %+v", 409, o.Payload)
}

func (o *UpdateCommentConflict) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *UpdateCommentConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header x-cloud-error-codes
	o.XCloudErrorCodes = response.GetHeader("x-cloud-error-codes")

	o.Payload = new(models.BasicFailedReply)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
