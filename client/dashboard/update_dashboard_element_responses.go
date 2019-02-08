// Code generated by go-swagger; DO NOT EDIT.

package dashboard

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/chenrui333/go-looker/models"
)

// UpdateDashboardElementReader is a Reader for the UpdateDashboardElement structure.
type UpdateDashboardElementReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateDashboardElementReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewUpdateDashboardElementOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewUpdateDashboardElementBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewUpdateDashboardElementNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 422:
		result := NewUpdateDashboardElementUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUpdateDashboardElementOK creates a UpdateDashboardElementOK with default headers values
func NewUpdateDashboardElementOK() *UpdateDashboardElementOK {
	return &UpdateDashboardElementOK{}
}

/*UpdateDashboardElementOK handles this case with default header values.

DashboardElement
*/
type UpdateDashboardElementOK struct {
	Payload *models.DashboardElement
}

func (o *UpdateDashboardElementOK) Error() string {
	return fmt.Sprintf("[PATCH /dashboard_elements/{dashboard_element_id}][%d] updateDashboardElementOK  %+v", 200, o.Payload)
}

func (o *UpdateDashboardElementOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DashboardElement)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateDashboardElementBadRequest creates a UpdateDashboardElementBadRequest with default headers values
func NewUpdateDashboardElementBadRequest() *UpdateDashboardElementBadRequest {
	return &UpdateDashboardElementBadRequest{}
}

/*UpdateDashboardElementBadRequest handles this case with default header values.

Bad Request
*/
type UpdateDashboardElementBadRequest struct {
	Payload *models.Error
}

func (o *UpdateDashboardElementBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /dashboard_elements/{dashboard_element_id}][%d] updateDashboardElementBadRequest  %+v", 400, o.Payload)
}

func (o *UpdateDashboardElementBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateDashboardElementNotFound creates a UpdateDashboardElementNotFound with default headers values
func NewUpdateDashboardElementNotFound() *UpdateDashboardElementNotFound {
	return &UpdateDashboardElementNotFound{}
}

/*UpdateDashboardElementNotFound handles this case with default header values.

Not Found
*/
type UpdateDashboardElementNotFound struct {
	Payload *models.Error
}

func (o *UpdateDashboardElementNotFound) Error() string {
	return fmt.Sprintf("[PATCH /dashboard_elements/{dashboard_element_id}][%d] updateDashboardElementNotFound  %+v", 404, o.Payload)
}

func (o *UpdateDashboardElementNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateDashboardElementUnprocessableEntity creates a UpdateDashboardElementUnprocessableEntity with default headers values
func NewUpdateDashboardElementUnprocessableEntity() *UpdateDashboardElementUnprocessableEntity {
	return &UpdateDashboardElementUnprocessableEntity{}
}

/*UpdateDashboardElementUnprocessableEntity handles this case with default header values.

Validation Error
*/
type UpdateDashboardElementUnprocessableEntity struct {
	Payload *models.ValidationError
}

func (o *UpdateDashboardElementUnprocessableEntity) Error() string {
	return fmt.Sprintf("[PATCH /dashboard_elements/{dashboard_element_id}][%d] updateDashboardElementUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *UpdateDashboardElementUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
