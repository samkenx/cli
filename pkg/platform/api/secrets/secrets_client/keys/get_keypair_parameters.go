// Code generated by go-swagger; DO NOT EDIT.

package keys

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetKeypairParams creates a new GetKeypairParams object
// with the default values initialized.
func NewGetKeypairParams() *GetKeypairParams {

	return &GetKeypairParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetKeypairParamsWithTimeout creates a new GetKeypairParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetKeypairParamsWithTimeout(timeout time.Duration) *GetKeypairParams {

	return &GetKeypairParams{

		timeout: timeout,
	}
}

// NewGetKeypairParamsWithContext creates a new GetKeypairParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetKeypairParamsWithContext(ctx context.Context) *GetKeypairParams {

	return &GetKeypairParams{

		Context: ctx,
	}
}

// NewGetKeypairParamsWithHTTPClient creates a new GetKeypairParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetKeypairParamsWithHTTPClient(client *http.Client) *GetKeypairParams {

	return &GetKeypairParams{
		HTTPClient: client,
	}
}

/*GetKeypairParams contains all the parameters to send to the API endpoint
for the get keypair operation typically these are written to a http.Request
*/
type GetKeypairParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get keypair params
func (o *GetKeypairParams) WithTimeout(timeout time.Duration) *GetKeypairParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get keypair params
func (o *GetKeypairParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get keypair params
func (o *GetKeypairParams) WithContext(ctx context.Context) *GetKeypairParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get keypair params
func (o *GetKeypairParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get keypair params
func (o *GetKeypairParams) WithHTTPClient(client *http.Client) *GetKeypairParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get keypair params
func (o *GetKeypairParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetKeypairParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}