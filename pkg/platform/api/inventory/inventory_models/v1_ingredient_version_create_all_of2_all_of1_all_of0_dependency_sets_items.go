// Code generated by go-swagger; DO NOT EDIT.

package inventory_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// V1IngredientVersionCreateAllOf2AllOf1AllOf0DependencySetsItems v1 ingredient version create all of2 all of1 all of0 dependency sets items
// swagger:model v1IngredientVersionCreateAllOf2AllOf1AllOf0DependencySetsItems
type V1IngredientVersionCreateAllOf2AllOf1AllOf0DependencySetsItems struct {

	// dependencies
	// Required: true
	// Min Items: 1
	Dependencies []*V1IngredientVersionCreateAllOf2AllOf1AllOf0DependencySetsItemsDependenciesItems `json:"dependencies"`

	// A description of this set.
	Description string `json:"description,omitempty"`

	// Whatever text or metadata was parsed to create this set.
	OriginalRequirement string `json:"original_requirement,omitempty"`

	// Recipe
	//
	// The different types of dependencies supported by the platform.
	// Required: true
	// Enum: [build runtime test]
	Type *string `json:"type"`
}

// Validate validates this v1 ingredient version create all of2 all of1 all of0 dependency sets items
func (m *V1IngredientVersionCreateAllOf2AllOf1AllOf0DependencySetsItems) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDependencies(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1IngredientVersionCreateAllOf2AllOf1AllOf0DependencySetsItems) validateDependencies(formats strfmt.Registry) error {

	if err := validate.Required("dependencies", "body", m.Dependencies); err != nil {
		return err
	}

	iDependenciesSize := int64(len(m.Dependencies))

	if err := validate.MinItems("dependencies", "body", iDependenciesSize, 1); err != nil {
		return err
	}

	for i := 0; i < len(m.Dependencies); i++ {
		if swag.IsZero(m.Dependencies[i]) { // not required
			continue
		}

		if m.Dependencies[i] != nil {
			if err := m.Dependencies[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("dependencies" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

var v1IngredientVersionCreateAllOf2AllOf1AllOf0DependencySetsItemsTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["build","runtime","test"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		v1IngredientVersionCreateAllOf2AllOf1AllOf0DependencySetsItemsTypeTypePropEnum = append(v1IngredientVersionCreateAllOf2AllOf1AllOf0DependencySetsItemsTypeTypePropEnum, v)
	}
}

const (

	// V1IngredientVersionCreateAllOf2AllOf1AllOf0DependencySetsItemsTypeBuild captures enum value "build"
	V1IngredientVersionCreateAllOf2AllOf1AllOf0DependencySetsItemsTypeBuild string = "build"

	// V1IngredientVersionCreateAllOf2AllOf1AllOf0DependencySetsItemsTypeRuntime captures enum value "runtime"
	V1IngredientVersionCreateAllOf2AllOf1AllOf0DependencySetsItemsTypeRuntime string = "runtime"

	// V1IngredientVersionCreateAllOf2AllOf1AllOf0DependencySetsItemsTypeTest captures enum value "test"
	V1IngredientVersionCreateAllOf2AllOf1AllOf0DependencySetsItemsTypeTest string = "test"
)

// prop value enum
func (m *V1IngredientVersionCreateAllOf2AllOf1AllOf0DependencySetsItems) validateTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, v1IngredientVersionCreateAllOf2AllOf1AllOf0DependencySetsItemsTypeTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *V1IngredientVersionCreateAllOf2AllOf1AllOf0DependencySetsItems) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", *m.Type); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1IngredientVersionCreateAllOf2AllOf1AllOf0DependencySetsItems) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1IngredientVersionCreateAllOf2AllOf1AllOf0DependencySetsItems) UnmarshalBinary(b []byte) error {
	var res V1IngredientVersionCreateAllOf2AllOf1AllOf0DependencySetsItems
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
