package serve

import (
	"github.com/gorilla/mux"
	"github.com/vectorman1/alaskalog"
	"github.com/vectorman1/saruman/src/web/middleware"
	"github.com/vectorman1/saruman/src/web/routes"
	"net/http"
)

func SetupRoutes() *mux.Router{
	r := mux.NewRouter()

	api := r.PathPrefix("/api/v1").Subrouter()

	registerGETRoutes(api)
	registerPOSTRoutes(api)

	return r
}

func registerGETRoutes(api *mux.Router) {
	alaskalog.Logger.Infoln("Registering GET route/handler mapping...")

	for route, handleFunc := range routes.GETRoutesMap {
		finalHandler := http.HandlerFunc(handleFunc)
		alaskalog.Logger.Infof("Registering route:%s", route)
		api.Handle(route,
			middleware.VerifyContentType(
				middleware.AuthorizeApiKey(finalHandler)))
	}
}

func registerPOSTRoutes(api *mux.Router) {
	alaskalog.Logger.Infoln("Registering POST route/handler mapping...")

	for route, handleFunc := range routes.POSTRoutesMap {
		finalHandler := http.HandlerFunc(handleFunc)
		alaskalog.Logger.Infof("Registering route:%s", route)
		api.Handle(route,
			middleware.VerifyContentType(
				middleware.AuthorizeApiKey(finalHandler)))
	}
}