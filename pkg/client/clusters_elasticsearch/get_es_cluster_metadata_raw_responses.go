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

// GetEsClusterMetadataRawReader is a Reader for the GetEsClusterMetadataRaw structure.
type GetEsClusterMetadataRawReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetEsClusterMetadataRawReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetEsClusterMetadataRawOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetEsClusterMetadataRawNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetEsClusterMetadataRawOK creates a GetEsClusterMetadataRawOK with default headers values
func NewGetEsClusterMetadataRawOK() *GetEsClusterMetadataRawOK {
	return &GetEsClusterMetadataRawOK{}
}

/*GetEsClusterMetadataRawOK handles this case with default header values.

The cluster metadata was successfully returned
*/
type GetEsClusterMetadataRawOK struct {
	/*The date-time when the resource was created (ISO format relative to UTC)
	 */
	XCloudResourceCreated string
	/*The date-time when the resource was last modified (ISO format relative to UTC)
	 */
	XCloudResourceLastModified string
	/*The resource version, which is used to avoid update conflicts with concurrent operations
	 */
	XCloudResourceVersion string

	Payload interface{}
}

func (o *GetEsClusterMetadataRawOK) Error() string {
	return fmt.Sprintf("[GET /clusters/elasticsearch/{cluster_id}/metadata/raw][%d] getEsClusterMetadataRawOK  %+v", 200, o.Payload)
}

func (o *GetEsClusterMetadataRawOK) GetPayload() interface{} {
	return o.Payload
}

func (o *GetEsClusterMetadataRawOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header x-cloud-resource-created
	o.XCloudResourceCreated = response.GetHeader("x-cloud-resource-created")

	// response header x-cloud-resource-last-modified
	o.XCloudResourceLastModified = response.GetHeader("x-cloud-resource-last-modified")

	// response header x-cloud-resource-version
	o.XCloudResourceVersion = response.GetHeader("x-cloud-resource-version")

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetEsClusterMetadataRawNotFound creates a GetEsClusterMetadataRawNotFound with default headers values
func NewGetEsClusterMetadataRawNotFound() *GetEsClusterMetadataRawNotFound {
	return &GetEsClusterMetadataRawNotFound{}
}

/*GetEsClusterMetadataRawNotFound handles this case with default header values.

The cluster specified by {cluster_id} cannot be found (code: 'clusters.cluster_not_found')
*/
type GetEsClusterMetadataRawNotFound struct {
	Payload *models.BasicFailedReply
}

func (o *GetEsClusterMetadataRawNotFound) Error() string {
	return fmt.Sprintf("[GET /clusters/elasticsearch/{cluster_id}/metadata/raw][%d] getEsClusterMetadataRawNotFound  %+v", 404, o.Payload)
}

func (o *GetEsClusterMetadataRawNotFound) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *GetEsClusterMetadataRawNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BasicFailedReply)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
