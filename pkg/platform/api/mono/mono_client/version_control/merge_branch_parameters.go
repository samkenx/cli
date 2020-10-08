// Code generated by go-swagger; DO NOT EDIT.

package version_control

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

// NewMergeBranchParams creates a new MergeBranchParams object
// with the default values initialized.
func NewMergeBranchParams() *MergeBranchParams {
	var (
		strategyDefault = string("preview")
	)
	return &MergeBranchParams{
		Strategy: strategyDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewMergeBranchParamsWithTimeout creates a new MergeBranchParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewMergeBranchParamsWithTimeout(timeout time.Duration) *MergeBranchParams {
	var (
		strategyDefault = string("preview")
	)
	return &MergeBranchParams{
		Strategy: strategyDefault,

		timeout: timeout,
	}
}

// NewMergeBranchParamsWithContext creates a new MergeBranchParams object
// with the default values initialized, and the ability to set a context for a request
func NewMergeBranchParamsWithContext(ctx context.Context) *MergeBranchParams {
	var (
		strategyDefault = string("preview")
	)
	return &MergeBranchParams{
		Strategy: strategyDefault,

		Context: ctx,
	}
}

// NewMergeBranchParamsWithHTTPClient creates a new MergeBranchParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewMergeBranchParamsWithHTTPClient(client *http.Client) *MergeBranchParams {
	var (
		strategyDefault = string("preview")
	)
	return &MergeBranchParams{
		Strategy:   strategyDefault,
		HTTPClient: client,
	}
}

/*MergeBranchParams contains all the parameters to send to the API endpoint
for the merge branch operation typically these are written to a http.Request
*/
type MergeBranchParams struct {

	/*BranchID*/
	BranchID strfmt.UUID
	/*Strategy*/
	Strategy string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the merge branch params
func (o *MergeBranchParams) WithTimeout(timeout time.Duration) *MergeBranchParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the merge branch params
func (o *MergeBranchParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the merge branch params
func (o *MergeBranchParams) WithContext(ctx context.Context) *MergeBranchParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the merge branch params
func (o *MergeBranchParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the merge branch params
func (o *MergeBranchParams) WithHTTPClient(client *http.Client) *MergeBranchParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the merge branch params
func (o *MergeBranchParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBranchID adds the branchID to the merge branch params
func (o *MergeBranchParams) WithBranchID(branchID strfmt.UUID) *MergeBranchParams {
	o.SetBranchID(branchID)
	return o
}

// SetBranchID adds the branchId to the merge branch params
func (o *MergeBranchParams) SetBranchID(branchID strfmt.UUID) {
	o.BranchID = branchID
}

// WithStrategy adds the strategy to the merge branch params
func (o *MergeBranchParams) WithStrategy(strategy string) *MergeBranchParams {
	o.SetStrategy(strategy)
	return o
}

// SetStrategy adds the strategy to the merge branch params
func (o *MergeBranchParams) SetStrategy(strategy string) {
	o.Strategy = strategy
}

// WriteToRequest writes these params to a swagger request
func (o *MergeBranchParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param branchID
	if err := r.SetPathParam("branchID", o.BranchID.String()); err != nil {
		return err
	}

	// query param strategy
	qrStrategy := o.Strategy
	qStrategy := qrStrategy
	if qStrategy != "" {
		if err := r.SetQueryParam("strategy", qStrategy); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
