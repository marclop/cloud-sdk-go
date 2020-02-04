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

package platform_configuration_templates

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/elastic/cloud-sdk-go/pkg/models"
)

// SetDeploymentTemplateReader is a Reader for the SetDeploymentTemplate structure.
type SetDeploymentTemplateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SetDeploymentTemplateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSetDeploymentTemplateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 201:
		result := NewSetDeploymentTemplateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewSetDeploymentTemplateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewSetDeploymentTemplateNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewSetDeploymentTemplateConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 449:
		result := NewSetDeploymentTemplateRetryWith()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewSetDeploymentTemplateOK creates a SetDeploymentTemplateOK with default headers values
func NewSetDeploymentTemplateOK() *SetDeploymentTemplateOK {
	return &SetDeploymentTemplateOK{}
}

/*SetDeploymentTemplateOK handles this case with default header values.

The deployment definition was valid and the template has been updated.
*/
type SetDeploymentTemplateOK struct {
	Payload *models.IDResponse
}

func (o *SetDeploymentTemplateOK) Error() string {
	return fmt.Sprintf("[PUT /platform/configuration/templates/deployments/{template_id}][%d] setDeploymentTemplateOK  %+v", 200, o.Payload)
}

func (o *SetDeploymentTemplateOK) GetPayload() *models.IDResponse {
	return o.Payload
}

func (o *SetDeploymentTemplateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IDResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSetDeploymentTemplateCreated creates a SetDeploymentTemplateCreated with default headers values
func NewSetDeploymentTemplateCreated() *SetDeploymentTemplateCreated {
	return &SetDeploymentTemplateCreated{}
}

/*SetDeploymentTemplateCreated handles this case with default header values.

The deployment definition was valid and the template was created.
*/
type SetDeploymentTemplateCreated struct {
	Payload *models.IDResponse
}

func (o *SetDeploymentTemplateCreated) Error() string {
	return fmt.Sprintf("[PUT /platform/configuration/templates/deployments/{template_id}][%d] setDeploymentTemplateCreated  %+v", 201, o.Payload)
}

func (o *SetDeploymentTemplateCreated) GetPayload() *models.IDResponse {
	return o.Payload
}

func (o *SetDeploymentTemplateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IDResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSetDeploymentTemplateBadRequest creates a SetDeploymentTemplateBadRequest with default headers values
func NewSetDeploymentTemplateBadRequest() *SetDeploymentTemplateBadRequest {
	return &SetDeploymentTemplateBadRequest{}
}

/*SetDeploymentTemplateBadRequest handles this case with default header values.

The template definition contained errors. (code: `templates.invalid_template`)
*/
type SetDeploymentTemplateBadRequest struct {
	/*The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

func (o *SetDeploymentTemplateBadRequest) Error() string {
	return fmt.Sprintf("[PUT /platform/configuration/templates/deployments/{template_id}][%d] setDeploymentTemplateBadRequest  %+v", 400, o.Payload)
}

func (o *SetDeploymentTemplateBadRequest) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *SetDeploymentTemplateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header x-cloud-error-codes
	o.XCloudErrorCodes = response.GetHeader("x-cloud-error-codes")

	o.Payload = new(models.BasicFailedReply)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSetDeploymentTemplateNotFound creates a SetDeploymentTemplateNotFound with default headers values
func NewSetDeploymentTemplateNotFound() *SetDeploymentTemplateNotFound {
	return &SetDeploymentTemplateNotFound{}
}

/*SetDeploymentTemplateNotFound handles this case with default header values.

The deployment template specified by {template_id} cannot be found. (code: `templates.template_not_found`)
*/
type SetDeploymentTemplateNotFound struct {
	/*The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

func (o *SetDeploymentTemplateNotFound) Error() string {
	return fmt.Sprintf("[PUT /platform/configuration/templates/deployments/{template_id}][%d] setDeploymentTemplateNotFound  %+v", 404, o.Payload)
}

func (o *SetDeploymentTemplateNotFound) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *SetDeploymentTemplateNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header x-cloud-error-codes
	o.XCloudErrorCodes = response.GetHeader("x-cloud-error-codes")

	o.Payload = new(models.BasicFailedReply)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSetDeploymentTemplateConflict creates a SetDeploymentTemplateConflict with default headers values
func NewSetDeploymentTemplateConflict() *SetDeploymentTemplateConflict {
	return &SetDeploymentTemplateConflict{}
}

/*SetDeploymentTemplateConflict handles this case with default header values.

The version supplied in the request conflicted with the version found on the server. (code: `templates.version_conflict`)
*/
type SetDeploymentTemplateConflict struct {
	/*The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

func (o *SetDeploymentTemplateConflict) Error() string {
	return fmt.Sprintf("[PUT /platform/configuration/templates/deployments/{template_id}][%d] setDeploymentTemplateConflict  %+v", 409, o.Payload)
}

func (o *SetDeploymentTemplateConflict) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *SetDeploymentTemplateConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header x-cloud-error-codes
	o.XCloudErrorCodes = response.GetHeader("x-cloud-error-codes")

	o.Payload = new(models.BasicFailedReply)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSetDeploymentTemplateRetryWith creates a SetDeploymentTemplateRetryWith with default headers values
func NewSetDeploymentTemplateRetryWith() *SetDeploymentTemplateRetryWith {
	return &SetDeploymentTemplateRetryWith{}
}

/*SetDeploymentTemplateRetryWith handles this case with default header values.

Elevated permissions are required. (code: `root.unauthorized.rbac.elevated_permissions_required`)
*/
type SetDeploymentTemplateRetryWith struct {
	/*The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

func (o *SetDeploymentTemplateRetryWith) Error() string {
	return fmt.Sprintf("[PUT /platform/configuration/templates/deployments/{template_id}][%d] setDeploymentTemplateRetryWith  %+v", 449, o.Payload)
}

func (o *SetDeploymentTemplateRetryWith) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *SetDeploymentTemplateRetryWith) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header x-cloud-error-codes
	o.XCloudErrorCodes = response.GetHeader("x-cloud-error-codes")

	o.Payload = new(models.BasicFailedReply)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
