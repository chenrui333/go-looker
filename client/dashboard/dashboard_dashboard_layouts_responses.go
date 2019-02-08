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

// DashboardDashboardLayoutsReader is a Reader for the DashboardDashboardLayouts structure.
type DashboardDashboardLayoutsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DashboardDashboardLayoutsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDashboardDashboardLayoutsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewDashboardDashboardLayoutsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewDashboardDashboardLayoutsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDashboardDashboardLayoutsOK creates a DashboardDashboardLayoutsOK with default headers values
func NewDashboardDashboardLayoutsOK() *DashboardDashboardLayoutsOK {
	return &DashboardDashboardLayoutsOK{}
}

/*DashboardDashboardLayoutsOK handles this case with default header values.

DashboardLayout
*/
type DashboardDashboardLayoutsOK struct {
	Payload []*models.DashboardLayout
}

func (o *DashboardDashboardLayoutsOK) Error() string {
	return fmt.Sprintf("[GET /dashboards/{dashboard_id}/dashboard_layouts][%d] dashboardDashboardLayoutsOK  %+v", 200, o.Payload)
}

func (o *DashboardDashboardLayoutsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDashboardDashboardLayoutsBadRequest creates a DashboardDashboardLayoutsBadRequest with default headers values
func NewDashboardDashboardLayoutsBadRequest() *DashboardDashboardLayoutsBadRequest {
	return &DashboardDashboardLayoutsBadRequest{}
}

/*DashboardDashboardLayoutsBadRequest handles this case with default header values.

Bad Request
*/
type DashboardDashboardLayoutsBadRequest struct {
	Payload *models.Error
}

func (o *DashboardDashboardLayoutsBadRequest) Error() string {
	return fmt.Sprintf("[GET /dashboards/{dashboard_id}/dashboard_layouts][%d] dashboardDashboardLayoutsBadRequest  %+v", 400, o.Payload)
}

func (o *DashboardDashboardLayoutsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDashboardDashboardLayoutsNotFound creates a DashboardDashboardLayoutsNotFound with default headers values
func NewDashboardDashboardLayoutsNotFound() *DashboardDashboardLayoutsNotFound {
	return &DashboardDashboardLayoutsNotFound{}
}

/*DashboardDashboardLayoutsNotFound handles this case with default header values.

Not Found
*/
type DashboardDashboardLayoutsNotFound struct {
	Payload *models.Error
}

func (o *DashboardDashboardLayoutsNotFound) Error() string {
	return fmt.Sprintf("[GET /dashboards/{dashboard_id}/dashboard_layouts][%d] dashboardDashboardLayoutsNotFound  %+v", 404, o.Payload)
}

func (o *DashboardDashboardLayoutsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}