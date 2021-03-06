// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// PushHandlerFunc turns a function with the right signature into a push handler
type PushHandlerFunc func(PushParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PushHandlerFunc) Handle(params PushParams) middleware.Responder {
	return fn(params)
}

// PushHandler interface for that can handle valid push params
type PushHandler interface {
	Handle(PushParams) middleware.Responder
}

// NewPush creates a new http.Handler for the push operation
func NewPush(ctx *middleware.Context, handler PushHandler) *Push {
	return &Push{Context: ctx, Handler: handler}
}

/*Push swagger:route POST /push push

Push some data

Push some data

*/
type Push struct {
	Context *middleware.Context
	Handler PushHandler
}

func (o *Push) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewPushParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
