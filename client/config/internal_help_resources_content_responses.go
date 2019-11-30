// Code generated by go-swagger; DO NOT EDIT.

package config

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/chenrui333/go-looker/models"
)

// InternalHelpResourcesContentReader is a Reader for the InternalHelpResourcesContent structure.
type InternalHelpResourcesContentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *InternalHelpResourcesContentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewInternalHelpResourcesContentOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewInternalHelpResourcesContentBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewInternalHelpResourcesContentNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewInternalHelpResourcesContentOK creates a InternalHelpResourcesContentOK with default headers values
func NewInternalHelpResourcesContentOK() *InternalHelpResourcesContentOK {
	return &InternalHelpResourcesContentOK{}
}

/*InternalHelpResourcesContentOK handles this case with default header values.

Internal Help Resources Content
*/
type InternalHelpResourcesContentOK struct {
	Payload *models.InternalHelpResourcesContent
}

func (o *InternalHelpResourcesContentOK) Error() string {
	return fmt.Sprintf("[GET /internal_help_resources_content][%d] internalHelpResourcesContentOK  %+v", 200, o.Payload)
}

func (o *InternalHelpResourcesContentOK) GetPayload() *models.InternalHelpResourcesContent {
	return o.Payload
}

func (o *InternalHelpResourcesContentOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InternalHelpResourcesContent)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewInternalHelpResourcesContentBadRequest creates a InternalHelpResourcesContentBadRequest with default headers values
func NewInternalHelpResourcesContentBadRequest() *InternalHelpResourcesContentBadRequest {
	return &InternalHelpResourcesContentBadRequest{}
}

/*InternalHelpResourcesContentBadRequest handles this case with default header values.

Bad Request
*/
type InternalHelpResourcesContentBadRequest struct {
	Payload *models.Error
}

func (o *InternalHelpResourcesContentBadRequest) Error() string {
	return fmt.Sprintf("[GET /internal_help_resources_content][%d] internalHelpResourcesContentBadRequest  %+v", 400, o.Payload)
}

func (o *InternalHelpResourcesContentBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *InternalHelpResourcesContentBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewInternalHelpResourcesContentNotFound creates a InternalHelpResourcesContentNotFound with default headers values
func NewInternalHelpResourcesContentNotFound() *InternalHelpResourcesContentNotFound {
	return &InternalHelpResourcesContentNotFound{}
}

/*InternalHelpResourcesContentNotFound handles this case with default header values.

Not Found
*/
type InternalHelpResourcesContentNotFound struct {
	Payload *models.Error
}

func (o *InternalHelpResourcesContentNotFound) Error() string {
	return fmt.Sprintf("[GET /internal_help_resources_content][%d] internalHelpResourcesContentNotFound  %+v", 404, o.Payload)
}

func (o *InternalHelpResourcesContentNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *InternalHelpResourcesContentNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
