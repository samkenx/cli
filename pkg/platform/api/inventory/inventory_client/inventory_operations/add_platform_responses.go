// Code generated by go-swagger; DO NOT EDIT.

package inventory_operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	inventory_models "github.com/ActiveState/cli/pkg/platform/api/inventory/inventory_models"
)

// AddPlatformReader is a Reader for the AddPlatform structure.
type AddPlatformReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddPlatformReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewAddPlatformCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewAddPlatformBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewAddPlatformDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAddPlatformCreated creates a AddPlatformCreated with default headers values
func NewAddPlatformCreated() *AddPlatformCreated {
	return &AddPlatformCreated{}
}

/*AddPlatformCreated handles this case with default header values.

The added platform
*/
type AddPlatformCreated struct {
	Payload *inventory_models.V1Platform
}

func (o *AddPlatformCreated) Error() string {
	return fmt.Sprintf("[POST /v1/platforms][%d] addPlatformCreated  %+v", 201, o.Payload)
}

func (o *AddPlatformCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(inventory_models.V1Platform)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddPlatformBadRequest creates a AddPlatformBadRequest with default headers values
func NewAddPlatformBadRequest() *AddPlatformBadRequest {
	return &AddPlatformBadRequest{}
}

/*AddPlatformBadRequest handles this case with default header values.

If the platform is invalid
*/
type AddPlatformBadRequest struct {
	Payload *inventory_models.RestAPIValidationError
}

func (o *AddPlatformBadRequest) Error() string {
	return fmt.Sprintf("[POST /v1/platforms][%d] addPlatformBadRequest  %+v", 400, o.Payload)
}

func (o *AddPlatformBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(inventory_models.RestAPIValidationError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddPlatformDefault creates a AddPlatformDefault with default headers values
func NewAddPlatformDefault(code int) *AddPlatformDefault {
	return &AddPlatformDefault{
		_statusCode: code,
	}
}

/*AddPlatformDefault handles this case with default header values.

If there is an error processing the request
*/
type AddPlatformDefault struct {
	_statusCode int

	Payload *inventory_models.RestAPIError
}

// Code gets the status code for the add platform default response
func (o *AddPlatformDefault) Code() int {
	return o._statusCode
}

func (o *AddPlatformDefault) Error() string {
	return fmt.Sprintf("[POST /v1/platforms][%d] addPlatform default  %+v", o._statusCode, o.Payload)
}

func (o *AddPlatformDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(inventory_models.RestAPIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
