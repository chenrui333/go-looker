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

// DashboardElement dashboard element
// swagger:model DashboardElement
type DashboardElement struct {

	// Text tile body text
	BodyText string `json:"body_text,omitempty"`

	// Text tile body text as Html
	// Read Only: true
	BodyTextAsHTML string `json:"body_text_as_html,omitempty"`

	// Operations the current user is able to perform on this object
	// Read Only: true
	Can map[string]bool `json:"can,omitempty"`

	// Id of Dashboard
	// Read Only: true
	DashboardID string `json:"dashboard_id,omitempty"`

	// Relative path of URI of LookML file to edit the dashboard element (LookML dashboard only).
	// Read Only: true
	// Format: uri
	EditURI strfmt.URI `json:"edit_uri,omitempty"`

	// Unique Id
	// Read Only: true
	ID string `json:"id,omitempty"`

	// Look
	// Read Only: true
	Look *LookWithQuery `json:"look,omitempty"`

	// Id Of Look
	LookID string `json:"look_id,omitempty"`

	// ID of merge result
	MergeResultID string `json:"merge_result_id,omitempty"`

	// Note Display
	NoteDisplay string `json:"note_display,omitempty"`

	// Note State
	NoteState string `json:"note_state,omitempty"`

	// Note Text
	NoteText string `json:"note_text,omitempty"`

	// Note Text as Html
	// Read Only: true
	NoteTextAsHTML string `json:"note_text_as_html,omitempty"`

	// Query
	// Read Only: true
	Query *Query `json:"query,omitempty"`

	// Id Of Query
	QueryID int64 `json:"query_id,omitempty"`

	// Refresh Interval
	RefreshInterval string `json:"refresh_interval,omitempty"`

	// Refresh Interval as integer
	// Read Only: true
	RefreshIntervalToI int64 `json:"refresh_interval_to_i,omitempty"`

	// Data about the result maker.
	// Read Only: true
	ResultMaker *ResultMakerWithIDVisConfigAndDynamicFields `json:"result_maker,omitempty"`

	// ID of the ResultMakerLookup entry.
	// Read Only: true
	ResultMakerID int64 `json:"result_maker_id,omitempty"`

	// Text tile subtitle text
	SubtitleText string `json:"subtitle_text,omitempty"`

	// Title of dashboard element
	Title string `json:"title,omitempty"`

	// Whether title is hidden
	TitleHidden bool `json:"title_hidden,omitempty"`

	// Text tile title
	TitleText string `json:"title_text,omitempty"`

	// Type
	Type string `json:"type,omitempty"`
}

// Validate validates this dashboard element
func (m *DashboardElement) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEditURI(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLook(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateQuery(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResultMaker(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DashboardElement) validateEditURI(formats strfmt.Registry) error {

	if swag.IsZero(m.EditURI) { // not required
		return nil
	}

	if err := validate.FormatOf("edit_uri", "body", "uri", m.EditURI.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *DashboardElement) validateLook(formats strfmt.Registry) error {

	if swag.IsZero(m.Look) { // not required
		return nil
	}

	if m.Look != nil {
		if err := m.Look.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("look")
			}
			return err
		}
	}

	return nil
}

func (m *DashboardElement) validateQuery(formats strfmt.Registry) error {

	if swag.IsZero(m.Query) { // not required
		return nil
	}

	if m.Query != nil {
		if err := m.Query.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("query")
			}
			return err
		}
	}

	return nil
}

func (m *DashboardElement) validateResultMaker(formats strfmt.Registry) error {

	if swag.IsZero(m.ResultMaker) { // not required
		return nil
	}

	if m.ResultMaker != nil {
		if err := m.ResultMaker.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("result_maker")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DashboardElement) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DashboardElement) UnmarshalBinary(b []byte) error {
	var res DashboardElement
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}