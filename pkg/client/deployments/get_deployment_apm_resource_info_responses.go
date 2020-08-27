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

package deployments

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/elastic/cloud-sdk-go/pkg/models"
)

// GetDeploymentApmResourceInfoReader is a Reader for the GetDeploymentApmResourceInfo structure.
type GetDeploymentApmResourceInfoReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDeploymentApmResourceInfoReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetDeploymentApmResourceInfoOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetDeploymentApmResourceInfoNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetDeploymentApmResourceInfoInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetDeploymentApmResourceInfoOK creates a GetDeploymentApmResourceInfoOK with default headers values
func NewGetDeploymentApmResourceInfoOK() *GetDeploymentApmResourceInfoOK {
	return &GetDeploymentApmResourceInfoOK{}
}

/*GetDeploymentApmResourceInfoOK handles this case with default header values.

Standard response.
*/
type GetDeploymentApmResourceInfoOK struct {
	Payload *models.ApmResourceInfo
}

func (o *GetDeploymentApmResourceInfoOK) Error() string {
	return fmt.Sprintf("[GET /deployments/{deployment_id}/apm/{ref_id}][%d] getDeploymentApmResourceInfoOK  %+v", 200, o.Payload)
}

func (o *GetDeploymentApmResourceInfoOK) GetPayload() *models.ApmResourceInfo {
	return o.Payload
}

func (o *GetDeploymentApmResourceInfoOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ApmResourceInfo)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDeploymentApmResourceInfoNotFound creates a GetDeploymentApmResourceInfoNotFound with default headers values
func NewGetDeploymentApmResourceInfoNotFound() *GetDeploymentApmResourceInfoNotFound {
	return &GetDeploymentApmResourceInfoNotFound{}
}

/*GetDeploymentApmResourceInfoNotFound handles this case with default header values.

The Deployment specified by {deployment_id} cannot be found. (code: `deployments.deployment_not_found`)
*/
type GetDeploymentApmResourceInfoNotFound struct {
	/*The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

func (o *GetDeploymentApmResourceInfoNotFound) Error() string {
	return fmt.Sprintf("[GET /deployments/{deployment_id}/apm/{ref_id}][%d] getDeploymentApmResourceInfoNotFound  %+v", 404, o.Payload)
}

func (o *GetDeploymentApmResourceInfoNotFound) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *GetDeploymentApmResourceInfoNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header x-cloud-error-codes
	o.XCloudErrorCodes = response.GetHeader("x-cloud-error-codes")

	o.Payload = new(models.BasicFailedReply)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDeploymentApmResourceInfoInternalServerError creates a GetDeploymentApmResourceInfoInternalServerError with default headers values
func NewGetDeploymentApmResourceInfoInternalServerError() *GetDeploymentApmResourceInfoInternalServerError {
	return &GetDeploymentApmResourceInfoInternalServerError{}
}

/*GetDeploymentApmResourceInfoInternalServerError handles this case with default header values.

We have failed you. (code: `deployments.deployment_resource_no_longer_exists`)
*/
type GetDeploymentApmResourceInfoInternalServerError struct {
	/*The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

func (o *GetDeploymentApmResourceInfoInternalServerError) Error() string {
	return fmt.Sprintf("[GET /deployments/{deployment_id}/apm/{ref_id}][%d] getDeploymentApmResourceInfoInternalServerError  %+v", 500, o.Payload)
}

func (o *GetDeploymentApmResourceInfoInternalServerError) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *GetDeploymentApmResourceInfoInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header x-cloud-error-codes
	o.XCloudErrorCodes = response.GetHeader("x-cloud-error-codes")

	o.Payload = new(models.BasicFailedReply)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
