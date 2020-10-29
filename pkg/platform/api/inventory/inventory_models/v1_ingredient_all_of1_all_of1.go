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

// V1IngredientAllOf1AllOf1 Ingredient Update
//
// The fields of an ingredient that can be updated
// swagger:model v1IngredientAllOf1AllOf1
type V1IngredientAllOf1AllOf1 struct {

	// A concise summary of what this ingredient can be used for
	// Required: true
	Description *string `json:"description"`

	// URL of the website about this ingredient (if any)
	// Format: uri
	Website strfmt.URI `json:"website,omitempty"`
}

// Validate validates this v1 ingredient all of1 all of1
func (m *V1IngredientAllOf1AllOf1) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWebsite(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1IngredientAllOf1AllOf1) validateDescription(formats strfmt.Registry) error {

	if err := validate.Required("description", "body", m.Description); err != nil {
		return err
	}

	return nil
}

func (m *V1IngredientAllOf1AllOf1) validateWebsite(formats strfmt.Registry) error {

	if swag.IsZero(m.Website) { // not required
		return nil
	}

	if err := validate.FormatOf("website", "body", "uri", m.Website.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1IngredientAllOf1AllOf1) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1IngredientAllOf1AllOf1) UnmarshalBinary(b []byte) error {
	var res V1IngredientAllOf1AllOf1
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
