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
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/digitalocean/go-netbox/netbox/models"
)

// IPAMVlansListReader is a Reader for the IPAMVlansList structure.
type IPAMVlansListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IPAMVlansListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewIPAMVlansListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewIPAMVlansListOK creates a IPAMVlansListOK with default headers values
func NewIPAMVlansListOK() *IPAMVlansListOK {
	return &IPAMVlansListOK{}
}

/*IPAMVlansListOK handles this case with default header values.

IPAMVlansListOK ipam vlans list o k
*/
type IPAMVlansListOK struct {
	Payload *IPAMVlansListOKBody
}

func (o *IPAMVlansListOK) Error() string {
	return fmt.Sprintf("[GET /ipam/vlans/][%d] ipamVlansListOK  %+v", 200, o.Payload)
}

func (o *IPAMVlansListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(IPAMVlansListOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*IPAMVlansListOKBody IP a m vlans list o k body
swagger:model IPAMVlansListOKBody
*/
type IPAMVlansListOKBody struct {

	// count
	// Required: true
	Count *int64 `json:"count"`

	// next
	// Format: uri
	Next *strfmt.URI `json:"next,omitempty"`

	// previous
	// Format: uri
	Previous *strfmt.URI `json:"previous,omitempty"`

	// results
	// Required: true
	Results []*models.VLAN `json:"results"`
}

// Validate validates this IP a m vlans list o k body
func (o *IPAMVlansListOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateCount(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateNext(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validatePrevious(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateResults(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *IPAMVlansListOKBody) validateCount(formats strfmt.Registry) error {

	if err := validate.Required("ipamVlansListOK"+"."+"count", "body", o.Count); err != nil {
		return err
	}

	return nil
}

func (o *IPAMVlansListOKBody) validateNext(formats strfmt.Registry) error {

	if swag.IsZero(o.Next) { // not required
		return nil
	}

	if err := validate.FormatOf("ipamVlansListOK"+"."+"next", "body", "uri", o.Next.String(), formats); err != nil {
		return err
	}

	return nil
}

func (o *IPAMVlansListOKBody) validatePrevious(formats strfmt.Registry) error {

	if swag.IsZero(o.Previous) { // not required
		return nil
	}

	if err := validate.FormatOf("ipamVlansListOK"+"."+"previous", "body", "uri", o.Previous.String(), formats); err != nil {
		return err
	}

	return nil
}

func (o *IPAMVlansListOKBody) validateResults(formats strfmt.Registry) error {

	if err := validate.Required("ipamVlansListOK"+"."+"results", "body", o.Results); err != nil {
		return err
	}

	for i := 0; i < len(o.Results); i++ {
		if swag.IsZero(o.Results[i]) { // not required
			continue
		}

		if o.Results[i] != nil {
			if err := o.Results[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("ipamVlansListOK" + "." + "results" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *IPAMVlansListOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *IPAMVlansListOKBody) UnmarshalBinary(b []byte) error {
	var res IPAMVlansListOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
