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

// GetOperatingSystemReader is a Reader for the GetOperatingSystem structure.
type GetOperatingSystemReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetOperatingSystemReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetOperatingSystemOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetOperatingSystemDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetOperatingSystemOK creates a GetOperatingSystemOK with default headers values
func NewGetOperatingSystemOK() *GetOperatingSystemOK {
	return &GetOperatingSystemOK{}
}

/*GetOperatingSystemOK handles this case with default header values.

The retrieved operating system
*/
type GetOperatingSystemOK struct {
	Payload *inventory_models.V1OperatingSystem
}

func (o *GetOperatingSystemOK) Error() string {
	return fmt.Sprintf("[GET /v1/operating-systems/{operating_system_id}][%d] getOperatingSystemOK  %+v", 200, o.Payload)
}

func (o *GetOperatingSystemOK) GetPayload() *inventory_models.V1OperatingSystem {
	return o.Payload
}

func (o *GetOperatingSystemOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(inventory_models.V1OperatingSystem)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOperatingSystemDefault creates a GetOperatingSystemDefault with default headers values
func NewGetOperatingSystemDefault(code int) *GetOperatingSystemDefault {
	return &GetOperatingSystemDefault{
		_statusCode: code,
	}
}

/*GetOperatingSystemDefault handles this case with default header values.

generic error response
*/
type GetOperatingSystemDefault struct {
	_statusCode int

	Payload *inventory_models.RestAPIError
}

// Code gets the status code for the get operating system default response
func (o *GetOperatingSystemDefault) Code() int {
	return o._statusCode
}

func (o *GetOperatingSystemDefault) Error() string {
	return fmt.Sprintf("[GET /v1/operating-systems/{operating_system_id}][%d] getOperatingSystem default  %+v", o._statusCode, o.Payload)
}

func (o *GetOperatingSystemDefault) GetPayload() *inventory_models.RestAPIError {
	return o.Payload
}

func (o *GetOperatingSystemDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(inventory_models.RestAPIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
