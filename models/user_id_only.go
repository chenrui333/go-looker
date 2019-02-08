// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// UserIDOnly user Id only
// swagger:model UserIdOnly
type UserIDOnly struct {

	// Unique Id
	// Read Only: true
	ID int64 `json:"id,omitempty"`
}

// Validate validates this user Id only
func (m *UserIDOnly) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *UserIDOnly) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UserIDOnly) UnmarshalBinary(b []byte) error {
	var res UserIDOnly
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}