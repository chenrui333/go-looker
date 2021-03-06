// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// LDAPGroupWrite l d a p group write
// swagger:model LDAPGroupWrite
type LDAPGroupWrite struct {

	// Name of group in LDAP
	Name string `json:"name,omitempty"`

	// Looker Role Ids
	RoleIds []int64 `json:"role_ids"`

	// Link to ldap config
	// Read Only: true
	// Format: uri
	URL strfmt.URI `json:"url,omitempty"`
}

// Validate validates this l d a p group write
func (m *LDAPGroupWrite) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateURL(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LDAPGroupWrite) validateURL(formats strfmt.Registry) error {

	if swag.IsZero(m.URL) { // not required
		return nil
	}

	if err := validate.FormatOf("url", "body", "uri", m.URL.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *LDAPGroupWrite) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LDAPGroupWrite) UnmarshalBinary(b []byte) error {
	var res LDAPGroupWrite
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
