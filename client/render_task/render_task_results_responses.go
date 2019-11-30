// Code generated by go-swagger; DO NOT EDIT.

package render_task

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/chenrui333/go-looker/models"
)

// RenderTaskResultsReader is a Reader for the RenderTaskResults structure.
type RenderTaskResultsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RenderTaskResultsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRenderTaskResultsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewRenderTaskResultsAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewRenderTaskResultsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewRenderTaskResultsOK creates a RenderTaskResultsOK with default headers values
func NewRenderTaskResultsOK() *RenderTaskResultsOK {
	return &RenderTaskResultsOK{}
}

/*RenderTaskResultsOK handles this case with default header values.

Document or image
*/
type RenderTaskResultsOK struct {
	Payload string
}

func (o *RenderTaskResultsOK) Error() string {
	return fmt.Sprintf("[GET /render_tasks/{render_task_id}/results][%d] renderTaskResultsOK  %+v", 200, o.Payload)
}

func (o *RenderTaskResultsOK) GetPayload() string {
	return o.Payload
}

func (o *RenderTaskResultsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRenderTaskResultsAccepted creates a RenderTaskResultsAccepted with default headers values
func NewRenderTaskResultsAccepted() *RenderTaskResultsAccepted {
	return &RenderTaskResultsAccepted{}
}

/*RenderTaskResultsAccepted handles this case with default header values.

Accepted
*/
type RenderTaskResultsAccepted struct {
}

func (o *RenderTaskResultsAccepted) Error() string {
	return fmt.Sprintf("[GET /render_tasks/{render_task_id}/results][%d] renderTaskResultsAccepted ", 202)
}

func (o *RenderTaskResultsAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRenderTaskResultsNotFound creates a RenderTaskResultsNotFound with default headers values
func NewRenderTaskResultsNotFound() *RenderTaskResultsNotFound {
	return &RenderTaskResultsNotFound{}
}

/*RenderTaskResultsNotFound handles this case with default header values.

Not Found
*/
type RenderTaskResultsNotFound struct {
	Payload *models.Error
}

func (o *RenderTaskResultsNotFound) Error() string {
	return fmt.Sprintf("[GET /render_tasks/{render_task_id}/results][%d] renderTaskResultsNotFound  %+v", 404, o.Payload)
}

func (o *RenderTaskResultsNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *RenderTaskResultsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
