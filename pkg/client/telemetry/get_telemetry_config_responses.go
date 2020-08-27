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

package telemetry

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/elastic/cloud-sdk-go/pkg/models"
)

// GetTelemetryConfigReader is a Reader for the GetTelemetryConfig structure.
type GetTelemetryConfigReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetTelemetryConfigReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetTelemetryConfigOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewGetTelemetryConfigForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetTelemetryConfigOK creates a GetTelemetryConfigOK with default headers values
func NewGetTelemetryConfigOK() *GetTelemetryConfigOK {
	return &GetTelemetryConfigOK{}
}

/*GetTelemetryConfigOK handles this case with default header values.

The current ECE telemetry configuration
*/
type GetTelemetryConfigOK struct {
	Payload *models.TelemetryConfig
}

func (o *GetTelemetryConfigOK) Error() string {
	return fmt.Sprintf("[GET /phone-home/config][%d] getTelemetryConfigOK  %+v", 200, o.Payload)
}

func (o *GetTelemetryConfigOK) GetPayload() *models.TelemetryConfig {
	return o.Payload
}

func (o *GetTelemetryConfigOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TelemetryConfig)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelemetryConfigForbidden creates a GetTelemetryConfigForbidden with default headers values
func NewGetTelemetryConfigForbidden() *GetTelemetryConfigForbidden {
	return &GetTelemetryConfigForbidden{}
}

/*GetTelemetryConfigForbidden handles this case with default header values.

User must have Platform level permissions. (code: `root.unauthorized.rbac`)
*/
type GetTelemetryConfigForbidden struct {
	/*The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

func (o *GetTelemetryConfigForbidden) Error() string {
	return fmt.Sprintf("[GET /phone-home/config][%d] getTelemetryConfigForbidden  %+v", 403, o.Payload)
}

func (o *GetTelemetryConfigForbidden) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *GetTelemetryConfigForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header x-cloud-error-codes
	o.XCloudErrorCodes = response.GetHeader("x-cloud-error-codes")

	o.Payload = new(models.BasicFailedReply)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
