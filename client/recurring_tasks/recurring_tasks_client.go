// Code generated by go-swagger; DO NOT EDIT.

package recurring_tasks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new recurring tasks API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for recurring tasks API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	CreateRecurringTask(params *CreateRecurringTaskParams) (*CreateRecurringTaskCreated, error)

	DeleteExistingRecurringTask(params *DeleteExistingRecurringTaskParams) (*DeleteExistingRecurringTaskOK, error)

	GetExistingRecurringTask(params *GetExistingRecurringTaskParams) (*GetExistingRecurringTaskOK, error)

	ListExistingRecurringTasks(params *ListExistingRecurringTasksParams) (*ListExistingRecurringTasksOK, error)

	UpdateExistingRecurringTask(params *UpdateExistingRecurringTaskParams) (*UpdateExistingRecurringTaskOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  CreateRecurringTask adds a new recurring task

  Creates a new Recurring Task
*/
func (a *Client) CreateRecurringTask(params *CreateRecurringTaskParams) (*CreateRecurringTaskCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateRecurringTaskParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "create-recurring-task",
		Method:             "POST",
		PathPattern:        "/recurring_tasques",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CreateRecurringTaskReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CreateRecurringTaskCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for create-recurring-task: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DeleteExistingRecurringTask deletes a recurring task

  Deletes a persisted Recurring Task
*/
func (a *Client) DeleteExistingRecurringTask(params *DeleteExistingRecurringTaskParams) (*DeleteExistingRecurringTaskOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteExistingRecurringTaskParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "delete-existing-recurring-task",
		Method:             "DELETE",
		PathPattern:        "/recurring_tasques/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteExistingRecurringTaskReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteExistingRecurringTaskOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for delete-existing-recurring-task: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetExistingRecurringTask gets a recurring task

  Retrieves a persisted Recurring Task
*/
func (a *Client) GetExistingRecurringTask(params *GetExistingRecurringTaskParams) (*GetExistingRecurringTaskOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetExistingRecurringTaskParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "get-existing-recurring-task",
		Method:             "GET",
		PathPattern:        "/recurring_tasques/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetExistingRecurringTaskReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetExistingRecurringTaskOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for get-existing-recurring-task: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ListExistingRecurringTasks lists recurring tasks

  Lists persisted Recurring Tasks
*/
func (a *Client) ListExistingRecurringTasks(params *ListExistingRecurringTasksParams) (*ListExistingRecurringTasksOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListExistingRecurringTasksParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "list-existing-recurring-tasks",
		Method:             "GET",
		PathPattern:        "/recurring_tasques",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ListExistingRecurringTasksReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ListExistingRecurringTasksOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for list-existing-recurring-tasks: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UpdateExistingRecurringTask updates a recurring task

  Updates a persisted Recurring Task
*/
func (a *Client) UpdateExistingRecurringTask(params *UpdateExistingRecurringTaskParams) (*UpdateExistingRecurringTaskOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateExistingRecurringTaskParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "update-existing-recurring-task",
		Method:             "PUT",
		PathPattern:        "/recurring_tasques/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UpdateExistingRecurringTaskReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UpdateExistingRecurringTaskOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for update-existing-recurring-task: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}