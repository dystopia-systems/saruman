package serve

import (
	"github.com/gorilla/mux"
	"github.com/vectorman1/saruman/routes"
)

func SetupRoutes() *mux.Router{
	r := mux.NewRouter()

	api := r.PathPrefix("/api/v1").Subrouter()

	registerGETRoutes(api)

	return r
}

func registerGETRoutes(api *mux.Router) {
	for route, handleFunc := range routes.GETRoutes {
		api.HandleFunc(route, handleFunc)
	}
}
