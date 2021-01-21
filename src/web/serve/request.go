package serve

import (
	"github.com/alexedwards/scs/v2"
	"github.com/dystopia-systems/alaskalog"
	"github.com/gorilla/mux"
	"net/http"
	"saruman/src/web/middleware"
	"saruman/src/web/routes"
	"time"
)

var sessionManager *scs.SessionManager

func SetupRoutes() *mux.Router {
	sessionManager = scs.New()
	sessionManager.Lifetime = time.Hour

	alaskalog.Logger.Info("Registering public routes...")

	r := mux.NewRouter().StrictSlash(false)

	publicApi := r.PathPrefix("/").Subrouter()

	registerRoutes(publicApi, routes.GetPublicMap, "GET")
	registerRoutes(publicApi, routes.PostPublicMap, "POST")

	alaskalog.Logger.Info("Registering secure routes...")

	secureApi := r.PathPrefix("/v1/api").Subrouter()

	registerRoutes(secureApi, routes.GetSecureMap, "GET")
	registerRoutes(secureApi, routes.PostSecureMap, "POST")

	secureApi.Use(middleware.AuthorizeApiKey)
	secureApi.Use(middleware.VerifyContentType)

	return r
}

func registerRoutes(
		r *mux.Router,
		routes map[string]func(w http.ResponseWriter, r *http.Request, s *scs.SessionManager),
		method string,
) {
	for route, handlerFunc := range routes {
		handler := sessionManager.LoadAndSave(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			handlerFunc(w, r, sessionManager)
		}))

		alaskalog.Logger.Infof("Registering route:%s:%s", method, route)

		r.Methods(method).Handler(handler)
	}
}
