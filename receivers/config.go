package receivers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/vectorman1/saruman/consts"
	"github.com/vectorman1/saruman/models"
	"github.com/vectorman1/saruman/service"
	"net/http"
)

func ConfigHandler(w http.ResponseWriter, r *http.Request) {
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
