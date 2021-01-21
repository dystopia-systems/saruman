package handlers

import (
	"net/http"
	"saruman/src/consts"
)

func IndexGet(w http.ResponseWriter, r *http.Request)  {
	http.Redirect(w, r, consts.RedirectUrl, http.StatusMovedPermanently)
}
