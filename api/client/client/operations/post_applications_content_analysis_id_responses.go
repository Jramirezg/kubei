// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cisco-open/kubei/api/client/models"
)

// PostApplicationsContentAnalysisIDReader is a Reader for the PostApplicationsContentAnalysisID structure.
type PostApplicationsContentAnalysisIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostApplicationsContentAnalysisIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPostApplicationsContentAnalysisIDCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewPostApplicationsContentAnalysisIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewPostApplicationsContentAnalysisIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPostApplicationsContentAnalysisIDCreated creates a PostApplicationsContentAnalysisIDCreated with default headers values
func NewPostApplicationsContentAnalysisIDCreated() *PostApplicationsContentAnalysisIDCreated {
	return &PostApplicationsContentAnalysisIDCreated{}
}

/* PostApplicationsContentAnalysisIDCreated describes a response with status code 201, with default header values.

Application content analysis successfully reported.
*/
type PostApplicationsContentAnalysisIDCreated struct {
}

func (o *PostApplicationsContentAnalysisIDCreated) Error() string {
	return fmt.Sprintf("[POST /applications/contentAnalysis/{id}][%d] postApplicationsContentAnalysisIdCreated ", 201)
}

func (o *PostApplicationsContentAnalysisIDCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostApplicationsContentAnalysisIDNotFound creates a PostApplicationsContentAnalysisIDNotFound with default headers values
func NewPostApplicationsContentAnalysisIDNotFound() *PostApplicationsContentAnalysisIDNotFound {
	return &PostApplicationsContentAnalysisIDNotFound{}
}

/* PostApplicationsContentAnalysisIDNotFound describes a response with status code 404, with default header values.

Application not found.
*/
type PostApplicationsContentAnalysisIDNotFound struct {
}

func (o *PostApplicationsContentAnalysisIDNotFound) Error() string {
	return fmt.Sprintf("[POST /applications/contentAnalysis/{id}][%d] postApplicationsContentAnalysisIdNotFound ", 404)
}

func (o *PostApplicationsContentAnalysisIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostApplicationsContentAnalysisIDDefault creates a PostApplicationsContentAnalysisIDDefault with default headers values
func NewPostApplicationsContentAnalysisIDDefault(code int) *PostApplicationsContentAnalysisIDDefault {
	return &PostApplicationsContentAnalysisIDDefault{
		_statusCode: code,
	}
}

/* PostApplicationsContentAnalysisIDDefault describes a response with status code -1, with default header values.

unknown error
*/
type PostApplicationsContentAnalysisIDDefault struct {
	_statusCode int

	Payload *models.APIResponse
}

// Code gets the status code for the post applications content analysis ID default response
func (o *PostApplicationsContentAnalysisIDDefault) Code() int {
	return o._statusCode
}

func (o *PostApplicationsContentAnalysisIDDefault) Error() string {
	return fmt.Sprintf("[POST /applications/contentAnalysis/{id}][%d] PostApplicationsContentAnalysisID default  %+v", o._statusCode, o.Payload)
}
func (o *PostApplicationsContentAnalysisIDDefault) GetPayload() *models.APIResponse {
	return o.Payload
}

func (o *PostApplicationsContentAnalysisIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
