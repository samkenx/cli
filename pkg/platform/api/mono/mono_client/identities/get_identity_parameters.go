// Code generated by go-swagger; DO NOT EDIT.

package identities

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

// NewGetIdentityParams creates a new GetIdentityParams object
// with the default values initialized.
func NewGetIdentityParams() *GetIdentityParams {
	var ()
	return &GetIdentityParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetIdentityParamsWithTimeout creates a new GetIdentityParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetIdentityParamsWithTimeout(timeout time.Duration) *GetIdentityParams {
	var ()
	return &GetIdentityParams{

		timeout: timeout,
	}
}

// NewGetIdentityParamsWithContext creates a new GetIdentityParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetIdentityParamsWithContext(ctx context.Context) *GetIdentityParams {
	var ()
	return &GetIdentityParams{

		Context: ctx,
	}
}

// NewGetIdentityParamsWithHTTPClient creates a new GetIdentityParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetIdentityParamsWithHTTPClient(client *http.Client) *GetIdentityParams {
	var ()
	return &GetIdentityParams{
		HTTPClient: client,
	}
}

/*GetIdentityParams contains all the parameters to send to the API endpoint
for the get identity operation typically these are written to a http.Request
*/
type GetIdentityParams struct {

	/*IdentityID
	  identityID of desired Identity

	*/
	IdentityID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get identity params
func (o *GetIdentityParams) WithTimeout(timeout time.Duration) *GetIdentityParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get identity params
func (o *GetIdentityParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get identity params
func (o *GetIdentityParams) WithContext(ctx context.Context) *GetIdentityParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get identity params
func (o *GetIdentityParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get identity params
func (o *GetIdentityParams) WithHTTPClient(client *http.Client) *GetIdentityParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get identity params
func (o *GetIdentityParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIdentityID adds the identityID to the get identity params
func (o *GetIdentityParams) WithIdentityID(identityID strfmt.UUID) *GetIdentityParams {
	o.SetIdentityID(identityID)
	return o
}

// SetIdentityID adds the identityId to the get identity params
func (o *GetIdentityParams) SetIdentityID(identityID strfmt.UUID) {
	o.IdentityID = identityID
}

// WriteToRequest writes these params to a swagger request
func (o *GetIdentityParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param identityID
	if err := r.SetPathParam("identityID", o.IdentityID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
