// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/cisco-open/kubei/runtime_scan/api/server/models"
)

// PostScanScanUUIDResultsCreatedCode is the HTTP code returned for type PostScanScanUUIDResultsCreated
const PostScanScanUUIDResultsCreatedCode int = 201

/*PostScanScanUUIDResultsCreated Image vulnerability scan successfully reported.

swagger:response postScanScanUuidResultsCreated
*/
type PostScanScanUUIDResultsCreated struct {
}

// NewPostScanScanUUIDResultsCreated creates PostScanScanUUIDResultsCreated with default headers values
func NewPostScanScanUUIDResultsCreated() *PostScanScanUUIDResultsCreated {

	return &PostScanScanUUIDResultsCreated{}
}

// WriteResponse to the client
func (o *PostScanScanUUIDResultsCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(201)
}

/*PostScanScanUUIDResultsDefault unknown error

swagger:response postScanScanUuidResultsDefault
*/
type PostScanScanUUIDResultsDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.APIResponse `json:"body,omitempty"`
}

// NewPostScanScanUUIDResultsDefault creates PostScanScanUUIDResultsDefault with default headers values
func NewPostScanScanUUIDResultsDefault(code int) *PostScanScanUUIDResultsDefault {
	if code <= 0 {
		code = 500
	}

	return &PostScanScanUUIDResultsDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the post scan scan UUID results default response
func (o *PostScanScanUUIDResultsDefault) WithStatusCode(code int) *PostScanScanUUIDResultsDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the post scan scan UUID results default response
func (o *PostScanScanUUIDResultsDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the post scan scan UUID results default response
func (o *PostScanScanUUIDResultsDefault) WithPayload(payload *models.APIResponse) *PostScanScanUUIDResultsDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post scan scan UUID results default response
func (o *PostScanScanUUIDResultsDefault) SetPayload(payload *models.APIResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostScanScanUUIDResultsDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
