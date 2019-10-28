// Code generated by go-swagger; DO NOT EDIT.

package inventory_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// V1OperatingSystemAllOf0Links Self Link
//
// A self link
// swagger:model v1OperatingSystemAllOf0Links
type V1OperatingSystemAllOf0Links struct {

	// The URI of this resource
	// Required: true
	// Format: uri
	Self *strfmt.URI `json:"self"`
}

// Validate validates this v1 operating system all of0 links
func (m *V1OperatingSystemAllOf0Links) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1OperatingSystemAllOf0Links) validateSelf(formats strfmt.Registry) error {

	if err := validate.Required("self", "body", m.Self); err != nil {
		return err
	}

	if err := validate.FormatOf("self", "body", "uri", m.Self.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1OperatingSystemAllOf0Links) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1OperatingSystemAllOf0Links) UnmarshalBinary(b []byte) error {
	var res V1OperatingSystemAllOf0Links
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
