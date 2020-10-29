// Code generated by go-swagger; DO NOT EDIT.

package inventory_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// V1SolutionRecipeRecipeResolvedIngredientsItemsIngredientVersionAllOf3AllOf1AllOf0ProvidedFeaturesItems Provided Feature
//
// A feature that is provided by a revisioned resource
// swagger:model v1SolutionRecipeRecipeResolvedIngredientsItemsIngredientVersionAllOf3AllOf1AllOf0ProvidedFeaturesItems
type V1SolutionRecipeRecipeResolvedIngredientsItemsIngredientVersionAllOf3AllOf1AllOf0ProvidedFeaturesItems struct {
	V1SolutionRecipeRecipeResolvedIngredientsItemsIngredientVersionAllOf3AllOf1AllOf0ProvidedFeaturesItemsAllOf0

	V1SolutionRecipeRecipeResolvedIngredientsItemsIngredientVersionAllOf3AllOf1AllOf0ProvidedFeaturesItemsAllOf1
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *V1SolutionRecipeRecipeResolvedIngredientsItemsIngredientVersionAllOf3AllOf1AllOf0ProvidedFeaturesItems) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 V1SolutionRecipeRecipeResolvedIngredientsItemsIngredientVersionAllOf3AllOf1AllOf0ProvidedFeaturesItemsAllOf0
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.V1SolutionRecipeRecipeResolvedIngredientsItemsIngredientVersionAllOf3AllOf1AllOf0ProvidedFeaturesItemsAllOf0 = aO0

	// AO1
	var aO1 V1SolutionRecipeRecipeResolvedIngredientsItemsIngredientVersionAllOf3AllOf1AllOf0ProvidedFeaturesItemsAllOf1
	if err := swag.ReadJSON(raw, &aO1); err != nil {
		return err
	}
	m.V1SolutionRecipeRecipeResolvedIngredientsItemsIngredientVersionAllOf3AllOf1AllOf0ProvidedFeaturesItemsAllOf1 = aO1

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m V1SolutionRecipeRecipeResolvedIngredientsItemsIngredientVersionAllOf3AllOf1AllOf0ProvidedFeaturesItems) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.V1SolutionRecipeRecipeResolvedIngredientsItemsIngredientVersionAllOf3AllOf1AllOf0ProvidedFeaturesItemsAllOf0)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	aO1, err := swag.WriteJSON(m.V1SolutionRecipeRecipeResolvedIngredientsItemsIngredientVersionAllOf3AllOf1AllOf0ProvidedFeaturesItemsAllOf1)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this v1 solution recipe recipe resolved ingredients items ingredient version all of3 all of1 all of0 provided features items
func (m *V1SolutionRecipeRecipeResolvedIngredientsItemsIngredientVersionAllOf3AllOf1AllOf0ProvidedFeaturesItems) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with V1SolutionRecipeRecipeResolvedIngredientsItemsIngredientVersionAllOf3AllOf1AllOf0ProvidedFeaturesItemsAllOf0
	if err := m.V1SolutionRecipeRecipeResolvedIngredientsItemsIngredientVersionAllOf3AllOf1AllOf0ProvidedFeaturesItemsAllOf0.Validate(formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with V1SolutionRecipeRecipeResolvedIngredientsItemsIngredientVersionAllOf3AllOf1AllOf0ProvidedFeaturesItemsAllOf1
	if err := m.V1SolutionRecipeRecipeResolvedIngredientsItemsIngredientVersionAllOf3AllOf1AllOf0ProvidedFeaturesItemsAllOf1.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *V1SolutionRecipeRecipeResolvedIngredientsItemsIngredientVersionAllOf3AllOf1AllOf0ProvidedFeaturesItems) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1SolutionRecipeRecipeResolvedIngredientsItemsIngredientVersionAllOf3AllOf1AllOf0ProvidedFeaturesItems) UnmarshalBinary(b []byte) error {
	var res V1SolutionRecipeRecipeResolvedIngredientsItemsIngredientVersionAllOf3AllOf1AllOf0ProvidedFeaturesItems
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
