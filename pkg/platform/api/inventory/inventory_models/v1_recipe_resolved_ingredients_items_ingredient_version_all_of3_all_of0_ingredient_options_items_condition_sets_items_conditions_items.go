// Code generated by go-swagger; DO NOT EDIT.

package inventory_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// V1RecipeResolvedIngredientsItemsIngredientVersionAllOf3AllOf0IngredientOptionsItemsConditionSetsItemsConditionsItems Condition Sub Schema
//
// A feature that must be present in a recipe for the containing entity to apply. If nothing in the recipe matches this condition, the containing entity is disable/cannot be used.
// swagger:model v1RecipeResolvedIngredientsItemsIngredientVersionAllOf3AllOf0IngredientOptionsItemsConditionSetsItemsConditionsItems
type V1RecipeResolvedIngredientsItemsIngredientVersionAllOf3AllOf0IngredientOptionsItemsConditionSetsItemsConditionsItems struct {

	// What feature must be present for the containing entity to apply
	// Required: true
	Feature *string `json:"feature"`

	// The namespace the conditional feature is contained in
	// Required: true
	Namespace *string `json:"namespace"`

	// Requirements Sub Schema
	//
	// The version constraints that an ingredient version's requirement or condition puts on a feature
	// Required: true
	// Min Length: 1
	Requirements []*V1RecipeResolvedIngredientsItemsIngredientVersionAllOf3AllOf0IngredientOptionsItemsConditionSetsItemsConditionsItemsRequirementsItems `json:"requirements"`
}

// Validate validates this v1 recipe resolved ingredients items ingredient version all of3 all of0 ingredient options items condition sets items conditions items
func (m *V1RecipeResolvedIngredientsItemsIngredientVersionAllOf3AllOf0IngredientOptionsItemsConditionSetsItemsConditionsItems) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFeature(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNamespace(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRequirements(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1RecipeResolvedIngredientsItemsIngredientVersionAllOf3AllOf0IngredientOptionsItemsConditionSetsItemsConditionsItems) validateFeature(formats strfmt.Registry) error {

	if err := validate.Required("feature", "body", m.Feature); err != nil {
		return err
	}

	return nil
}

func (m *V1RecipeResolvedIngredientsItemsIngredientVersionAllOf3AllOf0IngredientOptionsItemsConditionSetsItemsConditionsItems) validateNamespace(formats strfmt.Registry) error {

	if err := validate.Required("namespace", "body", m.Namespace); err != nil {
		return err
	}

	return nil
}

func (m *V1RecipeResolvedIngredientsItemsIngredientVersionAllOf3AllOf0IngredientOptionsItemsConditionSetsItemsConditionsItems) validateRequirements(formats strfmt.Registry) error {

	if err := validate.Required("requirements", "body", m.Requirements); err != nil {
		return err
	}

	for i := 0; i < len(m.Requirements); i++ {
		if swag.IsZero(m.Requirements[i]) { // not required
			continue
		}

		if m.Requirements[i] != nil {
			if err := m.Requirements[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("requirements" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1RecipeResolvedIngredientsItemsIngredientVersionAllOf3AllOf0IngredientOptionsItemsConditionSetsItemsConditionsItems) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1RecipeResolvedIngredientsItemsIngredientVersionAllOf3AllOf0IngredientOptionsItemsConditionSetsItemsConditionsItems) UnmarshalBinary(b []byte) error {
	var res V1RecipeResolvedIngredientsItemsIngredientVersionAllOf3AllOf0IngredientOptionsItemsConditionSetsItemsConditionsItems
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
