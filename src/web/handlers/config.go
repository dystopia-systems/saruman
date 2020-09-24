package handlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/vectorman1/saruman/src/consts"
	"github.com/vectorman1/saruman/src/models"
	"github.com/vectorman1/saruman/src/service"
	"github.com/vectorman1/saruman/src/web/requestmodels"
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

			errbytes, _ := json.Marshal(models.Error{
				Code:    http.StatusInternalServerError,
				Message: "Error reading config.",
			})

			_,_ = w.Write(errbytes)

			return
		}

		w.WriteHeader(http.StatusOK)
		_, _ = w.Write(res)

		break
	case "einstein":
		res, err := service.ReadFile(consts.EINSTEIN_CONFIG_PATH)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
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

func ConfigPostHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	var req requestmodels.CreateApiKeyRequest

	err := json.NewDecoder(r.Body).Decode(&req)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	key, success := service.CreateApiKey(req.Key)

	if !success {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	res, _ := json.Marshal(key)
	_, _ = w.Write(res)
	return
}