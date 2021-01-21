package serve

import (
	"github.com/alexedwards/scs/v2"
	"github.com/dystopia-systems/alaskalog"
	"github.com/gorilla/mux"
	"net/http"
	"saruman/src/consts"
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

	r.Path(consts.BasePath).HandlerFunc(handlers.IndexHandler)

	api := r.PathPrefix("/v1/api").Subrouter()

	apiKeyRoute := api.PathPrefix(consts.ApiKeyPath)
	apiKeyRoute.Methods(http.MethodGet).Path(consts.BasePath).HandlerFunc(handlers.ApiKeyBaseGetHandler)

	apiConfigRoute := api.PathPrefix(consts.ConfigPath)
	apiConfigRoute.Methods(http.MethodGet).Path(consts.IdVar).HandlerFunc(handlers.ConfigAppGetHandler)
	apiConfigRoute.Methods(http.MethodPost).Path(consts.IdVar).HandlerFunc(handlers.ConfigAppPostHandler)
	apiConfigRoute.Methods(http.MethodPut).Path(consts.IdVar).HandlerFunc(handlers.ConfigAppPutHandler)

	api.Use(middleware.AuthorizeApiKey)
	api.Use(middleware.VerifyContentType)

	return r
}
