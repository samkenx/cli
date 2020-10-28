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

// V1LibcVersionCoreAllOf1AllOf1 Revision Base
//
// Base properties of a revisioned resource which can be modified by a new revision. Does not include provided features, so for most usage v1-revision is actually the appropriate schema.
// swagger:model v1LibcVersionCoreAllOf1AllOf1
type V1LibcVersionCoreAllOf1AllOf1 struct {

	// The platform user_id for the author of the revision. This will be automatically populated for writes based on the credentials you provide to the API.
	// Format: uuid
	AuthorPlatformUserID strfmt.UUID `json:"author_platform_user_id,omitempty"`

	// A comment describing the revision.
	// Required: true
	Comment *string `json:"comment"`

	// Whether this revision should be considered 'stable'. When a new stable revision is created, it supercedes any existing stable revision and becomes the default revision of the revisioned resource going forward.
	IsStableRevision *bool `json:"is_stable_revision,omitempty"`
}

// Validate validates this v1 libc version core all of1 all of1
func (m *V1LibcVersionCoreAllOf1AllOf1) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAuthorPlatformUserID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateComment(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1LibcVersionCoreAllOf1AllOf1) validateAuthorPlatformUserID(formats strfmt.Registry) error {

	if swag.IsZero(m.AuthorPlatformUserID) { // not required
		return nil
	}

	if err := validate.FormatOf("author_platform_user_id", "body", "uuid", m.AuthorPlatformUserID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *V1LibcVersionCoreAllOf1AllOf1) validateComment(formats strfmt.Registry) error {

	if err := validate.Required("comment", "body", m.Comment); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1LibcVersionCoreAllOf1AllOf1) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1LibcVersionCoreAllOf1AllOf1) UnmarshalBinary(b []byte) error {
	var res V1LibcVersionCoreAllOf1AllOf1
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
