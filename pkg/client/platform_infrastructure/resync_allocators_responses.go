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

// ResyncAllocatorsReader is a Reader for the ResyncAllocators structure.
type ResyncAllocatorsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ResyncAllocatorsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewResyncAllocatorsAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 449:
		result := NewResyncAllocatorsRetryWith()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewResyncAllocatorsAccepted creates a ResyncAllocatorsAccepted with default headers values
func NewResyncAllocatorsAccepted() *ResyncAllocatorsAccepted {
	return &ResyncAllocatorsAccepted{}
}

/*ResyncAllocatorsAccepted handles this case with default header values.

The ids of documents, organized by model version, that will be synchronized.
*/
type ResyncAllocatorsAccepted struct {
	Payload *models.ModelVersionIndexSynchronizationResults
}

func (o *ResyncAllocatorsAccepted) Error() string {
	return fmt.Sprintf("[POST /platform/infrastructure/allocators/_resync][%d] resyncAllocatorsAccepted  %+v", 202, o.Payload)
}

func (o *ResyncAllocatorsAccepted) GetPayload() *models.ModelVersionIndexSynchronizationResults {
	return o.Payload
}

func (o *ResyncAllocatorsAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ModelVersionIndexSynchronizationResults)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewResyncAllocatorsRetryWith creates a ResyncAllocatorsRetryWith with default headers values
func NewResyncAllocatorsRetryWith() *ResyncAllocatorsRetryWith {
	return &ResyncAllocatorsRetryWith{}
}

/*ResyncAllocatorsRetryWith handles this case with default header values.

Elevated permissions are required. (code: `root.unauthorized.rbac.elevated_permissions_required`)
*/
type ResyncAllocatorsRetryWith struct {
	/*The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

func (o *ResyncAllocatorsRetryWith) Error() string {
	return fmt.Sprintf("[POST /platform/infrastructure/allocators/_resync][%d] resyncAllocatorsRetryWith  %+v", 449, o.Payload)
}

func (o *ResyncAllocatorsRetryWith) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *ResyncAllocatorsRetryWith) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header x-cloud-error-codes
	o.XCloudErrorCodes = response.GetHeader("x-cloud-error-codes")

	o.Payload = new(models.BasicFailedReply)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
