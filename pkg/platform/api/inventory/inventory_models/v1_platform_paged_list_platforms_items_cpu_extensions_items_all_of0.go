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

// V1PlatformPagedListPlatformsItemsCPUExtensionsItemsAllOf0 v1 platform paged list platforms items Cpu extensions items all of0
// swagger:model v1PlatformPagedListPlatformsItemsCpuExtensionsItemsAllOf0
type V1PlatformPagedListPlatformsItemsCPUExtensionsItemsAllOf0 struct {

	// cpu extension id
	// Required: true
	// Format: uuid
	CPUExtensionID *strfmt.UUID `json:"cpu_extension_id"`

	// links
	// Required: true
	Links *V1PlatformPagedListPlatformsItemsCPUExtensionsItemsAllOf0Links `json:"links"`
}

// Validate validates this v1 platform paged list platforms items Cpu extensions items all of0
func (m *V1PlatformPagedListPlatformsItemsCPUExtensionsItemsAllOf0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCPUExtensionID(formats); err != nil {
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

func (m *V1PlatformPagedListPlatformsItemsCPUExtensionsItemsAllOf0) validateCPUExtensionID(formats strfmt.Registry) error {

	if err := validate.Required("cpu_extension_id", "body", m.CPUExtensionID); err != nil {
		return err
	}

	if err := validate.FormatOf("cpu_extension_id", "body", "uuid", m.CPUExtensionID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *V1PlatformPagedListPlatformsItemsCPUExtensionsItemsAllOf0) validateLinks(formats strfmt.Registry) error {

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
func (m *V1PlatformPagedListPlatformsItemsCPUExtensionsItemsAllOf0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1PlatformPagedListPlatformsItemsCPUExtensionsItemsAllOf0) UnmarshalBinary(b []byte) error {
	var res V1PlatformPagedListPlatformsItemsCPUExtensionsItemsAllOf0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
