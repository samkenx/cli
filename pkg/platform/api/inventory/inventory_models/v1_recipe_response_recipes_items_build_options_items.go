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

// V1RecipeResponseRecipesItemsBuildOptionsItems Recipe Build Option
//
// A global build option for a selected for a build in a recipe
// swagger:model v1RecipeResponseRecipesItemsBuildOptionsItems
type V1RecipeResponseRecipesItemsBuildOptionsItems struct {

	// description
	// Required: true
	Description *string `json:"description"`

	// selected value
	// Required: true
	SelectedValue *V1RecipeResponseRecipesItemsBuildOptionsItemsSelectedValue `json:"selected_value"`

	// title
	// Required: true
	Title *string `json:"title"`
}

// Validate validates this v1 recipe response recipes items build options items
func (m *V1RecipeResponseRecipesItemsBuildOptionsItems) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSelectedValue(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTitle(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1RecipeResponseRecipesItemsBuildOptionsItems) validateDescription(formats strfmt.Registry) error {

	if err := validate.Required("description", "body", m.Description); err != nil {
		return err
	}

	return nil
}

func (m *V1RecipeResponseRecipesItemsBuildOptionsItems) validateSelectedValue(formats strfmt.Registry) error {

	if err := validate.Required("selected_value", "body", m.SelectedValue); err != nil {
		return err
	}

	if m.SelectedValue != nil {
		if err := m.SelectedValue.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("selected_value")
			}
			return err
		}
	}

	return nil
}

func (m *V1RecipeResponseRecipesItemsBuildOptionsItems) validateTitle(formats strfmt.Registry) error {

	if err := validate.Required("title", "body", m.Title); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1RecipeResponseRecipesItemsBuildOptionsItems) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1RecipeResponseRecipesItemsBuildOptionsItems) UnmarshalBinary(b []byte) error {
	var res V1RecipeResponseRecipesItemsBuildOptionsItems
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
