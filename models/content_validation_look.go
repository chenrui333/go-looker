// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// ContentValidationLook content validation look
// swagger:model ContentValidationLook
type ContentValidationLook struct {

	// Unique Id
	// Read Only: true
	ID int64 `json:"id,omitempty"`

	// Space of this look.
	// Read Only: true
	Space *ContentValidationSpace `json:"space,omitempty"`

	// Look Title
	// Read Only: true
	Title string `json:"title,omitempty"`
}

// Validate validates this content validation look
func (m *ContentValidationLook) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSpace(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ContentValidationLook) validateSpace(formats strfmt.Registry) error {

	if swag.IsZero(m.Space) { // not required
		return nil
	}

	if m.Space != nil {
		if err := m.Space.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("space")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ContentValidationLook) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ContentValidationLook) UnmarshalBinary(b []byte) error {
	var res ContentValidationLook
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}