// Code generated by go-swagger; DO NOT EDIT.

package recurring_tasks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/lloydmeta/tasques/models"
)

// GetExistingRecurringTaskReader is a Reader for the GetExistingRecurringTask structure.
type GetExistingRecurringTaskReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetExistingRecurringTaskReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetExistingRecurringTaskOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetExistingRecurringTaskNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetExistingRecurringTaskOK creates a GetExistingRecurringTaskOK with default headers values
func NewGetExistingRecurringTaskOK() *GetExistingRecurringTaskOK {
	return &GetExistingRecurringTaskOK{}
}

/*GetExistingRecurringTaskOK handles this case with default header values.

OK
*/
type GetExistingRecurringTaskOK struct {
	Payload *models.RecurringTask
}

func (o *GetExistingRecurringTaskOK) Error() string {
	return fmt.Sprintf("[GET /recurring_tasques/{id}][%d] getExistingRecurringTaskOK  %+v", 200, o.Payload)
}

func (o *GetExistingRecurringTaskOK) GetPayload() *models.RecurringTask {
	return o.Payload
}

func (o *GetExistingRecurringTaskOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RecurringTask)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetExistingRecurringTaskNotFound creates a GetExistingRecurringTaskNotFound with default headers values
func NewGetExistingRecurringTaskNotFound() *GetExistingRecurringTaskNotFound {
	return &GetExistingRecurringTaskNotFound{}
}

/*GetExistingRecurringTaskNotFound handles this case with default header values.

Recurring Task does not exist
*/
type GetExistingRecurringTaskNotFound struct {
	Payload *models.CommonBody
}

func (o *GetExistingRecurringTaskNotFound) Error() string {
	return fmt.Sprintf("[GET /recurring_tasques/{id}][%d] getExistingRecurringTaskNotFound  %+v", 404, o.Payload)
}

func (o *GetExistingRecurringTaskNotFound) GetPayload() *models.CommonBody {
	return o.Payload
}

func (o *GetExistingRecurringTaskNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CommonBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}