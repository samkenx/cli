// Code generated by go-swagger; DO NOT EDIT.

package mono_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// SolverRequirement solver requirement
//
// swagger:model SolverRequirement
type SolverRequirement struct {

	// feature
	Feature string `json:"feature,omitempty"`

	// ingredient version id
	IngredientVersionID string `json:"ingredient_version_id,omitempty"`

	// namespace
	Namespace string `json:"namespace,omitempty"`

	// revision
	Revision int64 `json:"revision,omitempty"`

	// version requirements
	VersionRequirements Constraints `json:"version_requirements,omitempty"`
}

// Validate validates this solver requirement
func (m *SolverRequirement) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateVersionRequirements(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SolverRequirement) validateVersionRequirements(formats strfmt.Registry) error {

	if swag.IsZero(m.VersionRequirements) { // not required
		return nil
	}

	if err := m.VersionRequirements.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("version_requirements")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SolverRequirement) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SolverRequirement) UnmarshalBinary(b []byte) error {
	var res SolverRequirement
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
