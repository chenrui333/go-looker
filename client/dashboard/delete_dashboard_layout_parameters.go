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

// NewDeleteDashboardLayoutParams creates a new DeleteDashboardLayoutParams object
// with the default values initialized.
func NewDeleteDashboardLayoutParams() *DeleteDashboardLayoutParams {
	var ()
	return &DeleteDashboardLayoutParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteDashboardLayoutParamsWithTimeout creates a new DeleteDashboardLayoutParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteDashboardLayoutParamsWithTimeout(timeout time.Duration) *DeleteDashboardLayoutParams {
	var ()
	return &DeleteDashboardLayoutParams{

		timeout: timeout,
	}
}

// NewDeleteDashboardLayoutParamsWithContext creates a new DeleteDashboardLayoutParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteDashboardLayoutParamsWithContext(ctx context.Context) *DeleteDashboardLayoutParams {
	var ()
	return &DeleteDashboardLayoutParams{

		Context: ctx,
	}
}

// NewDeleteDashboardLayoutParamsWithHTTPClient creates a new DeleteDashboardLayoutParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteDashboardLayoutParamsWithHTTPClient(client *http.Client) *DeleteDashboardLayoutParams {
	var ()
	return &DeleteDashboardLayoutParams{
		HTTPClient: client,
	}
}

/*DeleteDashboardLayoutParams contains all the parameters to send to the API endpoint
for the delete dashboard layout operation typically these are written to a http.Request
*/
type DeleteDashboardLayoutParams struct {

	/*DashboardLayoutID
	  Id of dashboard layout

	*/
	DashboardLayoutID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete dashboard layout params
func (o *DeleteDashboardLayoutParams) WithTimeout(timeout time.Duration) *DeleteDashboardLayoutParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete dashboard layout params
func (o *DeleteDashboardLayoutParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete dashboard layout params
func (o *DeleteDashboardLayoutParams) WithContext(ctx context.Context) *DeleteDashboardLayoutParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete dashboard layout params
func (o *DeleteDashboardLayoutParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete dashboard layout params
func (o *DeleteDashboardLayoutParams) WithHTTPClient(client *http.Client) *DeleteDashboardLayoutParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete dashboard layout params
func (o *DeleteDashboardLayoutParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDashboardLayoutID adds the dashboardLayoutID to the delete dashboard layout params
func (o *DeleteDashboardLayoutParams) WithDashboardLayoutID(dashboardLayoutID string) *DeleteDashboardLayoutParams {
	o.SetDashboardLayoutID(dashboardLayoutID)
	return o
}

// SetDashboardLayoutID adds the dashboardLayoutId to the delete dashboard layout params
func (o *DeleteDashboardLayoutParams) SetDashboardLayoutID(dashboardLayoutID string) {
	o.DashboardLayoutID = dashboardLayoutID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteDashboardLayoutParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param dashboard_layout_id
	if err := r.SetPathParam("dashboard_layout_id", o.DashboardLayoutID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}