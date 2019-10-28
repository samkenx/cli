// Code generated by go-swagger; DO NOT EDIT.

package inventory_operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetKernelParams creates a new GetKernelParams object
// with the default values initialized.
func NewGetKernelParams() *GetKernelParams {
	var ()
	return &GetKernelParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetKernelParamsWithTimeout creates a new GetKernelParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetKernelParamsWithTimeout(timeout time.Duration) *GetKernelParams {
	var ()
	return &GetKernelParams{

		timeout: timeout,
	}
}

// NewGetKernelParamsWithContext creates a new GetKernelParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetKernelParamsWithContext(ctx context.Context) *GetKernelParams {
	var ()
	return &GetKernelParams{

		Context: ctx,
	}
}

// NewGetKernelParamsWithHTTPClient creates a new GetKernelParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetKernelParamsWithHTTPClient(client *http.Client) *GetKernelParams {
	var ()
	return &GetKernelParams{
		HTTPClient: client,
	}
}

/*GetKernelParams contains all the parameters to send to the API endpoint
for the get kernel operation typically these are written to a http.Request
*/
type GetKernelParams struct {

	/*KernelID*/
	KernelID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get kernel params
func (o *GetKernelParams) WithTimeout(timeout time.Duration) *GetKernelParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get kernel params
func (o *GetKernelParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get kernel params
func (o *GetKernelParams) WithContext(ctx context.Context) *GetKernelParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get kernel params
func (o *GetKernelParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get kernel params
func (o *GetKernelParams) WithHTTPClient(client *http.Client) *GetKernelParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get kernel params
func (o *GetKernelParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithKernelID adds the kernelID to the get kernel params
func (o *GetKernelParams) WithKernelID(kernelID strfmt.UUID) *GetKernelParams {
	o.SetKernelID(kernelID)
	return o
}

// SetKernelID adds the kernelId to the get kernel params
func (o *GetKernelParams) SetKernelID(kernelID strfmt.UUID) {
	o.KernelID = kernelID
}

// WriteToRequest writes these params to a swagger request
func (o *GetKernelParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param kernel_id
	if err := r.SetPathParam("kernel_id", o.KernelID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
