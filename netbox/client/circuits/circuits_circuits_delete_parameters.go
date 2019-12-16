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

package circuits

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewCircuitsCircuitsDeleteParams creates a new CircuitsCircuitsDeleteParams object
// with the default values initialized.
func NewCircuitsCircuitsDeleteParams() *CircuitsCircuitsDeleteParams {
	var ()
	return &CircuitsCircuitsDeleteParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCircuitsCircuitsDeleteParamsWithTimeout creates a new CircuitsCircuitsDeleteParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCircuitsCircuitsDeleteParamsWithTimeout(timeout time.Duration) *CircuitsCircuitsDeleteParams {
	var ()
	return &CircuitsCircuitsDeleteParams{

		timeout: timeout,
	}
}

// NewCircuitsCircuitsDeleteParamsWithContext creates a new CircuitsCircuitsDeleteParams object
// with the default values initialized, and the ability to set a context for a request
func NewCircuitsCircuitsDeleteParamsWithContext(ctx context.Context) *CircuitsCircuitsDeleteParams {
	var ()
	return &CircuitsCircuitsDeleteParams{

		Context: ctx,
	}
}

// NewCircuitsCircuitsDeleteParamsWithHTTPClient creates a new CircuitsCircuitsDeleteParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCircuitsCircuitsDeleteParamsWithHTTPClient(client *http.Client) *CircuitsCircuitsDeleteParams {
	var ()
	return &CircuitsCircuitsDeleteParams{
		HTTPClient: client,
	}
}

/*CircuitsCircuitsDeleteParams contains all the parameters to send to the API endpoint
for the circuits circuits delete operation typically these are written to a http.Request
*/
type CircuitsCircuitsDeleteParams struct {

	/*ID
	  A unique integer value identifying this circuit.

	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the circuits circuits delete params
func (o *CircuitsCircuitsDeleteParams) WithTimeout(timeout time.Duration) *CircuitsCircuitsDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the circuits circuits delete params
func (o *CircuitsCircuitsDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the circuits circuits delete params
func (o *CircuitsCircuitsDeleteParams) WithContext(ctx context.Context) *CircuitsCircuitsDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the circuits circuits delete params
func (o *CircuitsCircuitsDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the circuits circuits delete params
func (o *CircuitsCircuitsDeleteParams) WithHTTPClient(client *http.Client) *CircuitsCircuitsDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the circuits circuits delete params
func (o *CircuitsCircuitsDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the circuits circuits delete params
func (o *CircuitsCircuitsDeleteParams) WithID(id int64) *CircuitsCircuitsDeleteParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the circuits circuits delete params
func (o *CircuitsCircuitsDeleteParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *CircuitsCircuitsDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
