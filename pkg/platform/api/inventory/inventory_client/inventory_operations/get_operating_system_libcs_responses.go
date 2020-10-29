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

// GetOperatingSystemLibcsReader is a Reader for the GetOperatingSystemLibcs structure.
type GetOperatingSystemLibcsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetOperatingSystemLibcsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetOperatingSystemLibcsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewGetOperatingSystemLibcsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetOperatingSystemLibcsOK creates a GetOperatingSystemLibcsOK with default headers values
func NewGetOperatingSystemLibcsOK() *GetOperatingSystemLibcsOK {
	return &GetOperatingSystemLibcsOK{}
}

/*GetOperatingSystemLibcsOK handles this case with default header values.

A paginated list of libcs
*/
type GetOperatingSystemLibcsOK struct {
	Payload *inventory_models.V1LibcPagedList
}

func (o *GetOperatingSystemLibcsOK) Error() string {
	return fmt.Sprintf("[GET /v1/operating-systems/{operating_system_id}/libcs][%d] getOperatingSystemLibcsOK  %+v", 200, o.Payload)
}

func (o *GetOperatingSystemLibcsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(inventory_models.V1LibcPagedList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOperatingSystemLibcsDefault creates a GetOperatingSystemLibcsDefault with default headers values
func NewGetOperatingSystemLibcsDefault(code int) *GetOperatingSystemLibcsDefault {
	return &GetOperatingSystemLibcsDefault{
		_statusCode: code,
	}
}

/*GetOperatingSystemLibcsDefault handles this case with default header values.

generic error response
*/
type GetOperatingSystemLibcsDefault struct {
	_statusCode int

	Payload *inventory_models.RestAPIError
}

// Code gets the status code for the get operating system libcs default response
func (o *GetOperatingSystemLibcsDefault) Code() int {
	return o._statusCode
}

func (o *GetOperatingSystemLibcsDefault) Error() string {
	return fmt.Sprintf("[GET /v1/operating-systems/{operating_system_id}/libcs][%d] getOperatingSystemLibcs default  %+v", o._statusCode, o.Payload)
}

func (o *GetOperatingSystemLibcsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(inventory_models.RestAPIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
