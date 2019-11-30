// Code generated by go-swagger; DO NOT EDIT.

package color_collection

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/chenrui333/go-looker/models"
)

// ColorCollectionReader is a Reader for the ColorCollection structure.
type ColorCollectionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ColorCollectionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewColorCollectionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewColorCollectionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewColorCollectionNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewColorCollectionOK creates a ColorCollectionOK with default headers values
func NewColorCollectionOK() *ColorCollectionOK {
	return &ColorCollectionOK{}
}

/*ColorCollectionOK handles this case with default header values.

ColorCollection
*/
type ColorCollectionOK struct {
	Payload *models.ColorCollection
}

func (o *ColorCollectionOK) Error() string {
	return fmt.Sprintf("[GET /color_collections/{collection_id}][%d] colorCollectionOK  %+v", 200, o.Payload)
}

func (o *ColorCollectionOK) GetPayload() *models.ColorCollection {
	return o.Payload
}

func (o *ColorCollectionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ColorCollection)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewColorCollectionBadRequest creates a ColorCollectionBadRequest with default headers values
func NewColorCollectionBadRequest() *ColorCollectionBadRequest {
	return &ColorCollectionBadRequest{}
}

/*ColorCollectionBadRequest handles this case with default header values.

Bad Request
*/
type ColorCollectionBadRequest struct {
	Payload *models.Error
}

func (o *ColorCollectionBadRequest) Error() string {
	return fmt.Sprintf("[GET /color_collections/{collection_id}][%d] colorCollectionBadRequest  %+v", 400, o.Payload)
}

func (o *ColorCollectionBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *ColorCollectionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewColorCollectionNotFound creates a ColorCollectionNotFound with default headers values
func NewColorCollectionNotFound() *ColorCollectionNotFound {
	return &ColorCollectionNotFound{}
}

/*ColorCollectionNotFound handles this case with default header values.

Not Found
*/
type ColorCollectionNotFound struct {
	Payload *models.Error
}

func (o *ColorCollectionNotFound) Error() string {
	return fmt.Sprintf("[GET /color_collections/{collection_id}][%d] colorCollectionNotFound  %+v", 404, o.Payload)
}

func (o *ColorCollectionNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *ColorCollectionNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
