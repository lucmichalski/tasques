// Code generated by go-swagger; DO NOT EDIT.

package recurring_tasks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewDeleteExistingRecurringTaskParams creates a new DeleteExistingRecurringTaskParams object
// with the default values initialized.
func NewDeleteExistingRecurringTaskParams() *DeleteExistingRecurringTaskParams {
	var ()
	return &DeleteExistingRecurringTaskParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteExistingRecurringTaskParamsWithTimeout creates a new DeleteExistingRecurringTaskParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteExistingRecurringTaskParamsWithTimeout(timeout time.Duration) *DeleteExistingRecurringTaskParams {
	var ()
	return &DeleteExistingRecurringTaskParams{

		timeout: timeout,
	}
}

// NewDeleteExistingRecurringTaskParamsWithContext creates a new DeleteExistingRecurringTaskParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteExistingRecurringTaskParamsWithContext(ctx context.Context) *DeleteExistingRecurringTaskParams {
	var ()
	return &DeleteExistingRecurringTaskParams{

		Context: ctx,
	}
}

// NewDeleteExistingRecurringTaskParamsWithHTTPClient creates a new DeleteExistingRecurringTaskParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteExistingRecurringTaskParamsWithHTTPClient(client *http.Client) *DeleteExistingRecurringTaskParams {
	var ()
	return &DeleteExistingRecurringTaskParams{
		HTTPClient: client,
	}
}

/*DeleteExistingRecurringTaskParams contains all the parameters to send to the API endpoint
for the delete existing recurring task operation typically these are written to a http.Request
*/
type DeleteExistingRecurringTaskParams struct {

	/*ID
	  The id of the Recurring Task

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete existing recurring task params
func (o *DeleteExistingRecurringTaskParams) WithTimeout(timeout time.Duration) *DeleteExistingRecurringTaskParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete existing recurring task params
func (o *DeleteExistingRecurringTaskParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete existing recurring task params
func (o *DeleteExistingRecurringTaskParams) WithContext(ctx context.Context) *DeleteExistingRecurringTaskParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete existing recurring task params
func (o *DeleteExistingRecurringTaskParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete existing recurring task params
func (o *DeleteExistingRecurringTaskParams) WithHTTPClient(client *http.Client) *DeleteExistingRecurringTaskParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete existing recurring task params
func (o *DeleteExistingRecurringTaskParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the delete existing recurring task params
func (o *DeleteExistingRecurringTaskParams) WithID(id string) *DeleteExistingRecurringTaskParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete existing recurring task params
func (o *DeleteExistingRecurringTaskParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteExistingRecurringTaskParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}