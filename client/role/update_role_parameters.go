// Code generated by go-swagger; DO NOT EDIT.

package role

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

	models "github.com/chenrui333/go-looker/models"
)

// NewUpdateRoleParams creates a new UpdateRoleParams object
// with the default values initialized.
func NewUpdateRoleParams() *UpdateRoleParams {
	var ()
	return &UpdateRoleParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateRoleParamsWithTimeout creates a new UpdateRoleParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateRoleParamsWithTimeout(timeout time.Duration) *UpdateRoleParams {
	var ()
	return &UpdateRoleParams{

		timeout: timeout,
	}
}

// NewUpdateRoleParamsWithContext creates a new UpdateRoleParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateRoleParamsWithContext(ctx context.Context) *UpdateRoleParams {
	var ()
	return &UpdateRoleParams{

		Context: ctx,
	}
}

// NewUpdateRoleParamsWithHTTPClient creates a new UpdateRoleParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateRoleParamsWithHTTPClient(client *http.Client) *UpdateRoleParams {
	var ()
	return &UpdateRoleParams{
		HTTPClient: client,
	}
}

/*UpdateRoleParams contains all the parameters to send to the API endpoint
for the update role operation typically these are written to a http.Request
*/
type UpdateRoleParams struct {

	/*Body
	  Role

	*/
	Body *models.Role
	/*RoleID
	  id of role

	*/
	RoleID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update role params
func (o *UpdateRoleParams) WithTimeout(timeout time.Duration) *UpdateRoleParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update role params
func (o *UpdateRoleParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update role params
func (o *UpdateRoleParams) WithContext(ctx context.Context) *UpdateRoleParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update role params
func (o *UpdateRoleParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update role params
func (o *UpdateRoleParams) WithHTTPClient(client *http.Client) *UpdateRoleParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update role params
func (o *UpdateRoleParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update role params
func (o *UpdateRoleParams) WithBody(body *models.Role) *UpdateRoleParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update role params
func (o *UpdateRoleParams) SetBody(body *models.Role) {
	o.Body = body
}

// WithRoleID adds the roleID to the update role params
func (o *UpdateRoleParams) WithRoleID(roleID int64) *UpdateRoleParams {
	o.SetRoleID(roleID)
	return o
}

// SetRoleID adds the roleId to the update role params
func (o *UpdateRoleParams) SetRoleID(roleID int64) {
	o.RoleID = roleID
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateRoleParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param role_id
	if err := r.SetPathParam("role_id", swag.FormatInt64(o.RoleID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}