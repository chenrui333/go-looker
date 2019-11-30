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

// SearchUserLoginLockoutsReader is a Reader for the SearchUserLoginLockouts structure.
type SearchUserLoginLockoutsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SearchUserLoginLockoutsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSearchUserLoginLockoutsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewSearchUserLoginLockoutsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewSearchUserLoginLockoutsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewSearchUserLoginLockoutsOK creates a SearchUserLoginLockoutsOK with default headers values
func NewSearchUserLoginLockoutsOK() *SearchUserLoginLockoutsOK {
	return &SearchUserLoginLockoutsOK{}
}

/*SearchUserLoginLockoutsOK handles this case with default header values.

User Login Lockout
*/
type SearchUserLoginLockoutsOK struct {
	Payload []*models.UserLoginLockout
}

func (o *SearchUserLoginLockoutsOK) Error() string {
	return fmt.Sprintf("[GET /user_login_lockouts/search][%d] searchUserLoginLockoutsOK  %+v", 200, o.Payload)
}

func (o *SearchUserLoginLockoutsOK) GetPayload() []*models.UserLoginLockout {
	return o.Payload
}

func (o *SearchUserLoginLockoutsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSearchUserLoginLockoutsBadRequest creates a SearchUserLoginLockoutsBadRequest with default headers values
func NewSearchUserLoginLockoutsBadRequest() *SearchUserLoginLockoutsBadRequest {
	return &SearchUserLoginLockoutsBadRequest{}
}

/*SearchUserLoginLockoutsBadRequest handles this case with default header values.

Bad Request
*/
type SearchUserLoginLockoutsBadRequest struct {
	Payload *models.Error
}

func (o *SearchUserLoginLockoutsBadRequest) Error() string {
	return fmt.Sprintf("[GET /user_login_lockouts/search][%d] searchUserLoginLockoutsBadRequest  %+v", 400, o.Payload)
}

func (o *SearchUserLoginLockoutsBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *SearchUserLoginLockoutsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSearchUserLoginLockoutsNotFound creates a SearchUserLoginLockoutsNotFound with default headers values
func NewSearchUserLoginLockoutsNotFound() *SearchUserLoginLockoutsNotFound {
	return &SearchUserLoginLockoutsNotFound{}
}

/*SearchUserLoginLockoutsNotFound handles this case with default header values.

Not Found
*/
type SearchUserLoginLockoutsNotFound struct {
	Payload *models.Error
}

func (o *SearchUserLoginLockoutsNotFound) Error() string {
	return fmt.Sprintf("[GET /user_login_lockouts/search][%d] searchUserLoginLockoutsNotFound  %+v", 404, o.Payload)
}

func (o *SearchUserLoginLockoutsNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *SearchUserLoginLockoutsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
