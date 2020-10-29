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

// V1RecipeResolvedIngredientsItemsIngredientVersionAllOf2AllOf0 Ingredient Version Core
//
// The fields of an ingredient version that can be updated
// swagger:model v1RecipeResolvedIngredientsItemsIngredientVersionAllOf2AllOf0
type V1RecipeResolvedIngredientsItemsIngredientVersionAllOf2AllOf0 struct {

	// The text from the license or elsewhere that declares the copyright holder(s) and year(s)
	// Required: true
	CopyrightText *string `json:"copyright_text"`

	// The URL of a webpage where the documentation for this ingredient version is hosted, if available
	// Format: uri
	DocumentationURI strfmt.URI `json:"documentation_uri,omitempty"`

	// Tells if this ingredient version consists of only a binary without any source.
	IsBinaryOnly *bool `json:"is_binary_only,omitempty"`

	// An SPDX 2.1 license expression describing the exact licensing for this ingredient version
	// Required: true
	LicenseExpression *string `json:"license_expression"`

	// The date and time this ingredient version was first released
	// Required: true
	// Format: date-time
	ReleaseTimestamp *strfmt.DateTime `json:"release_timestamp"`

	// The URL from which we initially retrieved the source for this ingredient version.
	// Format: uri
	SourceURI *strfmt.URI `json:"source_uri,omitempty"`
}

// Validate validates this v1 recipe resolved ingredients items ingredient version all of2 all of0
func (m *V1RecipeResolvedIngredientsItemsIngredientVersionAllOf2AllOf0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCopyrightText(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDocumentationURI(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLicenseExpression(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReleaseTimestamp(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSourceURI(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1RecipeResolvedIngredientsItemsIngredientVersionAllOf2AllOf0) validateCopyrightText(formats strfmt.Registry) error {

	if err := validate.Required("copyright_text", "body", m.CopyrightText); err != nil {
		return err
	}

	return nil
}

func (m *V1RecipeResolvedIngredientsItemsIngredientVersionAllOf2AllOf0) validateDocumentationURI(formats strfmt.Registry) error {

	if swag.IsZero(m.DocumentationURI) { // not required
		return nil
	}

	if err := validate.FormatOf("documentation_uri", "body", "uri", m.DocumentationURI.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *V1RecipeResolvedIngredientsItemsIngredientVersionAllOf2AllOf0) validateLicenseExpression(formats strfmt.Registry) error {

	if err := validate.Required("license_expression", "body", m.LicenseExpression); err != nil {
		return err
	}

	return nil
}

func (m *V1RecipeResolvedIngredientsItemsIngredientVersionAllOf2AllOf0) validateReleaseTimestamp(formats strfmt.Registry) error {

	if err := validate.Required("release_timestamp", "body", m.ReleaseTimestamp); err != nil {
		return err
	}

	if err := validate.FormatOf("release_timestamp", "body", "date-time", m.ReleaseTimestamp.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *V1RecipeResolvedIngredientsItemsIngredientVersionAllOf2AllOf0) validateSourceURI(formats strfmt.Registry) error {

	if swag.IsZero(m.SourceURI) { // not required
		return nil
	}

	if err := validate.FormatOf("source_uri", "body", "uri", m.SourceURI.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1RecipeResolvedIngredientsItemsIngredientVersionAllOf2AllOf0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1RecipeResolvedIngredientsItemsIngredientVersionAllOf2AllOf0) UnmarshalBinary(b []byte) error {
	var res V1RecipeResolvedIngredientsItemsIngredientVersionAllOf2AllOf0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
