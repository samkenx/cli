// Code generated by go-swagger; DO NOT EDIT.

package organizations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/ActiveState/cli/internal/api/models"
)

// EditOrganizationReader is a Reader for the EditOrganization structure.
type EditOrganizationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *EditOrganizationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewEditOrganizationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewEditOrganizationBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewEditOrganizationForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewEditOrganizationNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewEditOrganizationInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewEditOrganizationOK creates a EditOrganizationOK with default headers values
func NewEditOrganizationOK() *EditOrganizationOK {
	return &EditOrganizationOK{}
}

/*EditOrganizationOK handles this case with default header values.

Organization updated
*/
type EditOrganizationOK struct {
	Payload *models.Organization
}

func (o *EditOrganizationOK) Error() string {
	return fmt.Sprintf("[POST /organizations/{organizationName}][%d] editOrganizationOK  %+v", 200, o.Payload)
}

func (o *EditOrganizationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Organization)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEditOrganizationBadRequest creates a EditOrganizationBadRequest with default headers values
func NewEditOrganizationBadRequest() *EditOrganizationBadRequest {
	return &EditOrganizationBadRequest{}
}

/*EditOrganizationBadRequest handles this case with default header values.

Bad Request
*/
type EditOrganizationBadRequest struct {
	Payload *models.Message
}

func (o *EditOrganizationBadRequest) Error() string {
	return fmt.Sprintf("[POST /organizations/{organizationName}][%d] editOrganizationBadRequest  %+v", 400, o.Payload)
}

func (o *EditOrganizationBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Message)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEditOrganizationForbidden creates a EditOrganizationForbidden with default headers values
func NewEditOrganizationForbidden() *EditOrganizationForbidden {
	return &EditOrganizationForbidden{}
}

/*EditOrganizationForbidden handles this case with default header values.

Unauthorized
*/
type EditOrganizationForbidden struct {
	Payload *models.Message
}

func (o *EditOrganizationForbidden) Error() string {
	return fmt.Sprintf("[POST /organizations/{organizationName}][%d] editOrganizationForbidden  %+v", 403, o.Payload)
}

func (o *EditOrganizationForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Message)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEditOrganizationNotFound creates a EditOrganizationNotFound with default headers values
func NewEditOrganizationNotFound() *EditOrganizationNotFound {
	return &EditOrganizationNotFound{}
}

/*EditOrganizationNotFound handles this case with default header values.

Not Found
*/
type EditOrganizationNotFound struct {
	Payload *models.Message
}

func (o *EditOrganizationNotFound) Error() string {
	return fmt.Sprintf("[POST /organizations/{organizationName}][%d] editOrganizationNotFound  %+v", 404, o.Payload)
}

func (o *EditOrganizationNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Message)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEditOrganizationInternalServerError creates a EditOrganizationInternalServerError with default headers values
func NewEditOrganizationInternalServerError() *EditOrganizationInternalServerError {
	return &EditOrganizationInternalServerError{}
}

/*EditOrganizationInternalServerError handles this case with default header values.

Server Error
*/
type EditOrganizationInternalServerError struct {
	Payload *models.Message
}

func (o *EditOrganizationInternalServerError) Error() string {
	return fmt.Sprintf("[POST /organizations/{organizationName}][%d] editOrganizationInternalServerError  %+v", 500, o.Payload)
}

func (o *EditOrganizationInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Message)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}