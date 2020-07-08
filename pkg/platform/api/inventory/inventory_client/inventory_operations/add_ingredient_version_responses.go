// Code generated by go-swagger; DO NOT EDIT.

package inventory_operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/ActiveState/cli/pkg/platform/api/inventory/inventory_models"
)

// AddIngredientVersionReader is a Reader for the AddIngredientVersion structure.
type AddIngredientVersionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddIngredientVersionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewAddIngredientVersionCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewAddIngredientVersionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewAddIngredientVersionDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAddIngredientVersionCreated creates a AddIngredientVersionCreated with default headers values
func NewAddIngredientVersionCreated() *AddIngredientVersionCreated {
	return &AddIngredientVersionCreated{}
}

/*AddIngredientVersionCreated handles this case with default header values.

The added ingredient version
*/
type AddIngredientVersionCreated struct {
	Payload *inventory_models.V1IngredientVersion
}

func (o *AddIngredientVersionCreated) Error() string {
	return fmt.Sprintf("[POST /v1/ingredients/{ingredient_id}/versions][%d] addIngredientVersionCreated  %+v", 201, o.Payload)
}

func (o *AddIngredientVersionCreated) GetPayload() *inventory_models.V1IngredientVersion {
	return o.Payload
}

func (o *AddIngredientVersionCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(inventory_models.V1IngredientVersion)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddIngredientVersionBadRequest creates a AddIngredientVersionBadRequest with default headers values
func NewAddIngredientVersionBadRequest() *AddIngredientVersionBadRequest {
	return &AddIngredientVersionBadRequest{}
}

/*AddIngredientVersionBadRequest handles this case with default header values.

If the ingredient version is invalid
*/
type AddIngredientVersionBadRequest struct {
	Payload *inventory_models.RestAPIValidationError
}

func (o *AddIngredientVersionBadRequest) Error() string {
	return fmt.Sprintf("[POST /v1/ingredients/{ingredient_id}/versions][%d] addIngredientVersionBadRequest  %+v", 400, o.Payload)
}

func (o *AddIngredientVersionBadRequest) GetPayload() *inventory_models.RestAPIValidationError {
	return o.Payload
}

func (o *AddIngredientVersionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(inventory_models.RestAPIValidationError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddIngredientVersionDefault creates a AddIngredientVersionDefault with default headers values
func NewAddIngredientVersionDefault(code int) *AddIngredientVersionDefault {
	return &AddIngredientVersionDefault{
		_statusCode: code,
	}
}

/*AddIngredientVersionDefault handles this case with default header values.

If there is an error processing the request
*/
type AddIngredientVersionDefault struct {
	_statusCode int

	Payload *inventory_models.RestAPIError
}

// Code gets the status code for the add ingredient version default response
func (o *AddIngredientVersionDefault) Code() int {
	return o._statusCode
}

func (o *AddIngredientVersionDefault) Error() string {
	return fmt.Sprintf("[POST /v1/ingredients/{ingredient_id}/versions][%d] addIngredientVersion default  %+v", o._statusCode, o.Payload)
}

func (o *AddIngredientVersionDefault) GetPayload() *inventory_models.RestAPIError {
	return o.Payload
}

func (o *AddIngredientVersionDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(inventory_models.RestAPIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
