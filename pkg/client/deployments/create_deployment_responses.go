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

	strfmt "github.com/go-openapi/strfmt"

	"github.com/elastic/cloud-sdk-go/pkg/models"
)

// CreateDeploymentReader is a Reader for the CreateDeployment structure.
type CreateDeploymentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateDeploymentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateDeploymentCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewCreateDeploymentAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateDeploymentBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewCreateDeploymentUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCreateDeploymentCreated creates a CreateDeploymentCreated with default headers values
func NewCreateDeploymentCreated() *CreateDeploymentCreated {
	return &CreateDeploymentCreated{}
}

/*CreateDeploymentCreated handles this case with default header values.

The request was valid and a new deployment was created
*/
type CreateDeploymentCreated struct {
	Payload *models.DeploymentCreateResponse
}

func (o *CreateDeploymentCreated) Error() string {
	return fmt.Sprintf("[POST /deployments][%d] createDeploymentCreated  %+v", 201, o.Payload)
}

func (o *CreateDeploymentCreated) GetPayload() *models.DeploymentCreateResponse {
	return o.Payload
}

func (o *CreateDeploymentCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DeploymentCreateResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateDeploymentAccepted creates a CreateDeploymentAccepted with default headers values
func NewCreateDeploymentAccepted() *CreateDeploymentAccepted {
	return &CreateDeploymentAccepted{}
}

/*CreateDeploymentAccepted handles this case with default header values.

The request was valid
*/
type CreateDeploymentAccepted struct {
	Payload *models.DeploymentCreateResponse
}

func (o *CreateDeploymentAccepted) Error() string {
	return fmt.Sprintf("[POST /deployments][%d] createDeploymentAccepted  %+v", 202, o.Payload)
}

func (o *CreateDeploymentAccepted) GetPayload() *models.DeploymentCreateResponse {
	return o.Payload
}

func (o *CreateDeploymentAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DeploymentCreateResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateDeploymentBadRequest creates a CreateDeploymentBadRequest with default headers values
func NewCreateDeploymentBadRequest() *CreateDeploymentBadRequest {
	return &CreateDeploymentBadRequest{}
}

/*CreateDeploymentBadRequest handles this case with default header values.

The deployment request had errors
*/
type CreateDeploymentBadRequest struct {
	Payload *models.BasicFailedReply
}

func (o *CreateDeploymentBadRequest) Error() string {
	return fmt.Sprintf("[POST /deployments][%d] createDeploymentBadRequest  %+v", 400, o.Payload)
}

func (o *CreateDeploymentBadRequest) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *CreateDeploymentBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BasicFailedReply)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateDeploymentUnauthorized creates a CreateDeploymentUnauthorized with default headers values
func NewCreateDeploymentUnauthorized() *CreateDeploymentUnauthorized {
	return &CreateDeploymentUnauthorized{}
}

/*CreateDeploymentUnauthorized handles this case with default header values.

You are not authorized to perform this action
*/
type CreateDeploymentUnauthorized struct {
	Payload *models.BasicFailedReply
}

func (o *CreateDeploymentUnauthorized) Error() string {
	return fmt.Sprintf("[POST /deployments][%d] createDeploymentUnauthorized  %+v", 401, o.Payload)
}

func (o *CreateDeploymentUnauthorized) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *CreateDeploymentUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BasicFailedReply)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
