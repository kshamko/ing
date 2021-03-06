// Code generated by go-swagger; DO NOT EDIT.

package routes

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/kshamko/ing/internal/models"
)

// RoutesOKCode is the HTTP code returned for type RoutesOK
const RoutesOKCode int = 200

/*RoutesOK A successful response.

swagger:response routesOK
*/
type RoutesOK struct {

	/*
	  In: Body
	*/
	Payload *models.Routes `json:"body,omitempty"`
}

// NewRoutesOK creates RoutesOK with default headers values
func NewRoutesOK() *RoutesOK {

	return &RoutesOK{}
}

// WithPayload adds the payload to the routes o k response
func (o *RoutesOK) WithPayload(payload *models.Routes) *RoutesOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the routes o k response
func (o *RoutesOK) SetPayload(payload *models.Routes) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *RoutesOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// RoutesUnprocessableEntityCode is the HTTP code returned for type RoutesUnprocessableEntity
const RoutesUnprocessableEntityCode int = 422

/*RoutesUnprocessableEntity Unprocessable entity

swagger:response routesUnprocessableEntity
*/
type RoutesUnprocessableEntity struct {

	/*
	  In: Body
	*/
	Payload *models.APIInvalidResponse `json:"body,omitempty"`
}

// NewRoutesUnprocessableEntity creates RoutesUnprocessableEntity with default headers values
func NewRoutesUnprocessableEntity() *RoutesUnprocessableEntity {

	return &RoutesUnprocessableEntity{}
}

// WithPayload adds the payload to the routes unprocessable entity response
func (o *RoutesUnprocessableEntity) WithPayload(payload *models.APIInvalidResponse) *RoutesUnprocessableEntity {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the routes unprocessable entity response
func (o *RoutesUnprocessableEntity) SetPayload(payload *models.APIInvalidResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *RoutesUnprocessableEntity) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(422)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// RoutesInternalServerErrorCode is the HTTP code returned for type RoutesInternalServerError
const RoutesInternalServerErrorCode int = 500

/*RoutesInternalServerError Error response

swagger:response routesInternalServerError
*/
type RoutesInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.APIInvalidResponse `json:"body,omitempty"`
}

// NewRoutesInternalServerError creates RoutesInternalServerError with default headers values
func NewRoutesInternalServerError() *RoutesInternalServerError {

	return &RoutesInternalServerError{}
}

// WithPayload adds the payload to the routes internal server error response
func (o *RoutesInternalServerError) WithPayload(payload *models.APIInvalidResponse) *RoutesInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the routes internal server error response
func (o *RoutesInternalServerError) SetPayload(payload *models.APIInvalidResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *RoutesInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
