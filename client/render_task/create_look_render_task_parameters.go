// Code generated by go-swagger; DO NOT EDIT.

package render_task

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

// NewCreateLookRenderTaskParams creates a new CreateLookRenderTaskParams object
// with the default values initialized.
func NewCreateLookRenderTaskParams() *CreateLookRenderTaskParams {
	var ()
	return &CreateLookRenderTaskParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateLookRenderTaskParamsWithTimeout creates a new CreateLookRenderTaskParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateLookRenderTaskParamsWithTimeout(timeout time.Duration) *CreateLookRenderTaskParams {
	var ()
	return &CreateLookRenderTaskParams{

		timeout: timeout,
	}
}

// NewCreateLookRenderTaskParamsWithContext creates a new CreateLookRenderTaskParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateLookRenderTaskParamsWithContext(ctx context.Context) *CreateLookRenderTaskParams {
	var ()
	return &CreateLookRenderTaskParams{

		Context: ctx,
	}
}

// NewCreateLookRenderTaskParamsWithHTTPClient creates a new CreateLookRenderTaskParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateLookRenderTaskParamsWithHTTPClient(client *http.Client) *CreateLookRenderTaskParams {
	var ()
	return &CreateLookRenderTaskParams{
		HTTPClient: client,
	}
}

/*CreateLookRenderTaskParams contains all the parameters to send to the API endpoint
for the create look render task operation typically these are written to a http.Request
*/
type CreateLookRenderTaskParams struct {

	/*Fields
	  Requested fields.

	*/
	Fields *string
	/*Height
	  Output height in pixels

	*/
	Height int64
	/*LookID
	  Id of look to render

	*/
	LookID int64
	/*ResultFormat
	  Output type: png, or jpg

	*/
	ResultFormat string
	/*Width
	  Output width in pixels

	*/
	Width int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create look render task params
func (o *CreateLookRenderTaskParams) WithTimeout(timeout time.Duration) *CreateLookRenderTaskParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create look render task params
func (o *CreateLookRenderTaskParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create look render task params
func (o *CreateLookRenderTaskParams) WithContext(ctx context.Context) *CreateLookRenderTaskParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create look render task params
func (o *CreateLookRenderTaskParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create look render task params
func (o *CreateLookRenderTaskParams) WithHTTPClient(client *http.Client) *CreateLookRenderTaskParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create look render task params
func (o *CreateLookRenderTaskParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFields adds the fields to the create look render task params
func (o *CreateLookRenderTaskParams) WithFields(fields *string) *CreateLookRenderTaskParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the create look render task params
func (o *CreateLookRenderTaskParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithHeight adds the height to the create look render task params
func (o *CreateLookRenderTaskParams) WithHeight(height int64) *CreateLookRenderTaskParams {
	o.SetHeight(height)
	return o
}

// SetHeight adds the height to the create look render task params
func (o *CreateLookRenderTaskParams) SetHeight(height int64) {
	o.Height = height
}

// WithLookID adds the lookID to the create look render task params
func (o *CreateLookRenderTaskParams) WithLookID(lookID int64) *CreateLookRenderTaskParams {
	o.SetLookID(lookID)
	return o
}

// SetLookID adds the lookId to the create look render task params
func (o *CreateLookRenderTaskParams) SetLookID(lookID int64) {
	o.LookID = lookID
}

// WithResultFormat adds the resultFormat to the create look render task params
func (o *CreateLookRenderTaskParams) WithResultFormat(resultFormat string) *CreateLookRenderTaskParams {
	o.SetResultFormat(resultFormat)
	return o
}

// SetResultFormat adds the resultFormat to the create look render task params
func (o *CreateLookRenderTaskParams) SetResultFormat(resultFormat string) {
	o.ResultFormat = resultFormat
}

// WithWidth adds the width to the create look render task params
func (o *CreateLookRenderTaskParams) WithWidth(width int64) *CreateLookRenderTaskParams {
	o.SetWidth(width)
	return o
}

// SetWidth adds the width to the create look render task params
func (o *CreateLookRenderTaskParams) SetWidth(width int64) {
	o.Width = width
}

// WriteToRequest writes these params to a swagger request
func (o *CreateLookRenderTaskParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// query param height
	qrHeight := o.Height
	qHeight := swag.FormatInt64(qrHeight)
	if qHeight != "" {
		if err := r.SetQueryParam("height", qHeight); err != nil {
			return err
		}
	}

	// path param look_id
	if err := r.SetPathParam("look_id", swag.FormatInt64(o.LookID)); err != nil {
		return err
	}

	// path param result_format
	if err := r.SetPathParam("result_format", o.ResultFormat); err != nil {
		return err
	}

	// query param width
	qrWidth := o.Width
	qWidth := swag.FormatInt64(qrWidth)
	if qWidth != "" {
		if err := r.SetQueryParam("width", qWidth); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
