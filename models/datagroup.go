// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// Datagroup datagroup
// swagger:model Datagroup
type Datagroup struct {

	// Operations the current user is able to perform on this object
	// Read Only: true
	Can map[string]bool `json:"can,omitempty"`

	// UNIX timestamp at which this entry was created.
	// Read Only: true
	CreatedAt int64 `json:"created_at,omitempty"`

	// ID of the datagroup. Also used as the unique identifier.
	// Read Only: true
	ID int64 `json:"id,omitempty"`

	// Name of the model containing the datagroup. Unique when combined with name.
	// Read Only: true
	ModelName string `json:"model_name,omitempty"`

	// Name of the datagroup. Unique when combined with model_name.
	// Read Only: true
	Name string `json:"name,omitempty"`

	// UNIX timestamp before which cache entries are considered stale. Cannot be in the future.
	StaleBefore int64 `json:"stale_before,omitempty"`

	// UNIX timestamp at which this entry trigger was last checked.
	// Read Only: true
	TriggerCheckAt int64 `json:"trigger_check_at,omitempty"`

	// The message returned with the error of the last trigger check.
	// Read Only: true
	TriggerError string `json:"trigger_error,omitempty"`

	// The value of the trigger when last checked.
	// Read Only: true
	TriggerValue string `json:"trigger_value,omitempty"`

	// UNIX timestamp at which this entry became triggered. Cannot be in the future.
	TriggeredAt int64 `json:"triggered_at,omitempty"`
}

// Validate validates this datagroup
func (m *Datagroup) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Datagroup) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Datagroup) UnmarshalBinary(b []byte) error {
	var res Datagroup
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
