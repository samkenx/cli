// Code generated by go-swagger; DO NOT EDIT.

package mono_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// MemberEditable member editable
//
// swagger:model MemberEditable
type MemberEditable struct {

	// listed
	Listed *bool `json:"listed,omitempty"`

	// owner
	Owner *bool `json:"owner,omitempty"`
}

// Validate validates this member editable
func (m *MemberEditable) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *MemberEditable) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MemberEditable) UnmarshalBinary(b []byte) error {
	var res MemberEditable
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
