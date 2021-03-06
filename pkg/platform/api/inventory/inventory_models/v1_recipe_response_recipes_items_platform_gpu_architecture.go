// Code generated by go-swagger; DO NOT EDIT.

package inventory_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// V1RecipeResponseRecipesItemsPlatformGpuArchitecture GPU Architecture
//
// The full GPU architecture data model
// swagger:model v1RecipeResponseRecipesItemsPlatformGpuArchitecture
type V1RecipeResponseRecipesItemsPlatformGpuArchitecture struct {
	V1RecipeResponseRecipesItemsPlatformGpuArchitectureAllOf0

	V1RecipeResponseRecipesItemsPlatformGpuArchitectureAllOf1

	V1RecipeResponseRecipesItemsPlatformGpuArchitectureAllOf2
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *V1RecipeResponseRecipesItemsPlatformGpuArchitecture) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 V1RecipeResponseRecipesItemsPlatformGpuArchitectureAllOf0
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.V1RecipeResponseRecipesItemsPlatformGpuArchitectureAllOf0 = aO0

	// AO1
	var aO1 V1RecipeResponseRecipesItemsPlatformGpuArchitectureAllOf1
	if err := swag.ReadJSON(raw, &aO1); err != nil {
		return err
	}
	m.V1RecipeResponseRecipesItemsPlatformGpuArchitectureAllOf1 = aO1

	// AO2
	var aO2 V1RecipeResponseRecipesItemsPlatformGpuArchitectureAllOf2
	if err := swag.ReadJSON(raw, &aO2); err != nil {
		return err
	}
	m.V1RecipeResponseRecipesItemsPlatformGpuArchitectureAllOf2 = aO2

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m V1RecipeResponseRecipesItemsPlatformGpuArchitecture) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 3)

	aO0, err := swag.WriteJSON(m.V1RecipeResponseRecipesItemsPlatformGpuArchitectureAllOf0)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	aO1, err := swag.WriteJSON(m.V1RecipeResponseRecipesItemsPlatformGpuArchitectureAllOf1)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO1)

	aO2, err := swag.WriteJSON(m.V1RecipeResponseRecipesItemsPlatformGpuArchitectureAllOf2)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO2)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this v1 recipe response recipes items platform gpu architecture
func (m *V1RecipeResponseRecipesItemsPlatformGpuArchitecture) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with V1RecipeResponseRecipesItemsPlatformGpuArchitectureAllOf0
	if err := m.V1RecipeResponseRecipesItemsPlatformGpuArchitectureAllOf0.Validate(formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with V1RecipeResponseRecipesItemsPlatformGpuArchitectureAllOf1
	if err := m.V1RecipeResponseRecipesItemsPlatformGpuArchitectureAllOf1.Validate(formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with V1RecipeResponseRecipesItemsPlatformGpuArchitectureAllOf2
	if err := m.V1RecipeResponseRecipesItemsPlatformGpuArchitectureAllOf2.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *V1RecipeResponseRecipesItemsPlatformGpuArchitecture) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1RecipeResponseRecipesItemsPlatformGpuArchitecture) UnmarshalBinary(b []byte) error {
	var res V1RecipeResponseRecipesItemsPlatformGpuArchitecture
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
