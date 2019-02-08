// Code generated by go-swagger; DO NOT EDIT.

package connection

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

// NewDeleteConnectionOverrideParams creates a new DeleteConnectionOverrideParams object
// with the default values initialized.
func NewDeleteConnectionOverrideParams() *DeleteConnectionOverrideParams {
	var ()
	return &DeleteConnectionOverrideParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteConnectionOverrideParamsWithTimeout creates a new DeleteConnectionOverrideParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteConnectionOverrideParamsWithTimeout(timeout time.Duration) *DeleteConnectionOverrideParams {
	var ()
	return &DeleteConnectionOverrideParams{

		timeout: timeout,
	}
}

// NewDeleteConnectionOverrideParamsWithContext creates a new DeleteConnectionOverrideParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteConnectionOverrideParamsWithContext(ctx context.Context) *DeleteConnectionOverrideParams {
	var ()
	return &DeleteConnectionOverrideParams{

		Context: ctx,
	}
}

// NewDeleteConnectionOverrideParamsWithHTTPClient creates a new DeleteConnectionOverrideParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteConnectionOverrideParamsWithHTTPClient(client *http.Client) *DeleteConnectionOverrideParams {
	var ()
	return &DeleteConnectionOverrideParams{
		HTTPClient: client,
	}
}

/*DeleteConnectionOverrideParams contains all the parameters to send to the API endpoint
for the delete connection override operation typically these are written to a http.Request
*/
type DeleteConnectionOverrideParams struct {

	/*ConnectionName
	  Name of connection

	*/
	ConnectionName string
	/*OverrideContext
	  Context of connection override

	*/
	OverrideContext string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete connection override params
func (o *DeleteConnectionOverrideParams) WithTimeout(timeout time.Duration) *DeleteConnectionOverrideParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete connection override params
func (o *DeleteConnectionOverrideParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete connection override params
func (o *DeleteConnectionOverrideParams) WithContext(ctx context.Context) *DeleteConnectionOverrideParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete connection override params
func (o *DeleteConnectionOverrideParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete connection override params
func (o *DeleteConnectionOverrideParams) WithHTTPClient(client *http.Client) *DeleteConnectionOverrideParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete connection override params
func (o *DeleteConnectionOverrideParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithConnectionName adds the connectionName to the delete connection override params
func (o *DeleteConnectionOverrideParams) WithConnectionName(connectionName string) *DeleteConnectionOverrideParams {
	o.SetConnectionName(connectionName)
	return o
}

// SetConnectionName adds the connectionName to the delete connection override params
func (o *DeleteConnectionOverrideParams) SetConnectionName(connectionName string) {
	o.ConnectionName = connectionName
}

// WithOverrideContext adds the overrideContext to the delete connection override params
func (o *DeleteConnectionOverrideParams) WithOverrideContext(overrideContext string) *DeleteConnectionOverrideParams {
	o.SetOverrideContext(overrideContext)
	return o
}

// SetOverrideContext adds the overrideContext to the delete connection override params
func (o *DeleteConnectionOverrideParams) SetOverrideContext(overrideContext string) {
	o.OverrideContext = overrideContext
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteConnectionOverrideParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param connection_name
	if err := r.SetPathParam("connection_name", o.ConnectionName); err != nil {
		return err
	}

	// path param override_context
	if err := r.SetPathParam("override_context", o.OverrideContext); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}