// Code generated by go-swagger; DO NOT EDIT.

package headchef_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// V1BuildRequestRecipePlatformOperatingSystemVersionAllOf1 Operating System Version Core
//
// The properties of an operating system version needed to create a new one
// swagger:model v1BuildRequestRecipePlatformOperatingSystemVersionAllOf1
type V1BuildRequestRecipePlatformOperatingSystemVersionAllOf1 struct {
	V1BuildRequestRecipePlatformOperatingSystemVersionAllOf1AllOf0

	V1BuildRequestRecipePlatformOperatingSystemVersionAllOf1AllOf1
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *V1BuildRequestRecipePlatformOperatingSystemVersionAllOf1) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 V1BuildRequestRecipePlatformOperatingSystemVersionAllOf1AllOf0
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.V1BuildRequestRecipePlatformOperatingSystemVersionAllOf1AllOf0 = aO0

	// AO1
	var aO1 V1BuildRequestRecipePlatformOperatingSystemVersionAllOf1AllOf1
	if err := swag.ReadJSON(raw, &aO1); err != nil {
		return err
	}
	m.V1BuildRequestRecipePlatformOperatingSystemVersionAllOf1AllOf1 = aO1

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m V1BuildRequestRecipePlatformOperatingSystemVersionAllOf1) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.V1BuildRequestRecipePlatformOperatingSystemVersionAllOf1AllOf0)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	aO1, err := swag.WriteJSON(m.V1BuildRequestRecipePlatformOperatingSystemVersionAllOf1AllOf1)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this v1 build request recipe platform operating system version all of1
func (m *V1BuildRequestRecipePlatformOperatingSystemVersionAllOf1) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with V1BuildRequestRecipePlatformOperatingSystemVersionAllOf1AllOf0
	if err := m.V1BuildRequestRecipePlatformOperatingSystemVersionAllOf1AllOf0.Validate(formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with V1BuildRequestRecipePlatformOperatingSystemVersionAllOf1AllOf1
	if err := m.V1BuildRequestRecipePlatformOperatingSystemVersionAllOf1AllOf1.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *V1BuildRequestRecipePlatformOperatingSystemVersionAllOf1) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1BuildRequestRecipePlatformOperatingSystemVersionAllOf1) UnmarshalBinary(b []byte) error {
	var res V1BuildRequestRecipePlatformOperatingSystemVersionAllOf1
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
