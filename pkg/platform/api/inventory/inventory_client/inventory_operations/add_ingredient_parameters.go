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

// NewAddIngredientParams creates a new AddIngredientParams object
// with the default values initialized.
func NewAddIngredientParams() *AddIngredientParams {
	var ()
	return &AddIngredientParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewAddIngredientParamsWithTimeout creates a new AddIngredientParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAddIngredientParamsWithTimeout(timeout time.Duration) *AddIngredientParams {
	var ()
	return &AddIngredientParams{

		timeout: timeout,
	}
}

// NewAddIngredientParamsWithContext creates a new AddIngredientParams object
// with the default values initialized, and the ability to set a context for a request
func NewAddIngredientParamsWithContext(ctx context.Context) *AddIngredientParams {
	var ()
	return &AddIngredientParams{

		Context: ctx,
	}
}

// NewAddIngredientParamsWithHTTPClient creates a new AddIngredientParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAddIngredientParamsWithHTTPClient(client *http.Client) *AddIngredientParams {
	var ()
	return &AddIngredientParams{
		HTTPClient: client,
	}
}

/*AddIngredientParams contains all the parameters to send to the API endpoint
for the add ingredient operation typically these are written to a http.Request
*/
type AddIngredientParams struct {

	/*Ingredient*/
	Ingredient *inventory_models.V1IngredientCreate
	/*IsDefaultProviderBehavior
	  How setting of is_default_provider should be handled for the features provided by the thing being created. The options are: 'override' - if is_default_provider is set to true for a provided feature then replace the existing default provider of that feature (if one exists); 'if-new' - set is_default_provider to true if this is an entirely new feature, otherwise false

	*/
	IsDefaultProviderBehavior *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the add ingredient params
func (o *AddIngredientParams) WithTimeout(timeout time.Duration) *AddIngredientParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the add ingredient params
func (o *AddIngredientParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the add ingredient params
func (o *AddIngredientParams) WithContext(ctx context.Context) *AddIngredientParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the add ingredient params
func (o *AddIngredientParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the add ingredient params
func (o *AddIngredientParams) WithHTTPClient(client *http.Client) *AddIngredientParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the add ingredient params
func (o *AddIngredientParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIngredient adds the ingredient to the add ingredient params
func (o *AddIngredientParams) WithIngredient(ingredient *inventory_models.V1IngredientCreate) *AddIngredientParams {
	o.SetIngredient(ingredient)
	return o
}

// SetIngredient adds the ingredient to the add ingredient params
func (o *AddIngredientParams) SetIngredient(ingredient *inventory_models.V1IngredientCreate) {
	o.Ingredient = ingredient
}

// WithIsDefaultProviderBehavior adds the isDefaultProviderBehavior to the add ingredient params
func (o *AddIngredientParams) WithIsDefaultProviderBehavior(isDefaultProviderBehavior *string) *AddIngredientParams {
	o.SetIsDefaultProviderBehavior(isDefaultProviderBehavior)
	return o
}

// SetIsDefaultProviderBehavior adds the isDefaultProviderBehavior to the add ingredient params
func (o *AddIngredientParams) SetIsDefaultProviderBehavior(isDefaultProviderBehavior *string) {
	o.IsDefaultProviderBehavior = isDefaultProviderBehavior
}

// WriteToRequest writes these params to a swagger request
func (o *AddIngredientParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Ingredient != nil {
		if err := r.SetBodyParam(o.Ingredient); err != nil {
			return err
		}
	}

	if o.IsDefaultProviderBehavior != nil {

		// query param is_default_provider_behavior
		var qrIsDefaultProviderBehavior string
		if o.IsDefaultProviderBehavior != nil {
			qrIsDefaultProviderBehavior = *o.IsDefaultProviderBehavior
		}
		qIsDefaultProviderBehavior := qrIsDefaultProviderBehavior
		if qIsDefaultProviderBehavior != "" {
			if err := r.SetQueryParam("is_default_provider_behavior", qIsDefaultProviderBehavior); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
