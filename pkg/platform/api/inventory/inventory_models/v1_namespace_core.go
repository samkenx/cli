// Code generated by go-swagger; DO NOT EDIT.

package inventory_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// V1NamespaceCore Namespace Core
//
// The properties of a namespace needed to create a new one
// swagger:model v1NamespaceCore
type V1NamespaceCore struct {

	// The name of the language the ingredients in this namespace are for. Will be set based on namespace type.
	// Enum: [perl python tcl]
	ForLanguage string `json:"for_language,omitempty"`

	// is public
	// Required: true
	IsPublic *bool `json:"is_public"`

	// The algorithm to use for name normalization in this namespace
	// Required: true
	// Enum: [none python]
	NameNormalizationAlgorithm *string `json:"name_normalization_algorithm"`

	// namespace
	// Required: true
	Namespace *string `json:"namespace"`

	// owner platform organization id
	// Required: true
	// Format: uuid
	OwnerPlatformOrganizationID *strfmt.UUID `json:"owner_platform_organization_id"`

	// The type of the namespace. A namespace type tells what the namespace contains.
	// Required: true
	// Enum: [bundle internal language-core language-ingredient platform-component shared-ingredient]
	Type *string `json:"type"`

	// The algorithm to use for version parsing in this namespace
	// Required: true
	// Enum: [feature generic perl python semver]
	VersionParsingAlgorithm *string `json:"version_parsing_algorithm"`
}

// Validate validates this v1 namespace core
func (m *V1NamespaceCore) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateForLanguage(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIsPublic(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNameNormalizationAlgorithm(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNamespace(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOwnerPlatformOrganizationID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVersionParsingAlgorithm(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var v1NamespaceCoreTypeForLanguagePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["perl","python","tcl"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		v1NamespaceCoreTypeForLanguagePropEnum = append(v1NamespaceCoreTypeForLanguagePropEnum, v)
	}
}

const (

	// V1NamespaceCoreForLanguagePerl captures enum value "perl"
	V1NamespaceCoreForLanguagePerl string = "perl"

	// V1NamespaceCoreForLanguagePython captures enum value "python"
	V1NamespaceCoreForLanguagePython string = "python"

	// V1NamespaceCoreForLanguageTcl captures enum value "tcl"
	V1NamespaceCoreForLanguageTcl string = "tcl"
)

// prop value enum
func (m *V1NamespaceCore) validateForLanguageEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, v1NamespaceCoreTypeForLanguagePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *V1NamespaceCore) validateForLanguage(formats strfmt.Registry) error {

	if swag.IsZero(m.ForLanguage) { // not required
		return nil
	}

	// value enum
	if err := m.validateForLanguageEnum("for_language", "body", m.ForLanguage); err != nil {
		return err
	}

	return nil
}

func (m *V1NamespaceCore) validateIsPublic(formats strfmt.Registry) error {

	if err := validate.Required("is_public", "body", m.IsPublic); err != nil {
		return err
	}

	return nil
}

var v1NamespaceCoreTypeNameNormalizationAlgorithmPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["none","python"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		v1NamespaceCoreTypeNameNormalizationAlgorithmPropEnum = append(v1NamespaceCoreTypeNameNormalizationAlgorithmPropEnum, v)
	}
}

const (

	// V1NamespaceCoreNameNormalizationAlgorithmNone captures enum value "none"
	V1NamespaceCoreNameNormalizationAlgorithmNone string = "none"

	// V1NamespaceCoreNameNormalizationAlgorithmPython captures enum value "python"
	V1NamespaceCoreNameNormalizationAlgorithmPython string = "python"
)

// prop value enum
func (m *V1NamespaceCore) validateNameNormalizationAlgorithmEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, v1NamespaceCoreTypeNameNormalizationAlgorithmPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *V1NamespaceCore) validateNameNormalizationAlgorithm(formats strfmt.Registry) error {

	if err := validate.Required("name_normalization_algorithm", "body", m.NameNormalizationAlgorithm); err != nil {
		return err
	}

	// value enum
	if err := m.validateNameNormalizationAlgorithmEnum("name_normalization_algorithm", "body", *m.NameNormalizationAlgorithm); err != nil {
		return err
	}

	return nil
}

func (m *V1NamespaceCore) validateNamespace(formats strfmt.Registry) error {

	if err := validate.Required("namespace", "body", m.Namespace); err != nil {
		return err
	}

	return nil
}

func (m *V1NamespaceCore) validateOwnerPlatformOrganizationID(formats strfmt.Registry) error {

	if err := validate.Required("owner_platform_organization_id", "body", m.OwnerPlatformOrganizationID); err != nil {
		return err
	}

	if err := validate.FormatOf("owner_platform_organization_id", "body", "uuid", m.OwnerPlatformOrganizationID.String(), formats); err != nil {
		return err
	}

	return nil
}

var v1NamespaceCoreTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["bundle","internal","language-core","language-ingredient","platform-component","shared-ingredient"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		v1NamespaceCoreTypeTypePropEnum = append(v1NamespaceCoreTypeTypePropEnum, v)
	}
}

const (

	// V1NamespaceCoreTypeBundle captures enum value "bundle"
	V1NamespaceCoreTypeBundle string = "bundle"

	// V1NamespaceCoreTypeInternal captures enum value "internal"
	V1NamespaceCoreTypeInternal string = "internal"

	// V1NamespaceCoreTypeLanguageCore captures enum value "language-core"
	V1NamespaceCoreTypeLanguageCore string = "language-core"

	// V1NamespaceCoreTypeLanguageIngredient captures enum value "language-ingredient"
	V1NamespaceCoreTypeLanguageIngredient string = "language-ingredient"

	// V1NamespaceCoreTypePlatformComponent captures enum value "platform-component"
	V1NamespaceCoreTypePlatformComponent string = "platform-component"

	// V1NamespaceCoreTypeSharedIngredient captures enum value "shared-ingredient"
	V1NamespaceCoreTypeSharedIngredient string = "shared-ingredient"
)

// prop value enum
func (m *V1NamespaceCore) validateTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, v1NamespaceCoreTypeTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *V1NamespaceCore) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", *m.Type); err != nil {
		return err
	}

	return nil
}

var v1NamespaceCoreTypeVersionParsingAlgorithmPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["feature","generic","perl","python","semver"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		v1NamespaceCoreTypeVersionParsingAlgorithmPropEnum = append(v1NamespaceCoreTypeVersionParsingAlgorithmPropEnum, v)
	}
}

const (

	// V1NamespaceCoreVersionParsingAlgorithmFeature captures enum value "feature"
	V1NamespaceCoreVersionParsingAlgorithmFeature string = "feature"

	// V1NamespaceCoreVersionParsingAlgorithmGeneric captures enum value "generic"
	V1NamespaceCoreVersionParsingAlgorithmGeneric string = "generic"

	// V1NamespaceCoreVersionParsingAlgorithmPerl captures enum value "perl"
	V1NamespaceCoreVersionParsingAlgorithmPerl string = "perl"

	// V1NamespaceCoreVersionParsingAlgorithmPython captures enum value "python"
	V1NamespaceCoreVersionParsingAlgorithmPython string = "python"

	// V1NamespaceCoreVersionParsingAlgorithmSemver captures enum value "semver"
	V1NamespaceCoreVersionParsingAlgorithmSemver string = "semver"
)

// prop value enum
func (m *V1NamespaceCore) validateVersionParsingAlgorithmEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, v1NamespaceCoreTypeVersionParsingAlgorithmPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *V1NamespaceCore) validateVersionParsingAlgorithm(formats strfmt.Registry) error {

	if err := validate.Required("version_parsing_algorithm", "body", m.VersionParsingAlgorithm); err != nil {
		return err
	}

	// value enum
	if err := m.validateVersionParsingAlgorithmEnum("version_parsing_algorithm", "body", *m.VersionParsingAlgorithm); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1NamespaceCore) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1NamespaceCore) UnmarshalBinary(b []byte) error {
	var res V1NamespaceCore
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
