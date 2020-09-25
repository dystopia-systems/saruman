package main

import (
	"context"
	"github.com/vectorman1/alaskalog"
	"github.com/vectorman1/saruman/src/core/db/mysql"
	"github.com/vectorman1/saruman/src/web/routes"
	"github.com/vectorman1/saruman/src/web/serve"
	"golang.org/x/sync/errgroup"
	"net/http"
)

func main(){
	routes.InitializeMap()
	r := serve.SetupRoutes()

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