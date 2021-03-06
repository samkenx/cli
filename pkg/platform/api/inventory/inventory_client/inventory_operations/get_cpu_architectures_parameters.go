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
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetCPUArchitecturesParams creates a new GetCPUArchitecturesParams object
// with the default values initialized.
func NewGetCPUArchitecturesParams() *GetCPUArchitecturesParams {
	var (
		allowUnstableDefault = bool(false)
		limitDefault         = int64(50)
		pageDefault          = int64(1)
	)
	return &GetCPUArchitecturesParams{
		AllowUnstable: &allowUnstableDefault,
		Limit:         &limitDefault,
		Page:          &pageDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetCPUArchitecturesParamsWithTimeout creates a new GetCPUArchitecturesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetCPUArchitecturesParamsWithTimeout(timeout time.Duration) *GetCPUArchitecturesParams {
	var (
		allowUnstableDefault = bool(false)
		limitDefault         = int64(50)
		pageDefault          = int64(1)
	)
	return &GetCPUArchitecturesParams{
		AllowUnstable: &allowUnstableDefault,
		Limit:         &limitDefault,
		Page:          &pageDefault,

		timeout: timeout,
	}
}

// NewGetCPUArchitecturesParamsWithContext creates a new GetCPUArchitecturesParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetCPUArchitecturesParamsWithContext(ctx context.Context) *GetCPUArchitecturesParams {
	var (
		allowUnstableDefault = bool(false)
		limitDefault         = int64(50)
		pageDefault          = int64(1)
	)
	return &GetCPUArchitecturesParams{
		AllowUnstable: &allowUnstableDefault,
		Limit:         &limitDefault,
		Page:          &pageDefault,

		Context: ctx,
	}
}

// NewGetCPUArchitecturesParamsWithHTTPClient creates a new GetCPUArchitecturesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetCPUArchitecturesParamsWithHTTPClient(client *http.Client) *GetCPUArchitecturesParams {
	var (
		allowUnstableDefault = bool(false)
		limitDefault         = int64(50)
		pageDefault          = int64(1)
	)
	return &GetCPUArchitecturesParams{
		AllowUnstable: &allowUnstableDefault,
		Limit:         &limitDefault,
		Page:          &pageDefault,
		HTTPClient:    client,
	}
}

/*GetCPUArchitecturesParams contains all the parameters to send to the API endpoint
for the get Cpu architectures operation typically these are written to a http.Request
*/
type GetCPUArchitecturesParams struct {

	/*AllowUnstable
	  Whether to show an unstable revision of a resource if there is an available unstable version newer than the newest available stable version

	*/
	AllowUnstable *bool
	/*Limit
	  The maximum number of items returned per page

	*/
	Limit *int64
	/*Page
	  The page number returned

	*/
	Page *int64
	/*StateAt
	  Show the state of a resource as it was at the specified timestamp. If omitted, shows the current state of the resource.

	*/
	StateAt *strfmt.DateTime

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get Cpu architectures params
func (o *GetCPUArchitecturesParams) WithTimeout(timeout time.Duration) *GetCPUArchitecturesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get Cpu architectures params
func (o *GetCPUArchitecturesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get Cpu architectures params
func (o *GetCPUArchitecturesParams) WithContext(ctx context.Context) *GetCPUArchitecturesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get Cpu architectures params
func (o *GetCPUArchitecturesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get Cpu architectures params
func (o *GetCPUArchitecturesParams) WithHTTPClient(client *http.Client) *GetCPUArchitecturesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get Cpu architectures params
func (o *GetCPUArchitecturesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAllowUnstable adds the allowUnstable to the get Cpu architectures params
func (o *GetCPUArchitecturesParams) WithAllowUnstable(allowUnstable *bool) *GetCPUArchitecturesParams {
	o.SetAllowUnstable(allowUnstable)
	return o
}

// SetAllowUnstable adds the allowUnstable to the get Cpu architectures params
func (o *GetCPUArchitecturesParams) SetAllowUnstable(allowUnstable *bool) {
	o.AllowUnstable = allowUnstable
}

// WithLimit adds the limit to the get Cpu architectures params
func (o *GetCPUArchitecturesParams) WithLimit(limit *int64) *GetCPUArchitecturesParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the get Cpu architectures params
func (o *GetCPUArchitecturesParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithPage adds the page to the get Cpu architectures params
func (o *GetCPUArchitecturesParams) WithPage(page *int64) *GetCPUArchitecturesParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the get Cpu architectures params
func (o *GetCPUArchitecturesParams) SetPage(page *int64) {
	o.Page = page
}

// WithStateAt adds the stateAt to the get Cpu architectures params
func (o *GetCPUArchitecturesParams) WithStateAt(stateAt *strfmt.DateTime) *GetCPUArchitecturesParams {
	o.SetStateAt(stateAt)
	return o
}

// SetStateAt adds the stateAt to the get Cpu architectures params
func (o *GetCPUArchitecturesParams) SetStateAt(stateAt *strfmt.DateTime) {
	o.StateAt = stateAt
}

// WriteToRequest writes these params to a swagger request
func (o *GetCPUArchitecturesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.AllowUnstable != nil {

		// query param allow_unstable
		var qrAllowUnstable bool
		if o.AllowUnstable != nil {
			qrAllowUnstable = *o.AllowUnstable
		}
		qAllowUnstable := swag.FormatBool(qrAllowUnstable)
		if qAllowUnstable != "" {
			if err := r.SetQueryParam("allow_unstable", qAllowUnstable); err != nil {
				return err
			}
		}

	}

	if o.Limit != nil {

		// query param limit
		var qrLimit int64
		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := swag.FormatInt64(qrLimit)
		if qLimit != "" {
			if err := r.SetQueryParam("limit", qLimit); err != nil {
				return err
			}
		}

	}

	if o.Page != nil {

		// query param page
		var qrPage int64
		if o.Page != nil {
			qrPage = *o.Page
		}
		qPage := swag.FormatInt64(qrPage)
		if qPage != "" {
			if err := r.SetQueryParam("page", qPage); err != nil {
				return err
			}
		}

	}

	if o.StateAt != nil {

		// query param state_at
		var qrStateAt strfmt.DateTime
		if o.StateAt != nil {
			qrStateAt = *o.StateAt
		}
		qStateAt := qrStateAt.String()
		if qStateAt != "" {
			if err := r.SetQueryParam("state_at", qStateAt); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
