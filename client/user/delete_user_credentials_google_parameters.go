// Code generated by go-swagger; DO NOT EDIT.

package user

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

// NewDeleteUserCredentialsGoogleParams creates a new DeleteUserCredentialsGoogleParams object
// with the default values initialized.
func NewDeleteUserCredentialsGoogleParams() *DeleteUserCredentialsGoogleParams {
	var ()
	return &DeleteUserCredentialsGoogleParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteUserCredentialsGoogleParamsWithTimeout creates a new DeleteUserCredentialsGoogleParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteUserCredentialsGoogleParamsWithTimeout(timeout time.Duration) *DeleteUserCredentialsGoogleParams {
	var ()
	return &DeleteUserCredentialsGoogleParams{

		timeout: timeout,
	}
}

// NewDeleteUserCredentialsGoogleParamsWithContext creates a new DeleteUserCredentialsGoogleParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteUserCredentialsGoogleParamsWithContext(ctx context.Context) *DeleteUserCredentialsGoogleParams {
	var ()
	return &DeleteUserCredentialsGoogleParams{

		Context: ctx,
	}
}

// NewDeleteUserCredentialsGoogleParamsWithHTTPClient creates a new DeleteUserCredentialsGoogleParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteUserCredentialsGoogleParamsWithHTTPClient(client *http.Client) *DeleteUserCredentialsGoogleParams {
	var ()
	return &DeleteUserCredentialsGoogleParams{
		HTTPClient: client,
	}
}

/*DeleteUserCredentialsGoogleParams contains all the parameters to send to the API endpoint
for the delete user credentials google operation typically these are written to a http.Request
*/
type DeleteUserCredentialsGoogleParams struct {

	/*UserID
	  id of user

	*/
	UserID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete user credentials google params
func (o *DeleteUserCredentialsGoogleParams) WithTimeout(timeout time.Duration) *DeleteUserCredentialsGoogleParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete user credentials google params
func (o *DeleteUserCredentialsGoogleParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete user credentials google params
func (o *DeleteUserCredentialsGoogleParams) WithContext(ctx context.Context) *DeleteUserCredentialsGoogleParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete user credentials google params
func (o *DeleteUserCredentialsGoogleParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete user credentials google params
func (o *DeleteUserCredentialsGoogleParams) WithHTTPClient(client *http.Client) *DeleteUserCredentialsGoogleParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete user credentials google params
func (o *DeleteUserCredentialsGoogleParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUserID adds the userID to the delete user credentials google params
func (o *DeleteUserCredentialsGoogleParams) WithUserID(userID int64) *DeleteUserCredentialsGoogleParams {
	o.SetUserID(userID)
	return o
}

// SetUserID adds the userId to the delete user credentials google params
func (o *DeleteUserCredentialsGoogleParams) SetUserID(userID int64) {
	o.UserID = userID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteUserCredentialsGoogleParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param user_id
	if err := r.SetPathParam("user_id", swag.FormatInt64(o.UserID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
