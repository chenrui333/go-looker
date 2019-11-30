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

// NewSearchUsersParams creates a new SearchUsersParams object
// with the default values initialized.
func NewSearchUsersParams() *SearchUsersParams {
	var ()
	return &SearchUsersParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewSearchUsersParamsWithTimeout creates a new SearchUsersParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewSearchUsersParamsWithTimeout(timeout time.Duration) *SearchUsersParams {
	var ()
	return &SearchUsersParams{

		timeout: timeout,
	}
}

// NewSearchUsersParamsWithContext creates a new SearchUsersParams object
// with the default values initialized, and the ability to set a context for a request
func NewSearchUsersParamsWithContext(ctx context.Context) *SearchUsersParams {
	var ()
	return &SearchUsersParams{

		Context: ctx,
	}
}

// NewSearchUsersParamsWithHTTPClient creates a new SearchUsersParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewSearchUsersParamsWithHTTPClient(client *http.Client) *SearchUsersParams {
	var ()
	return &SearchUsersParams{
		HTTPClient: client,
	}
}

/*SearchUsersParams contains all the parameters to send to the API endpoint
for the search users operation typically these are written to a http.Request
*/
type SearchUsersParams struct {

	/*ContentMetadataID
	  Search for users who have access to this content_metadata item

	*/
	ContentMetadataID *int64
	/*Email
	  Search for the user with this email address

	*/
	Email *string
	/*Fields
	  Include only these fields in the response

	*/
	Fields *string
	/*FilterOr
	  Combine given search criteria in a boolean OR expression

	*/
	FilterOr *bool
	/*FirstName
	  Match First name.

	*/
	FirstName *string
	/*GroupID
	  Search for users who are direct members of this group

	*/
	GroupID *int64
	/*ID
	  Match User Id.

	*/
	ID *int64
	/*IsDisabled
	  Search for disabled user accounts

	*/
	IsDisabled *bool
	/*LastName
	  Match Last name.

	*/
	LastName *string
	/*Page
	  Return only page N of paginated results

	*/
	Page *int64
	/*PerPage
	  Return N rows of data per page

	*/
	PerPage *int64
	/*Sorts
	  Fields to sort by.

	*/
	Sorts *string
	/*VerifiedLookerEmployee
	  Search for user accounts associated with Looker employees

	*/
	VerifiedLookerEmployee *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the search users params
func (o *SearchUsersParams) WithTimeout(timeout time.Duration) *SearchUsersParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the search users params
func (o *SearchUsersParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the search users params
func (o *SearchUsersParams) WithContext(ctx context.Context) *SearchUsersParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the search users params
func (o *SearchUsersParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the search users params
func (o *SearchUsersParams) WithHTTPClient(client *http.Client) *SearchUsersParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the search users params
func (o *SearchUsersParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContentMetadataID adds the contentMetadataID to the search users params
func (o *SearchUsersParams) WithContentMetadataID(contentMetadataID *int64) *SearchUsersParams {
	o.SetContentMetadataID(contentMetadataID)
	return o
}

// SetContentMetadataID adds the contentMetadataId to the search users params
func (o *SearchUsersParams) SetContentMetadataID(contentMetadataID *int64) {
	o.ContentMetadataID = contentMetadataID
}

// WithEmail adds the email to the search users params
func (o *SearchUsersParams) WithEmail(email *string) *SearchUsersParams {
	o.SetEmail(email)
	return o
}

// SetEmail adds the email to the search users params
func (o *SearchUsersParams) SetEmail(email *string) {
	o.Email = email
}

// WithFields adds the fields to the search users params
func (o *SearchUsersParams) WithFields(fields *string) *SearchUsersParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the search users params
func (o *SearchUsersParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithFilterOr adds the filterOr to the search users params
func (o *SearchUsersParams) WithFilterOr(filterOr *bool) *SearchUsersParams {
	o.SetFilterOr(filterOr)
	return o
}

// SetFilterOr adds the filterOr to the search users params
func (o *SearchUsersParams) SetFilterOr(filterOr *bool) {
	o.FilterOr = filterOr
}

// WithFirstName adds the firstName to the search users params
func (o *SearchUsersParams) WithFirstName(firstName *string) *SearchUsersParams {
	o.SetFirstName(firstName)
	return o
}

// SetFirstName adds the firstName to the search users params
func (o *SearchUsersParams) SetFirstName(firstName *string) {
	o.FirstName = firstName
}

// WithGroupID adds the groupID to the search users params
func (o *SearchUsersParams) WithGroupID(groupID *int64) *SearchUsersParams {
	o.SetGroupID(groupID)
	return o
}

// SetGroupID adds the groupId to the search users params
func (o *SearchUsersParams) SetGroupID(groupID *int64) {
	o.GroupID = groupID
}

// WithID adds the id to the search users params
func (o *SearchUsersParams) WithID(id *int64) *SearchUsersParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the search users params
func (o *SearchUsersParams) SetID(id *int64) {
	o.ID = id
}

// WithIsDisabled adds the isDisabled to the search users params
func (o *SearchUsersParams) WithIsDisabled(isDisabled *bool) *SearchUsersParams {
	o.SetIsDisabled(isDisabled)
	return o
}

// SetIsDisabled adds the isDisabled to the search users params
func (o *SearchUsersParams) SetIsDisabled(isDisabled *bool) {
	o.IsDisabled = isDisabled
}

// WithLastName adds the lastName to the search users params
func (o *SearchUsersParams) WithLastName(lastName *string) *SearchUsersParams {
	o.SetLastName(lastName)
	return o
}

// SetLastName adds the lastName to the search users params
func (o *SearchUsersParams) SetLastName(lastName *string) {
	o.LastName = lastName
}

// WithPage adds the page to the search users params
func (o *SearchUsersParams) WithPage(page *int64) *SearchUsersParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the search users params
func (o *SearchUsersParams) SetPage(page *int64) {
	o.Page = page
}

// WithPerPage adds the perPage to the search users params
func (o *SearchUsersParams) WithPerPage(perPage *int64) *SearchUsersParams {
	o.SetPerPage(perPage)
	return o
}

// SetPerPage adds the perPage to the search users params
func (o *SearchUsersParams) SetPerPage(perPage *int64) {
	o.PerPage = perPage
}

// WithSorts adds the sorts to the search users params
func (o *SearchUsersParams) WithSorts(sorts *string) *SearchUsersParams {
	o.SetSorts(sorts)
	return o
}

// SetSorts adds the sorts to the search users params
func (o *SearchUsersParams) SetSorts(sorts *string) {
	o.Sorts = sorts
}

// WithVerifiedLookerEmployee adds the verifiedLookerEmployee to the search users params
func (o *SearchUsersParams) WithVerifiedLookerEmployee(verifiedLookerEmployee *bool) *SearchUsersParams {
	o.SetVerifiedLookerEmployee(verifiedLookerEmployee)
	return o
}

// SetVerifiedLookerEmployee adds the verifiedLookerEmployee to the search users params
func (o *SearchUsersParams) SetVerifiedLookerEmployee(verifiedLookerEmployee *bool) {
	o.VerifiedLookerEmployee = verifiedLookerEmployee
}

// WriteToRequest writes these params to a swagger request
func (o *SearchUsersParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.ContentMetadataID != nil {

		// query param content_metadata_id
		var qrContentMetadataID int64
		if o.ContentMetadataID != nil {
			qrContentMetadataID = *o.ContentMetadataID
		}
		qContentMetadataID := swag.FormatInt64(qrContentMetadataID)
		if qContentMetadataID != "" {
			if err := r.SetQueryParam("content_metadata_id", qContentMetadataID); err != nil {
				return err
			}
		}

	}

	if o.Email != nil {

		// query param email
		var qrEmail string
		if o.Email != nil {
			qrEmail = *o.Email
		}
		qEmail := qrEmail
		if qEmail != "" {
			if err := r.SetQueryParam("email", qEmail); err != nil {
				return err
			}
		}

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

	if o.FilterOr != nil {

		// query param filter_or
		var qrFilterOr bool
		if o.FilterOr != nil {
			qrFilterOr = *o.FilterOr
		}
		qFilterOr := swag.FormatBool(qrFilterOr)
		if qFilterOr != "" {
			if err := r.SetQueryParam("filter_or", qFilterOr); err != nil {
				return err
			}
		}

	}

	if o.FirstName != nil {

		// query param first_name
		var qrFirstName string
		if o.FirstName != nil {
			qrFirstName = *o.FirstName
		}
		qFirstName := qrFirstName
		if qFirstName != "" {
			if err := r.SetQueryParam("first_name", qFirstName); err != nil {
				return err
			}
		}

	}

	if o.GroupID != nil {

		// query param group_id
		var qrGroupID int64
		if o.GroupID != nil {
			qrGroupID = *o.GroupID
		}
		qGroupID := swag.FormatInt64(qrGroupID)
		if qGroupID != "" {
			if err := r.SetQueryParam("group_id", qGroupID); err != nil {
				return err
			}
		}

	}

	if o.ID != nil {

		// query param id
		var qrID int64
		if o.ID != nil {
			qrID = *o.ID
		}
		qID := swag.FormatInt64(qrID)
		if qID != "" {
			if err := r.SetQueryParam("id", qID); err != nil {
				return err
			}
		}

	}

	if o.IsDisabled != nil {

		// query param is_disabled
		var qrIsDisabled bool
		if o.IsDisabled != nil {
			qrIsDisabled = *o.IsDisabled
		}
		qIsDisabled := swag.FormatBool(qrIsDisabled)
		if qIsDisabled != "" {
			if err := r.SetQueryParam("is_disabled", qIsDisabled); err != nil {
				return err
			}
		}

	}

	if o.LastName != nil {

		// query param last_name
		var qrLastName string
		if o.LastName != nil {
			qrLastName = *o.LastName
		}
		qLastName := qrLastName
		if qLastName != "" {
			if err := r.SetQueryParam("last_name", qLastName); err != nil {
				return err
			}
		}

	}

	if o.Page != nil {

		// query param page
		var qrPage int64
		if o.Page != nil {
			qrPage = *o.Page
		}
		qPage := swag.FormatInt64(qrPage)
		if qPage != "" {
			if err := r.SetQueryParam("page", qPage); err != nil {
				return err
			}
		}

	}

	if o.PerPage != nil {

		// query param per_page
		var qrPerPage int64
		if o.PerPage != nil {
			qrPerPage = *o.PerPage
		}
		qPerPage := swag.FormatInt64(qrPerPage)
		if qPerPage != "" {
			if err := r.SetQueryParam("per_page", qPerPage); err != nil {
				return err
			}
		}

	}

	if o.Sorts != nil {

		// query param sorts
		var qrSorts string
		if o.Sorts != nil {
			qrSorts = *o.Sorts
		}
		qSorts := qrSorts
		if qSorts != "" {
			if err := r.SetQueryParam("sorts", qSorts); err != nil {
				return err
			}
		}

	}

	if o.VerifiedLookerEmployee != nil {

		// query param verified_looker_employee
		var qrVerifiedLookerEmployee bool
		if o.VerifiedLookerEmployee != nil {
			qrVerifiedLookerEmployee = *o.VerifiedLookerEmployee
		}
		qVerifiedLookerEmployee := swag.FormatBool(qrVerifiedLookerEmployee)
		if qVerifiedLookerEmployee != "" {
			if err := r.SetQueryParam("verified_looker_employee", qVerifiedLookerEmployee); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
