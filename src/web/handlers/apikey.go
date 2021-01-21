package handlers

import (
	"encoding/json"
	"net/http"
	"saruman/src/service"
)

func ApiKeyGet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	key, success := service.CreateApiKey()

	if !success {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	res, _ := json.Marshal(key)
	_, _ = w.Write(res)
	return
}