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

// SetLicenseReader is a Reader for the SetLicense structure.
type SetLicenseReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SetLicenseReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSetLicenseOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewSetLicenseBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 449:
		result := NewSetLicenseRetryWith()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewSetLicenseOK creates a SetLicenseOK with default headers values
func NewSetLicenseOK() *SetLicenseOK {
	return &SetLicenseOK{}
}

/*SetLicenseOK handles this case with default header values.

The license was updated.
*/
type SetLicenseOK struct {
	Payload models.EmptyResponse
}

func (o *SetLicenseOK) Error() string {
	return fmt.Sprintf("[PUT /platform/license][%d] setLicenseOK  %+v", 200, o.Payload)
}

func (o *SetLicenseOK) GetPayload() models.EmptyResponse {
	return o.Payload
}

func (o *SetLicenseOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSetLicenseBadRequest creates a SetLicenseBadRequest with default headers values
func NewSetLicenseBadRequest() *SetLicenseBadRequest {
	return &SetLicenseBadRequest{}
}

/*SetLicenseBadRequest handles this case with default header values.

The license could not be updated. (code: `license.invalid_license`)
*/
type SetLicenseBadRequest struct {
	/*The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

func (o *SetLicenseBadRequest) Error() string {
	return fmt.Sprintf("[PUT /platform/license][%d] setLicenseBadRequest  %+v", 400, o.Payload)
}

func (o *SetLicenseBadRequest) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *SetLicenseBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header x-cloud-error-codes
	o.XCloudErrorCodes = response.GetHeader("x-cloud-error-codes")

	o.Payload = new(models.BasicFailedReply)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSetLicenseRetryWith creates a SetLicenseRetryWith with default headers values
func NewSetLicenseRetryWith() *SetLicenseRetryWith {
	return &SetLicenseRetryWith{}
}

/*SetLicenseRetryWith handles this case with default header values.

Elevated permissions are required. (code: `root.unauthorized.rbac.elevated_permissions_required`)
*/
type SetLicenseRetryWith struct {
	/*The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

func (o *SetLicenseRetryWith) Error() string {
	return fmt.Sprintf("[PUT /platform/license][%d] setLicenseRetryWith  %+v", 449, o.Payload)
}

func (o *SetLicenseRetryWith) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *SetLicenseRetryWith) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header x-cloud-error-codes
	o.XCloudErrorCodes = response.GetHeader("x-cloud-error-codes")

	o.Payload = new(models.BasicFailedReply)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
