// Code generated by go-swagger; DO NOT EDIT.

package status

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/ActiveState/cli/internal/api/models"
)

// ListActivitiesReader is a Reader for the ListActivities structure.
type ListActivitiesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListActivitiesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewListActivitiesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewListActivitiesOK creates a ListActivitiesOK with default headers values
func NewListActivitiesOK() *ListActivitiesOK {
	return &ListActivitiesOK{}
}

/*ListActivitiesOK handles this case with default header values.

Success
*/
type ListActivitiesOK struct {
	Payload []*models.Activity
}

func (o *ListActivitiesOK) Error() string {
	return fmt.Sprintf("[GET /activities][%d] listActivitiesOK  %+v", 200, o.Payload)
}

func (o *ListActivitiesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}