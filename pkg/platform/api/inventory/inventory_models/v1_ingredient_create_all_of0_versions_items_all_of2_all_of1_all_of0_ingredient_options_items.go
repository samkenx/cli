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

// V1IngredientCreateAllOf0VersionsItemsAllOf2AllOf1AllOf0IngredientOptionsItems Ingredient Option
//
// A string of command line arguments that should be passed to the builder used for this ingredient if this ingredient option's condition sets are satisfied
// swagger:model v1IngredientCreateAllOf0VersionsItemsAllOf2AllOf1AllOf0IngredientOptionsItems
type V1IngredientCreateAllOf0VersionsItemsAllOf2AllOf1AllOf0IngredientOptionsItems struct {

	// The command-line arguments to append to the build invocation
	// Required: true
	// Min Items: 1
	CommandLineArgs []string `json:"command_line_args"`

	// At least one condition set from this list must be satisfied for this ingredient option to be applied in a recipe (i.e condition sets are ORed together)
	ConditionSets []*V1IngredientCreateAllOf0VersionsItemsAllOf2AllOf1AllOf0IngredientOptionsItemsConditionSetsItems `json:"condition_sets"`
}

// Validate validates this v1 ingredient create all of0 versions items all of2 all of1 all of0 ingredient options items
func (m *V1IngredientCreateAllOf0VersionsItemsAllOf2AllOf1AllOf0IngredientOptionsItems) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCommandLineArgs(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateConditionSets(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1IngredientCreateAllOf0VersionsItemsAllOf2AllOf1AllOf0IngredientOptionsItems) validateCommandLineArgs(formats strfmt.Registry) error {

	if err := validate.Required("command_line_args", "body", m.CommandLineArgs); err != nil {
		return err
	}

	iCommandLineArgsSize := int64(len(m.CommandLineArgs))

	if err := validate.MinItems("command_line_args", "body", iCommandLineArgsSize, 1); err != nil {
		return err
	}

	for i := 0; i < len(m.CommandLineArgs); i++ {

		if err := validate.MinLength("command_line_args"+"."+strconv.Itoa(i), "body", string(m.CommandLineArgs[i]), 1); err != nil {
			return err
		}

	}

	return nil
}

func (m *V1IngredientCreateAllOf0VersionsItemsAllOf2AllOf1AllOf0IngredientOptionsItems) validateConditionSets(formats strfmt.Registry) error {

	if swag.IsZero(m.ConditionSets) { // not required
		return nil
	}

	for i := 0; i < len(m.ConditionSets); i++ {
		if swag.IsZero(m.ConditionSets[i]) { // not required
			continue
		}

		if m.ConditionSets[i] != nil {
			if err := m.ConditionSets[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("condition_sets" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1IngredientCreateAllOf0VersionsItemsAllOf2AllOf1AllOf0IngredientOptionsItems) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1IngredientCreateAllOf0VersionsItemsAllOf2AllOf1AllOf0IngredientOptionsItems) UnmarshalBinary(b []byte) error {
	var res V1IngredientCreateAllOf0VersionsItemsAllOf2AllOf1AllOf0IngredientOptionsItems
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
