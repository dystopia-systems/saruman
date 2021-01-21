package handlers

import (
	"github.com/alexedwards/scs/v2"
	"github.com/vectorman1/saruman/src/consts"
	"net/http"
)

func IndexHandler(w http.ResponseWriter, r *http.Request, s *scs.SessionManager)  {
	http.Redirect(w, r, consts.RedirectUrl, http.StatusMovedPermanently)
}
