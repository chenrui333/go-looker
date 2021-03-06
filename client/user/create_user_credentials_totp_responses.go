// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/chenrui333/go-looker/models"
)

// CreateUserCredentialsTotpReader is a Reader for the CreateUserCredentialsTotp structure.
type CreateUserCredentialsTotpReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateUserCredentialsTotpReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateUserCredentialsTotpOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateUserCredentialsTotpBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCreateUserCredentialsTotpNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewCreateUserCredentialsTotpConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewCreateUserCredentialsTotpUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCreateUserCredentialsTotpOK creates a CreateUserCredentialsTotpOK with default headers values
func NewCreateUserCredentialsTotpOK() *CreateUserCredentialsTotpOK {
	return &CreateUserCredentialsTotpOK{}
}

/*CreateUserCredentialsTotpOK handles this case with default header values.

Two-Factor Credential
*/
type CreateUserCredentialsTotpOK struct {
	Payload *models.CredentialsTotp
}

func (o *CreateUserCredentialsTotpOK) Error() string {
	return fmt.Sprintf("[POST /users/{user_id}/credentials_totp][%d] createUserCredentialsTotpOK  %+v", 200, o.Payload)
}

func (o *CreateUserCredentialsTotpOK) GetPayload() *models.CredentialsTotp {
	return o.Payload
}

func (o *CreateUserCredentialsTotpOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CredentialsTotp)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateUserCredentialsTotpBadRequest creates a CreateUserCredentialsTotpBadRequest with default headers values
func NewCreateUserCredentialsTotpBadRequest() *CreateUserCredentialsTotpBadRequest {
	return &CreateUserCredentialsTotpBadRequest{}
}

/*CreateUserCredentialsTotpBadRequest handles this case with default header values.

Bad Request
*/
type CreateUserCredentialsTotpBadRequest struct {
	Payload *models.Error
}

func (o *CreateUserCredentialsTotpBadRequest) Error() string {
	return fmt.Sprintf("[POST /users/{user_id}/credentials_totp][%d] createUserCredentialsTotpBadRequest  %+v", 400, o.Payload)
}

func (o *CreateUserCredentialsTotpBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateUserCredentialsTotpBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateUserCredentialsTotpNotFound creates a CreateUserCredentialsTotpNotFound with default headers values
func NewCreateUserCredentialsTotpNotFound() *CreateUserCredentialsTotpNotFound {
	return &CreateUserCredentialsTotpNotFound{}
}

/*CreateUserCredentialsTotpNotFound handles this case with default header values.

Not Found
*/
type CreateUserCredentialsTotpNotFound struct {
	Payload *models.Error
}

func (o *CreateUserCredentialsTotpNotFound) Error() string {
	return fmt.Sprintf("[POST /users/{user_id}/credentials_totp][%d] createUserCredentialsTotpNotFound  %+v", 404, o.Payload)
}

func (o *CreateUserCredentialsTotpNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateUserCredentialsTotpNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateUserCredentialsTotpConflict creates a CreateUserCredentialsTotpConflict with default headers values
func NewCreateUserCredentialsTotpConflict() *CreateUserCredentialsTotpConflict {
	return &CreateUserCredentialsTotpConflict{}
}

/*CreateUserCredentialsTotpConflict handles this case with default header values.

Resource Already Exists
*/
type CreateUserCredentialsTotpConflict struct {
	Payload *models.Error
}

func (o *CreateUserCredentialsTotpConflict) Error() string {
	return fmt.Sprintf("[POST /users/{user_id}/credentials_totp][%d] createUserCredentialsTotpConflict  %+v", 409, o.Payload)
}

func (o *CreateUserCredentialsTotpConflict) GetPayload() *models.Error {
	return o.Payload
}

func (o *CreateUserCredentialsTotpConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateUserCredentialsTotpUnprocessableEntity creates a CreateUserCredentialsTotpUnprocessableEntity with default headers values
func NewCreateUserCredentialsTotpUnprocessableEntity() *CreateUserCredentialsTotpUnprocessableEntity {
	return &CreateUserCredentialsTotpUnprocessableEntity{}
}

/*CreateUserCredentialsTotpUnprocessableEntity handles this case with default header values.

Validation Error
*/
type CreateUserCredentialsTotpUnprocessableEntity struct {
	Payload *models.ValidationError
}

func (o *CreateUserCredentialsTotpUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /users/{user_id}/credentials_totp][%d] createUserCredentialsTotpUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *CreateUserCredentialsTotpUnprocessableEntity) GetPayload() *models.ValidationError {
	return o.Payload
}

func (o *CreateUserCredentialsTotpUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
