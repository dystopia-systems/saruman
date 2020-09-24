package serve

import (
	"github.com/gorilla/mux"
	"github.com/vectorman1/saruman/src/web/middleware"
	"github.com/vectorman1/saruman/src/web/routes"
	"net/http"
)

func SetupRoutes() *mux.Router{
	r := mux.NewRouter()

	api := r.PathPrefix("/api/v1").Subrouter()

	registerGETRoutes(api)

	return r
}

func registerGETRoutes(api *mux.Router) {
	for route, handleFunc := range routes.GETRoutesMap {
		finalHandler := http.HandlerFunc(handleFunc)

		api.Handle(route,
			middleware.VerifyContentType(
					finalHandler))
	}
}
