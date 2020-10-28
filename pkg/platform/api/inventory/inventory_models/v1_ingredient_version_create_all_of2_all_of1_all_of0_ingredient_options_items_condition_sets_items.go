// Code generated by go-swagger; DO NOT EDIT.

package inventory_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// V1IngredientVersionCreateAllOf2AllOf1AllOf0IngredientOptionsItemsConditionSetsItems All conditions and build flag conditions in a condition set must be satisfied in order for the condition set to be considered satisfied (i.e. the union of both condition lists are ANDed together). Must contain at least one condition or build flag condition.
// swagger:model v1IngredientVersionCreateAllOf2AllOf1AllOf0IngredientOptionsItemsConditionSetsItems
type V1IngredientVersionCreateAllOf2AllOf1AllOf0IngredientOptionsItemsConditionSetsItems struct {

	// This set's conditions upon build flags
	BuildFlagConditions []*V1IngredientVersionCreateAllOf2AllOf1AllOf0IngredientOptionsItemsConditionSetsItemsBuildFlagConditionsItems `json:"build_flag_conditions"`

	// This set's conditions upon features
	Conditions []*V1IngredientVersionCreateAllOf2AllOf1AllOf0IngredientOptionsItemsConditionSetsItemsConditionsItems `json:"conditions"`
}

// Validate validates this v1 ingredient version create all of2 all of1 all of0 ingredient options items condition sets items
func (m *V1IngredientVersionCreateAllOf2AllOf1AllOf0IngredientOptionsItemsConditionSetsItems) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBuildFlagConditions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateConditions(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1IngredientVersionCreateAllOf2AllOf1AllOf0IngredientOptionsItemsConditionSetsItems) validateBuildFlagConditions(formats strfmt.Registry) error {

	if swag.IsZero(m.BuildFlagConditions) { // not required
		return nil
	}

	for i := 0; i < len(m.BuildFlagConditions); i++ {
		if swag.IsZero(m.BuildFlagConditions[i]) { // not required
			continue
		}

		if m.BuildFlagConditions[i] != nil {
			if err := m.BuildFlagConditions[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("build_flag_conditions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *V1IngredientVersionCreateAllOf2AllOf1AllOf0IngredientOptionsItemsConditionSetsItems) validateConditions(formats strfmt.Registry) error {

	if swag.IsZero(m.Conditions) { // not required
		return nil
	}

	for i := 0; i < len(m.Conditions); i++ {
		if swag.IsZero(m.Conditions[i]) { // not required
			continue
		}

		if m.Conditions[i] != nil {
			if err := m.Conditions[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("conditions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1IngredientVersionCreateAllOf2AllOf1AllOf0IngredientOptionsItemsConditionSetsItems) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1IngredientVersionCreateAllOf2AllOf1AllOf0IngredientOptionsItemsConditionSetsItems) UnmarshalBinary(b []byte) error {
	var res V1IngredientVersionCreateAllOf2AllOf1AllOf0IngredientOptionsItemsConditionSetsItems
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
