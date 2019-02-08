// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// SQLQueryCreate Sql query create
// swagger:model SqlQueryCreate
type SQLQueryCreate struct {

	// Operations the current user is able to perform on this object
	// Read Only: true
	Can map[string]bool `json:"can,omitempty"`

	// Id of Connection (this or "model_name" required)
	ConnectionID string `json:"connection_id,omitempty"`

	// Name of LookML Model (this or "connection_id" required)
	ModelName string `json:"model_name,omitempty"`

	// SQL query
	SQL string `json:"sql,omitempty"`
}

// Validate validates this Sql query create
func (m *SQLQueryCreate) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SQLQueryCreate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SQLQueryCreate) UnmarshalBinary(b []byte) error {
	var res SQLQueryCreate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}