// Code generated by go-swagger; DO NOT EDIT.

package inventory_operations

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

	"github.com/ActiveState/cli/pkg/platform/api/inventory/inventory_models"
)

// NewAddCPUArchitectureCPUExtensionParams creates a new AddCPUArchitectureCPUExtensionParams object
// with the default values initialized.
func NewAddCPUArchitectureCPUExtensionParams() *AddCPUArchitectureCPUExtensionParams {
	var ()
	return &AddCPUArchitectureCPUExtensionParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewAddCPUArchitectureCPUExtensionParamsWithTimeout creates a new AddCPUArchitectureCPUExtensionParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAddCPUArchitectureCPUExtensionParamsWithTimeout(timeout time.Duration) *AddCPUArchitectureCPUExtensionParams {
	var ()
	return &AddCPUArchitectureCPUExtensionParams{

		timeout: timeout,
	}
}

// NewAddCPUArchitectureCPUExtensionParamsWithContext creates a new AddCPUArchitectureCPUExtensionParams object
// with the default values initialized, and the ability to set a context for a request
func NewAddCPUArchitectureCPUExtensionParamsWithContext(ctx context.Context) *AddCPUArchitectureCPUExtensionParams {
	var ()
	return &AddCPUArchitectureCPUExtensionParams{

		Context: ctx,
	}
}

// NewAddCPUArchitectureCPUExtensionParamsWithHTTPClient creates a new AddCPUArchitectureCPUExtensionParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAddCPUArchitectureCPUExtensionParamsWithHTTPClient(client *http.Client) *AddCPUArchitectureCPUExtensionParams {
	var ()
	return &AddCPUArchitectureCPUExtensionParams{
		HTTPClient: client,
	}
}

/*AddCPUArchitectureCPUExtensionParams contains all the parameters to send to the API endpoint
for the add Cpu architecture Cpu extension operation typically these are written to a http.Request
*/
type AddCPUArchitectureCPUExtensionParams struct {

	/*CPUArchitectureID*/
	CPUArchitectureID strfmt.UUID
	/*CPUExtensionID*/
	CPUExtensionID *inventory_models.AddCPUArchitectureCPUExtensionParamsBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the add Cpu architecture Cpu extension params
func (o *AddCPUArchitectureCPUExtensionParams) WithTimeout(timeout time.Duration) *AddCPUArchitectureCPUExtensionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the add Cpu architecture Cpu extension params
func (o *AddCPUArchitectureCPUExtensionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the add Cpu architecture Cpu extension params
func (o *AddCPUArchitectureCPUExtensionParams) WithContext(ctx context.Context) *AddCPUArchitectureCPUExtensionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the add Cpu architecture Cpu extension params
func (o *AddCPUArchitectureCPUExtensionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the add Cpu architecture Cpu extension params
func (o *AddCPUArchitectureCPUExtensionParams) WithHTTPClient(client *http.Client) *AddCPUArchitectureCPUExtensionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the add Cpu architecture Cpu extension params
func (o *AddCPUArchitectureCPUExtensionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCPUArchitectureID adds the cPUArchitectureID to the add Cpu architecture Cpu extension params
func (o *AddCPUArchitectureCPUExtensionParams) WithCPUArchitectureID(cPUArchitectureID strfmt.UUID) *AddCPUArchitectureCPUExtensionParams {
	o.SetCPUArchitectureID(cPUArchitectureID)
	return o
}

// SetCPUArchitectureID adds the cpuArchitectureId to the add Cpu architecture Cpu extension params
func (o *AddCPUArchitectureCPUExtensionParams) SetCPUArchitectureID(cPUArchitectureID strfmt.UUID) {
	o.CPUArchitectureID = cPUArchitectureID
}

// WithCPUExtensionID adds the cPUExtensionID to the add Cpu architecture Cpu extension params
func (o *AddCPUArchitectureCPUExtensionParams) WithCPUExtensionID(cPUExtensionID *inventory_models.AddCPUArchitectureCPUExtensionParamsBody) *AddCPUArchitectureCPUExtensionParams {
	o.SetCPUExtensionID(cPUExtensionID)
	return o
}

// SetCPUExtensionID adds the cpuExtensionId to the add Cpu architecture Cpu extension params
func (o *AddCPUArchitectureCPUExtensionParams) SetCPUExtensionID(cPUExtensionID *inventory_models.AddCPUArchitectureCPUExtensionParamsBody) {
	o.CPUExtensionID = cPUExtensionID
}

// WriteToRequest writes these params to a swagger request
func (o *AddCPUArchitectureCPUExtensionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param cpu_architecture_id
	if err := r.SetPathParam("cpu_architecture_id", o.CPUArchitectureID.String()); err != nil {
		return err
	}

	if o.CPUExtensionID != nil {
		if err := r.SetBodyParam(o.CPUExtensionID); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
