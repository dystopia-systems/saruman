package routes

import (
	"github.com/dystopia-systems/alaskalog"
	"net/http"
	"saruman/src/consts"
	"saruman/src/web/handlers"
)

var postPublicRoutes []string
var postSecureRoutes []string

var postPublicHandlers []func(http.ResponseWriter, *http.Request)
var postSecureHandlers []func(http.ResponseWriter, *http.Request)

var getPublicRoutes = []string {
	consts.IndexUrl,
}
var getPublicHandlers = []func(http.ResponseWriter, *http.Request) {
	handlers.IndexHandler,
}

var getSecureRoutes =  []string {
	consts.ConfigAppUrl,
	consts.ApiKeyBaseUrl,
}
var getSecureHandlers = []func(http.ResponseWriter, *http.Request) {
	handlers.ConfigAppGetHandler,
	handlers.ApiKeyBaseGetHandler,
}

var (
	PostPublicMap map[string]func(http.ResponseWriter, *http.Request)
	GetPublicMap  map[string]func(http.ResponseWriter, *http.Request)
	PostSecureMap map[string]func(http.ResponseWriter, *http.Request)
	GetSecureMap  map[string]func(http.ResponseWriter, *http.Request)
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
	handlers []func(http.ResponseWriter, *http.Request)) map[string]func(http.ResponseWriter, *http.Request) {
	res := make(map[string]func(http.ResponseWriter, *http.Request))

	for i, route := range routes {
		res[route] = handlers[i]
	}

	return res
}
