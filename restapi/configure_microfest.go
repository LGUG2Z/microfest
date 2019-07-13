// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"log"
	"net/http"
	"os"

	"github.com/LGUG2Z/microfest/handlers"
	"github.com/LGUG2Z/microfest/models"
	"github.com/LGUG2Z/microfest/restapi/operations"
	interpose "github.com/carbocation/interpose/middleware"
	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
)

//go:generate swagger generate server --target ../../microfest --name Microfest --spec ../swagger.yml --principal models.Principal

func configureFlags(api *operations.MicrofestAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.MicrofestAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError
	api.GetHealthcheckHandler = operations.GetHealthcheckHandlerFunc(handlers.GetHealthcheck)
	api.GetConfigurationHandler = operations.GetConfigurationHandlerFunc(handlers.GetConfiguration)
	api.PostConfigurationHandler = operations.PostConfigurationHandlerFunc(handlers.PostConfiguration)
	api.GetManifestHandler = operations.GetManifestHandlerFunc(handlers.GetManifest)
	api.PostManifestHandler = operations.PostManifestHandlerFunc(handlers.PostManifest)
	api.PutManifestHandler = operations.PutManifestHandlerFunc(handlers.PutManifest)

	ApiKey := os.Getenv("API_KEY")
	if len(ApiKey) == 0 {
		log.Fatal("API_KEY must be set")
	}

	handlers.BoltPath = os.Getenv("BOLT_PATH")
	if len(handlers.BoltPath) == 0 {
		log.Fatal("BOLT_PATH must be set")
	}

	api.JSONConsumer = runtime.JSONConsumer()
	api.JSONProducer = runtime.JSONProducer()

	// Applies when the "X-API-KEY" header is set
	api.APIKeyHeaderAuth = func(token string) (*models.Principal, error) {
		if token == ApiKey {
			prin := models.Principal(token)
			return &prin, nil
		}
		api.Logger("Access attempt with incorrect api key auth: %s", token)
		return nil, errors.New(401, "incorrect api key auth")
	}

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix"
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	logViaLogrus := interpose.NegroniLogrus()
	return logViaLogrus(handler)
}
