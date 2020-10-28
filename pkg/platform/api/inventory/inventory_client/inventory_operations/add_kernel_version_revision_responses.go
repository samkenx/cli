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

// AddKernelVersionRevisionReader is a Reader for the AddKernelVersionRevision structure.
type AddKernelVersionRevisionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddKernelVersionRevisionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewAddKernelVersionRevisionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewAddKernelVersionRevisionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewAddKernelVersionRevisionDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAddKernelVersionRevisionOK creates a AddKernelVersionRevisionOK with default headers values
func NewAddKernelVersionRevisionOK() *AddKernelVersionRevisionOK {
	return &AddKernelVersionRevisionOK{}
}

/*AddKernelVersionRevisionOK handles this case with default header values.

The updated state of the kernel version
*/
type AddKernelVersionRevisionOK struct {
	Payload *inventory_models.V1KernelVersion
}

func (o *AddKernelVersionRevisionOK) Error() string {
	return fmt.Sprintf("[POST /v1/kernels/{kernel_id}/versions/{kernel_version_id}/revisions][%d] addKernelVersionRevisionOK  %+v", 200, o.Payload)
}

func (o *AddKernelVersionRevisionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(inventory_models.V1KernelVersion)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddKernelVersionRevisionBadRequest creates a AddKernelVersionRevisionBadRequest with default headers values
func NewAddKernelVersionRevisionBadRequest() *AddKernelVersionRevisionBadRequest {
	return &AddKernelVersionRevisionBadRequest{}
}

/*AddKernelVersionRevisionBadRequest handles this case with default header values.

If the kernel version revision is invalid
*/
type AddKernelVersionRevisionBadRequest struct {
	Payload *inventory_models.RestAPIValidationError
}

func (o *AddKernelVersionRevisionBadRequest) Error() string {
	return fmt.Sprintf("[POST /v1/kernels/{kernel_id}/versions/{kernel_version_id}/revisions][%d] addKernelVersionRevisionBadRequest  %+v", 400, o.Payload)
}

func (o *AddKernelVersionRevisionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(inventory_models.RestAPIValidationError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddKernelVersionRevisionDefault creates a AddKernelVersionRevisionDefault with default headers values
func NewAddKernelVersionRevisionDefault(code int) *AddKernelVersionRevisionDefault {
	return &AddKernelVersionRevisionDefault{
		_statusCode: code,
	}
}

/*AddKernelVersionRevisionDefault handles this case with default header values.

If there is an error processing the request
*/
type AddKernelVersionRevisionDefault struct {
	_statusCode int

	Payload *inventory_models.RestAPIError
}

// Code gets the status code for the add kernel version revision default response
func (o *AddKernelVersionRevisionDefault) Code() int {
	return o._statusCode
}

func (o *AddKernelVersionRevisionDefault) Error() string {
	return fmt.Sprintf("[POST /v1/kernels/{kernel_id}/versions/{kernel_version_id}/revisions][%d] addKernelVersionRevision default  %+v", o._statusCode, o.Payload)
}

func (o *AddKernelVersionRevisionDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(inventory_models.RestAPIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
