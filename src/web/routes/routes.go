package routes

import (
	"github.com/vectorman1/alaskalog"
	"github.com/vectorman1/saruman/src/consts"
	"github.com/vectorman1/saruman/src/web/handlers"
	"net/http"
)

var GETRoutesMap = make(map[string]func(http.ResponseWriter, *http.Request))
var POSTRoutesMap = make(map[string]func(w http.ResponseWriter, r *http.Request))

var GetRouters = [...]string {
	consts.IndexUrl,
	consts.ConfigAppUrl,
	consts.ApiKeyBaseUrl,
}

var PostRouters = [...]string {
}

var GetHandlers = [...]func(http.ResponseWriter, *http.Request) {
	handlers.IndexHandler,
	handlers.ConfigAppGetHandler,
	handlers.ApiKeyBaseGetHandler,
}

var PostHandlers = [...]func(w http.ResponseWriter, r *http.Request) {
}

func InitializeMap() {
	alaskalog.Logger.Infoln("Initializing GET route/handler mapping...")
	for i, route := range GetRouters {
		GETRoutesMap[route] = GetHandlers[i]
	}
	alaskalog.Logger.Infoln("Initializing POST route/handler mapping...")
	for i, route := range PostRouters {
		POSTRoutesMap[route] = PostHandlers[i]
	}
}