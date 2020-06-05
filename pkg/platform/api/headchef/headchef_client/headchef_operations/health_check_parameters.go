// Code generated by go-swagger; DO NOT EDIT.

package headchef_operations

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

// NewHealthCheckParams creates a new HealthCheckParams object
// with the default values initialized.
func NewHealthCheckParams() *HealthCheckParams {

	return &HealthCheckParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewHealthCheckParamsWithTimeout creates a new HealthCheckParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewHealthCheckParamsWithTimeout(timeout time.Duration) *HealthCheckParams {

	return &HealthCheckParams{

		timeout: timeout,
	}
}

// NewHealthCheckParamsWithContext creates a new HealthCheckParams object
// with the default values initialized, and the ability to set a context for a request
func NewHealthCheckParamsWithContext(ctx context.Context) *HealthCheckParams {

	return &HealthCheckParams{

		Context: ctx,
	}
}

// NewHealthCheckParamsWithHTTPClient creates a new HealthCheckParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewHealthCheckParamsWithHTTPClient(client *http.Client) *HealthCheckParams {

	return &HealthCheckParams{
		HTTPClient: client,
	}
}

/*HealthCheckParams contains all the parameters to send to the API endpoint
for the health check operation typically these are written to a http.Request
*/
type HealthCheckParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the health check params
func (o *HealthCheckParams) WithTimeout(timeout time.Duration) *HealthCheckParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the health check params
func (o *HealthCheckParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the health check params
func (o *HealthCheckParams) WithContext(ctx context.Context) *HealthCheckParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the health check params
func (o *HealthCheckParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the health check params
func (o *HealthCheckParams) WithHTTPClient(client *http.Client) *HealthCheckParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the health check params
func (o *HealthCheckParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *HealthCheckParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
