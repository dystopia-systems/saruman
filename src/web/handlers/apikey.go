package handlers

import (
	"encoding/json"
	"github.com/alexedwards/scs/v2"
	"net/http"
	"saruman/src/service"
)

func ApiKeyBaseGetHandler(w http.ResponseWriter, r *http.Request, s *scs.SessionManager) {
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