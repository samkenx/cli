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

// NewValidateRecipeParams creates a new ValidateRecipeParams object
// with the default values initialized.
func NewValidateRecipeParams() *ValidateRecipeParams {
	var ()
	return &ValidateRecipeParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewValidateRecipeParamsWithTimeout creates a new ValidateRecipeParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewValidateRecipeParamsWithTimeout(timeout time.Duration) *ValidateRecipeParams {
	var ()
	return &ValidateRecipeParams{

		timeout: timeout,
	}
}

// NewValidateRecipeParamsWithContext creates a new ValidateRecipeParams object
// with the default values initialized, and the ability to set a context for a request
func NewValidateRecipeParamsWithContext(ctx context.Context) *ValidateRecipeParams {
	var ()
	return &ValidateRecipeParams{

		Context: ctx,
	}
}

// NewValidateRecipeParamsWithHTTPClient creates a new ValidateRecipeParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewValidateRecipeParamsWithHTTPClient(client *http.Client) *ValidateRecipeParams {
	var ()
	return &ValidateRecipeParams{
		HTTPClient: client,
	}
}

/*ValidateRecipeParams contains all the parameters to send to the API endpoint
for the validate recipe operation typically these are written to a http.Request
*/
type ValidateRecipeParams struct {

	/*Recipe*/
	Recipe *inventory_models.V1Recipe

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the validate recipe params
func (o *ValidateRecipeParams) WithTimeout(timeout time.Duration) *ValidateRecipeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the validate recipe params
func (o *ValidateRecipeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the validate recipe params
func (o *ValidateRecipeParams) WithContext(ctx context.Context) *ValidateRecipeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the validate recipe params
func (o *ValidateRecipeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the validate recipe params
func (o *ValidateRecipeParams) WithHTTPClient(client *http.Client) *ValidateRecipeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the validate recipe params
func (o *ValidateRecipeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRecipe adds the recipe to the validate recipe params
func (o *ValidateRecipeParams) WithRecipe(recipe *inventory_models.V1Recipe) *ValidateRecipeParams {
	o.SetRecipe(recipe)
	return o
}

// SetRecipe adds the recipe to the validate recipe params
func (o *ValidateRecipeParams) SetRecipe(recipe *inventory_models.V1Recipe) {
	o.Recipe = recipe
}

// WriteToRequest writes these params to a swagger request
func (o *ValidateRecipeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Recipe != nil {
		if err := r.SetBodyParam(o.Recipe); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
