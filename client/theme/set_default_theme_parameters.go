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

// NewSetDefaultThemeParams creates a new SetDefaultThemeParams object
// with the default values initialized.
func NewSetDefaultThemeParams() *SetDefaultThemeParams {
	var ()
	return &SetDefaultThemeParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewSetDefaultThemeParamsWithTimeout creates a new SetDefaultThemeParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewSetDefaultThemeParamsWithTimeout(timeout time.Duration) *SetDefaultThemeParams {
	var ()
	return &SetDefaultThemeParams{

		timeout: timeout,
	}
}

// NewSetDefaultThemeParamsWithContext creates a new SetDefaultThemeParams object
// with the default values initialized, and the ability to set a context for a request
func NewSetDefaultThemeParamsWithContext(ctx context.Context) *SetDefaultThemeParams {
	var ()
	return &SetDefaultThemeParams{

		Context: ctx,
	}
}

// NewSetDefaultThemeParamsWithHTTPClient creates a new SetDefaultThemeParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewSetDefaultThemeParamsWithHTTPClient(client *http.Client) *SetDefaultThemeParams {
	var ()
	return &SetDefaultThemeParams{
		HTTPClient: client,
	}
}

/*SetDefaultThemeParams contains all the parameters to send to the API endpoint
for the set default theme operation typically these are written to a http.Request
*/
type SetDefaultThemeParams struct {

	/*Name
	  Name of theme to set as default

	*/
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the set default theme params
func (o *SetDefaultThemeParams) WithTimeout(timeout time.Duration) *SetDefaultThemeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the set default theme params
func (o *SetDefaultThemeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the set default theme params
func (o *SetDefaultThemeParams) WithContext(ctx context.Context) *SetDefaultThemeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the set default theme params
func (o *SetDefaultThemeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the set default theme params
func (o *SetDefaultThemeParams) WithHTTPClient(client *http.Client) *SetDefaultThemeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the set default theme params
func (o *SetDefaultThemeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the set default theme params
func (o *SetDefaultThemeParams) WithName(name string) *SetDefaultThemeParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the set default theme params
func (o *SetDefaultThemeParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *SetDefaultThemeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
