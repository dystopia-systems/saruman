package routes

import (
	"github.com/alexedwards/scs/v2"
	"github.com/dystopia-systems/alaskalog"
	"net/http"
	"saruman/src/consts"
	"saruman/src/web/handlers"
)

var postPublicRoutes []string
var postSecureRoutes []string

var postPublicHandlers []func(http.ResponseWriter, *http.Request, *scs.SessionManager)
var postSecureHandlers []func(http.ResponseWriter, *http.Request, *scs.SessionManager)

var getPublicRoutes = []string {
	consts.IndexUrl,
	consts.EveEsiBaseUrl,
	consts.EveEsiCallBackUrl,
}
var getPublicHandlers = []func(http.ResponseWriter, *http.Request, *scs.SessionManager) {
	handlers.IndexHandler,
	handlers.EveEsiGetBase,
	handlers.EveEsiGetCallback,
}

var getSecureRoutes =  []string {
	consts.ConfigAppUrl,
	consts.ApiKeyBaseUrl,
}
var getSecureHandlers = []func(http.ResponseWriter, *http.Request, *scs.SessionManager) {
	handlers.ConfigAppGetHandler,
	handlers.ApiKeyBaseGetHandler,
}

var (
	PostPublicMap map[string]func(http.ResponseWriter, *http.Request, *scs.SessionManager)
	GetPublicMap  map[string]func(http.ResponseWriter, *http.Request, *scs.SessionManager)
	PostSecureMap map[string]func(http.ResponseWriter, *http.Request, *scs.SessionManager)
	GetSecureMap  map[string]func(http.ResponseWriter, *http.Request, *scs.SessionManager)
)

func InitializeRouteMappings() {
	alaskalog.Logger.Infoln("Initializing route/handler mapping...")

	GetSecureMap = initializeMap(getSecureRoutes, getSecureHandlers)
	PostSecureMap = initializeMap(postSecureRoutes, postSecureHandlers)
	GetPublicMap = initializeMap(getPublicRoutes, getPublicHandlers)
	PostPublicMap = initializeMap(postPublicRoutes, postPublicHandlers)
}

func initializeMap(
	routes []string,
	handlers []func(http.ResponseWriter, *http.Request, *scs.SessionManager)) map[string]func(http.ResponseWriter, *http.Request, *scs.SessionManager){
	res := make(map[string]func(http.ResponseWriter, *http.Request, *scs.SessionManager))

	for i, route := range routes {
		res[route] = handlers[i]
	}

	return res
}
