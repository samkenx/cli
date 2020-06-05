// Code generated by go-swagger; DO NOT EDIT.

package headchef_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// V1CPUArchitectureCoreAllOf0 v1 Cpu architecture core all of0
//
// swagger:model v1CpuArchitectureCoreAllOf0
type V1CPUArchitectureCoreAllOf0 struct {

	// bit width
	// Enum: [32 64]
	BitWidth string `json:"bit_width,omitempty"`

	// The name of the CPU architecture
	// Required: true
	Name *string `json:"name"`
}

// Validate validates this v1 Cpu architecture core all of0
func (m *V1CPUArchitectureCoreAllOf0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBitWidth(formats); err != nil {
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

var v1CpuArchitectureCoreAllOf0TypeBitWidthPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["32","64"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		v1CpuArchitectureCoreAllOf0TypeBitWidthPropEnum = append(v1CpuArchitectureCoreAllOf0TypeBitWidthPropEnum, v)
	}
}

const (

	// V1CPUArchitectureCoreAllOf0BitWidthNr32 captures enum value "32"
	V1CPUArchitectureCoreAllOf0BitWidthNr32 string = "32"

	// V1CPUArchitectureCoreAllOf0BitWidthNr64 captures enum value "64"
	V1CPUArchitectureCoreAllOf0BitWidthNr64 string = "64"
)

// prop value enum
func (m *V1CPUArchitectureCoreAllOf0) validateBitWidthEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, v1CpuArchitectureCoreAllOf0TypeBitWidthPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *V1CPUArchitectureCoreAllOf0) validateBitWidth(formats strfmt.Registry) error {

	if swag.IsZero(m.BitWidth) { // not required
		return nil
	}

	// value enum
	if err := m.validateBitWidthEnum("bit_width", "body", m.BitWidth); err != nil {
		return err
	}

	return nil
}

func (m *V1CPUArchitectureCoreAllOf0) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1CPUArchitectureCoreAllOf0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1CPUArchitectureCoreAllOf0) UnmarshalBinary(b []byte) error {
	var res V1CPUArchitectureCoreAllOf0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
