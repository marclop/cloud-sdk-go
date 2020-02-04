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

// GetBlueprinterRoleReader is a Reader for the GetBlueprinterRole structure.
type GetBlueprinterRoleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetBlueprinterRoleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetBlueprinterRoleOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetBlueprinterRoleNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetBlueprinterRoleOK creates a GetBlueprinterRoleOK with default headers values
func NewGetBlueprinterRoleOK() *GetBlueprinterRoleOK {
	return &GetBlueprinterRoleOK{}
}

/*GetBlueprinterRoleOK handles this case with default header values.

The role aggregate definition.
*/
type GetBlueprinterRoleOK struct {
	Payload *models.RoleAggregate
}

func (o *GetBlueprinterRoleOK) Error() string {
	return fmt.Sprintf("[GET /platform/infrastructure/blueprinter/roles/{blueprinter_role_id}][%d] getBlueprinterRoleOK  %+v", 200, o.Payload)
}

func (o *GetBlueprinterRoleOK) GetPayload() *models.RoleAggregate {
	return o.Payload
}

func (o *GetBlueprinterRoleOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RoleAggregate)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetBlueprinterRoleNotFound creates a GetBlueprinterRoleNotFound with default headers values
func NewGetBlueprinterRoleNotFound() *GetBlueprinterRoleNotFound {
	return &GetBlueprinterRoleNotFound{}
}

/*GetBlueprinterRoleNotFound handles this case with default header values.

The role can't be found. (code: `roles.not_found`)
*/
type GetBlueprinterRoleNotFound struct {
	/*The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

func (o *GetBlueprinterRoleNotFound) Error() string {
	return fmt.Sprintf("[GET /platform/infrastructure/blueprinter/roles/{blueprinter_role_id}][%d] getBlueprinterRoleNotFound  %+v", 404, o.Payload)
}

func (o *GetBlueprinterRoleNotFound) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *GetBlueprinterRoleNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header x-cloud-error-codes
	o.XCloudErrorCodes = response.GetHeader("x-cloud-error-codes")

	o.Payload = new(models.BasicFailedReply)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
