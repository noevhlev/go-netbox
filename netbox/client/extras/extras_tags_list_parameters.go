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

package extras

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

// NewExtrasTagsListParams creates a new ExtrasTagsListParams object
// with the default values initialized.
func NewExtrasTagsListParams() *ExtrasTagsListParams {
	var ()
	return &ExtrasTagsListParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewExtrasTagsListParamsWithTimeout creates a new ExtrasTagsListParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewExtrasTagsListParamsWithTimeout(timeout time.Duration) *ExtrasTagsListParams {
	var ()
	return &ExtrasTagsListParams{

		timeout: timeout,
	}
}

// NewExtrasTagsListParamsWithContext creates a new ExtrasTagsListParams object
// with the default values initialized, and the ability to set a context for a request
func NewExtrasTagsListParamsWithContext(ctx context.Context) *ExtrasTagsListParams {
	var ()
	return &ExtrasTagsListParams{

		Context: ctx,
	}
}

// NewExtrasTagsListParamsWithHTTPClient creates a new ExtrasTagsListParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewExtrasTagsListParamsWithHTTPClient(client *http.Client) *ExtrasTagsListParams {
	var ()
	return &ExtrasTagsListParams{
		HTTPClient: client,
	}
}

/*ExtrasTagsListParams contains all the parameters to send to the API endpoint
for the extras tags list operation typically these are written to a http.Request
*/
type ExtrasTagsListParams struct {

	/*Limit
	  Number of results to return per page.

	*/
	Limit *int64
	/*Name*/
	Name *string
	/*Offset
	  The initial index from which to return the results.

	*/
	Offset *int64
	/*Q*/
	Q *string
	/*Slug*/
	Slug *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the extras tags list params
func (o *ExtrasTagsListParams) WithTimeout(timeout time.Duration) *ExtrasTagsListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the extras tags list params
func (o *ExtrasTagsListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the extras tags list params
func (o *ExtrasTagsListParams) WithContext(ctx context.Context) *ExtrasTagsListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the extras tags list params
func (o *ExtrasTagsListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the extras tags list params
func (o *ExtrasTagsListParams) WithHTTPClient(client *http.Client) *ExtrasTagsListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the extras tags list params
func (o *ExtrasTagsListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLimit adds the limit to the extras tags list params
func (o *ExtrasTagsListParams) WithLimit(limit *int64) *ExtrasTagsListParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the extras tags list params
func (o *ExtrasTagsListParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithName adds the name to the extras tags list params
func (o *ExtrasTagsListParams) WithName(name *string) *ExtrasTagsListParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the extras tags list params
func (o *ExtrasTagsListParams) SetName(name *string) {
	o.Name = name
}

// WithOffset adds the offset to the extras tags list params
func (o *ExtrasTagsListParams) WithOffset(offset *int64) *ExtrasTagsListParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the extras tags list params
func (o *ExtrasTagsListParams) SetOffset(offset *int64) {
	o.Offset = offset
}

// WithQ adds the q to the extras tags list params
func (o *ExtrasTagsListParams) WithQ(q *string) *ExtrasTagsListParams {
	o.SetQ(q)
	return o
}

// SetQ adds the q to the extras tags list params
func (o *ExtrasTagsListParams) SetQ(q *string) {
	o.Q = q
}

// WithSlug adds the slug to the extras tags list params
func (o *ExtrasTagsListParams) WithSlug(slug *string) *ExtrasTagsListParams {
	o.SetSlug(slug)
	return o
}

// SetSlug adds the slug to the extras tags list params
func (o *ExtrasTagsListParams) SetSlug(slug *string) {
	o.Slug = slug
}

// WriteToRequest writes these params to a swagger request
func (o *ExtrasTagsListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Limit != nil {

		// query param limit
		var qrLimit int64
		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := swag.FormatInt64(qrLimit)
		if qLimit != "" {
			if err := r.SetQueryParam("limit", qLimit); err != nil {
				return err
			}
		}

	}

	if o.Name != nil {

		// query param name
		var qrName string
		if o.Name != nil {
			qrName = *o.Name
		}
		qName := qrName
		if qName != "" {
			if err := r.SetQueryParam("name", qName); err != nil {
				return err
			}
		}

	}

	if o.Offset != nil {

		// query param offset
		var qrOffset int64
		if o.Offset != nil {
			qrOffset = *o.Offset
		}
		qOffset := swag.FormatInt64(qrOffset)
		if qOffset != "" {
			if err := r.SetQueryParam("offset", qOffset); err != nil {
				return err
			}
		}

	}

	if o.Q != nil {

		// query param q
		var qrQ string
		if o.Q != nil {
			qrQ = *o.Q
		}
		qQ := qrQ
		if qQ != "" {
			if err := r.SetQueryParam("q", qQ); err != nil {
				return err
			}
		}

	}

	if o.Slug != nil {

		// query param slug
		var qrSlug string
		if o.Slug != nil {
			qrSlug = *o.Slug
		}
		qSlug := qrSlug
		if qSlug != "" {
			if err := r.SetQueryParam("slug", qSlug); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
