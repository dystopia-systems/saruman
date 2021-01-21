package main

import (
	"context"
	"github.com/dystopia-systems/alaskalog"
	"github.com/gorilla/mux"
	"golang.org/x/sync/errgroup"
	"net/http"
	"saruman/src/core/db/mysql"
	"saruman/src/web/serve"
)

func main(){
	r := serve.SetupRoutes()
	http.Handle("/", r)

	_ = r.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
		t, perr := route.GetPathTemplate()

		if perr != nil {
			alaskalog.Logger.Warnln(perr)

			return perr
		}

		alaskalog.Logger.Infof("%s", t)

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

	alaskalog.Logger.Infoln("Saruman is now running. Listening on port :3000...")

	alaskalog.Logger.Fatal(http.ListenAndServe(":3000", r))
}