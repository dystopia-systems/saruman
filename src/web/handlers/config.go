package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/alexedwards/scs/v2"
	"github.com/gorilla/mux"
	"net/http"
	"saruman/src/consts"
	"saruman/src/models"
	"saruman/src/service"
)

func ConfigAppGetHandler(w http.ResponseWriter, r *http.Request, s *scs.SessionManager) {
	pathParams := mux.Vars(r)

	w.Header().Set("Content-type", "application/json")

	appId := pathParams["app-id"]

	if appId == "" {
		w.WriteHeader(http.StatusBadRequest)
		res, _ := json.Marshal(models.Error{Code: http.StatusBadRequest, Message: "Provide app-id"})
		_, _ = w.Write(res)
		return
	}

	res, err := service.ReadFile(fmt.Sprintf(consts.CONFIG_PATH, appId))

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)

		errBytes, _ := json.Marshal(models.Error{
			Code:    http.StatusInternalServerError,
			Message: fmt.Sprintf("Error reading config. %v", err),
		})

		_,_ = w.Write(errBytes)

		return
	}

	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(res)

	return
}

