package handlers

import (
	"github.com/LGUG2Z/microfest/restapi/operations"
	"github.com/go-openapi/runtime/middleware"
)

func GetHealthcheck(params operations.GetHealthcheckParams) middleware.Responder {
	return operations.NewGetHealthcheckOK()
}
