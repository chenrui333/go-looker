// Code generated by go-swagger; DO NOT EDIT.

package integration

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/chenrui333/go-looker/models"
)

// TestIntegrationReader is a Reader for the TestIntegration structure.
type TestIntegrationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TestIntegrationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewTestIntegrationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewTestIntegrationBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewTestIntegrationNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewTestIntegrationUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewTestIntegrationOK creates a TestIntegrationOK with default headers values
func NewTestIntegrationOK() *TestIntegrationOK {
	return &TestIntegrationOK{}
}

/*TestIntegrationOK handles this case with default header values.

Test Result
*/
type TestIntegrationOK struct {
	Payload *models.IntegrationTestResult
}

func (o *TestIntegrationOK) Error() string {
	return fmt.Sprintf("[POST /integrations/{integration_id}/test][%d] testIntegrationOK  %+v", 200, o.Payload)
}

func (o *TestIntegrationOK) GetPayload() *models.IntegrationTestResult {
	return o.Payload
}

func (o *TestIntegrationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IntegrationTestResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTestIntegrationBadRequest creates a TestIntegrationBadRequest with default headers values
func NewTestIntegrationBadRequest() *TestIntegrationBadRequest {
	return &TestIntegrationBadRequest{}
}

/*TestIntegrationBadRequest handles this case with default header values.

Bad Request
*/
type TestIntegrationBadRequest struct {
	Payload *models.Error
}

func (o *TestIntegrationBadRequest) Error() string {
	return fmt.Sprintf("[POST /integrations/{integration_id}/test][%d] testIntegrationBadRequest  %+v", 400, o.Payload)
}

func (o *TestIntegrationBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *TestIntegrationBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTestIntegrationNotFound creates a TestIntegrationNotFound with default headers values
func NewTestIntegrationNotFound() *TestIntegrationNotFound {
	return &TestIntegrationNotFound{}
}

/*TestIntegrationNotFound handles this case with default header values.

Not Found
*/
type TestIntegrationNotFound struct {
	Payload *models.Error
}

func (o *TestIntegrationNotFound) Error() string {
	return fmt.Sprintf("[POST /integrations/{integration_id}/test][%d] testIntegrationNotFound  %+v", 404, o.Payload)
}

func (o *TestIntegrationNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *TestIntegrationNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTestIntegrationUnprocessableEntity creates a TestIntegrationUnprocessableEntity with default headers values
func NewTestIntegrationUnprocessableEntity() *TestIntegrationUnprocessableEntity {
	return &TestIntegrationUnprocessableEntity{}
}

/*TestIntegrationUnprocessableEntity handles this case with default header values.

Validation Error
*/
type TestIntegrationUnprocessableEntity struct {
	Payload *models.ValidationError
}

func (o *TestIntegrationUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /integrations/{integration_id}/test][%d] testIntegrationUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *TestIntegrationUnprocessableEntity) GetPayload() *models.ValidationError {
	return o.Payload
}

func (o *TestIntegrationUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
