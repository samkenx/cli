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

	inventory_models "github.com/ActiveState/cli/pkg/platform/api/inventory/inventory_models"
)

// NewAddCPUExtensionParams creates a new AddCPUExtensionParams object
// with the default values initialized.
func NewAddCPUExtensionParams() *AddCPUExtensionParams {
	var ()
	return &AddCPUExtensionParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewAddCPUExtensionParamsWithTimeout creates a new AddCPUExtensionParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAddCPUExtensionParamsWithTimeout(timeout time.Duration) *AddCPUExtensionParams {
	var ()
	return &AddCPUExtensionParams{

		timeout: timeout,
	}
}

// NewAddCPUExtensionParamsWithContext creates a new AddCPUExtensionParams object
// with the default values initialized, and the ability to set a context for a request
func NewAddCPUExtensionParamsWithContext(ctx context.Context) *AddCPUExtensionParams {
	var ()
	return &AddCPUExtensionParams{

		Context: ctx,
	}
}

// NewAddCPUExtensionParamsWithHTTPClient creates a new AddCPUExtensionParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAddCPUExtensionParamsWithHTTPClient(client *http.Client) *AddCPUExtensionParams {
	var ()
	return &AddCPUExtensionParams{
		HTTPClient: client,
	}
}

/*AddCPUExtensionParams contains all the parameters to send to the API endpoint
for the add Cpu extension operation typically these are written to a http.Request
*/
type AddCPUExtensionParams struct {

	/*CPUExtension*/
	CPUExtension *inventory_models.V1CPUExtensionCore

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the add Cpu extension params
func (o *AddCPUExtensionParams) WithTimeout(timeout time.Duration) *AddCPUExtensionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the add Cpu extension params
func (o *AddCPUExtensionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the add Cpu extension params
func (o *AddCPUExtensionParams) WithContext(ctx context.Context) *AddCPUExtensionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the add Cpu extension params
func (o *AddCPUExtensionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the add Cpu extension params
func (o *AddCPUExtensionParams) WithHTTPClient(client *http.Client) *AddCPUExtensionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the add Cpu extension params
func (o *AddCPUExtensionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCPUExtension adds the cPUExtension to the add Cpu extension params
func (o *AddCPUExtensionParams) WithCPUExtension(cPUExtension *inventory_models.V1CPUExtensionCore) *AddCPUExtensionParams {
	o.SetCPUExtension(cPUExtension)
	return o
}

// SetCPUExtension adds the cpuExtension to the add Cpu extension params
func (o *AddCPUExtensionParams) SetCPUExtension(cPUExtension *inventory_models.V1CPUExtensionCore) {
	o.CPUExtension = cPUExtension
}

// WriteToRequest writes these params to a swagger request
func (o *AddCPUExtensionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.CPUExtension != nil {
		if err := r.SetBodyParam(o.CPUExtension); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
