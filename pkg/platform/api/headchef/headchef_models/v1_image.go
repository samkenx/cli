// Code generated by go-swagger; DO NOT EDIT.

package headchef_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1Image Image
//
// The full image data model
//
// swagger:model v1Image
type V1Image struct {
	V1ImageAllOf0

	V1ImageCore

	V1SubSchemaRevisionedResource
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *V1Image) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 V1ImageAllOf0
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.V1ImageAllOf0 = aO0

	// AO1
	var aO1 V1ImageCore
	if err := swag.ReadJSON(raw, &aO1); err != nil {
		return err
	}
	m.V1ImageCore = aO1

	// AO2
	var aO2 V1SubSchemaRevisionedResource
	if err := swag.ReadJSON(raw, &aO2); err != nil {
		return err
	}
	m.V1SubSchemaRevisionedResource = aO2

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m V1Image) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 3)

	aO0, err := swag.WriteJSON(m.V1ImageAllOf0)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	aO1, err := swag.WriteJSON(m.V1ImageCore)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO1)

	aO2, err := swag.WriteJSON(m.V1SubSchemaRevisionedResource)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO2)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this v1 image
func (m *V1Image) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with V1ImageAllOf0
	if err := m.V1ImageAllOf0.Validate(formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with V1ImageCore
	if err := m.V1ImageCore.Validate(formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with V1SubSchemaRevisionedResource
	if err := m.V1SubSchemaRevisionedResource.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *V1Image) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1Image) UnmarshalBinary(b []byte) error {
	var res V1Image
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
