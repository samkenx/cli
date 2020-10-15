// Code generated by go-swagger; DO NOT EDIT.

package version_control

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/ActiveState/cli/pkg/platform/api/mono/mono_models"
)

// GetOrderFromCheckpointReader is a Reader for the GetOrderFromCheckpoint structure.
type GetOrderFromCheckpointReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetOrderFromCheckpointReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetOrderFromCheckpointOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewGetOrderFromCheckpointInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetOrderFromCheckpointOK creates a GetOrderFromCheckpointOK with default headers values
func NewGetOrderFromCheckpointOK() *GetOrderFromCheckpointOK {
	return &GetOrderFromCheckpointOK{}
}

/*GetOrderFromCheckpointOK handles this case with default header values.

Generate a solver order for the provided checkpoint data
*/
type GetOrderFromCheckpointOK struct {
	Payload *mono_models.Order
}

func (o *GetOrderFromCheckpointOK) Error() string {
	return fmt.Sprintf("[POST /vcs/order][%d] getOrderFromCheckpointOK  %+v", 200, o.Payload)
}

func (o *GetOrderFromCheckpointOK) GetPayload() *mono_models.Order {
	return o.Payload
}

func (o *GetOrderFromCheckpointOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(mono_models.Order)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrderFromCheckpointInternalServerError creates a GetOrderFromCheckpointInternalServerError with default headers values
func NewGetOrderFromCheckpointInternalServerError() *GetOrderFromCheckpointInternalServerError {
	return &GetOrderFromCheckpointInternalServerError{}
}

/*GetOrderFromCheckpointInternalServerError handles this case with default header values.

Error generating order
*/
type GetOrderFromCheckpointInternalServerError struct {
	Payload *mono_models.Message
}

func (o *GetOrderFromCheckpointInternalServerError) Error() string {
	return fmt.Sprintf("[POST /vcs/order][%d] getOrderFromCheckpointInternalServerError  %+v", 500, o.Payload)
}

func (o *GetOrderFromCheckpointInternalServerError) GetPayload() *mono_models.Message {
	return o.Payload
}

func (o *GetOrderFromCheckpointInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(mono_models.Message)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
