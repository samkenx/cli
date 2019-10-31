// Code generated by go-swagger; DO NOT EDIT.

package headchef_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// BuildRequestRecipeResolvedRequirementsItemsIngredientVersion Ingredient Version Core Properties Sub Schema
//
// The core properties of an ingredient version. This is split out for sharing between multiple schemas.
// swagger:model buildRequestRecipeResolvedRequirementsItemsIngredientVersion
type BuildRequestRecipeResolvedRequirementsItemsIngredientVersion struct {

	// One or more optional build flags that can be applied to this ingredient when it is built.
	BuildFlags []*BuildRequestRecipeResolvedRequirementsItemsIngredientVersionBuildFlagsItems `json:"build_flags"`

	// description
	// Required: true
	Description *string `json:"description"`

	// ingredient id
	// Required: true
	// Format: uuid
	IngredientID *strfmt.UUID `json:"ingredient_id"`

	// ingredient version id
	// Required: true
	// Format: uuid
	IngredientVersionID *strfmt.UUID `json:"ingredient_version_id"`

	// A release is not stable if it is an alpha, beta, developer test release, etc.
	// Required: true
	IsStableRelease *bool `json:"is_stable_release"`

	// The features provided by this ingredient version.
	Provides map[string]BuildRequestRecipeResolvedRequirementsItemsIngredientVersionProvidesAdditionalProperties `json:"provides,omitempty"`

	// The release date for this version of the ingredient.
	// Required: true
	// Format: date-time
	ReleaseDate *strfmt.DateTime `json:"release_date"`

	// An internal version that starts at 1 and is incremented if any change is made to how this ingredient version is built.
	// Required: true
	// Minimum: 1
	Revision *int32 `json:"revision"`

	// A link to this ingredient's source code on the public Internet.
	// Required: true
	// Format: uri
	SourceURI *strfmt.URI `json:"source_uri"`

	// version
	// Required: true
	Version *string `json:"version"`
}

// Validate validates this build request recipe resolved requirements items ingredient version
func (m *BuildRequestRecipeResolvedRequirementsItemsIngredientVersion) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBuildFlags(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIngredientID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIngredientVersionID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIsStableRelease(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProvides(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReleaseDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRevision(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSourceURI(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVersion(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BuildRequestRecipeResolvedRequirementsItemsIngredientVersion) validateBuildFlags(formats strfmt.Registry) error {

	if swag.IsZero(m.BuildFlags) { // not required
		return nil
	}

	for i := 0; i < len(m.BuildFlags); i++ {
		if swag.IsZero(m.BuildFlags[i]) { // not required
			continue
		}

		if m.BuildFlags[i] != nil {
			if err := m.BuildFlags[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("build_flags" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *BuildRequestRecipeResolvedRequirementsItemsIngredientVersion) validateDescription(formats strfmt.Registry) error {

	if err := validate.Required("description", "body", m.Description); err != nil {
		return err
	}

	return nil
}

func (m *BuildRequestRecipeResolvedRequirementsItemsIngredientVersion) validateIngredientID(formats strfmt.Registry) error {

	if err := validate.Required("ingredient_id", "body", m.IngredientID); err != nil {
		return err
	}

	if err := validate.FormatOf("ingredient_id", "body", "uuid", m.IngredientID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *BuildRequestRecipeResolvedRequirementsItemsIngredientVersion) validateIngredientVersionID(formats strfmt.Registry) error {

	if err := validate.Required("ingredient_version_id", "body", m.IngredientVersionID); err != nil {
		return err
	}

	if err := validate.FormatOf("ingredient_version_id", "body", "uuid", m.IngredientVersionID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *BuildRequestRecipeResolvedRequirementsItemsIngredientVersion) validateIsStableRelease(formats strfmt.Registry) error {

	if err := validate.Required("is_stable_release", "body", m.IsStableRelease); err != nil {
		return err
	}

	return nil
}

func (m *BuildRequestRecipeResolvedRequirementsItemsIngredientVersion) validateProvides(formats strfmt.Registry) error {

	if swag.IsZero(m.Provides) { // not required
		return nil
	}

	for k := range m.Provides {

		if err := validate.Required("provides"+"."+k, "body", m.Provides[k]); err != nil {
			return err
		}
		if val, ok := m.Provides[k]; ok {
			if err := val.Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}

func (m *BuildRequestRecipeResolvedRequirementsItemsIngredientVersion) validateReleaseDate(formats strfmt.Registry) error {

	if err := validate.Required("release_date", "body", m.ReleaseDate); err != nil {
		return err
	}

	if err := validate.FormatOf("release_date", "body", "date-time", m.ReleaseDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *BuildRequestRecipeResolvedRequirementsItemsIngredientVersion) validateRevision(formats strfmt.Registry) error {

	if err := validate.Required("revision", "body", m.Revision); err != nil {
		return err
	}

	if err := validate.MinimumInt("revision", "body", int64(*m.Revision), 1, false); err != nil {
		return err
	}

	return nil
}

func (m *BuildRequestRecipeResolvedRequirementsItemsIngredientVersion) validateSourceURI(formats strfmt.Registry) error {

	if err := validate.Required("source_uri", "body", m.SourceURI); err != nil {
		return err
	}

	if err := validate.FormatOf("source_uri", "body", "uri", m.SourceURI.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *BuildRequestRecipeResolvedRequirementsItemsIngredientVersion) validateVersion(formats strfmt.Registry) error {

	if err := validate.Required("version", "body", m.Version); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *BuildRequestRecipeResolvedRequirementsItemsIngredientVersion) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BuildRequestRecipeResolvedRequirementsItemsIngredientVersion) UnmarshalBinary(b []byte) error {
	var res BuildRequestRecipeResolvedRequirementsItemsIngredientVersion
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}