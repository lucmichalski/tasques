// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// TaskNewTask task new task
// swagger:model task.NewTask
type TaskNewTask struct {

	// Arguments for this Task
	Args interface{} `json:"args,omitempty"`

	// Context for this Task
	Context interface{} `json:"context,omitempty"`

	// The kind of Task; corresponds roughly with a function name
	// Required: true
	Kind *string `json:"kind"`

	// The priority of this Task (higher means higher priority)
	// If not passed, defaults to zero (neutral)
	Priority int64 `json:"priority,omitempty"`

	// How long a Worker has upon claiming this Task to finish or report back before it gets timed out by the Tasques server
	// If not passed, falls back to a server-side configured default
	ProcessingTimeout string `json:"processing_timeout,omitempty"`

	// The queue that a Task will be inserted into
	// Required: true
	Queue *string `json:"queue"`

	// The number of times that a Task will be retried if it fails
	// If not passed, falls back to a server-side configured default
	RetryTimes int64 `json:"retry_times,omitempty"`

	// If defined, when this Task should run
	// If not passed, falls back to now.
	// Format: date-time
	RunAt strfmt.DateTime `json:"run_at,omitempty"`
}

// Validate validates this task new task
func (m *TaskNewTask) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateKind(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateQueue(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRunAt(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TaskNewTask) validateKind(formats strfmt.Registry) error {

	if err := validate.Required("kind", "body", m.Kind); err != nil {
		return err
	}

	return nil
}

func (m *TaskNewTask) validateQueue(formats strfmt.Registry) error {

	if err := validate.Required("queue", "body", m.Queue); err != nil {
		return err
	}

	return nil
}

func (m *TaskNewTask) validateRunAt(formats strfmt.Registry) error {

	if swag.IsZero(m.RunAt) { // not required
		return nil
	}

	if err := validate.FormatOf("run_at", "body", "date-time", m.RunAt.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TaskNewTask) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TaskNewTask) UnmarshalBinary(b []byte) error {
	var res TaskNewTask
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
