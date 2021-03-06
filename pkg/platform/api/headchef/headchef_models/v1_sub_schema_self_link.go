// Code generated by go-swagger; DO NOT EDIT.

package headchef_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// V1SubSchemaSelfLink Self Link
//
// A self link
//
// swagger:model v1SubSchemaSelfLink
type V1SubSchemaSelfLink struct {

	// The URI of this resource
	// Required: true
	// Format: uri
	Self *strfmt.URI `json:"self"`
}

// Validate validates this v1 sub schema self link
func (m *V1SubSchemaSelfLink) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1SubSchemaSelfLink) validateSelf(formats strfmt.Registry) error {

	if err := validate.Required("self", "body", m.Self); err != nil {
		return err
	}

	if err := validate.FormatOf("self", "body", "uri", m.Self.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1SubSchemaSelfLink) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1SubSchemaSelfLink) UnmarshalBinary(b []byte) error {
	var res V1SubSchemaSelfLink
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
