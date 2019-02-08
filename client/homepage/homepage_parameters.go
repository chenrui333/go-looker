// Code generated by go-swagger; DO NOT EDIT.

package homepage

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewHomepageParams creates a new HomepageParams object
// with the default values initialized.
func NewHomepageParams() *HomepageParams {
	var ()
	return &HomepageParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewHomepageParamsWithTimeout creates a new HomepageParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewHomepageParamsWithTimeout(timeout time.Duration) *HomepageParams {
	var ()
	return &HomepageParams{

		timeout: timeout,
	}
}

// NewHomepageParamsWithContext creates a new HomepageParams object
// with the default values initialized, and the ability to set a context for a request
func NewHomepageParamsWithContext(ctx context.Context) *HomepageParams {
	var ()
	return &HomepageParams{

		Context: ctx,
	}
}

// NewHomepageParamsWithHTTPClient creates a new HomepageParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewHomepageParamsWithHTTPClient(client *http.Client) *HomepageParams {
	var ()
	return &HomepageParams{
		HTTPClient: client,
	}
}

/*HomepageParams contains all the parameters to send to the API endpoint
for the homepage operation typically these are written to a http.Request
*/
type HomepageParams struct {

	/*Fields
	  Requested fields.

	*/
	Fields *string
	/*HomepageID
	  Id of homepage

	*/
	HomepageID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the homepage params
func (o *HomepageParams) WithTimeout(timeout time.Duration) *HomepageParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the homepage params
func (o *HomepageParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the homepage params
func (o *HomepageParams) WithContext(ctx context.Context) *HomepageParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the homepage params
func (o *HomepageParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the homepage params
func (o *HomepageParams) WithHTTPClient(client *http.Client) *HomepageParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the homepage params
func (o *HomepageParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFields adds the fields to the homepage params
func (o *HomepageParams) WithFields(fields *string) *HomepageParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the homepage params
func (o *HomepageParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithHomepageID adds the homepageID to the homepage params
func (o *HomepageParams) WithHomepageID(homepageID int64) *HomepageParams {
	o.SetHomepageID(homepageID)
	return o
}

// SetHomepageID adds the homepageId to the homepage params
func (o *HomepageParams) SetHomepageID(homepageID int64) {
	o.HomepageID = homepageID
}

// WriteToRequest writes these params to a swagger request
func (o *HomepageParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param homepage_id
	if err := r.SetPathParam("homepage_id", swag.FormatInt64(o.HomepageID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
