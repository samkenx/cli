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

// AddBuildFlagReader is a Reader for the AddBuildFlag structure.
type AddBuildFlagReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddBuildFlagReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewAddBuildFlagCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewAddBuildFlagBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewAddBuildFlagDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAddBuildFlagCreated creates a AddBuildFlagCreated with default headers values
func NewAddBuildFlagCreated() *AddBuildFlagCreated {
	return &AddBuildFlagCreated{}
}

/*AddBuildFlagCreated handles this case with default header values.

The added build flag
*/
type AddBuildFlagCreated struct {
	Payload *inventory_models.V1BuildFlag
}

func (o *AddBuildFlagCreated) Error() string {
	return fmt.Sprintf("[POST /v1/build-flags][%d] addBuildFlagCreated  %+v", 201, o.Payload)
}

func (o *AddBuildFlagCreated) GetPayload() *inventory_models.V1BuildFlag {
	return o.Payload
}

func (o *AddBuildFlagCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(inventory_models.V1BuildFlag)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddBuildFlagBadRequest creates a AddBuildFlagBadRequest with default headers values
func NewAddBuildFlagBadRequest() *AddBuildFlagBadRequest {
	return &AddBuildFlagBadRequest{}
}

/*AddBuildFlagBadRequest handles this case with default header values.

If the build flag is invalid
*/
type AddBuildFlagBadRequest struct {
	Payload *inventory_models.RestAPIValidationError
}

func (o *AddBuildFlagBadRequest) Error() string {
	return fmt.Sprintf("[POST /v1/build-flags][%d] addBuildFlagBadRequest  %+v", 400, o.Payload)
}

func (o *AddBuildFlagBadRequest) GetPayload() *inventory_models.RestAPIValidationError {
	return o.Payload
}

func (o *AddBuildFlagBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(inventory_models.RestAPIValidationError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddBuildFlagDefault creates a AddBuildFlagDefault with default headers values
func NewAddBuildFlagDefault(code int) *AddBuildFlagDefault {
	return &AddBuildFlagDefault{
		_statusCode: code,
	}
}

/*AddBuildFlagDefault handles this case with default header values.

If there is an error processing the build flag
*/
type AddBuildFlagDefault struct {
	_statusCode int

	Payload *inventory_models.RestAPIError
}

// Code gets the status code for the add build flag default response
func (o *AddBuildFlagDefault) Code() int {
	return o._statusCode
}

func (o *AddBuildFlagDefault) Error() string {
	return fmt.Sprintf("[POST /v1/build-flags][%d] addBuildFlag default  %+v", o._statusCode, o.Payload)
}

func (o *AddBuildFlagDefault) GetPayload() *inventory_models.RestAPIError {
	return o.Payload
}

func (o *AddBuildFlagDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(inventory_models.RestAPIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
