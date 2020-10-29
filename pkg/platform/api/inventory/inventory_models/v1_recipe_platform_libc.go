// Code generated by go-swagger; DO NOT EDIT.

package inventory_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// V1RecipePlatformLibc Libc
//
// The full libc data model
// swagger:model v1RecipePlatformLibc
type V1RecipePlatformLibc struct {
	V1RecipePlatformLibcAllOf0

	V1RecipePlatformLibcAllOf1
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *V1RecipePlatformLibc) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 V1RecipePlatformLibcAllOf0
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.V1RecipePlatformLibcAllOf0 = aO0

	// AO1
	var aO1 V1RecipePlatformLibcAllOf1
	if err := swag.ReadJSON(raw, &aO1); err != nil {
		return err
	}
	m.V1RecipePlatformLibcAllOf1 = aO1

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m V1RecipePlatformLibc) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.V1RecipePlatformLibcAllOf0)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	aO1, err := swag.WriteJSON(m.V1RecipePlatformLibcAllOf1)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this v1 recipe platform libc
func (m *V1RecipePlatformLibc) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with V1RecipePlatformLibcAllOf0
	if err := m.V1RecipePlatformLibcAllOf0.Validate(formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with V1RecipePlatformLibcAllOf1
	if err := m.V1RecipePlatformLibcAllOf1.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *V1RecipePlatformLibc) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1RecipePlatformLibc) UnmarshalBinary(b []byte) error {
	var res V1RecipePlatformLibc
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
