// Code generated by go-swagger; DO NOT EDIT.

package inventory_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// V1SubSchemaDependencyType Recipe
//
// The different types of dependencies supported by the platform.
//
// swagger:model v1SubSchemaDependencyType
type V1SubSchemaDependencyType string

const (

	// V1SubSchemaDependencyTypeBuild captures enum value "build"
	V1SubSchemaDependencyTypeBuild V1SubSchemaDependencyType = "build"

	// V1SubSchemaDependencyTypeRuntime captures enum value "runtime"
	V1SubSchemaDependencyTypeRuntime V1SubSchemaDependencyType = "runtime"

	// V1SubSchemaDependencyTypeTest captures enum value "test"
	V1SubSchemaDependencyTypeTest V1SubSchemaDependencyType = "test"
)

// for schema
var v1SubSchemaDependencyTypeEnum []interface{}

func init() {
	var res []V1SubSchemaDependencyType
	if err := json.Unmarshal([]byte(`["build","runtime","test"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		v1SubSchemaDependencyTypeEnum = append(v1SubSchemaDependencyTypeEnum, v)
	}
}

func (m V1SubSchemaDependencyType) validateV1SubSchemaDependencyTypeEnum(path, location string, value V1SubSchemaDependencyType) error {
	if err := validate.Enum(path, location, value, v1SubSchemaDependencyTypeEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this v1 sub schema dependency type
func (m V1SubSchemaDependencyType) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateV1SubSchemaDependencyTypeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
