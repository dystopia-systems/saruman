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
		res, _ := json.Marshal(models.Error{Code: http.StatusBadRequest, Message: "Invalid app-id provided"})
		_, _ = w.Write(res)
		return
	}

	switch appId {
	case "gosniff":
		res, err := service.ReadFile(consts.GOSNIFF_CONFIG_PATH)

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

		break
	case "einstein":
		res, err := service.ReadFile(consts.EINSTEIN_CONFIG_PATH)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			errBytes, _ := json.Marshal(err)
			_, _ = w.Write(errBytes)
			return
		}

		w.WriteHeader(http.StatusOK)
		_, _ = w.Write(res)

		break
	default:
		w.WriteHeader(http.StatusNotFound)
		break
	}
}
