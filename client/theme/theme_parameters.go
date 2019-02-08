// Code generated by go-swagger; DO NOT EDIT.

package theme

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewThemeParams creates a new ThemeParams object
// with the default values initialized.
func NewThemeParams() *ThemeParams {
	var ()
	return &ThemeParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewThemeParamsWithTimeout creates a new ThemeParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewThemeParamsWithTimeout(timeout time.Duration) *ThemeParams {
	var ()
	return &ThemeParams{

		timeout: timeout,
	}
}

// NewThemeParamsWithContext creates a new ThemeParams object
// with the default values initialized, and the ability to set a context for a request
func NewThemeParamsWithContext(ctx context.Context) *ThemeParams {
	var ()
	return &ThemeParams{

		Context: ctx,
	}
}

// NewThemeParamsWithHTTPClient creates a new ThemeParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewThemeParamsWithHTTPClient(client *http.Client) *ThemeParams {
	var ()
	return &ThemeParams{
		HTTPClient: client,
	}
}

/*ThemeParams contains all the parameters to send to the API endpoint
for the theme operation typically these are written to a http.Request
*/
type ThemeParams struct {

	/*Fields
	  Requested fields.

	*/
	Fields *string
	/*ThemeID
	  Id of Theme

	*/
	ThemeID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the theme params
func (o *ThemeParams) WithTimeout(timeout time.Duration) *ThemeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the theme params
func (o *ThemeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the theme params
func (o *ThemeParams) WithContext(ctx context.Context) *ThemeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the theme params
func (o *ThemeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the theme params
func (o *ThemeParams) WithHTTPClient(client *http.Client) *ThemeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the theme params
func (o *ThemeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFields adds the fields to the theme params
func (o *ThemeParams) WithFields(fields *string) *ThemeParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the theme params
func (o *ThemeParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithThemeID adds the themeID to the theme params
func (o *ThemeParams) WithThemeID(themeID string) *ThemeParams {
	o.SetThemeID(themeID)
	return o
}

// SetThemeID adds the themeId to the theme params
func (o *ThemeParams) SetThemeID(themeID string) {
	o.ThemeID = themeID
}

// WriteToRequest writes these params to a swagger request
func (o *ThemeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param theme_id
	if err := r.SetPathParam("theme_id", o.ThemeID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
