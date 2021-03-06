// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// DashboardSettings dashboard settings
// swagger:model DashboardSettings
type DashboardSettings struct {

	// Key color
	KeyColor string `json:"key_color,omitempty"`

	// Background color for the dashboard
	PageBackgroundColor string `json:"page_background_color,omitempty"`

	// Page margin (side) width
	PageSideMargins int64 `json:"page_side_margins,omitempty"`

	// Background color for tiles
	TileBackgroundColor string `json:"tile_background_color,omitempty"`

	// Tile shadow on/off
	TileShadow bool `json:"tile_shadow,omitempty"`

	// Space between tiles
	TileSpaceBetween int64 `json:"tile_space_between,omitempty"`

	// Title alignment on dashboard tiles
	TileTitleAlignment string `json:"tile_title_alignment,omitempty"`
}

// Validate validates this dashboard settings
func (m *DashboardSettings) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DashboardSettings) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DashboardSettings) UnmarshalBinary(b []byte) error {
	var res DashboardSettings
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
