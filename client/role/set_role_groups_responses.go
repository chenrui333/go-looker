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

// SetRoleGroupsReader is a Reader for the SetRoleGroups structure.
type SetRoleGroupsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SetRoleGroupsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSetRoleGroupsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewSetRoleGroupsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewSetRoleGroupsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewSetRoleGroupsUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewSetRoleGroupsOK creates a SetRoleGroupsOK with default headers values
func NewSetRoleGroupsOK() *SetRoleGroupsOK {
	return &SetRoleGroupsOK{}
}

/*SetRoleGroupsOK handles this case with default header values.

Groups with role.
*/
type SetRoleGroupsOK struct {
	Payload []*models.Group
}

func (o *SetRoleGroupsOK) Error() string {
	return fmt.Sprintf("[PUT /roles/{role_id}/groups][%d] setRoleGroupsOK  %+v", 200, o.Payload)
}

func (o *SetRoleGroupsOK) GetPayload() []*models.Group {
	return o.Payload
}

func (o *SetRoleGroupsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSetRoleGroupsBadRequest creates a SetRoleGroupsBadRequest with default headers values
func NewSetRoleGroupsBadRequest() *SetRoleGroupsBadRequest {
	return &SetRoleGroupsBadRequest{}
}

/*SetRoleGroupsBadRequest handles this case with default header values.

Bad Request
*/
type SetRoleGroupsBadRequest struct {
	Payload *models.Error
}

func (o *SetRoleGroupsBadRequest) Error() string {
	return fmt.Sprintf("[PUT /roles/{role_id}/groups][%d] setRoleGroupsBadRequest  %+v", 400, o.Payload)
}

func (o *SetRoleGroupsBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *SetRoleGroupsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSetRoleGroupsNotFound creates a SetRoleGroupsNotFound with default headers values
func NewSetRoleGroupsNotFound() *SetRoleGroupsNotFound {
	return &SetRoleGroupsNotFound{}
}

/*SetRoleGroupsNotFound handles this case with default header values.

Not Found
*/
type SetRoleGroupsNotFound struct {
	Payload *models.Error
}

func (o *SetRoleGroupsNotFound) Error() string {
	return fmt.Sprintf("[PUT /roles/{role_id}/groups][%d] setRoleGroupsNotFound  %+v", 404, o.Payload)
}

func (o *SetRoleGroupsNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *SetRoleGroupsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSetRoleGroupsUnprocessableEntity creates a SetRoleGroupsUnprocessableEntity with default headers values
func NewSetRoleGroupsUnprocessableEntity() *SetRoleGroupsUnprocessableEntity {
	return &SetRoleGroupsUnprocessableEntity{}
}

/*SetRoleGroupsUnprocessableEntity handles this case with default header values.

Validation Error
*/
type SetRoleGroupsUnprocessableEntity struct {
	Payload *models.ValidationError
}

func (o *SetRoleGroupsUnprocessableEntity) Error() string {
	return fmt.Sprintf("[PUT /roles/{role_id}/groups][%d] setRoleGroupsUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *SetRoleGroupsUnprocessableEntity) GetPayload() *models.ValidationError {
	return o.Payload
}

func (o *SetRoleGroupsUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
