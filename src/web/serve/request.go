package serve

import (
	"github.com/alexedwards/scs/v2"
	"github.com/dystopia-systems/alaskalog"
	"github.com/gorilla/mux"
	"net/http"
	"saruman/src/web/handlers"
	"saruman/src/web/middleware"
	"time"
)

var sessionManager *scs.SessionManager

func SetupRoutes() *mux.Router {
	sessionManager = scs.New()
	sessionManager.Lifetime = time.Hour

	r := mux.NewRouter()

	alaskalog.Logger.Info("Registering secure routes...")

	api := r.PathPrefix("/v1/api").Subrouter()

	apiKeyRoute := api.PathPrefix("/api-key/")
	apiKeyRoute.Methods(http.MethodGet).Path("/").HandlerFunc(handlers.ApiKeyBaseGetHandler)

	apiConfigRoute := api.PathPrefix("/config/")
	apiConfigRoute.Methods(http.MethodGet).Path("/{id}/").HandlerFunc(handlers.ConfigAppGetHandler)
	apiConfigRoute.Methods(http.MethodPost).Path("/{id}/").HandlerFunc(handlers.ConfigAppPostHandler)
	apiConfigRoute.Methods(http.MethodPut).Path("/{id}/").HandlerFunc(handlers.ConfigAppPutHandler)

	api.Use(middleware.AuthorizeApiKey)
	api.Use(middleware.VerifyContentType)

	return r
}
