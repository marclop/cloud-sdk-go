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

// GetEsClusterCurationSettingsReader is a Reader for the GetEsClusterCurationSettings structure.
type GetEsClusterCurationSettingsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetEsClusterCurationSettingsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetEsClusterCurationSettingsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetEsClusterCurationSettingsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetEsClusterCurationSettingsOK creates a GetEsClusterCurationSettingsOK with default headers values
func NewGetEsClusterCurationSettingsOK() *GetEsClusterCurationSettingsOK {
	return &GetEsClusterCurationSettingsOK{}
}

/*GetEsClusterCurationSettingsOK handles this case with default header values.

The cluster curation settings were successfully returned
*/
type GetEsClusterCurationSettingsOK struct {
	/*The date-time when the resource was created (ISO format relative to UTC)
	 */
	XCloudResourceCreated string
	/*The date-time when the resource was last modified (ISO format relative to UTC)
	 */
	XCloudResourceLastModified string
	/*The resource version, which is used to avoid update conflicts with concurrent operations
	 */
	XCloudResourceVersion string

	Payload *models.ClusterCurationSettings
}

func (o *GetEsClusterCurationSettingsOK) Error() string {
	return fmt.Sprintf("[GET /clusters/elasticsearch/{cluster_id}/curation/settings][%d] getEsClusterCurationSettingsOK  %+v", 200, o.Payload)
}

func (o *GetEsClusterCurationSettingsOK) GetPayload() *models.ClusterCurationSettings {
	return o.Payload
}

func (o *GetEsClusterCurationSettingsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header x-cloud-resource-created
	o.XCloudResourceCreated = response.GetHeader("x-cloud-resource-created")

	// response header x-cloud-resource-last-modified
	o.XCloudResourceLastModified = response.GetHeader("x-cloud-resource-last-modified")

	// response header x-cloud-resource-version
	o.XCloudResourceVersion = response.GetHeader("x-cloud-resource-version")

	o.Payload = new(models.ClusterCurationSettings)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetEsClusterCurationSettingsNotFound creates a GetEsClusterCurationSettingsNotFound with default headers values
func NewGetEsClusterCurationSettingsNotFound() *GetEsClusterCurationSettingsNotFound {
	return &GetEsClusterCurationSettingsNotFound{}
}

/*GetEsClusterCurationSettingsNotFound handles this case with default header values.

The cluster specified by {cluster_id} cannot be found (code: 'clusters.cluster_not_found')
*/
type GetEsClusterCurationSettingsNotFound struct {
	Payload *models.BasicFailedReply
}

func (o *GetEsClusterCurationSettingsNotFound) Error() string {
	return fmt.Sprintf("[GET /clusters/elasticsearch/{cluster_id}/curation/settings][%d] getEsClusterCurationSettingsNotFound  %+v", 404, o.Payload)
}

func (o *GetEsClusterCurationSettingsNotFound) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *GetEsClusterCurationSettingsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BasicFailedReply)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
