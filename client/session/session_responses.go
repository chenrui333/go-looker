// Code generated by go-swagger; DO NOT EDIT.

package session

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/chenrui333/go-looker/models"
)

// SessionReader is a Reader for the Session structure.
type SessionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SessionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewSessionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewSessionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewSessionNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewSessionOK creates a SessionOK with default headers values
func NewSessionOK() *SessionOK {
	return &SessionOK{}
}

/*SessionOK handles this case with default header values.

Session
*/
type SessionOK struct {
	Payload *models.APISession
}

func (o *SessionOK) Error() string {
	return fmt.Sprintf("[GET /session][%d] sessionOK  %+v", 200, o.Payload)
}

func (o *SessionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APISession)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSessionBadRequest creates a SessionBadRequest with default headers values
func NewSessionBadRequest() *SessionBadRequest {
	return &SessionBadRequest{}
}

/*SessionBadRequest handles this case with default header values.

Bad Request
*/
type SessionBadRequest struct {
	Payload *models.Error
}

func (o *SessionBadRequest) Error() string {
	return fmt.Sprintf("[GET /session][%d] sessionBadRequest  %+v", 400, o.Payload)
}

func (o *SessionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSessionNotFound creates a SessionNotFound with default headers values
func NewSessionNotFound() *SessionNotFound {
	return &SessionNotFound{}
}

/*SessionNotFound handles this case with default header values.

Not Found
*/
type SessionNotFound struct {
	Payload *models.Error
}

func (o *SessionNotFound) Error() string {
	return fmt.Sprintf("[GET /session][%d] sessionNotFound  %+v", 404, o.Payload)
}

func (o *SessionNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}