// Code generated by go-swagger; DO NOT EDIT.

package buildlogstream_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// LogGeneral General log message
//
// A general status message
//
// swagger:model logGeneral
type LogGeneral struct {

	// facility
	// Required: true
	// Enum: [DIE DUMP ERROR INFO LOG WARN]
	Facility *string `json:"facility"`

	// message
	// Required: true
	Message *string `json:"message"`
}

// Validate validates this log general
func (m *LogGeneral) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFacility(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMessage(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var logGeneralTypeFacilityPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["DIE","DUMP","ERROR","INFO","LOG","WARN"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		logGeneralTypeFacilityPropEnum = append(logGeneralTypeFacilityPropEnum, v)
	}
}

const (

	// LogGeneralFacilityDIE captures enum value "DIE"
	LogGeneralFacilityDIE string = "DIE"

	// LogGeneralFacilityDUMP captures enum value "DUMP"
	LogGeneralFacilityDUMP string = "DUMP"

	// LogGeneralFacilityERROR captures enum value "ERROR"
	LogGeneralFacilityERROR string = "ERROR"

	// LogGeneralFacilityINFO captures enum value "INFO"
	LogGeneralFacilityINFO string = "INFO"

	// LogGeneralFacilityLOG captures enum value "LOG"
	LogGeneralFacilityLOG string = "LOG"

	// LogGeneralFacilityWARN captures enum value "WARN"
	LogGeneralFacilityWARN string = "WARN"
)

// prop value enum
func (m *LogGeneral) validateFacilityEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, logGeneralTypeFacilityPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *LogGeneral) validateFacility(formats strfmt.Registry) error {

	if err := validate.Required("facility", "body", m.Facility); err != nil {
		return err
	}

	// value enum
	if err := m.validateFacilityEnum("facility", "body", *m.Facility); err != nil {
		return err
	}

	return nil
}

func (m *LogGeneral) validateMessage(formats strfmt.Registry) error {

	if err := validate.Required("message", "body", m.Message); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *LogGeneral) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LogGeneral) UnmarshalBinary(b []byte) error {
	var res LogGeneral
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
