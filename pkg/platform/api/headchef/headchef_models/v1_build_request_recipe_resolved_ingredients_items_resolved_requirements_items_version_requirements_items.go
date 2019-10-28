// Code generated by go-swagger; DO NOT EDIT.

package headchef_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// V1BuildRequestRecipeResolvedIngredientsItemsResolvedRequirementsItemsVersionRequirementsItems Version Requirement Sub Schema
//
// The version constraint for a feature
// swagger:model v1BuildRequestRecipeResolvedIngredientsItemsResolvedRequirementsItemsVersionRequirementsItems
type V1BuildRequestRecipeResolvedIngredientsItemsResolvedRequirementsItemsVersionRequirementsItems struct {

	// The operator used to compare the version against a given provided feature to determine if it meets this requirement
	// Required: true
	// Enum: [eq gt gte lt lte ne]
	Comparator *string `json:"comparator"`

	// The required version in its original form
	// Required: true
	// Min Length: 1
	Version *string `json:"version"`
}

// Validate validates this v1 build request recipe resolved ingredients items resolved requirements items version requirements items
func (m *V1BuildRequestRecipeResolvedIngredientsItemsResolvedRequirementsItemsVersionRequirementsItems) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateComparator(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVersion(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var v1BuildRequestRecipeResolvedIngredientsItemsResolvedRequirementsItemsVersionRequirementsItemsTypeComparatorPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["eq","gt","gte","lt","lte","ne"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		v1BuildRequestRecipeResolvedIngredientsItemsResolvedRequirementsItemsVersionRequirementsItemsTypeComparatorPropEnum = append(v1BuildRequestRecipeResolvedIngredientsItemsResolvedRequirementsItemsVersionRequirementsItemsTypeComparatorPropEnum, v)
	}
}

const (

	// V1BuildRequestRecipeResolvedIngredientsItemsResolvedRequirementsItemsVersionRequirementsItemsComparatorEq captures enum value "eq"
	V1BuildRequestRecipeResolvedIngredientsItemsResolvedRequirementsItemsVersionRequirementsItemsComparatorEq string = "eq"

	// V1BuildRequestRecipeResolvedIngredientsItemsResolvedRequirementsItemsVersionRequirementsItemsComparatorGt captures enum value "gt"
	V1BuildRequestRecipeResolvedIngredientsItemsResolvedRequirementsItemsVersionRequirementsItemsComparatorGt string = "gt"

	// V1BuildRequestRecipeResolvedIngredientsItemsResolvedRequirementsItemsVersionRequirementsItemsComparatorGte captures enum value "gte"
	V1BuildRequestRecipeResolvedIngredientsItemsResolvedRequirementsItemsVersionRequirementsItemsComparatorGte string = "gte"

	// V1BuildRequestRecipeResolvedIngredientsItemsResolvedRequirementsItemsVersionRequirementsItemsComparatorLt captures enum value "lt"
	V1BuildRequestRecipeResolvedIngredientsItemsResolvedRequirementsItemsVersionRequirementsItemsComparatorLt string = "lt"

	// V1BuildRequestRecipeResolvedIngredientsItemsResolvedRequirementsItemsVersionRequirementsItemsComparatorLte captures enum value "lte"
	V1BuildRequestRecipeResolvedIngredientsItemsResolvedRequirementsItemsVersionRequirementsItemsComparatorLte string = "lte"

	// V1BuildRequestRecipeResolvedIngredientsItemsResolvedRequirementsItemsVersionRequirementsItemsComparatorNe captures enum value "ne"
	V1BuildRequestRecipeResolvedIngredientsItemsResolvedRequirementsItemsVersionRequirementsItemsComparatorNe string = "ne"
)

// prop value enum
func (m *V1BuildRequestRecipeResolvedIngredientsItemsResolvedRequirementsItemsVersionRequirementsItems) validateComparatorEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, v1BuildRequestRecipeResolvedIngredientsItemsResolvedRequirementsItemsVersionRequirementsItemsTypeComparatorPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *V1BuildRequestRecipeResolvedIngredientsItemsResolvedRequirementsItemsVersionRequirementsItems) validateComparator(formats strfmt.Registry) error {

	if err := validate.Required("comparator", "body", m.Comparator); err != nil {
		return err
	}

	// value enum
	if err := m.validateComparatorEnum("comparator", "body", *m.Comparator); err != nil {
		return err
	}

	return nil
}

func (m *V1BuildRequestRecipeResolvedIngredientsItemsResolvedRequirementsItemsVersionRequirementsItems) validateVersion(formats strfmt.Registry) error {

	if err := validate.Required("version", "body", m.Version); err != nil {
		return err
	}

	if err := validate.MinLength("version", "body", string(*m.Version), 1); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1BuildRequestRecipeResolvedIngredientsItemsResolvedRequirementsItemsVersionRequirementsItems) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1BuildRequestRecipeResolvedIngredientsItemsResolvedRequirementsItemsVersionRequirementsItems) UnmarshalBinary(b []byte) error {
	var res V1BuildRequestRecipeResolvedIngredientsItemsResolvedRequirementsItemsVersionRequirementsItems
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
