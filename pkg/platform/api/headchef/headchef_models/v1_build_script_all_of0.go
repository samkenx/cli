// Code generated by go-swagger; DO NOT EDIT.

package headchef_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// V1BuildScriptAllOf0 v1 build script all of0
//
// swagger:model v1BuildScriptAllOf0
type V1BuildScriptAllOf0 struct {

	// build script id
	// Required: true
	// Format: uuid
	BuildScriptID *strfmt.UUID `json:"build_script_id"`

	// creation timestamp
	// Required: true
	// Format: date-time
	CreationTimestamp *strfmt.DateTime `json:"creation_timestamp"`

	// links
	// Required: true
	Links *V1SubSchemaSelfLink `json:"links"`
}

// Validate validates this v1 build script all of0
func (m *V1BuildScriptAllOf0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBuildScriptID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreationTimestamp(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1BuildScriptAllOf0) validateBuildScriptID(formats strfmt.Registry) error {

	if err := validate.Required("build_script_id", "body", m.BuildScriptID); err != nil {
		return err
	}

	if err := validate.FormatOf("build_script_id", "body", "uuid", m.BuildScriptID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *V1BuildScriptAllOf0) validateCreationTimestamp(formats strfmt.Registry) error {

	if err := validate.Required("creation_timestamp", "body", m.CreationTimestamp); err != nil {
		return err
	}

	if err := validate.FormatOf("creation_timestamp", "body", "date-time", m.CreationTimestamp.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *V1BuildScriptAllOf0) validateLinks(formats strfmt.Registry) error {

	if err := validate.Required("links", "body", m.Links); err != nil {
		return err
	}

	if m.Links != nil {
		if err := m.Links.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("links")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1BuildScriptAllOf0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1BuildScriptAllOf0) UnmarshalBinary(b []byte) error {
	var res V1BuildScriptAllOf0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
