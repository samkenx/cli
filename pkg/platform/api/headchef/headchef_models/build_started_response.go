// Code generated by go-swagger; DO NOT EDIT.

package headchef_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// BuildStartedResponse Build Started Response
//
// A response indicating that a requested build has been started but has yet to complete.
// swagger:model BuildStartedResponse
type BuildStartedResponse struct {

	// A user-facing message describing the build event.
	Message string `json:"message,omitempty"`

	// Build Request UUID Sub Schema
	//
	// A unique identifier for a build request.
	// Required: true
	// Format: uuid
	BuildRequestID *strfmt.UUID `json:"build_request_id"`

	// recipe id
	// Required: true
	// Format: uuid
	RecipeID *strfmt.UUID `json:"recipe_id"`

	// The timestamp for the message.
	// Required: true
	// Format: date-time
	Timestamp *strfmt.DateTime `json:"timestamp"`

	// Indicates the type of the contained message.
	// Required: true
	// Enum: [build_completed build_failed build_started]
	Type *string `json:"type"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *BuildStartedResponse) UnmarshalJSON(raw []byte) error {
	// BuildStarted
	var dataBuildStarted struct {
		Message string `json:"message,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataBuildStarted); err != nil {
		return err
	}

	m.Message = dataBuildStarted.Message

	// BuildStatus
	var dataBuildStatus struct {
		BuildRequestID *strfmt.UUID `json:"build_request_id"`

		RecipeID *strfmt.UUID `json:"recipe_id"`

		Timestamp *strfmt.DateTime `json:"timestamp"`

		Type *string `json:"type"`
	}
	if err := swag.ReadJSON(raw, &dataBuildStatus); err != nil {
		return err
	}

	m.BuildRequestID = dataBuildStatus.BuildRequestID

	m.RecipeID = dataBuildStatus.RecipeID

	m.Timestamp = dataBuildStatus.Timestamp

	m.Type = dataBuildStatus.Type

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m BuildStartedResponse) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	var dataBuildStarted struct {
		Message string `json:"message,omitempty"`
	}

	dataBuildStarted.Message = m.Message

	jsonDataBuildStarted, errBuildStarted := swag.WriteJSON(dataBuildStarted)
	if errBuildStarted != nil {
		return nil, errBuildStarted
	}
	_parts = append(_parts, jsonDataBuildStarted)

	var dataBuildStatus struct {
		BuildRequestID *strfmt.UUID `json:"build_request_id"`

		RecipeID *strfmt.UUID `json:"recipe_id"`

		Timestamp *strfmt.DateTime `json:"timestamp"`

		Type *string `json:"type"`
	}

	dataBuildStatus.BuildRequestID = m.BuildRequestID

	dataBuildStatus.RecipeID = m.RecipeID

	dataBuildStatus.Timestamp = m.Timestamp

	dataBuildStatus.Type = m.Type

	jsonDataBuildStatus, errBuildStatus := swag.WriteJSON(dataBuildStatus)
	if errBuildStatus != nil {
		return nil, errBuildStatus
	}
	_parts = append(_parts, jsonDataBuildStatus)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this build started response
func (m *BuildStartedResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBuildRequestID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRecipeID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTimestamp(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BuildStartedResponse) validateBuildRequestID(formats strfmt.Registry) error {

	if err := validate.Required("build_request_id", "body", m.BuildRequestID); err != nil {
		return err
	}

	if err := validate.FormatOf("build_request_id", "body", "uuid", m.BuildRequestID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *BuildStartedResponse) validateRecipeID(formats strfmt.Registry) error {

	if err := validate.Required("recipe_id", "body", m.RecipeID); err != nil {
		return err
	}

	if err := validate.FormatOf("recipe_id", "body", "uuid", m.RecipeID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *BuildStartedResponse) validateTimestamp(formats strfmt.Registry) error {

	if err := validate.Required("timestamp", "body", m.Timestamp); err != nil {
		return err
	}

	if err := validate.FormatOf("timestamp", "body", "date-time", m.Timestamp.String(), formats); err != nil {
		return err
	}

	return nil
}

var buildStartedResponseTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["build_completed","build_failed","build_started"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		buildStartedResponseTypeTypePropEnum = append(buildStartedResponseTypeTypePropEnum, v)
	}
}

// property enum
func (m *BuildStartedResponse) validateTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, buildStartedResponseTypeTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *BuildStartedResponse) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", *m.Type); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *BuildStartedResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BuildStartedResponse) UnmarshalBinary(b []byte) error {
	var res BuildStartedResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
