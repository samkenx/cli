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

// V1AuthorAllOf0 v1 author all of0
// swagger:model v1AuthorAllOf0
type V1AuthorAllOf0 struct {

	// The UUID for this author
	// Required: true
	// Format: uuid
	AuthorID *strfmt.UUID `json:"author_id"`

	// The timestamp of this author's creation
	// Required: true
	// Format: date-time
	CreationTimestamp *strfmt.DateTime `json:"creation_timestamp"`

	// links
	// Required: true
	Links *V1AuthorAllOf0Links `json:"links"`
}

// Validate validates this v1 author all of0
func (m *V1AuthorAllOf0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAuthorID(formats); err != nil {
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

func (m *V1AuthorAllOf0) validateAuthorID(formats strfmt.Registry) error {

	if err := validate.Required("author_id", "body", m.AuthorID); err != nil {
		return err
	}

	if err := validate.FormatOf("author_id", "body", "uuid", m.AuthorID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *V1AuthorAllOf0) validateCreationTimestamp(formats strfmt.Registry) error {

	if err := validate.Required("creation_timestamp", "body", m.CreationTimestamp); err != nil {
		return err
	}

	if err := validate.FormatOf("creation_timestamp", "body", "date-time", m.CreationTimestamp.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *V1AuthorAllOf0) validateLinks(formats strfmt.Registry) error {

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
func (m *V1AuthorAllOf0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1AuthorAllOf0) UnmarshalBinary(b []byte) error {
	var res V1AuthorAllOf0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
