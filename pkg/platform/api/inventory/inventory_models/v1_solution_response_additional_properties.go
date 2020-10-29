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

// V1SolutionResponseAdditionalProperties v1 solution response additional properties
// swagger:model v1SolutionResponseAdditionalProperties
type V1SolutionResponseAdditionalProperties struct {

	// The location of the recipe
	// Required: true
	// Format: uri
	Link *strfmt.URI `json:"link"`

	// recipe id
	// Required: true
	// Format: uuid
	RecipeID *strfmt.UUID `json:"recipe_id"`
}

// Validate validates this v1 solution response additional properties
func (m *V1SolutionResponseAdditionalProperties) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLink(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRecipeID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1SolutionResponseAdditionalProperties) validateLink(formats strfmt.Registry) error {

	if err := validate.Required("link", "body", m.Link); err != nil {
		return err
	}

	if err := validate.FormatOf("link", "body", "uri", m.Link.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *V1SolutionResponseAdditionalProperties) validateRecipeID(formats strfmt.Registry) error {

	if err := validate.Required("recipe_id", "body", m.RecipeID); err != nil {
		return err
	}

	if err := validate.FormatOf("recipe_id", "body", "uuid", m.RecipeID.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1SolutionResponseAdditionalProperties) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1SolutionResponseAdditionalProperties) UnmarshalBinary(b []byte) error {
	var res V1SolutionResponseAdditionalProperties
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
