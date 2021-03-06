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

// UserCredentialsTotpReader is a Reader for the UserCredentialsTotp structure.
type UserCredentialsTotpReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UserCredentialsTotpReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUserCredentialsTotpOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUserCredentialsTotpBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUserCredentialsTotpNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUserCredentialsTotpOK creates a UserCredentialsTotpOK with default headers values
func NewUserCredentialsTotpOK() *UserCredentialsTotpOK {
	return &UserCredentialsTotpOK{}
}

/*UserCredentialsTotpOK handles this case with default header values.

Two-Factor Credential
*/
type UserCredentialsTotpOK struct {
	Payload *models.CredentialsTotp
}

func (o *UserCredentialsTotpOK) Error() string {
	return fmt.Sprintf("[GET /users/{user_id}/credentials_totp][%d] userCredentialsTotpOK  %+v", 200, o.Payload)
}

func (o *UserCredentialsTotpOK) GetPayload() *models.CredentialsTotp {
	return o.Payload
}

func (o *UserCredentialsTotpOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CredentialsTotp)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUserCredentialsTotpBadRequest creates a UserCredentialsTotpBadRequest with default headers values
func NewUserCredentialsTotpBadRequest() *UserCredentialsTotpBadRequest {
	return &UserCredentialsTotpBadRequest{}
}

/*UserCredentialsTotpBadRequest handles this case with default header values.

Bad Request
*/
type UserCredentialsTotpBadRequest struct {
	Payload *models.Error
}

func (o *UserCredentialsTotpBadRequest) Error() string {
	return fmt.Sprintf("[GET /users/{user_id}/credentials_totp][%d] userCredentialsTotpBadRequest  %+v", 400, o.Payload)
}

func (o *UserCredentialsTotpBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *UserCredentialsTotpBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUserCredentialsTotpNotFound creates a UserCredentialsTotpNotFound with default headers values
func NewUserCredentialsTotpNotFound() *UserCredentialsTotpNotFound {
	return &UserCredentialsTotpNotFound{}
}

/*UserCredentialsTotpNotFound handles this case with default header values.

Not Found
*/
type UserCredentialsTotpNotFound struct {
	Payload *models.Error
}

func (o *UserCredentialsTotpNotFound) Error() string {
	return fmt.Sprintf("[GET /users/{user_id}/credentials_totp][%d] userCredentialsTotpNotFound  %+v", 404, o.Payload)
}

func (o *UserCredentialsTotpNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *UserCredentialsTotpNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
