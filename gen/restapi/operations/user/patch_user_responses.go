// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"user_server/gen/models"
)

// PatchUserOKCode is the HTTP code returned for type PatchUserOK
const PatchUserOKCode int = 200

/*PatchUserOK Success

swagger:response patchUserOK
*/
type PatchUserOK struct {

	/*
	  In: Body
	*/
	Payload *models.User `json:"body,omitempty"`
}

// NewPatchUserOK creates PatchUserOK with default headers values
func NewPatchUserOK() *PatchUserOK {

	return &PatchUserOK{}
}

// WithPayload adds the payload to the patch user o k response
func (o *PatchUserOK) WithPayload(payload *models.User) *PatchUserOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the patch user o k response
func (o *PatchUserOK) SetPayload(payload *models.User) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PatchUserOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PatchUserBadRequestCode is the HTTP code returned for type PatchUserBadRequest
const PatchUserBadRequestCode int = 400

/*PatchUserBadRequest Invalid format (the scheme depends on the implementation of the code generator)

swagger:response patchUserBadRequest
*/
type PatchUserBadRequest struct {
}

// NewPatchUserBadRequest creates PatchUserBadRequest with default headers values
func NewPatchUserBadRequest() *PatchUserBadRequest {

	return &PatchUserBadRequest{}
}

// WriteResponse to the client
func (o *PatchUserBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(400)
}

// PatchUserNotFoundCode is the HTTP code returned for type PatchUserNotFound
const PatchUserNotFoundCode int = 404

/*PatchUserNotFound User not found

swagger:response patchUserNotFound
*/
type PatchUserNotFound struct {
}

// NewPatchUserNotFound creates PatchUserNotFound with default headers values
func NewPatchUserNotFound() *PatchUserNotFound {

	return &PatchUserNotFound{}
}

// WriteResponse to the client
func (o *PatchUserNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}

// PatchUserInternalServerErrorCode is the HTTP code returned for type PatchUserInternalServerError
const PatchUserInternalServerErrorCode int = 500

/*PatchUserInternalServerError Internal error

swagger:response patchUserInternalServerError
*/
type PatchUserInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewPatchUserInternalServerError creates PatchUserInternalServerError with default headers values
func NewPatchUserInternalServerError() *PatchUserInternalServerError {

	return &PatchUserInternalServerError{}
}

// WithPayload adds the payload to the patch user internal server error response
func (o *PatchUserInternalServerError) WithPayload(payload *models.Error) *PatchUserInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the patch user internal server error response
func (o *PatchUserInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PatchUserInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
