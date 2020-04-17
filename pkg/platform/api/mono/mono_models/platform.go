// Code generated by go-swagger; DO NOT EDIT.

package mono_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Platform A specific Architecture and Operating System combination
//
//
// swagger:model Platform
type Platform struct {

	// added
	// Format: date-time
	Added strfmt.DateTime `json:"added,omitempty"`

	// display name
	DisplayName *string `json:"displayName,omitempty"`

	// os name
	OsName string `json:"osName,omitempty"`

	// os version
	OsVersion *string `json:"osVersion,omitempty"`

	// platform ID
	// Format: uuid
	PlatformID strfmt.UUID `json:"platformID,omitempty"`
}

// Validate validates this platform
func (m *Platform) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAdded(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePlatformID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Platform) validateAdded(formats strfmt.Registry) error {

	if swag.IsZero(m.Added) { // not required
		return nil
	}

	if err := validate.FormatOf("added", "body", "date-time", m.Added.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Platform) validatePlatformID(formats strfmt.Registry) error {

	if swag.IsZero(m.PlatformID) { // not required
		return nil
	}

	if err := validate.FormatOf("platformID", "body", "uuid", m.PlatformID.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Platform) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Platform) UnmarshalBinary(b []byte) error {
	var res Platform
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
