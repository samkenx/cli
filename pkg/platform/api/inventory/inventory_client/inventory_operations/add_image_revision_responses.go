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

// AddImageRevisionReader is a Reader for the AddImageRevision structure.
type AddImageRevisionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddImageRevisionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewAddImageRevisionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewAddImageRevisionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewAddImageRevisionDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAddImageRevisionOK creates a AddImageRevisionOK with default headers values
func NewAddImageRevisionOK() *AddImageRevisionOK {
	return &AddImageRevisionOK{}
}

/*AddImageRevisionOK handles this case with default header values.

The updated state of the image
*/
type AddImageRevisionOK struct {
	Payload *inventory_models.V1Image
}

func (o *AddImageRevisionOK) Error() string {
	return fmt.Sprintf("[POST /v1/images/{image_id}/revisions][%d] addImageRevisionOK  %+v", 200, o.Payload)
}

func (o *AddImageRevisionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(inventory_models.V1Image)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddImageRevisionBadRequest creates a AddImageRevisionBadRequest with default headers values
func NewAddImageRevisionBadRequest() *AddImageRevisionBadRequest {
	return &AddImageRevisionBadRequest{}
}

/*AddImageRevisionBadRequest handles this case with default header values.

If the image revision is invalid
*/
type AddImageRevisionBadRequest struct {
	Payload *inventory_models.RestAPIValidationError
}

func (o *AddImageRevisionBadRequest) Error() string {
	return fmt.Sprintf("[POST /v1/images/{image_id}/revisions][%d] addImageRevisionBadRequest  %+v", 400, o.Payload)
}

func (o *AddImageRevisionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(inventory_models.RestAPIValidationError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddImageRevisionDefault creates a AddImageRevisionDefault with default headers values
func NewAddImageRevisionDefault(code int) *AddImageRevisionDefault {
	return &AddImageRevisionDefault{
		_statusCode: code,
	}
}

/*AddImageRevisionDefault handles this case with default header values.

If there is an error processing the request
*/
type AddImageRevisionDefault struct {
	_statusCode int

	Payload *inventory_models.RestAPIError
}

// Code gets the status code for the add image revision default response
func (o *AddImageRevisionDefault) Code() int {
	return o._statusCode
}

func (o *AddImageRevisionDefault) Error() string {
	return fmt.Sprintf("[POST /v1/images/{image_id}/revisions][%d] addImageRevision default  %+v", o._statusCode, o.Payload)
}

func (o *AddImageRevisionDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(inventory_models.RestAPIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
