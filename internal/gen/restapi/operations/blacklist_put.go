// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"context"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// BlacklistPutHandlerFunc turns a function with the right signature into a blacklist put handler
type BlacklistPutHandlerFunc func(BlacklistPutParams) middleware.Responder

// Handle executing the request and returning a response
func (fn BlacklistPutHandlerFunc) Handle(params BlacklistPutParams) middleware.Responder {
	return fn(params)
}

// BlacklistPutHandler interface for that can handle valid blacklist put params
type BlacklistPutHandler interface {
	Handle(BlacklistPutParams) middleware.Responder
}

// NewBlacklistPut creates a new http.Handler for the blacklist put operation
func NewBlacklistPut(ctx *middleware.Context, handler BlacklistPutHandler) *BlacklistPut {
	return &BlacklistPut{Context: ctx, Handler: handler}
}

/* BlacklistPut swagger:route PUT /blacklist blacklistPut

Request to add IP in blacklist

*/
type BlacklistPut struct {
	Context *middleware.Context
	Handler BlacklistPutHandler
}

func (o *BlacklistPut) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewBlacklistPutParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}

// BlacklistPutBody blacklist put body
//
// swagger:model BlacklistPutBody
type BlacklistPutBody struct {

	// ip
	// Required: true
	IP *string `json:"ip"`
}

// Validate validates this blacklist put body
func (o *BlacklistPutBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateIP(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *BlacklistPutBody) validateIP(formats strfmt.Registry) error {

	if err := validate.Required("body"+"."+"ip", "body", o.IP); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this blacklist put body based on context it is used
func (o *BlacklistPutBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *BlacklistPutBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *BlacklistPutBody) UnmarshalBinary(b []byte) error {
	var res BlacklistPutBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// BlacklistPutInternalServerErrorBody blacklist put internal server error body
//
// swagger:model BlacklistPutInternalServerErrorBody
type BlacklistPutInternalServerErrorBody struct {

	// ok
	Ok string `json:"ok,omitempty"`
}

// Validate validates this blacklist put internal server error body
func (o *BlacklistPutInternalServerErrorBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this blacklist put internal server error body based on context it is used
func (o *BlacklistPutInternalServerErrorBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *BlacklistPutInternalServerErrorBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *BlacklistPutInternalServerErrorBody) UnmarshalBinary(b []byte) error {
	var res BlacklistPutInternalServerErrorBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// BlacklistPutOKBody blacklist put o k body
//
// swagger:model BlacklistPutOKBody
type BlacklistPutOKBody struct {

	// ok
	Ok string `json:"ok,omitempty"`
}

// Validate validates this blacklist put o k body
func (o *BlacklistPutOKBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this blacklist put o k body based on context it is used
func (o *BlacklistPutOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *BlacklistPutOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *BlacklistPutOKBody) UnmarshalBinary(b []byte) error {
	var res BlacklistPutOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
