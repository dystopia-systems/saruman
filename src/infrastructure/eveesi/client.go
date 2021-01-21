package eveesi

import (
	"github.com/antihax/goesi"
	"net/http"
	"saruman/src"
	"saruman/src/service"
)

func InitClient() {
	var config = service.GetConfig()
	var httpClient = http.Client{}

	ctx := src.GetAppContext()
	ctx.ESI = goesi.NewAPIClient(&httpClient, config.UserAgent)
	ctx.SSOAuthenticator = goesi.NewSSOAuthenticatorV2(&httpClient, config.ClientID, config.SecretKey, config.RedirectUrl, config.Scopes)
}

func GetClient() *goesi.APIClient {
	return src.GetAppContext().ESI
}