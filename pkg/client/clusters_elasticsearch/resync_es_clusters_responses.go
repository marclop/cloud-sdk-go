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

package clusters_elasticsearch

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/elastic/cloud-sdk-go/pkg/models"
)

// ResyncEsClustersReader is a Reader for the ResyncEsClusters structure.
type ResyncEsClustersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ResyncEsClustersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewResyncEsClustersAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 449:
		result := NewResyncEsClustersRetryWith()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewResyncEsClustersAccepted creates a ResyncEsClustersAccepted with default headers values
func NewResyncEsClustersAccepted() *ResyncEsClustersAccepted {
	return &ResyncEsClustersAccepted{}
}

/*ResyncEsClustersAccepted handles this case with default header values.

The ids of documents, organized by model version, that will be synchronized.
*/
type ResyncEsClustersAccepted struct {
	Payload *models.ModelVersionIndexSynchronizationResults
}

func (o *ResyncEsClustersAccepted) Error() string {
	return fmt.Sprintf("[POST /clusters/elasticsearch/_resync][%d] resyncEsClustersAccepted  %+v", 202, o.Payload)
}

func (o *ResyncEsClustersAccepted) GetPayload() *models.ModelVersionIndexSynchronizationResults {
	return o.Payload
}

func (o *ResyncEsClustersAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ModelVersionIndexSynchronizationResults)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewResyncEsClustersRetryWith creates a ResyncEsClustersRetryWith with default headers values
func NewResyncEsClustersRetryWith() *ResyncEsClustersRetryWith {
	return &ResyncEsClustersRetryWith{}
}

/*ResyncEsClustersRetryWith handles this case with default header values.

Elevated permissions are required. (code: `root.unauthorized.rbac.elevated_permissions_required`)
*/
type ResyncEsClustersRetryWith struct {
	/*The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

func (o *ResyncEsClustersRetryWith) Error() string {
	return fmt.Sprintf("[POST /clusters/elasticsearch/_resync][%d] resyncEsClustersRetryWith  %+v", 449, o.Payload)
}

func (o *ResyncEsClustersRetryWith) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *ResyncEsClustersRetryWith) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header x-cloud-error-codes
	o.XCloudErrorCodes = response.GetHeader("x-cloud-error-codes")

	o.Payload = new(models.BasicFailedReply)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
