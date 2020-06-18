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

// GetKernelGpuArchitecturesReader is a Reader for the GetKernelGpuArchitectures structure.
type GetKernelGpuArchitecturesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetKernelGpuArchitecturesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetKernelGpuArchitecturesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetKernelGpuArchitecturesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetKernelGpuArchitecturesOK creates a GetKernelGpuArchitecturesOK with default headers values
func NewGetKernelGpuArchitecturesOK() *GetKernelGpuArchitecturesOK {
	return &GetKernelGpuArchitecturesOK{}
}

/*GetKernelGpuArchitecturesOK handles this case with default header values.

A paginated list of GPU architectures
*/
type GetKernelGpuArchitecturesOK struct {
	Payload *inventory_models.V1GpuArchitecturePagedList
}

func (o *GetKernelGpuArchitecturesOK) Error() string {
	return fmt.Sprintf("[GET /v1/kernels/{kernel_id}/gpu-architectures][%d] getKernelGpuArchitecturesOK  %+v", 200, o.Payload)
}

func (o *GetKernelGpuArchitecturesOK) GetPayload() *inventory_models.V1GpuArchitecturePagedList {
	return o.Payload
}

func (o *GetKernelGpuArchitecturesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(inventory_models.V1GpuArchitecturePagedList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetKernelGpuArchitecturesDefault creates a GetKernelGpuArchitecturesDefault with default headers values
func NewGetKernelGpuArchitecturesDefault(code int) *GetKernelGpuArchitecturesDefault {
	return &GetKernelGpuArchitecturesDefault{
		_statusCode: code,
	}
}

/*GetKernelGpuArchitecturesDefault handles this case with default header values.

generic error response
*/
type GetKernelGpuArchitecturesDefault struct {
	_statusCode int

	Payload *inventory_models.RestAPIError
}

// Code gets the status code for the get kernel gpu architectures default response
func (o *GetKernelGpuArchitecturesDefault) Code() int {
	return o._statusCode
}

func (o *GetKernelGpuArchitecturesDefault) Error() string {
	return fmt.Sprintf("[GET /v1/kernels/{kernel_id}/gpu-architectures][%d] getKernelGpuArchitectures default  %+v", o._statusCode, o.Payload)
}

func (o *GetKernelGpuArchitecturesDefault) GetPayload() *inventory_models.RestAPIError {
	return o.Payload
}

func (o *GetKernelGpuArchitecturesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(inventory_models.RestAPIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
