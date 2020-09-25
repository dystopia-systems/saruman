package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/vectorman1/saruman/src/consts"
	"github.com/vectorman1/saruman/src/models"
	"github.com/vectorman1/saruman/src/service"
	"net/http"
)

func ConfigAppGetHandler(w http.ResponseWriter, r *http.Request) {
	pathParams := mux.Vars(r)

	w.Header().Set("Content-type", "application/json")

	appId := pathParams["app-id"]

	if appId == "" {
		w.WriteHeader(http.StatusBadRequest)
		res, _ := json.Marshal(models.Error{Code: http.StatusBadRequest, Message: "Provide app-id"})
		_, _ = w.Write(res)
		return
	}

	bytes, err := service.ReadFile(fmt.Sprintf("%s/%s.json",consts.CONFIG_STORE_PATH, appId))

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		res, _ := json.Marshal(models.Error{Code: http.StatusNotFound, Message: fmt.Sprintf("%v", err)})
		_, _ = w.Write(res)
		return
	}

	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(bytes)
	return
}
