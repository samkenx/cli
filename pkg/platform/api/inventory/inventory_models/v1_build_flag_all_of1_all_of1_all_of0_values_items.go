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

// V1BuildFlagAllOf1AllOf1AllOf0ValuesItems A description of this set.
// swagger:model v1BuildFlagAllOf1AllOf1AllOf0ValuesItems
type V1BuildFlagAllOf1AllOf1AllOf0ValuesItems struct {

	// At least one condition set from this list must be satisfied for this build flag value to be applied in a recipe (i.e condition sets are ORed together)
	ConditionSets []*V1BuildFlagAllOf1AllOf1AllOf0ValuesItemsConditionSetsItems `json:"condition_sets"`

	// The name for this build flag value
	// Required: true
	Name *string `json:"name"`
}

// Validate validates this v1 build flag all of1 all of1 all of0 values items
func (m *V1BuildFlagAllOf1AllOf1AllOf0ValuesItems) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateConditionSets(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1BuildFlagAllOf1AllOf1AllOf0ValuesItems) validateConditionSets(formats strfmt.Registry) error {

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

func (m *V1BuildFlagAllOf1AllOf1AllOf0ValuesItems) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1BuildFlagAllOf1AllOf1AllOf0ValuesItems) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1BuildFlagAllOf1AllOf1AllOf0ValuesItems) UnmarshalBinary(b []byte) error {
	var res V1BuildFlagAllOf1AllOf1AllOf0ValuesItems
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
