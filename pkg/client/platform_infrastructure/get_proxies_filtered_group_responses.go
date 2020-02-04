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

// GetProxiesFilteredGroupReader is a Reader for the GetProxiesFilteredGroup structure.
type GetProxiesFilteredGroupReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetProxiesFilteredGroupReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetProxiesFilteredGroupOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetProxiesFilteredGroupNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetProxiesFilteredGroupOK creates a GetProxiesFilteredGroupOK with default headers values
func NewGetProxiesFilteredGroupOK() *GetProxiesFilteredGroupOK {
	return &GetProxiesFilteredGroupOK{}
}

/*GetProxiesFilteredGroupOK handles this case with default header values.

Data for the filtered group of proxies identified by {proxies_filtered_group_id}
*/
type GetProxiesFilteredGroupOK struct {
	/*The date-time when the resource was created (ISO format relative to UTC)
	 */
	XCloudResourceCreated string
	/*The date-time when the resource was last modified (ISO format relative to UTC)
	 */
	XCloudResourceLastModified string
	/*The resource version, which is used to avoid update conflicts with concurrent operations
	 */
	XCloudResourceVersion string

	Payload *models.ProxiesFilteredGroup
}

func (o *GetProxiesFilteredGroupOK) Error() string {
	return fmt.Sprintf("[GET /platform/infrastructure/proxies/filtered-groups/{proxies_filtered_group_id}][%d] getProxiesFilteredGroupOK  %+v", 200, o.Payload)
}

func (o *GetProxiesFilteredGroupOK) GetPayload() *models.ProxiesFilteredGroup {
	return o.Payload
}

func (o *GetProxiesFilteredGroupOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header x-cloud-resource-created
	o.XCloudResourceCreated = response.GetHeader("x-cloud-resource-created")

	// response header x-cloud-resource-last-modified
	o.XCloudResourceLastModified = response.GetHeader("x-cloud-resource-last-modified")

	// response header x-cloud-resource-version
	o.XCloudResourceVersion = response.GetHeader("x-cloud-resource-version")

	o.Payload = new(models.ProxiesFilteredGroup)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetProxiesFilteredGroupNotFound creates a GetProxiesFilteredGroupNotFound with default headers values
func NewGetProxiesFilteredGroupNotFound() *GetProxiesFilteredGroupNotFound {
	return &GetProxiesFilteredGroupNotFound{}
}

/*GetProxiesFilteredGroupNotFound handles this case with default header values.

Unable to find the {proxies_filtered_group_id} specified filtered group of proxies. Edit your request, then try again. (code: `proxies.proxies_filtered_group_not_found`)
*/
type GetProxiesFilteredGroupNotFound struct {
	/*The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

func (o *GetProxiesFilteredGroupNotFound) Error() string {
	return fmt.Sprintf("[GET /platform/infrastructure/proxies/filtered-groups/{proxies_filtered_group_id}][%d] getProxiesFilteredGroupNotFound  %+v", 404, o.Payload)
}

func (o *GetProxiesFilteredGroupNotFound) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *GetProxiesFilteredGroupNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header x-cloud-error-codes
	o.XCloudErrorCodes = response.GetHeader("x-cloud-error-codes")

	o.Payload = new(models.BasicFailedReply)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
