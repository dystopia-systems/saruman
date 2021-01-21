package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"saruman/src/consts"
	"saruman/src/models"
	"saruman/src/service"
)

func ConfigAppGetHandler(w http.ResponseWriter, r *http.Request) {
	pathParams := mux.Vars(r)

	w.Header().Set("Content-type", "application/json")

	appId := pathParams["id"]

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

func ConfigAppPostHandler(w http.ResponseWriter, r *http.Request) {
	pathParams := mux.Vars(r)

	w.Header().Set("Content-type", "application/json")

	appId := pathParams["id"]

	if appId == "" {
		w.WriteHeader(http.StatusBadRequest)
		res, _ := json.Marshal(models.Error{Code: http.StatusBadRequest, Message: "Provide app-id"})
		_, _ = w.Write(res)

		return
	}

	_, err := service.ReadFile(fmt.Sprintf(consts.CONFIG_PATH, appId))

	if err == nil {
		w.WriteHeader(http.StatusConflict)

		res, _ := json.Marshal(models.Error{Code: http.StatusBadRequest, Message: "App-id already exists"})
		_, _ = w.Write(res)

		return
	}

	bodyBytes, _ := ioutil.ReadAll(r.Body)

	_, createErr := service.WriteFile(fmt.Sprintf(consts.CONFIG_PATH, appId), bodyBytes)

	if createErr != nil {
		w.WriteHeader(http.StatusInternalServerError)

		errBytes, _ := json.Marshal(models.Error{
			Code:    http.StatusInternalServerError,
			Message: fmt.Sprintf("Error writing config. %v", err),
		})

		_,_ = w.Write(errBytes)

		return
	}

	w.WriteHeader(http.StatusOK)

	return
}

func ConfigAppPutHandler(w http.ResponseWriter, r *http.Request) {
	pathParams := mux.Vars(r)

	w.Header().Set("Content-type", "application/json")

	appId := pathParams["id"]

	if appId == "" {
		w.WriteHeader(http.StatusBadRequest)
		res, _ := json.Marshal(models.Error{Code: http.StatusBadRequest, Message: "Provide app-id"})
		_, _ = w.Write(res)

		return
	}

	_, err := service.ReadFile(fmt.Sprintf(consts.CONFIG_PATH, appId))

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)

		res, _ := json.Marshal(models.Error{Code: http.StatusBadRequest, Message: "App-id doesn't exist"})
		_, _ = w.Write(res)

		return
	}

	bodyBytes, _ := ioutil.ReadAll(r.Body)

	_, createErr := service.WriteFile(fmt.Sprintf(consts.CONFIG_PATH, appId), bodyBytes)

	if createErr != nil {
		w.WriteHeader(http.StatusInternalServerError)

		errBytes, _ := json.Marshal(models.Error{
			Code:    http.StatusInternalServerError,
			Message: fmt.Sprintf("Error writing config. %v", createErr),
		})

		_,_ = w.Write(errBytes)

		return
	}

	w.WriteHeader(http.StatusOK)

	return
}