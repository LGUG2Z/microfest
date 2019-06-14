// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/validate"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetInfoParams creates a new GetInfoParams object
// no default values defined in spec.
func NewGetInfoParams() GetInfoParams {

	return GetInfoParams{}
}

// GetInfoParams contains all the bound params for the get info operation
// typically these are obtained from a http.Request
//
// swagger:parameters GetInfo
type GetInfoParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*The environment hostname
	  Required: true
	  In: query
	*/
	Host string
	/*The manifest key
	  Required: true
	  In: query
	*/
	Key string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewGetInfoParams() beforehand.
func (o *GetInfoParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	qHost, qhkHost, _ := qs.GetOK("host")
	if err := o.bindHost(qHost, qhkHost, route.Formats); err != nil {
		res = append(res, err)
	}

	qKey, qhkKey, _ := qs.GetOK("key")
	if err := o.bindKey(qKey, qhkKey, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindHost binds and validates parameter Host from query.
func (o *GetInfoParams) bindHost(rawData []string, hasKey bool, formats strfmt.Registry) error {
	if !hasKey {
		return errors.Required("host", "query")
	}
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// AllowEmptyValue: false
	if err := validate.RequiredString("host", "query", raw); err != nil {
		return err
	}

	o.Host = raw

	return nil
}

// bindKey binds and validates parameter Key from query.
func (o *GetInfoParams) bindKey(rawData []string, hasKey bool, formats strfmt.Registry) error {
	if !hasKey {
		return errors.Required("key", "query")
	}
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// AllowEmptyValue: false
	if err := validate.RequiredString("key", "query", raw); err != nil {
		return err
	}

	o.Key = raw

	return nil
}
