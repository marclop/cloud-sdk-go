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

package deployments_ip_filtering

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/elastic/cloud-sdk-go/pkg/models"
)

// DeleteIPFilterRulesetAssociationReader is a Reader for the DeleteIPFilterRulesetAssociation structure.
type DeleteIPFilterRulesetAssociationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteIPFilterRulesetAssociationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteIPFilterRulesetAssociationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 449:
		result := NewDeleteIPFilterRulesetAssociationRetryWith()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteIPFilterRulesetAssociationInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteIPFilterRulesetAssociationOK creates a DeleteIPFilterRulesetAssociationOK with default headers values
func NewDeleteIPFilterRulesetAssociationOK() *DeleteIPFilterRulesetAssociationOK {
	return &DeleteIPFilterRulesetAssociationOK{}
}

/*DeleteIPFilterRulesetAssociationOK handles this case with default header values.

Delete association request was valid and the association has been deleted
*/
type DeleteIPFilterRulesetAssociationOK struct {
	Payload models.EmptyResponse
}

func (o *DeleteIPFilterRulesetAssociationOK) Error() string {
	return fmt.Sprintf("[DELETE /deployments/ip-filtering/rulesets/{ruleset_id}/associations/{association_type}/{associated_entity_id}][%d] deleteIpFilterRulesetAssociationOK  %+v", 200, o.Payload)
}

func (o *DeleteIPFilterRulesetAssociationOK) GetPayload() models.EmptyResponse {
	return o.Payload
}

func (o *DeleteIPFilterRulesetAssociationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteIPFilterRulesetAssociationRetryWith creates a DeleteIPFilterRulesetAssociationRetryWith with default headers values
func NewDeleteIPFilterRulesetAssociationRetryWith() *DeleteIPFilterRulesetAssociationRetryWith {
	return &DeleteIPFilterRulesetAssociationRetryWith{}
}

/*DeleteIPFilterRulesetAssociationRetryWith handles this case with default header values.

Elevated permissions are required. (code: '"root.unauthorized.rbac.elevated_permissions_required"')
*/
type DeleteIPFilterRulesetAssociationRetryWith struct {
	Payload *models.BasicFailedReply
}

func (o *DeleteIPFilterRulesetAssociationRetryWith) Error() string {
	return fmt.Sprintf("[DELETE /deployments/ip-filtering/rulesets/{ruleset_id}/associations/{association_type}/{associated_entity_id}][%d] deleteIpFilterRulesetAssociationRetryWith  %+v", 449, o.Payload)
}

func (o *DeleteIPFilterRulesetAssociationRetryWith) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *DeleteIPFilterRulesetAssociationRetryWith) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BasicFailedReply)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteIPFilterRulesetAssociationInternalServerError creates a DeleteIPFilterRulesetAssociationInternalServerError with default headers values
func NewDeleteIPFilterRulesetAssociationInternalServerError() *DeleteIPFilterRulesetAssociationInternalServerError {
	return &DeleteIPFilterRulesetAssociationInternalServerError{}
}

/*DeleteIPFilterRulesetAssociationInternalServerError handles this case with default header values.

Request execution failed (code: 'ip_filtering.request_execution_failed')
*/
type DeleteIPFilterRulesetAssociationInternalServerError struct {
	Payload *models.BasicFailedReply
}

func (o *DeleteIPFilterRulesetAssociationInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /deployments/ip-filtering/rulesets/{ruleset_id}/associations/{association_type}/{associated_entity_id}][%d] deleteIpFilterRulesetAssociationInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteIPFilterRulesetAssociationInternalServerError) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *DeleteIPFilterRulesetAssociationInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BasicFailedReply)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
