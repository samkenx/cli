// Code generated by go-swagger; DO NOT EDIT.

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewAddEmailParams creates a new AddEmailParams object
// with the default values initialized.
func NewAddEmailParams() *AddEmailParams {
	var ()
	return &AddEmailParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewAddEmailParamsWithTimeout creates a new AddEmailParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAddEmailParamsWithTimeout(timeout time.Duration) *AddEmailParams {
	var ()
	return &AddEmailParams{

		timeout: timeout,
	}
}

// NewAddEmailParamsWithContext creates a new AddEmailParams object
// with the default values initialized, and the ability to set a context for a request
func NewAddEmailParamsWithContext(ctx context.Context) *AddEmailParams {
	var ()
	return &AddEmailParams{

		Context: ctx,
	}
}

// NewAddEmailParamsWithHTTPClient creates a new AddEmailParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAddEmailParamsWithHTTPClient(client *http.Client) *AddEmailParams {
	var ()
	return &AddEmailParams{
		HTTPClient: client,
	}
}

/*AddEmailParams contains all the parameters to send to the API endpoint
for the add email operation typically these are written to a http.Request
*/
type AddEmailParams struct {

	/*Email
	  Email to add

	*/
	Email AddEmailBody
	/*Username
	  username of desired User

	*/
	Username string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the add email params
func (o *AddEmailParams) WithTimeout(timeout time.Duration) *AddEmailParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the add email params
func (o *AddEmailParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the add email params
func (o *AddEmailParams) WithContext(ctx context.Context) *AddEmailParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the add email params
func (o *AddEmailParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the add email params
func (o *AddEmailParams) WithHTTPClient(client *http.Client) *AddEmailParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the add email params
func (o *AddEmailParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEmail adds the email to the add email params
func (o *AddEmailParams) WithEmail(email AddEmailBody) *AddEmailParams {
	o.SetEmail(email)
	return o
}

// SetEmail adds the email to the add email params
func (o *AddEmailParams) SetEmail(email AddEmailBody) {
	o.Email = email
}

// WithUsername adds the username to the add email params
func (o *AddEmailParams) WithUsername(username string) *AddEmailParams {
	o.SetUsername(username)
	return o
}

// SetUsername adds the username to the add email params
func (o *AddEmailParams) SetUsername(username string) {
	o.Username = username
}

// WriteToRequest writes these params to a swagger request
func (o *AddEmailParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if err := r.SetBodyParam(o.Email); err != nil {
		return err
	}

	// path param username
	if err := r.SetPathParam("username", o.Username); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
