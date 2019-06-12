// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// PostManifestCreatedCode is the HTTP code returned for type PostManifestCreated
const PostManifestCreatedCode int = 201

/*PostManifestCreated Manifest Created

swagger:response postManifestCreated
*/
type PostManifestCreated struct {

	/*
	  In: Body
	*/
	Payload string `json:"body,omitempty"`
}

// NewPostManifestCreated creates PostManifestCreated with default headers values
func NewPostManifestCreated() *PostManifestCreated {

	return &PostManifestCreated{}
}

// WithPayload adds the payload to the post manifest created response
func (o *PostManifestCreated) WithPayload(payload string) *PostManifestCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post manifest created response
func (o *PostManifestCreated) SetPayload(payload string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostManifestCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// PostManifestBadRequestCode is the HTTP code returned for type PostManifestBadRequest
const PostManifestBadRequestCode int = 400

/*PostManifestBadRequest Bad Request Body

swagger:response postManifestBadRequest
*/
type PostManifestBadRequest struct {
}

// NewPostManifestBadRequest creates PostManifestBadRequest with default headers values
func NewPostManifestBadRequest() *PostManifestBadRequest {

	return &PostManifestBadRequest{}
}

// WriteResponse to the client
func (o *PostManifestBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(400)
}

// PostManifestUnauthorizedCode is the HTTP code returned for type PostManifestUnauthorized
const PostManifestUnauthorizedCode int = 401

/*PostManifestUnauthorized Unauthorized

swagger:response postManifestUnauthorized
*/
type PostManifestUnauthorized struct {
	/*Authorization information is missing or invalid

	 */
	WWWAuthenticate string `json:"WWW-Authenticate"`
}

// NewPostManifestUnauthorized creates PostManifestUnauthorized with default headers values
func NewPostManifestUnauthorized() *PostManifestUnauthorized {

	return &PostManifestUnauthorized{}
}

// WithWWWAuthenticate adds the wWWAuthenticate to the post manifest unauthorized response
func (o *PostManifestUnauthorized) WithWWWAuthenticate(wWWAuthenticate string) *PostManifestUnauthorized {
	o.WWWAuthenticate = wWWAuthenticate
	return o
}

// SetWWWAuthenticate sets the wWWAuthenticate to the post manifest unauthorized response
func (o *PostManifestUnauthorized) SetWWWAuthenticate(wWWAuthenticate string) {
	o.WWWAuthenticate = wWWAuthenticate
}

// WriteResponse to the client
func (o *PostManifestUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header WWW-Authenticate

	wWWAuthenticate := o.WWWAuthenticate
	if wWWAuthenticate != "" {
		rw.Header().Set("WWW-Authenticate", wWWAuthenticate)
	}

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}

// PostManifestInternalServerErrorCode is the HTTP code returned for type PostManifestInternalServerError
const PostManifestInternalServerErrorCode int = 500

/*PostManifestInternalServerError Internal Server Error

swagger:response postManifestInternalServerError
*/
type PostManifestInternalServerError struct {

	/*
	  In: Body
	*/
	Payload string `json:"body,omitempty"`
}

// NewPostManifestInternalServerError creates PostManifestInternalServerError with default headers values
func NewPostManifestInternalServerError() *PostManifestInternalServerError {

	return &PostManifestInternalServerError{}
}

// WithPayload adds the payload to the post manifest internal server error response
func (o *PostManifestInternalServerError) WithPayload(payload string) *PostManifestInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post manifest internal server error response
func (o *PostManifestInternalServerError) SetPayload(payload string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostManifestInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}
