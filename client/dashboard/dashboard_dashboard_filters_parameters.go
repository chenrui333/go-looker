// Code generated by go-swagger; DO NOT EDIT.

package dashboard

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

// NewDashboardDashboardFiltersParams creates a new DashboardDashboardFiltersParams object
// with the default values initialized.
func NewDashboardDashboardFiltersParams() *DashboardDashboardFiltersParams {
	var ()
	return &DashboardDashboardFiltersParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDashboardDashboardFiltersParamsWithTimeout creates a new DashboardDashboardFiltersParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDashboardDashboardFiltersParamsWithTimeout(timeout time.Duration) *DashboardDashboardFiltersParams {
	var ()
	return &DashboardDashboardFiltersParams{

		timeout: timeout,
	}
}

// NewDashboardDashboardFiltersParamsWithContext creates a new DashboardDashboardFiltersParams object
// with the default values initialized, and the ability to set a context for a request
func NewDashboardDashboardFiltersParamsWithContext(ctx context.Context) *DashboardDashboardFiltersParams {
	var ()
	return &DashboardDashboardFiltersParams{

		Context: ctx,
	}
}

// NewDashboardDashboardFiltersParamsWithHTTPClient creates a new DashboardDashboardFiltersParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDashboardDashboardFiltersParamsWithHTTPClient(client *http.Client) *DashboardDashboardFiltersParams {
	var ()
	return &DashboardDashboardFiltersParams{
		HTTPClient: client,
	}
}

/*DashboardDashboardFiltersParams contains all the parameters to send to the API endpoint
for the dashboard dashboard filters operation typically these are written to a http.Request
*/
type DashboardDashboardFiltersParams struct {

	/*DashboardID
	  Id of dashboard

	*/
	DashboardID string
	/*Fields
	  Requested fields.

	*/
	Fields *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the dashboard dashboard filters params
func (o *DashboardDashboardFiltersParams) WithTimeout(timeout time.Duration) *DashboardDashboardFiltersParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dashboard dashboard filters params
func (o *DashboardDashboardFiltersParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dashboard dashboard filters params
func (o *DashboardDashboardFiltersParams) WithContext(ctx context.Context) *DashboardDashboardFiltersParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dashboard dashboard filters params
func (o *DashboardDashboardFiltersParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dashboard dashboard filters params
func (o *DashboardDashboardFiltersParams) WithHTTPClient(client *http.Client) *DashboardDashboardFiltersParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dashboard dashboard filters params
func (o *DashboardDashboardFiltersParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDashboardID adds the dashboardID to the dashboard dashboard filters params
func (o *DashboardDashboardFiltersParams) WithDashboardID(dashboardID string) *DashboardDashboardFiltersParams {
	o.SetDashboardID(dashboardID)
	return o
}

// SetDashboardID adds the dashboardId to the dashboard dashboard filters params
func (o *DashboardDashboardFiltersParams) SetDashboardID(dashboardID string) {
	o.DashboardID = dashboardID
}

// WithFields adds the fields to the dashboard dashboard filters params
func (o *DashboardDashboardFiltersParams) WithFields(fields *string) *DashboardDashboardFiltersParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the dashboard dashboard filters params
func (o *DashboardDashboardFiltersParams) SetFields(fields *string) {
	o.Fields = fields
}

// WriteToRequest writes these params to a swagger request
func (o *DashboardDashboardFiltersParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param dashboard_id
	if err := r.SetPathParam("dashboard_id", o.DashboardID); err != nil {
		return err
	}

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
