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

package clusters_kibana

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/elastic/cloud-sdk-go/pkg/models"
)

// StartKibanaClusterInstancesAllMaintenanceModeReader is a Reader for the StartKibanaClusterInstancesAllMaintenanceMode structure.
type StartKibanaClusterInstancesAllMaintenanceModeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StartKibanaClusterInstancesAllMaintenanceModeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewStartKibanaClusterInstancesAllMaintenanceModeAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewStartKibanaClusterInstancesAllMaintenanceModeForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewStartKibanaClusterInstancesAllMaintenanceModeNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 449:
		result := NewStartKibanaClusterInstancesAllMaintenanceModeRetryWith()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewStartKibanaClusterInstancesAllMaintenanceModeAccepted creates a StartKibanaClusterInstancesAllMaintenanceModeAccepted with default headers values
func NewStartKibanaClusterInstancesAllMaintenanceModeAccepted() *StartKibanaClusterInstancesAllMaintenanceModeAccepted {
	return &StartKibanaClusterInstancesAllMaintenanceModeAccepted{}
}

/*StartKibanaClusterInstancesAllMaintenanceModeAccepted handles this case with default header values.

The start maintenance mode command was issued successfully, use the "GET" command on the /{cluster_id} resource to monitor progress
*/
type StartKibanaClusterInstancesAllMaintenanceModeAccepted struct {
	Payload *models.ClusterCommandResponse
}

func (o *StartKibanaClusterInstancesAllMaintenanceModeAccepted) Error() string {
	return fmt.Sprintf("[POST /clusters/kibana/{cluster_id}/instances/maintenance-mode/_start][%d] startKibanaClusterInstancesAllMaintenanceModeAccepted  %+v", 202, o.Payload)
}

func (o *StartKibanaClusterInstancesAllMaintenanceModeAccepted) GetPayload() *models.ClusterCommandResponse {
	return o.Payload
}

func (o *StartKibanaClusterInstancesAllMaintenanceModeAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ClusterCommandResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStartKibanaClusterInstancesAllMaintenanceModeForbidden creates a StartKibanaClusterInstancesAllMaintenanceModeForbidden with default headers values
func NewStartKibanaClusterInstancesAllMaintenanceModeForbidden() *StartKibanaClusterInstancesAllMaintenanceModeForbidden {
	return &StartKibanaClusterInstancesAllMaintenanceModeForbidden{}
}

/*StartKibanaClusterInstancesAllMaintenanceModeForbidden handles this case with default header values.

The start maintenance mode command was prohibited for the given cluster. (code: `clusters.command_prohibited`)
*/
type StartKibanaClusterInstancesAllMaintenanceModeForbidden struct {
	/*The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

func (o *StartKibanaClusterInstancesAllMaintenanceModeForbidden) Error() string {
	return fmt.Sprintf("[POST /clusters/kibana/{cluster_id}/instances/maintenance-mode/_start][%d] startKibanaClusterInstancesAllMaintenanceModeForbidden  %+v", 403, o.Payload)
}

func (o *StartKibanaClusterInstancesAllMaintenanceModeForbidden) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *StartKibanaClusterInstancesAllMaintenanceModeForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header x-cloud-error-codes
	o.XCloudErrorCodes = response.GetHeader("x-cloud-error-codes")

	o.Payload = new(models.BasicFailedReply)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStartKibanaClusterInstancesAllMaintenanceModeNotFound creates a StartKibanaClusterInstancesAllMaintenanceModeNotFound with default headers values
func NewStartKibanaClusterInstancesAllMaintenanceModeNotFound() *StartKibanaClusterInstancesAllMaintenanceModeNotFound {
	return &StartKibanaClusterInstancesAllMaintenanceModeNotFound{}
}

/*StartKibanaClusterInstancesAllMaintenanceModeNotFound handles this case with default header values.

The cluster specified by {cluster_id} cannot be found. (code: `clusters.cluster_not_found`)
*/
type StartKibanaClusterInstancesAllMaintenanceModeNotFound struct {
	/*The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

func (o *StartKibanaClusterInstancesAllMaintenanceModeNotFound) Error() string {
	return fmt.Sprintf("[POST /clusters/kibana/{cluster_id}/instances/maintenance-mode/_start][%d] startKibanaClusterInstancesAllMaintenanceModeNotFound  %+v", 404, o.Payload)
}

func (o *StartKibanaClusterInstancesAllMaintenanceModeNotFound) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *StartKibanaClusterInstancesAllMaintenanceModeNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header x-cloud-error-codes
	o.XCloudErrorCodes = response.GetHeader("x-cloud-error-codes")

	o.Payload = new(models.BasicFailedReply)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStartKibanaClusterInstancesAllMaintenanceModeRetryWith creates a StartKibanaClusterInstancesAllMaintenanceModeRetryWith with default headers values
func NewStartKibanaClusterInstancesAllMaintenanceModeRetryWith() *StartKibanaClusterInstancesAllMaintenanceModeRetryWith {
	return &StartKibanaClusterInstancesAllMaintenanceModeRetryWith{}
}

/*StartKibanaClusterInstancesAllMaintenanceModeRetryWith handles this case with default header values.

Elevated permissions are required. (code: `root.unauthorized.rbac.elevated_permissions_required`)
*/
type StartKibanaClusterInstancesAllMaintenanceModeRetryWith struct {
	/*The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

func (o *StartKibanaClusterInstancesAllMaintenanceModeRetryWith) Error() string {
	return fmt.Sprintf("[POST /clusters/kibana/{cluster_id}/instances/maintenance-mode/_start][%d] startKibanaClusterInstancesAllMaintenanceModeRetryWith  %+v", 449, o.Payload)
}

func (o *StartKibanaClusterInstancesAllMaintenanceModeRetryWith) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *StartKibanaClusterInstancesAllMaintenanceModeRetryWith) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header x-cloud-error-codes
	o.XCloudErrorCodes = response.GetHeader("x-cloud-error-codes")

	o.Payload = new(models.BasicFailedReply)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
