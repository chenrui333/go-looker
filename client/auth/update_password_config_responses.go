// Code generated by go-swagger; DO NOT EDIT.

package auth

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/chenrui333/go-looker/models"
)

// UpdatePasswordConfigReader is a Reader for the UpdatePasswordConfig structure.
type UpdatePasswordConfigReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdatePasswordConfigReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdatePasswordConfigOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdatePasswordConfigBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdatePasswordConfigNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewUpdatePasswordConfigUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUpdatePasswordConfigOK creates a UpdatePasswordConfigOK with default headers values
func NewUpdatePasswordConfigOK() *UpdatePasswordConfigOK {
	return &UpdatePasswordConfigOK{}
}

/*UpdatePasswordConfigOK handles this case with default header values.

Password Config
*/
type UpdatePasswordConfigOK struct {
	Payload *models.PasswordConfig
}

func (o *UpdatePasswordConfigOK) Error() string {
	return fmt.Sprintf("[PATCH /password_config][%d] updatePasswordConfigOK  %+v", 200, o.Payload)
}

func (o *UpdatePasswordConfigOK) GetPayload() *models.PasswordConfig {
	return o.Payload
}

func (o *UpdatePasswordConfigOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PasswordConfig)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdatePasswordConfigBadRequest creates a UpdatePasswordConfigBadRequest with default headers values
func NewUpdatePasswordConfigBadRequest() *UpdatePasswordConfigBadRequest {
	return &UpdatePasswordConfigBadRequest{}
}

/*UpdatePasswordConfigBadRequest handles this case with default header values.

Bad Request
*/
type UpdatePasswordConfigBadRequest struct {
	Payload *models.Error
}

func (o *UpdatePasswordConfigBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /password_config][%d] updatePasswordConfigBadRequest  %+v", 400, o.Payload)
}

func (o *UpdatePasswordConfigBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdatePasswordConfigBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdatePasswordConfigNotFound creates a UpdatePasswordConfigNotFound with default headers values
func NewUpdatePasswordConfigNotFound() *UpdatePasswordConfigNotFound {
	return &UpdatePasswordConfigNotFound{}
}

/*UpdatePasswordConfigNotFound handles this case with default header values.

Not Found
*/
type UpdatePasswordConfigNotFound struct {
	Payload *models.Error
}

func (o *UpdatePasswordConfigNotFound) Error() string {
	return fmt.Sprintf("[PATCH /password_config][%d] updatePasswordConfigNotFound  %+v", 404, o.Payload)
}

func (o *UpdatePasswordConfigNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdatePasswordConfigNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdatePasswordConfigUnprocessableEntity creates a UpdatePasswordConfigUnprocessableEntity with default headers values
func NewUpdatePasswordConfigUnprocessableEntity() *UpdatePasswordConfigUnprocessableEntity {
	return &UpdatePasswordConfigUnprocessableEntity{}
}

/*UpdatePasswordConfigUnprocessableEntity handles this case with default header values.

Validation Error
*/
type UpdatePasswordConfigUnprocessableEntity struct {
	Payload *models.ValidationError
}

func (o *UpdatePasswordConfigUnprocessableEntity) Error() string {
	return fmt.Sprintf("[PATCH /password_config][%d] updatePasswordConfigUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *UpdatePasswordConfigUnprocessableEntity) GetPayload() *models.ValidationError {
	return o.Payload
}

func (o *UpdatePasswordConfigUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
