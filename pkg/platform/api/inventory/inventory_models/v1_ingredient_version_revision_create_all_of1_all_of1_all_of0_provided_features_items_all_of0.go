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

// V1IngredientVersionRevisionCreateAllOf1AllOf1AllOf0ProvidedFeaturesItemsAllOf0 v1 ingredient version revision create all of1 all of1 all of0 provided features items all of0
// swagger:model v1IngredientVersionRevisionCreateAllOf1AllOf1AllOf0ProvidedFeaturesItemsAllOf0
type V1IngredientVersionRevisionCreateAllOf1AllOf1AllOf0ProvidedFeaturesItemsAllOf0 struct {

	// feature
	// Required: true
	Feature *string `json:"feature"`

	// If this is true then it means that we assigned a version to this feature ourselves rather than getting it directly from metadata in the source ingredient.
	IsActivestateVersion *bool `json:"is_activestate_version,omitempty"`

	// Whether the provider of this feature is the default provider. There can only be one default provider per feature namespace, name, and version.
	// Required: true
	IsDefaultProvider *bool `json:"is_default_provider"`

	// namespace
	// Required: true
	Namespace *string `json:"namespace"`
}

// Validate validates this v1 ingredient version revision create all of1 all of1 all of0 provided features items all of0
func (m *V1IngredientVersionRevisionCreateAllOf1AllOf1AllOf0ProvidedFeaturesItemsAllOf0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFeature(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIsDefaultProvider(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNamespace(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1IngredientVersionRevisionCreateAllOf1AllOf1AllOf0ProvidedFeaturesItemsAllOf0) validateFeature(formats strfmt.Registry) error {

	if err := validate.Required("feature", "body", m.Feature); err != nil {
		return err
	}

	return nil
}

func (m *V1IngredientVersionRevisionCreateAllOf1AllOf1AllOf0ProvidedFeaturesItemsAllOf0) validateIsDefaultProvider(formats strfmt.Registry) error {

	if err := validate.Required("is_default_provider", "body", m.IsDefaultProvider); err != nil {
		return err
	}

	return nil
}

func (m *V1IngredientVersionRevisionCreateAllOf1AllOf1AllOf0ProvidedFeaturesItemsAllOf0) validateNamespace(formats strfmt.Registry) error {

	if err := validate.Required("namespace", "body", m.Namespace); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1IngredientVersionRevisionCreateAllOf1AllOf1AllOf0ProvidedFeaturesItemsAllOf0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1IngredientVersionRevisionCreateAllOf1AllOf1AllOf0ProvidedFeaturesItemsAllOf0) UnmarshalBinary(b []byte) error {
	var res V1IngredientVersionRevisionCreateAllOf1AllOf1AllOf0ProvidedFeaturesItemsAllOf0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
