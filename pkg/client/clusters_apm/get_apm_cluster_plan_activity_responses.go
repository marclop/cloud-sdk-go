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

package clusters_apm

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/elastic/cloud-sdk-go/pkg/models"
)

// GetApmClusterPlanActivityReader is a Reader for the GetApmClusterPlanActivity structure.
type GetApmClusterPlanActivityReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetApmClusterPlanActivityReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetApmClusterPlanActivityOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetApmClusterPlanActivityNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetApmClusterPlanActivityOK creates a GetApmClusterPlanActivityOK with default headers values
func NewGetApmClusterPlanActivityOK() *GetApmClusterPlanActivityOK {
	return &GetApmClusterPlanActivityOK{}
}

/*GetApmClusterPlanActivityOK handles this case with default header values.

Returning the plan activity for the specified APM cluster
*/
type GetApmClusterPlanActivityOK struct {
	Payload *models.ApmPlansInfo
}

func (o *GetApmClusterPlanActivityOK) Error() string {
	return fmt.Sprintf("[GET /clusters/apm/{cluster_id}/plan/activity][%d] getApmClusterPlanActivityOK  %+v", 200, o.Payload)
}

func (o *GetApmClusterPlanActivityOK) GetPayload() *models.ApmPlansInfo {
	return o.Payload
}

func (o *GetApmClusterPlanActivityOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ApmPlansInfo)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetApmClusterPlanActivityNotFound creates a GetApmClusterPlanActivityNotFound with default headers values
func NewGetApmClusterPlanActivityNotFound() *GetApmClusterPlanActivityNotFound {
	return &GetApmClusterPlanActivityNotFound{}
}

/*GetApmClusterPlanActivityNotFound handles this case with default header values.

The cluster specified by {cluster_id} cannot be found. (code: `clusters.cluster_not_found`)
*/
type GetApmClusterPlanActivityNotFound struct {
	/*The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

func (o *GetApmClusterPlanActivityNotFound) Error() string {
	return fmt.Sprintf("[GET /clusters/apm/{cluster_id}/plan/activity][%d] getApmClusterPlanActivityNotFound  %+v", 404, o.Payload)
}

func (o *GetApmClusterPlanActivityNotFound) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *GetApmClusterPlanActivityNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header x-cloud-error-codes
	o.XCloudErrorCodes = response.GetHeader("x-cloud-error-codes")

	o.Payload = new(models.BasicFailedReply)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
