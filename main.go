package main

import (
	"context"
	"github.com/dystopia-systems/alaskalog"
	"github.com/gorilla/mux"
	"golang.org/x/sync/errgroup"
	"net/http"
	"os"
	"saruman/src/core/db/mysql"
	"saruman/src/service"
	"saruman/src/web/serve"
)

func main(){
	r := serve.SetupRoutes()
	http.Handle("/", r)

	_ = r.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
		t, perr := route.GetPathTemplate()
		m, merr := route.GetMethods()

		if perr != nil {
			alaskalog.Logger.Warnln(perr)

			return perr
		}

		if merr != nil {
			return nil
		}

		alaskalog.Logger.Infof("%s: %s", m, t)

		return nil
	})

	g, _ := errgroup.WithContext(context.Background())

	g.Go(mysql.InitDb)

	err := g.Wait()

	if err != nil {
		alaskalog.Logger.Fatalf("Failed to init db. %v", err)
		return
	}

	migrationErr := mysql.InitMigration()

	if migrationErr != nil {
		alaskalog.Logger.Fatalf("Failed to migrate db. %v", migrationErr)
		return
	}

	initKey := os.Getenv("INITIAL_API_KEY")
	success := service.VerifyApiKey(initKey)
	if !success {
		key, ok := service.CreateInitialKey()

		if !ok {
			alaskalog.Logger.Fatalf("failed to create initial key")
		}

		alaskalog.Logger.Infoln("initial key created: ", key.Key)
	}

	alaskalog.Logger.Infoln("Saruman is now running. Listening on port :3000...")

	alaskalog.Logger.Fatal(http.ListenAndServe(":3000", r))
}