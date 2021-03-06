// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// Locale locale
// swagger:model Locale
type Locale struct {

	// Code for Locale
	// Read Only: true
	Code string `json:"code,omitempty"`

	// Name of Locale in English
	// Read Only: true
	EnglishName string `json:"english_name,omitempty"`

	// Name of Locale in its own language
	// Read Only: true
	NativeName string `json:"native_name,omitempty"`
}

// Validate validates this locale
func (m *Locale) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Locale) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Locale) UnmarshalBinary(b []byte) error {
	var res Locale
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
