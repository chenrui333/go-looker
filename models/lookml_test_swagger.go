// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// LookmlTest lookml test
// swagger:model LookmlTest
type LookmlTest struct {

	// Operations the current user is able to perform on this object
	// Read Only: true
	Can map[string]bool `json:"can,omitempty"`

	// Name of the explore this test runs a query against
	// Read Only: true
	ExploreName string `json:"explore_name,omitempty"`

	// Name of the LookML file containing this test.
	// Read Only: true
	File string `json:"file,omitempty"`

	// Line number of this test in LookML.
	// Read Only: true
	Line int64 `json:"line,omitempty"`

	// Name of model containing this test.
	// Read Only: true
	ModelName string `json:"model_name,omitempty"`

	// Name of this test.
	// Read Only: true
	Name string `json:"name,omitempty"`

	// The url parameters that can be used to reproduce this test's query on an explore.
	// Read Only: true
	QueryURLParams string `json:"query_url_params,omitempty"`
}

// Validate validates this lookml test
func (m *LookmlTest) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LookmlTest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LookmlTest) UnmarshalBinary(b []byte) error {
	var res LookmlTest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
