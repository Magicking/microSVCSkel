// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/Magicking/microSVCSkel/models"
)

// PushOKCode is the HTTP code returned for type PushOK
const PushOKCode int = 200

/*PushOK Return some string

swagger:response pushOK
*/
type PushOK struct {

	/*
	  In: Body
	*/
	Payload string `json:"body,omitempty"`
}

// NewPushOK creates PushOK with default headers values
func NewPushOK() *PushOK {
	return &PushOK{}
}

// WithPayload adds the payload to the push o k response
func (o *PushOK) WithPayload(payload string) *PushOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the push o k response
func (o *PushOK) SetPayload(payload string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PushOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}

}

/*PushDefault Unexpected error

swagger:response pushDefault
*/
type PushDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewPushDefault creates PushDefault with default headers values
func NewPushDefault(code int) *PushDefault {
	if code <= 0 {
		code = 500
	}

	return &PushDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the push default response
func (o *PushDefault) WithStatusCode(code int) *PushDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the push default response
func (o *PushDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the push default response
func (o *PushDefault) WithPayload(payload *models.Error) *PushDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the push default response
func (o *PushDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PushDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
