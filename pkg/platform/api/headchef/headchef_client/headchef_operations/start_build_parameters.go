// Code generated by go-swagger; DO NOT EDIT.

package headchef_operations

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

	headchef_models "github.com/ActiveState/cli/pkg/platform/api/headchef/headchef_models"
)

// NewStartBuildParams creates a new StartBuildParams object
// with the default values initialized.
func NewStartBuildParams() *StartBuildParams {
	var ()
	return &StartBuildParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewStartBuildParamsWithTimeout creates a new StartBuildParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewStartBuildParamsWithTimeout(timeout time.Duration) *StartBuildParams {
	var ()
	return &StartBuildParams{

		timeout: timeout,
	}
}

// NewStartBuildParamsWithContext creates a new StartBuildParams object
// with the default values initialized, and the ability to set a context for a request
func NewStartBuildParamsWithContext(ctx context.Context) *StartBuildParams {
	var ()
	return &StartBuildParams{

		Context: ctx,
	}
}

// NewStartBuildParamsWithHTTPClient creates a new StartBuildParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewStartBuildParamsWithHTTPClient(client *http.Client) *StartBuildParams {
	var ()
	return &StartBuildParams{
		HTTPClient: client,
	}
}

/*StartBuildParams contains all the parameters to send to the API endpoint
for the start build operation typically these are written to a http.Request
*/
type StartBuildParams struct {

	/*BuildRequest*/
	BuildRequest *headchef_models.BuildRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the start build params
func (o *StartBuildParams) WithTimeout(timeout time.Duration) *StartBuildParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the start build params
func (o *StartBuildParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the start build params
func (o *StartBuildParams) WithContext(ctx context.Context) *StartBuildParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the start build params
func (o *StartBuildParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the start build params
func (o *StartBuildParams) WithHTTPClient(client *http.Client) *StartBuildParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the start build params
func (o *StartBuildParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBuildRequest adds the buildRequest to the start build params
func (o *StartBuildParams) WithBuildRequest(buildRequest *headchef_models.BuildRequest) *StartBuildParams {
	o.SetBuildRequest(buildRequest)
	return o
}

// SetBuildRequest adds the buildRequest to the start build params
func (o *StartBuildParams) SetBuildRequest(buildRequest *headchef_models.BuildRequest) {
	o.BuildRequest = buildRequest
}

// WriteToRequest writes these params to a swagger request
func (o *StartBuildParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.BuildRequest != nil {
		if err := r.SetBodyParam(o.BuildRequest); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
