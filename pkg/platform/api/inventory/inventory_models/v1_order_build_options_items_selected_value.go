// Code generated by go-swagger; DO NOT EDIT.

package inventory_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// V1OrderBuildOptionsItemsSelectedValue v1 order build options items selected value
// swagger:model v1OrderBuildOptionsItemsSelectedValue
type V1OrderBuildOptionsItemsSelectedValue struct {

	// camel flag
	// Required: true
	CamelFlag *string `json:"camel_flag"`

	// is default
	// Required: true
	IsDefault *bool `json:"is_default"`

	// value
	// Required: true
	Value *string `json:"value"`
}

// Validate validates this v1 order build options items selected value
func (m *V1OrderBuildOptionsItemsSelectedValue) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCamelFlag(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIsDefault(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateValue(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1OrderBuildOptionsItemsSelectedValue) validateCamelFlag(formats strfmt.Registry) error {

	if err := validate.Required("camel_flag", "body", m.CamelFlag); err != nil {
		return err
	}

	return nil
}

func (m *V1OrderBuildOptionsItemsSelectedValue) validateIsDefault(formats strfmt.Registry) error {

	if err := validate.Required("is_default", "body", m.IsDefault); err != nil {
		return err
	}

	return nil
}

func (m *V1OrderBuildOptionsItemsSelectedValue) validateValue(formats strfmt.Registry) error {

	if err := validate.Required("value", "body", m.Value); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1OrderBuildOptionsItemsSelectedValue) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1OrderBuildOptionsItemsSelectedValue) UnmarshalBinary(b []byte) error {
	var res V1OrderBuildOptionsItemsSelectedValue
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}