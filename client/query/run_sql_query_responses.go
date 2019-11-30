// Code generated by go-swagger; DO NOT EDIT.

package query

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/chenrui333/go-looker/models"
)

// RunSQLQueryReader is a Reader for the RunSQLQuery structure.
type RunSQLQueryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RunSQLQueryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRunSQLQueryOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewRunSQLQueryBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewRunSQLQueryNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewRunSQLQueryUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewRunSQLQueryOK creates a RunSQLQueryOK with default headers values
func NewRunSQLQueryOK() *RunSQLQueryOK {
	return &RunSQLQueryOK{}
}

/*RunSQLQueryOK handles this case with default header values.

SQL Runner Query
*/
type RunSQLQueryOK struct {
	Payload string
}

func (o *RunSQLQueryOK) Error() string {
	return fmt.Sprintf("[POST /sql_queries/{slug}/run/{result_format}][%d] runSqlQueryOK  %+v", 200, o.Payload)
}

func (o *RunSQLQueryOK) GetPayload() string {
	return o.Payload
}

func (o *RunSQLQueryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRunSQLQueryBadRequest creates a RunSQLQueryBadRequest with default headers values
func NewRunSQLQueryBadRequest() *RunSQLQueryBadRequest {
	return &RunSQLQueryBadRequest{}
}

/*RunSQLQueryBadRequest handles this case with default header values.

Bad Request
*/
type RunSQLQueryBadRequest struct {
	Payload *models.Error
}

func (o *RunSQLQueryBadRequest) Error() string {
	return fmt.Sprintf("[POST /sql_queries/{slug}/run/{result_format}][%d] runSqlQueryBadRequest  %+v", 400, o.Payload)
}

func (o *RunSQLQueryBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *RunSQLQueryBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRunSQLQueryNotFound creates a RunSQLQueryNotFound with default headers values
func NewRunSQLQueryNotFound() *RunSQLQueryNotFound {
	return &RunSQLQueryNotFound{}
}

/*RunSQLQueryNotFound handles this case with default header values.

Not Found
*/
type RunSQLQueryNotFound struct {
	Payload *models.Error
}

func (o *RunSQLQueryNotFound) Error() string {
	return fmt.Sprintf("[POST /sql_queries/{slug}/run/{result_format}][%d] runSqlQueryNotFound  %+v", 404, o.Payload)
}

func (o *RunSQLQueryNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *RunSQLQueryNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRunSQLQueryUnprocessableEntity creates a RunSQLQueryUnprocessableEntity with default headers values
func NewRunSQLQueryUnprocessableEntity() *RunSQLQueryUnprocessableEntity {
	return &RunSQLQueryUnprocessableEntity{}
}

/*RunSQLQueryUnprocessableEntity handles this case with default header values.

Validation Error
*/
type RunSQLQueryUnprocessableEntity struct {
	Payload *models.ValidationError
}

func (o *RunSQLQueryUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /sql_queries/{slug}/run/{result_format}][%d] runSqlQueryUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *RunSQLQueryUnprocessableEntity) GetPayload() *models.ValidationError {
	return o.Payload
}

func (o *RunSQLQueryUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
