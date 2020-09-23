package routes

import (
	"github.com/vectorman1/saruman/consts"
	"github.com/vectorman1/saruman/receivers"
	"net/http"
)

var GETRoutes = make(map[string]func(http.ResponseWriter, *http.Request))

var ROUTES = [...]string {
	consts.IndexUrl,
	consts.ConfigUrl,
}

var HANDLERS = [...]func(http.ResponseWriter, *http.Request) {
	receivers.IndexHandler,
	receivers.ConfigHandler,
}

func InitializeMap() {
	for i, route := range ROUTES {
		GETRoutes[route] = HANDLERS[i]
	}
}