// Code generated by go-swagger; DO NOT EDIT.

package theme

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

// NewThemeOrDefaultParams creates a new ThemeOrDefaultParams object
// with the default values initialized.
func NewThemeOrDefaultParams() *ThemeOrDefaultParams {
	var ()
	return &ThemeOrDefaultParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewThemeOrDefaultParamsWithTimeout creates a new ThemeOrDefaultParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewThemeOrDefaultParamsWithTimeout(timeout time.Duration) *ThemeOrDefaultParams {
	var ()
	return &ThemeOrDefaultParams{

		timeout: timeout,
	}
}

// NewThemeOrDefaultParamsWithContext creates a new ThemeOrDefaultParams object
// with the default values initialized, and the ability to set a context for a request
func NewThemeOrDefaultParamsWithContext(ctx context.Context) *ThemeOrDefaultParams {
	var ()
	return &ThemeOrDefaultParams{

		Context: ctx,
	}
}

// NewThemeOrDefaultParamsWithHTTPClient creates a new ThemeOrDefaultParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewThemeOrDefaultParamsWithHTTPClient(client *http.Client) *ThemeOrDefaultParams {
	var ()
	return &ThemeOrDefaultParams{
		HTTPClient: client,
	}
}

/*ThemeOrDefaultParams contains all the parameters to send to the API endpoint
for the theme or default operation typically these are written to a http.Request
*/
type ThemeOrDefaultParams struct {

	/*Name
	  Name of theme

	*/
	Name string
	/*Ts
	  Timestamp representing the target datetime for the active period. Defaults to 'now'

	*/
	Ts *strfmt.DateTime

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the theme or default params
func (o *ThemeOrDefaultParams) WithTimeout(timeout time.Duration) *ThemeOrDefaultParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the theme or default params
func (o *ThemeOrDefaultParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the theme or default params
func (o *ThemeOrDefaultParams) WithContext(ctx context.Context) *ThemeOrDefaultParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the theme or default params
func (o *ThemeOrDefaultParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the theme or default params
func (o *ThemeOrDefaultParams) WithHTTPClient(client *http.Client) *ThemeOrDefaultParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the theme or default params
func (o *ThemeOrDefaultParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the theme or default params
func (o *ThemeOrDefaultParams) WithName(name string) *ThemeOrDefaultParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the theme or default params
func (o *ThemeOrDefaultParams) SetName(name string) {
	o.Name = name
}

// WithTs adds the ts to the theme or default params
func (o *ThemeOrDefaultParams) WithTs(ts *strfmt.DateTime) *ThemeOrDefaultParams {
	o.SetTs(ts)
	return o
}

// SetTs adds the ts to the theme or default params
func (o *ThemeOrDefaultParams) SetTs(ts *strfmt.DateTime) {
	o.Ts = ts
}

// WriteToRequest writes these params to a swagger request
func (o *ThemeOrDefaultParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param name
	qrName := o.Name
	qName := qrName
	if qName != "" {
		if err := r.SetQueryParam("name", qName); err != nil {
			return err
		}
	}

	if o.Ts != nil {

		// query param ts
		var qrTs strfmt.DateTime
		if o.Ts != nil {
			qrTs = *o.Ts
		}
		qTs := qrTs.String()
		if qTs != "" {
			if err := r.SetQueryParam("ts", qTs); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
