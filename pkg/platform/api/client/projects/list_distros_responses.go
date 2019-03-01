// Code generated by go-swagger; DO NOT EDIT.

package projects

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/ActiveState/cli/pkg/platform/api/models"
)

// ListDistrosReader is a Reader for the ListDistros structure.
type ListDistrosReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListDistrosReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewListDistrosOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewListDistrosNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewListDistrosInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewListDistrosOK creates a ListDistrosOK with default headers values
func NewListDistrosOK() *ListDistrosOK {
	return &ListDistrosOK{}
}

/*ListDistrosOK handles this case with default header values.

Success
*/
type ListDistrosOK struct {
	Payload []*models.Distro
}

func (o *ListDistrosOK) Error() string {
	return fmt.Sprintf("[GET /organizations/{organizationName}/projects/{projectName}/releases/{releaseID}/distros][%d] listDistrosOK  %+v", 200, o.Payload)
}

func (o *ListDistrosOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListDistrosNotFound creates a ListDistrosNotFound with default headers values
func NewListDistrosNotFound() *ListDistrosNotFound {
	return &ListDistrosNotFound{}
}

/*ListDistrosNotFound handles this case with default header values.

Not Found
*/
type ListDistrosNotFound struct {
	Payload *models.Message
}

func (o *ListDistrosNotFound) Error() string {
	return fmt.Sprintf("[GET /organizations/{organizationName}/projects/{projectName}/releases/{releaseID}/distros][%d] listDistrosNotFound  %+v", 404, o.Payload)
}

func (o *ListDistrosNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Message)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListDistrosInternalServerError creates a ListDistrosInternalServerError with default headers values
func NewListDistrosInternalServerError() *ListDistrosInternalServerError {
	return &ListDistrosInternalServerError{}
}

/*ListDistrosInternalServerError handles this case with default header values.

Server Error
*/
type ListDistrosInternalServerError struct {
	Payload *models.Message
}

func (o *ListDistrosInternalServerError) Error() string {
	return fmt.Sprintf("[GET /organizations/{organizationName}/projects/{projectName}/releases/{releaseID}/distros][%d] listDistrosInternalServerError  %+v", 500, o.Payload)
}

func (o *ListDistrosInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Message)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}