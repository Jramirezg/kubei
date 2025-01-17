// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"io"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	"github.com/cisco-open/kubei/api/client/models"
)

// GetPackagesIDApplicationResourcesReader is a Reader for the GetPackagesIDApplicationResources structure.
type GetPackagesIDApplicationResourcesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPackagesIDApplicationResourcesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetPackagesIDApplicationResourcesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetPackagesIDApplicationResourcesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetPackagesIDApplicationResourcesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetPackagesIDApplicationResourcesOK creates a GetPackagesIDApplicationResourcesOK with default headers values
func NewGetPackagesIDApplicationResourcesOK() *GetPackagesIDApplicationResourcesOK {
	return &GetPackagesIDApplicationResourcesOK{}
}

/* GetPackagesIDApplicationResourcesOK describes a response with status code 200, with default header values.

Success
*/
type GetPackagesIDApplicationResourcesOK struct {
	Payload *GetPackagesIDApplicationResourcesOKBody
}

func (o *GetPackagesIDApplicationResourcesOK) Error() string {
	return fmt.Sprintf("[GET /packages/{id}/applicationResources][%d] getPackagesIdApplicationResourcesOK  %+v", 200, o.Payload)
}
func (o *GetPackagesIDApplicationResourcesOK) GetPayload() *GetPackagesIDApplicationResourcesOKBody {
	return o.Payload
}

func (o *GetPackagesIDApplicationResourcesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetPackagesIDApplicationResourcesOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPackagesIDApplicationResourcesNotFound creates a GetPackagesIDApplicationResourcesNotFound with default headers values
func NewGetPackagesIDApplicationResourcesNotFound() *GetPackagesIDApplicationResourcesNotFound {
	return &GetPackagesIDApplicationResourcesNotFound{}
}

/* GetPackagesIDApplicationResourcesNotFound describes a response with status code 404, with default header values.

Package ID not found.
*/
type GetPackagesIDApplicationResourcesNotFound struct {
}

func (o *GetPackagesIDApplicationResourcesNotFound) Error() string {
	return fmt.Sprintf("[GET /packages/{id}/applicationResources][%d] getPackagesIdApplicationResourcesNotFound ", 404)
}

func (o *GetPackagesIDApplicationResourcesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetPackagesIDApplicationResourcesDefault creates a GetPackagesIDApplicationResourcesDefault with default headers values
func NewGetPackagesIDApplicationResourcesDefault(code int) *GetPackagesIDApplicationResourcesDefault {
	return &GetPackagesIDApplicationResourcesDefault{
		_statusCode: code,
	}
}

/* GetPackagesIDApplicationResourcesDefault describes a response with status code -1, with default header values.

unknown error
*/
type GetPackagesIDApplicationResourcesDefault struct {
	_statusCode int

	Payload *models.APIResponse
}

// Code gets the status code for the get packages ID application resources default response
func (o *GetPackagesIDApplicationResourcesDefault) Code() int {
	return o._statusCode
}

func (o *GetPackagesIDApplicationResourcesDefault) Error() string {
	return fmt.Sprintf("[GET /packages/{id}/applicationResources][%d] GetPackagesIDApplicationResources default  %+v", o._statusCode, o.Payload)
}
func (o *GetPackagesIDApplicationResourcesDefault) GetPayload() *models.APIResponse {
	return o.Payload
}

func (o *GetPackagesIDApplicationResourcesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetPackagesIDApplicationResourcesOKBody get packages ID application resources o k body
swagger:model GetPackagesIDApplicationResourcesOKBody
*/
type GetPackagesIDApplicationResourcesOKBody struct {

	// List of package application resources in the given filters and page. List length must be lower or equal to pageSize
	Items []*models.PackageApplicationResources `json:"items"`

	// Total package application resources count under the given filters
	// Required: true
	Total *int64 `json:"total"`
}

// Validate validates this get packages ID application resources o k body
func (o *GetPackagesIDApplicationResourcesOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateItems(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateTotal(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetPackagesIDApplicationResourcesOKBody) validateItems(formats strfmt.Registry) error {
	if swag.IsZero(o.Items) { // not required
		return nil
	}

	for i := 0; i < len(o.Items); i++ {
		if swag.IsZero(o.Items[i]) { // not required
			continue
		}

		if o.Items[i] != nil {
			if err := o.Items[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getPackagesIdApplicationResourcesOK" + "." + "items" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *GetPackagesIDApplicationResourcesOKBody) validateTotal(formats strfmt.Registry) error {

	if err := validate.Required("getPackagesIdApplicationResourcesOK"+"."+"total", "body", o.Total); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this get packages ID application resources o k body based on the context it is used
func (o *GetPackagesIDApplicationResourcesOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateItems(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetPackagesIDApplicationResourcesOKBody) contextValidateItems(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Items); i++ {

		if o.Items[i] != nil {
			if err := o.Items[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getPackagesIdApplicationResourcesOK" + "." + "items" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetPackagesIDApplicationResourcesOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetPackagesIDApplicationResourcesOKBody) UnmarshalBinary(b []byte) error {
	var res GetPackagesIDApplicationResourcesOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
