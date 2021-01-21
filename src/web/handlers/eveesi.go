package handlers

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/alexedwards/scs/v2"
	"math/rand"
	"net/http"
	"saruman/src"
	"saruman/src/models"
	"saruman/src/service"
)

func EveEsiGetCallback(w http.ResponseWriter, r *http.Request, s *scs.SessionManager) {
	code := r.FormValue("code")
	state := r.FormValue("state")

	if s.GetString(r.Context(), "state") != state {
		w.WriteHeader(http.StatusInternalServerError)

		res, _ := json.Marshal(models.Error{
			Code:    http.StatusInternalServerError,
			Message: "Invalid state",
		})

		_, _ = w.Write(res)
		_ = s.Destroy(r.Context())

		token, err := src.GetAppContext().SSOAuthenticator.TokenExchange(code)

		if err != nil {
			res, _ := json.Marshal(models.Error{
				Code:    http.StatusInternalServerError,
				Message: fmt.Sprintf("%v", err),
			})

			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write(res)
		}

		tokenSrc := src.GetAppContext().SSOAuthenticator.TokenSource(token)

		v, err := src.GetAppContext().SSOAuthenticator.Verify(tokenSrc)

		if err != nil {
			res, _ := json.Marshal(models.Error{
				Code:    http.StatusInternalServerError,
				Message: fmt.Sprintf("%v", err),
			})

			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write(res)
		}

		s.Put(r.Context(), "characters", v)
		_, _, err = s.Commit(r.Context())

		if err != nil {
			res, _ := json.Marshal(models.Error{
				Code:    http.StatusInternalServerError,
				Message: fmt.Sprintf("%v", err),
			})

			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write(res)
		}

		http.Redirect(w, r, "/account", http.StatusMovedPermanently)
		return
	}
}

func EveEsiGetBase(w http.ResponseWriter, r *http.Request, s *scs.SessionManager) {
	// Generate a random state string
	b := make([]byte, 16)
	rand.Read(b)
	state := base64.URLEncoding.EncodeToString(b)

	// Save the state on the session
	s.Put(r.Context(), "state", state)
	_, _, err := s.Commit(r.Context())

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Generate the SSO URL with the state string
	url := src.GetAppContext().SSOAuthenticator.AuthorizeURL(state, true, service.GetConfig().Scopes)

	// Send the user to the URL
	http.Redirect(w, r, url, http.StatusMovedPermanently)
}