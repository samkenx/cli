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

// AddPatchReader is a Reader for the AddPatch structure.
type AddPatchReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddPatchReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewAddPatchCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewAddPatchBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewAddPatchDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAddPatchCreated creates a AddPatchCreated with default headers values
func NewAddPatchCreated() *AddPatchCreated {
	return &AddPatchCreated{}
}

/*AddPatchCreated handles this case with default header values.

The added patch
*/
type AddPatchCreated struct {
	Payload *inventory_models.V1Patch
}

func (o *AddPatchCreated) Error() string {
	return fmt.Sprintf("[POST /v1/patches][%d] addPatchCreated  %+v", 201, o.Payload)
}

func (o *AddPatchCreated) GetPayload() *inventory_models.V1Patch {
	return o.Payload
}

func (o *AddPatchCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(inventory_models.V1Patch)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddPatchBadRequest creates a AddPatchBadRequest with default headers values
func NewAddPatchBadRequest() *AddPatchBadRequest {
	return &AddPatchBadRequest{}
}

/*AddPatchBadRequest handles this case with default header values.

If the patch is invalid
*/
type AddPatchBadRequest struct {
	Payload *inventory_models.RestAPIValidationError
}

func (o *AddPatchBadRequest) Error() string {
	return fmt.Sprintf("[POST /v1/patches][%d] addPatchBadRequest  %+v", 400, o.Payload)
}

func (o *AddPatchBadRequest) GetPayload() *inventory_models.RestAPIValidationError {
	return o.Payload
}

func (o *AddPatchBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(inventory_models.RestAPIValidationError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddPatchDefault creates a AddPatchDefault with default headers values
func NewAddPatchDefault(code int) *AddPatchDefault {
	return &AddPatchDefault{
		_statusCode: code,
	}
}

/*AddPatchDefault handles this case with default header values.

If there is an error processing the request
*/
type AddPatchDefault struct {
	_statusCode int

	Payload *inventory_models.RestAPIError
}

// Code gets the status code for the add patch default response
func (o *AddPatchDefault) Code() int {
	return o._statusCode
}

func (o *AddPatchDefault) Error() string {
	return fmt.Sprintf("[POST /v1/patches][%d] addPatch default  %+v", o._statusCode, o.Payload)
}

func (o *AddPatchDefault) GetPayload() *inventory_models.RestAPIError {
	return o.Payload
}

func (o *AddPatchDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(inventory_models.RestAPIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
