// Code generated by go-swagger; DO NOT EDIT.

package inventory_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// V1BuildFlag Build Flag
//
// The full build flag data model
// swagger:model v1BuildFlag
type V1BuildFlag struct {
	V1BuildFlagAllOf0

	V1BuildFlagAllOf1

	V1BuildFlagAllOf2
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *V1BuildFlag) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 V1BuildFlagAllOf0
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.V1BuildFlagAllOf0 = aO0

	// AO1
	var aO1 V1BuildFlagAllOf1
	if err := swag.ReadJSON(raw, &aO1); err != nil {
		return err
	}
	m.V1BuildFlagAllOf1 = aO1

	// AO2
	var aO2 V1BuildFlagAllOf2
	if err := swag.ReadJSON(raw, &aO2); err != nil {
		return err
	}
	m.V1BuildFlagAllOf2 = aO2

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m V1BuildFlag) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 3)

	aO0, err := swag.WriteJSON(m.V1BuildFlagAllOf0)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	aO1, err := swag.WriteJSON(m.V1BuildFlagAllOf1)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO1)

	aO2, err := swag.WriteJSON(m.V1BuildFlagAllOf2)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO2)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this v1 build flag
func (m *V1BuildFlag) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with V1BuildFlagAllOf0
	if err := m.V1BuildFlagAllOf0.Validate(formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with V1BuildFlagAllOf1
	if err := m.V1BuildFlagAllOf1.Validate(formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with V1BuildFlagAllOf2
	if err := m.V1BuildFlagAllOf2.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *V1BuildFlag) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1BuildFlag) UnmarshalBinary(b []byte) error {
	var res V1BuildFlag
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
