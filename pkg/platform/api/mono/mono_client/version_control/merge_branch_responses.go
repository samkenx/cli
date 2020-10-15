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

// MergeBranchReader is a Reader for the MergeBranch structure.
type MergeBranchReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *MergeBranchReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewMergeBranchOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewMergeBranchBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewMergeBranchForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewMergeBranchNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewMergeBranchConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewMergeBranchInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewMergeBranchOK creates a MergeBranchOK with default headers values
func NewMergeBranchOK() *MergeBranchOK {
	return &MergeBranchOK{}
}

/*MergeBranchOK handles this case with default header values.

Merge the branch with the branch it was forked from using the given strategy or preview options
*/
type MergeBranchOK struct {
	Payload *mono_models.BranchMergeStrategies
}

func (o *MergeBranchOK) Error() string {
	return fmt.Sprintf("[POST /vcs/branch/{branchID}/merge][%d] mergeBranchOK  %+v", 200, o.Payload)
}

func (o *MergeBranchOK) GetPayload() *mono_models.BranchMergeStrategies {
	return o.Payload
}

func (o *MergeBranchOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(mono_models.BranchMergeStrategies)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewMergeBranchBadRequest creates a MergeBranchBadRequest with default headers values
func NewMergeBranchBadRequest() *MergeBranchBadRequest {
	return &MergeBranchBadRequest{}
}

/*MergeBranchBadRequest handles this case with default header values.

Bad Request
*/
type MergeBranchBadRequest struct {
	Payload *mono_models.Message
}

func (o *MergeBranchBadRequest) Error() string {
	return fmt.Sprintf("[POST /vcs/branch/{branchID}/merge][%d] mergeBranchBadRequest  %+v", 400, o.Payload)
}

func (o *MergeBranchBadRequest) GetPayload() *mono_models.Message {
	return o.Payload
}

func (o *MergeBranchBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(mono_models.Message)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewMergeBranchForbidden creates a MergeBranchForbidden with default headers values
func NewMergeBranchForbidden() *MergeBranchForbidden {
	return &MergeBranchForbidden{}
}

/*MergeBranchForbidden handles this case with default header values.

Forbidden
*/
type MergeBranchForbidden struct {
	Payload *mono_models.Message
}

func (o *MergeBranchForbidden) Error() string {
	return fmt.Sprintf("[POST /vcs/branch/{branchID}/merge][%d] mergeBranchForbidden  %+v", 403, o.Payload)
}

func (o *MergeBranchForbidden) GetPayload() *mono_models.Message {
	return o.Payload
}

func (o *MergeBranchForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(mono_models.Message)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewMergeBranchNotFound creates a MergeBranchNotFound with default headers values
func NewMergeBranchNotFound() *MergeBranchNotFound {
	return &MergeBranchNotFound{}
}

/*MergeBranchNotFound handles this case with default header values.

branch was not found
*/
type MergeBranchNotFound struct {
	Payload *mono_models.Message
}

func (o *MergeBranchNotFound) Error() string {
	return fmt.Sprintf("[POST /vcs/branch/{branchID}/merge][%d] mergeBranchNotFound  %+v", 404, o.Payload)
}

func (o *MergeBranchNotFound) GetPayload() *mono_models.Message {
	return o.Payload
}

func (o *MergeBranchNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(mono_models.Message)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewMergeBranchConflict creates a MergeBranchConflict with default headers values
func NewMergeBranchConflict() *MergeBranchConflict {
	return &MergeBranchConflict{}
}

/*MergeBranchConflict handles this case with default header values.

Conflict
*/
type MergeBranchConflict struct {
	Payload *mono_models.Message
}

func (o *MergeBranchConflict) Error() string {
	return fmt.Sprintf("[POST /vcs/branch/{branchID}/merge][%d] mergeBranchConflict  %+v", 409, o.Payload)
}

func (o *MergeBranchConflict) GetPayload() *mono_models.Message {
	return o.Payload
}

func (o *MergeBranchConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(mono_models.Message)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewMergeBranchInternalServerError creates a MergeBranchInternalServerError with default headers values
func NewMergeBranchInternalServerError() *MergeBranchInternalServerError {
	return &MergeBranchInternalServerError{}
}

/*MergeBranchInternalServerError handles this case with default header values.

Server Error
*/
type MergeBranchInternalServerError struct {
	Payload *mono_models.Message
}

func (o *MergeBranchInternalServerError) Error() string {
	return fmt.Sprintf("[POST /vcs/branch/{branchID}/merge][%d] mergeBranchInternalServerError  %+v", 500, o.Payload)
}

func (o *MergeBranchInternalServerError) GetPayload() *mono_models.Message {
	return o.Payload
}

func (o *MergeBranchInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(mono_models.Message)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
