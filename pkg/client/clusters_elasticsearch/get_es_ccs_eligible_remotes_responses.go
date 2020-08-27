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
	"github.com/go-openapi/strfmt"

	"github.com/elastic/cloud-sdk-go/pkg/models"
)

// GetEsCcsEligibleRemotesReader is a Reader for the GetEsCcsEligibleRemotes structure.
type GetEsCcsEligibleRemotesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetEsCcsEligibleRemotesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetEsCcsEligibleRemotesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetEsCcsEligibleRemotesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetEsCcsEligibleRemotesOK creates a GetEsCcsEligibleRemotesOK with default headers values
func NewGetEsCcsEligibleRemotesOK() *GetEsCcsEligibleRemotesOK {
	return &GetEsCcsEligibleRemotesOK{}
}

/*GetEsCcsEligibleRemotesOK handles this case with default header values.

A list of Elasticsearch clusters that can be used as cross-cluster search remotes in deployments with the provided version
*/
type GetEsCcsEligibleRemotesOK struct {
	Payload *models.ElasticsearchClustersInfo
}

func (o *GetEsCcsEligibleRemotesOK) Error() string {
	return fmt.Sprintf("[GET /clusters/elasticsearch/ccs/eligible_remotes][%d] getEsCcsEligibleRemotesOK  %+v", 200, o.Payload)
}

func (o *GetEsCcsEligibleRemotesOK) GetPayload() *models.ElasticsearchClustersInfo {
	return o.Payload
}

func (o *GetEsCcsEligibleRemotesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ElasticsearchClustersInfo)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetEsCcsEligibleRemotesBadRequest creates a GetEsCcsEligibleRemotesBadRequest with default headers values
func NewGetEsCcsEligibleRemotesBadRequest() *GetEsCcsEligibleRemotesBadRequest {
	return &GetEsCcsEligibleRemotesBadRequest{}
}

/*GetEsCcsEligibleRemotesBadRequest handles this case with default header values.

The search request failed.
*/
type GetEsCcsEligibleRemotesBadRequest struct {
	Payload *models.BasicFailedReply
}

func (o *GetEsCcsEligibleRemotesBadRequest) Error() string {
	return fmt.Sprintf("[GET /clusters/elasticsearch/ccs/eligible_remotes][%d] getEsCcsEligibleRemotesBadRequest  %+v", 400, o.Payload)
}

func (o *GetEsCcsEligibleRemotesBadRequest) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *GetEsCcsEligibleRemotesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BasicFailedReply)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
