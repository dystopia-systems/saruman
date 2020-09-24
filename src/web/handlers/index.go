package handlers

import (
	"github.com/vectorman1/saruman/src/consts"
	"net/http"
)

func IndexHandler(w http.ResponseWriter, r *http.Request)  {
	http.Redirect(w, r, consts.RedirectUrl, http.StatusMovedPermanently)
}
