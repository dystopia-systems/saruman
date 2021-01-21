package handlers

import (
	"github.com/alexedwards/scs/v2"
	"net/http"
	"saruman/src/consts"
)

func IndexHandler(w http.ResponseWriter, r *http.Request, s *scs.SessionManager)  {
	http.Redirect(w, r, consts.RedirectUrl, http.StatusMovedPermanently)
}
