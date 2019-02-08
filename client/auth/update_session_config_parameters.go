// Code generated by go-swagger; DO NOT EDIT.

package auth

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

	models "github.com/chenrui333/go-looker/models"
)

// NewUpdateSessionConfigParams creates a new UpdateSessionConfigParams object
// with the default values initialized.
func NewUpdateSessionConfigParams() *UpdateSessionConfigParams {
	var ()
	return &UpdateSessionConfigParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateSessionConfigParamsWithTimeout creates a new UpdateSessionConfigParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateSessionConfigParamsWithTimeout(timeout time.Duration) *UpdateSessionConfigParams {
	var ()
	return &UpdateSessionConfigParams{

		timeout: timeout,
	}
}

// NewUpdateSessionConfigParamsWithContext creates a new UpdateSessionConfigParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateSessionConfigParamsWithContext(ctx context.Context) *UpdateSessionConfigParams {
	var ()
	return &UpdateSessionConfigParams{

		Context: ctx,
	}
}

// NewUpdateSessionConfigParamsWithHTTPClient creates a new UpdateSessionConfigParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateSessionConfigParamsWithHTTPClient(client *http.Client) *UpdateSessionConfigParams {
	var ()
	return &UpdateSessionConfigParams{
		HTTPClient: client,
	}
}

/*UpdateSessionConfigParams contains all the parameters to send to the API endpoint
for the update session config operation typically these are written to a http.Request
*/
type UpdateSessionConfigParams struct {

	/*Body
	  Session Config

	*/
	Body *models.SessionConfig

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update session config params
func (o *UpdateSessionConfigParams) WithTimeout(timeout time.Duration) *UpdateSessionConfigParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update session config params
func (o *UpdateSessionConfigParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update session config params
func (o *UpdateSessionConfigParams) WithContext(ctx context.Context) *UpdateSessionConfigParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update session config params
func (o *UpdateSessionConfigParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update session config params
func (o *UpdateSessionConfigParams) WithHTTPClient(client *http.Client) *UpdateSessionConfigParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update session config params
func (o *UpdateSessionConfigParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update session config params
func (o *UpdateSessionConfigParams) WithBody(body *models.SessionConfig) *UpdateSessionConfigParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update session config params
func (o *UpdateSessionConfigParams) SetBody(body *models.SessionConfig) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateSessionConfigParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}