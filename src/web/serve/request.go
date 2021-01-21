package serve

import (
	"github.com/dystopia-systems/alaskalog"
	"github.com/gorilla/mux"
	"net/http"
	"saruman/src/consts"
	"saruman/src/web/handlers"
	"saruman/src/web/middleware"
)

func SetupRoutes() *mux.Router {
	r := mux.NewRouter()

	alaskalog.Logger.Info("Registering routes...")

	r.Path(consts.BasePath).HandlerFunc(handlers.IndexGet)

	api := r.PathPrefix("/v1/api").Subrouter()

	apiKeyRoute := api.PathPrefix(consts.ApiKeyPath).Subrouter()
	apiKeyRoute.Methods(http.MethodGet).Path(consts.BasePath).HandlerFunc(handlers.ApiKeyGet)

	apiConfigRoute := api.PathPrefix(consts.ConfigPath).Subrouter()
	apiConfigRoute.Methods(http.MethodGet).Path(consts.IdVar).HandlerFunc(handlers.ConfigGet)
	apiConfigRoute.Methods(http.MethodPost).Path(consts.IdVar).HandlerFunc(handlers.ConfigPost)
	apiConfigRoute.Methods(http.MethodPut).Path(consts.IdVar).HandlerFunc(handlers.ConfigPut)

	apiQuoteRoute := api.PathPrefix(consts.QuotePath).Subrouter()
	yahooQuoteRoute := apiQuoteRoute.PathPrefix(consts.YahooPath).Subrouter()
	yahooQuoteRoute.Methods(http.MethodGet).PathPrefix(consts.IdVar).HandlerFunc(handlers.QuoteYahooGet)

	api.Use(middleware.AuthorizeApiKey)
	api.Use(middleware.VerifyContentType)

	return r
}
