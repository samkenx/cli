// Code generated by go-swagger; DO NOT EDIT.

package authentication

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

// NewGetLoginJwtTokenParams creates a new GetLoginJwtTokenParams object
// with the default values initialized.
func NewGetLoginJwtTokenParams() *GetLoginJwtTokenParams {
	var ()
	return &GetLoginJwtTokenParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetLoginJwtTokenParamsWithTimeout creates a new GetLoginJwtTokenParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetLoginJwtTokenParamsWithTimeout(timeout time.Duration) *GetLoginJwtTokenParams {
	var ()
	return &GetLoginJwtTokenParams{

		timeout: timeout,
	}
}

// NewGetLoginJwtTokenParamsWithContext creates a new GetLoginJwtTokenParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetLoginJwtTokenParamsWithContext(ctx context.Context) *GetLoginJwtTokenParams {
	var ()
	return &GetLoginJwtTokenParams{

		Context: ctx,
	}
}

// NewGetLoginJwtTokenParamsWithHTTPClient creates a new GetLoginJwtTokenParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetLoginJwtTokenParamsWithHTTPClient(client *http.Client) *GetLoginJwtTokenParams {
	var ()
	return &GetLoginJwtTokenParams{
		HTTPClient: client,
	}
}

/*GetLoginJwtTokenParams contains all the parameters to send to the API endpoint
for the get login jwt token operation typically these are written to a http.Request
*/
type GetLoginJwtTokenParams struct {

	/*RedirectURL
	  redirectURL after login

	*/
	RedirectURL strfmt.URI
	/*Token
	  token to login with

	*/
	Token string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get login jwt token params
func (o *GetLoginJwtTokenParams) WithTimeout(timeout time.Duration) *GetLoginJwtTokenParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get login jwt token params
func (o *GetLoginJwtTokenParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get login jwt token params
func (o *GetLoginJwtTokenParams) WithContext(ctx context.Context) *GetLoginJwtTokenParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get login jwt token params
func (o *GetLoginJwtTokenParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get login jwt token params
func (o *GetLoginJwtTokenParams) WithHTTPClient(client *http.Client) *GetLoginJwtTokenParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get login jwt token params
func (o *GetLoginJwtTokenParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRedirectURL adds the redirectURL to the get login jwt token params
func (o *GetLoginJwtTokenParams) WithRedirectURL(redirectURL strfmt.URI) *GetLoginJwtTokenParams {
	o.SetRedirectURL(redirectURL)
	return o
}

// SetRedirectURL adds the redirectUrl to the get login jwt token params
func (o *GetLoginJwtTokenParams) SetRedirectURL(redirectURL strfmt.URI) {
	o.RedirectURL = redirectURL
}

// WithToken adds the token to the get login jwt token params
func (o *GetLoginJwtTokenParams) WithToken(token string) *GetLoginJwtTokenParams {
	o.SetToken(token)
	return o
}

// SetToken adds the token to the get login jwt token params
func (o *GetLoginJwtTokenParams) SetToken(token string) {
	o.Token = token
}

// WriteToRequest writes these params to a swagger request
func (o *GetLoginJwtTokenParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param redirectURL
	qrRedirectURL := o.RedirectURL
	qRedirectURL := qrRedirectURL.String()
	if qRedirectURL != "" {
		if err := r.SetQueryParam("redirectURL", qRedirectURL); err != nil {
			return err
		}
	}

	// path param token
	if err := r.SetPathParam("token", o.Token); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}