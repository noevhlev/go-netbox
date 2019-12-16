// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2018 The go-netbox Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package ipam

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/digitalocean/go-netbox/netbox/models"
)

// IPAMPrefixesAvailablePrefixesCreateReader is a Reader for the IPAMPrefixesAvailablePrefixesCreate structure.
type IPAMPrefixesAvailablePrefixesCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IPAMPrefixesAvailablePrefixesCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewIPAMPrefixesAvailablePrefixesCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewIPAMPrefixesAvailablePrefixesCreateCreated creates a IPAMPrefixesAvailablePrefixesCreateCreated with default headers values
func NewIPAMPrefixesAvailablePrefixesCreateCreated() *IPAMPrefixesAvailablePrefixesCreateCreated {
	return &IPAMPrefixesAvailablePrefixesCreateCreated{}
}

/*IPAMPrefixesAvailablePrefixesCreateCreated handles this case with default header values.

IPAMPrefixesAvailablePrefixesCreateCreated ipam prefixes available prefixes create created
*/
type IPAMPrefixesAvailablePrefixesCreateCreated struct {
	Payload *models.Prefix
}

func (o *IPAMPrefixesAvailablePrefixesCreateCreated) Error() string {
	return fmt.Sprintf("[POST /ipam/prefixes/{id}/available-prefixes/][%d] ipamPrefixesAvailablePrefixesCreateCreated  %+v", 201, o.Payload)
}

func (o *IPAMPrefixesAvailablePrefixesCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Prefix)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
