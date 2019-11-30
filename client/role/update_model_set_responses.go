// Code generated by go-swagger; DO NOT EDIT.

package role

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/chenrui333/go-looker/models"
)

// UpdateModelSetReader is a Reader for the UpdateModelSet structure.
type UpdateModelSetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateModelSetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateModelSetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateModelSetBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateModelSetNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 405:
		result := NewUpdateModelSetMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewUpdateModelSetUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUpdateModelSetOK creates a UpdateModelSetOK with default headers values
func NewUpdateModelSetOK() *UpdateModelSetOK {
	return &UpdateModelSetOK{}
}

/*UpdateModelSetOK handles this case with default header values.

New state for specified model set.
*/
type UpdateModelSetOK struct {
	Payload *models.ModelSet
}

func (o *UpdateModelSetOK) Error() string {
	return fmt.Sprintf("[PATCH /model_sets/{model_set_id}][%d] updateModelSetOK  %+v", 200, o.Payload)
}

func (o *UpdateModelSetOK) GetPayload() *models.ModelSet {
	return o.Payload
}

func (o *UpdateModelSetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ModelSet)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateModelSetBadRequest creates a UpdateModelSetBadRequest with default headers values
func NewUpdateModelSetBadRequest() *UpdateModelSetBadRequest {
	return &UpdateModelSetBadRequest{}
}

/*UpdateModelSetBadRequest handles this case with default header values.

Bad Request
*/
type UpdateModelSetBadRequest struct {
	Payload *models.Error
}

func (o *UpdateModelSetBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /model_sets/{model_set_id}][%d] updateModelSetBadRequest  %+v", 400, o.Payload)
}

func (o *UpdateModelSetBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdateModelSetBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateModelSetNotFound creates a UpdateModelSetNotFound with default headers values
func NewUpdateModelSetNotFound() *UpdateModelSetNotFound {
	return &UpdateModelSetNotFound{}
}

/*UpdateModelSetNotFound handles this case with default header values.

Not Found
*/
type UpdateModelSetNotFound struct {
	Payload *models.Error
}

func (o *UpdateModelSetNotFound) Error() string {
	return fmt.Sprintf("[PATCH /model_sets/{model_set_id}][%d] updateModelSetNotFound  %+v", 404, o.Payload)
}

func (o *UpdateModelSetNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdateModelSetNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateModelSetMethodNotAllowed creates a UpdateModelSetMethodNotAllowed with default headers values
func NewUpdateModelSetMethodNotAllowed() *UpdateModelSetMethodNotAllowed {
	return &UpdateModelSetMethodNotAllowed{}
}

/*UpdateModelSetMethodNotAllowed handles this case with default header values.

Resource Can't Be Modified
*/
type UpdateModelSetMethodNotAllowed struct {
	Payload *models.Error
}

func (o *UpdateModelSetMethodNotAllowed) Error() string {
	return fmt.Sprintf("[PATCH /model_sets/{model_set_id}][%d] updateModelSetMethodNotAllowed  %+v", 405, o.Payload)
}

func (o *UpdateModelSetMethodNotAllowed) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdateModelSetMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateModelSetUnprocessableEntity creates a UpdateModelSetUnprocessableEntity with default headers values
func NewUpdateModelSetUnprocessableEntity() *UpdateModelSetUnprocessableEntity {
	return &UpdateModelSetUnprocessableEntity{}
}

/*UpdateModelSetUnprocessableEntity handles this case with default header values.

Validation Error
*/
type UpdateModelSetUnprocessableEntity struct {
	Payload *models.ValidationError
}

func (o *UpdateModelSetUnprocessableEntity) Error() string {
	return fmt.Sprintf("[PATCH /model_sets/{model_set_id}][%d] updateModelSetUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *UpdateModelSetUnprocessableEntity) GetPayload() *models.ValidationError {
	return o.Payload
}

func (o *UpdateModelSetUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
