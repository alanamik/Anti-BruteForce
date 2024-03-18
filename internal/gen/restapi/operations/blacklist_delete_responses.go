// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// BlacklistDeleteOKCode is the HTTP code returned for type BlacklistDeleteOK
const BlacklistDeleteOKCode int = 200

/*BlacklistDeleteOK The IP has been successfully deleted from blacklist

swagger:response blacklistDeleteOK
*/
type BlacklistDeleteOK struct {

	/*
	  In: Body
	*/
	Payload *BlacklistDeleteOKBody `json:"body,omitempty"`
}

// NewBlacklistDeleteOK creates BlacklistDeleteOK with default headers values
func NewBlacklistDeleteOK() *BlacklistDeleteOK {

	return &BlacklistDeleteOK{}
}

// WithPayload adds the payload to the blacklist delete o k response
func (o *BlacklistDeleteOK) WithPayload(payload *BlacklistDeleteOKBody) *BlacklistDeleteOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the blacklist delete o k response
func (o *BlacklistDeleteOK) SetPayload(payload *BlacklistDeleteOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *BlacklistDeleteOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// BlacklistDeleteInternalServerErrorCode is the HTTP code returned for type BlacklistDeleteInternalServerError
const BlacklistDeleteInternalServerErrorCode int = 500

/*BlacklistDeleteInternalServerError Internal Server Error

swagger:response blacklistDeleteInternalServerError
*/
type BlacklistDeleteInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *BlacklistDeleteInternalServerErrorBody `json:"body,omitempty"`
}

// NewBlacklistDeleteInternalServerError creates BlacklistDeleteInternalServerError with default headers values
func NewBlacklistDeleteInternalServerError() *BlacklistDeleteInternalServerError {

	return &BlacklistDeleteInternalServerError{}
}

// WithPayload adds the payload to the blacklist delete internal server error response
func (o *BlacklistDeleteInternalServerError) WithPayload(payload *BlacklistDeleteInternalServerErrorBody) *BlacklistDeleteInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the blacklist delete internal server error response
func (o *BlacklistDeleteInternalServerError) SetPayload(payload *BlacklistDeleteInternalServerErrorBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *BlacklistDeleteInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}