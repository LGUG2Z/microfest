// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	errors "github.com/go-openapi/errors"
	middleware "github.com/go-openapi/runtime/middleware"
	strfmt "github.com/go-openapi/strfmt"
	swag "github.com/go-openapi/swag"
	validate "github.com/go-openapi/validate"

	models "github.com/LGUG2Z/microfest/models"
)

// PutManifestHandlerFunc turns a function with the right signature into a put manifest handler
type PutManifestHandlerFunc func(PutManifestParams, *models.Principal) middleware.Responder

// Handle executing the request and returning a response
func (fn PutManifestHandlerFunc) Handle(params PutManifestParams, principal *models.Principal) middleware.Responder {
	return fn(params, principal)
}

// PutManifestHandler interface for that can handle valid put manifest params
type PutManifestHandler interface {
	Handle(PutManifestParams, *models.Principal) middleware.Responder
}

// NewPutManifest creates a new http.Handler for the put manifest operation
func NewPutManifest(ctx *middleware.Context, handler PutManifestHandler) *PutManifest {
	return &PutManifest{Context: ctx, Handler: handler}
}

/*PutManifest swagger:route PUT /manifest putManifest

Submits a patch to create a new manifest

*/
type PutManifest struct {
	Context *middleware.Context
	Handler PutManifestHandler
}

func (o *PutManifest) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewPutManifestParams()

	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		r = aCtx
	}
	var principal *models.Principal
	if uprinc != nil {
		principal = uprinc.(*models.Principal) // this is really a models.Principal, I promise
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}

// PutManifestBody put manifest body
// swagger:model PutManifestBody
type PutManifestBody struct {

	// manifest
	// Required: true
	Manifest interface{} `json:"manifest"`

	// release
	// Required: true
	Release *string `json:"release"`

	// updated
	// Required: true
	Updated []string `json:"updated"`
}

// Validate validates this put manifest body
func (o *PutManifestBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateManifest(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateRelease(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateUpdated(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PutManifestBody) validateManifest(formats strfmt.Registry) error {

	if err := validate.Required("microfest"+"."+"manifest", "body", o.Manifest); err != nil {
		return err
	}

	return nil
}

func (o *PutManifestBody) validateRelease(formats strfmt.Registry) error {

	if err := validate.Required("microfest"+"."+"release", "body", o.Release); err != nil {
		return err
	}

	return nil
}

func (o *PutManifestBody) validateUpdated(formats strfmt.Registry) error {

	if err := validate.Required("microfest"+"."+"updated", "body", o.Updated); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *PutManifestBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PutManifestBody) UnmarshalBinary(b []byte) error {
	var res PutManifestBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}