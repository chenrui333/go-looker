// Code generated by go-swagger; DO NOT EDIT.

package connection

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewAllDialectInfosParams creates a new AllDialectInfosParams object
// with the default values initialized.
func NewAllDialectInfosParams() *AllDialectInfosParams {
	var ()
	return &AllDialectInfosParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewAllDialectInfosParamsWithTimeout creates a new AllDialectInfosParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAllDialectInfosParamsWithTimeout(timeout time.Duration) *AllDialectInfosParams {
	var ()
	return &AllDialectInfosParams{

		timeout: timeout,
	}
}

// NewAllDialectInfosParamsWithContext creates a new AllDialectInfosParams object
// with the default values initialized, and the ability to set a context for a request
func NewAllDialectInfosParamsWithContext(ctx context.Context) *AllDialectInfosParams {
	var ()
	return &AllDialectInfosParams{

		Context: ctx,
	}
}

// NewAllDialectInfosParamsWithHTTPClient creates a new AllDialectInfosParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAllDialectInfosParamsWithHTTPClient(client *http.Client) *AllDialectInfosParams {
	var ()
	return &AllDialectInfosParams{
		HTTPClient: client,
	}
}

/*AllDialectInfosParams contains all the parameters to send to the API endpoint
for the all dialect infos operation typically these are written to a http.Request
*/
type AllDialectInfosParams struct {

	/*Fields
	  Requested fields.

	*/
	Fields *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the all dialect infos params
func (o *AllDialectInfosParams) WithTimeout(timeout time.Duration) *AllDialectInfosParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the all dialect infos params
func (o *AllDialectInfosParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the all dialect infos params
func (o *AllDialectInfosParams) WithContext(ctx context.Context) *AllDialectInfosParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the all dialect infos params
func (o *AllDialectInfosParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the all dialect infos params
func (o *AllDialectInfosParams) WithHTTPClient(client *http.Client) *AllDialectInfosParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the all dialect infos params
func (o *AllDialectInfosParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFields adds the fields to the all dialect infos params
func (o *AllDialectInfosParams) WithFields(fields *string) *AllDialectInfosParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the all dialect infos params
func (o *AllDialectInfosParams) SetFields(fields *string) {
	o.Fields = fields
}

// WriteToRequest writes these params to a swagger request
func (o *AllDialectInfosParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Fields != nil {

		// query param fields
		var qrFields string
		if o.Fields != nil {
			qrFields = *o.Fields
		}
		qFields := qrFields
		if qFields != "" {
			if err := r.SetQueryParam("fields", qFields); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
