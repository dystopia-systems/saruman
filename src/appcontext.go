package src

import "sync"
import "github.com/antihax/goesi"

type appContext struct {
	sync.RWMutex

	ESI *goesi.APIClient
	SSOAuthenticator *goesi.SSOAuthenticator
}

var ctx appContext

func GetAppContext() *appContext {
	return &ctx
}

