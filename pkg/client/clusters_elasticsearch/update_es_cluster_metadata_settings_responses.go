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

// UpdateEsClusterMetadataSettingsReader is a Reader for the UpdateEsClusterMetadataSettings structure.
type UpdateEsClusterMetadataSettingsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateEsClusterMetadataSettingsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateEsClusterMetadataSettingsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewUpdateEsClusterMetadataSettingsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateEsClusterMetadataSettingsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 449:
		result := NewUpdateEsClusterMetadataSettingsRetryWith()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUpdateEsClusterMetadataSettingsOK creates a UpdateEsClusterMetadataSettingsOK with default headers values
func NewUpdateEsClusterMetadataSettingsOK() *UpdateEsClusterMetadataSettingsOK {
	return &UpdateEsClusterMetadataSettingsOK{}
}

/*UpdateEsClusterMetadataSettingsOK handles this case with default header values.

The cluster metadata was successfully updated
*/
type UpdateEsClusterMetadataSettingsOK struct {
	/*The date-time when the resource was created (ISO format relative to UTC)
	 */
	XCloudResourceCreated string
	/*The date-time when the resource was last modified (ISO format relative to UTC)
	 */
	XCloudResourceLastModified string
	/*The resource version, which is used to avoid update conflicts with concurrent operations
	 */
	XCloudResourceVersion string

	Payload *models.ClusterMetadataSettings
}

func (o *UpdateEsClusterMetadataSettingsOK) Error() string {
	return fmt.Sprintf("[PATCH /clusters/elasticsearch/{cluster_id}/metadata/settings][%d] updateEsClusterMetadataSettingsOK  %+v", 200, o.Payload)
}

func (o *UpdateEsClusterMetadataSettingsOK) GetPayload() *models.ClusterMetadataSettings {
	return o.Payload
}

func (o *UpdateEsClusterMetadataSettingsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header x-cloud-resource-created
	o.XCloudResourceCreated = response.GetHeader("x-cloud-resource-created")

	// response header x-cloud-resource-last-modified
	o.XCloudResourceLastModified = response.GetHeader("x-cloud-resource-last-modified")

	// response header x-cloud-resource-version
	o.XCloudResourceVersion = response.GetHeader("x-cloud-resource-version")

	o.Payload = new(models.ClusterMetadataSettings)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateEsClusterMetadataSettingsForbidden creates a UpdateEsClusterMetadataSettingsForbidden with default headers values
func NewUpdateEsClusterMetadataSettingsForbidden() *UpdateEsClusterMetadataSettingsForbidden {
	return &UpdateEsClusterMetadataSettingsForbidden{}
}

/*UpdateEsClusterMetadataSettingsForbidden handles this case with default header values.

The provided action was prohibited for the given cluster
*/
type UpdateEsClusterMetadataSettingsForbidden struct {
	Payload *models.BasicFailedReply
}

func (o *UpdateEsClusterMetadataSettingsForbidden) Error() string {
	return fmt.Sprintf("[PATCH /clusters/elasticsearch/{cluster_id}/metadata/settings][%d] updateEsClusterMetadataSettingsForbidden  %+v", 403, o.Payload)
}

func (o *UpdateEsClusterMetadataSettingsForbidden) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *UpdateEsClusterMetadataSettingsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BasicFailedReply)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateEsClusterMetadataSettingsNotFound creates a UpdateEsClusterMetadataSettingsNotFound with default headers values
func NewUpdateEsClusterMetadataSettingsNotFound() *UpdateEsClusterMetadataSettingsNotFound {
	return &UpdateEsClusterMetadataSettingsNotFound{}
}

/*UpdateEsClusterMetadataSettingsNotFound handles this case with default header values.

The cluster specified by {cluster_id} cannot be found (code: 'clusters.cluster_not_found')
*/
type UpdateEsClusterMetadataSettingsNotFound struct {
	Payload *models.BasicFailedReply
}

func (o *UpdateEsClusterMetadataSettingsNotFound) Error() string {
	return fmt.Sprintf("[PATCH /clusters/elasticsearch/{cluster_id}/metadata/settings][%d] updateEsClusterMetadataSettingsNotFound  %+v", 404, o.Payload)
}

func (o *UpdateEsClusterMetadataSettingsNotFound) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *UpdateEsClusterMetadataSettingsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BasicFailedReply)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateEsClusterMetadataSettingsRetryWith creates a UpdateEsClusterMetadataSettingsRetryWith with default headers values
func NewUpdateEsClusterMetadataSettingsRetryWith() *UpdateEsClusterMetadataSettingsRetryWith {
	return &UpdateEsClusterMetadataSettingsRetryWith{}
}

/*UpdateEsClusterMetadataSettingsRetryWith handles this case with default header values.

Elevated permissions are required. (code: '"root.unauthorized.rbac.elevated_permissions_required"')
*/
type UpdateEsClusterMetadataSettingsRetryWith struct {
	Payload *models.BasicFailedReply
}

func (o *UpdateEsClusterMetadataSettingsRetryWith) Error() string {
	return fmt.Sprintf("[PATCH /clusters/elasticsearch/{cluster_id}/metadata/settings][%d] updateEsClusterMetadataSettingsRetryWith  %+v", 449, o.Payload)
}

func (o *UpdateEsClusterMetadataSettingsRetryWith) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *UpdateEsClusterMetadataSettingsRetryWith) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BasicFailedReply)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
