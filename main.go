package main

import (
	"context"
	"github.com/alexedwards/scs/v2"
	"github.com/dystopia-systems/alaskalog"
	"golang.org/x/sync/errgroup"
	"net/http"
	"saruman/src/core/db/mysql"
	"saruman/src/service"
	"saruman/src/web/routes"
	"saruman/src/web/serve"
	"time"
)

func main(){
	service.InitConfig()

	sessionManager := scs.New()
	sessionManager.Lifetime = time.Hour

	routes.InitializeRouteMappings()

	dbErr := mysql.InitDb()

	if dbErr != nil {
		alaskalog.Logger.Fatalln("Failed to initialize db.")
	}

	mux := serve.SetupRoutes()

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

	alaskalog.Logger.Fatal(http.ListenAndServe(":3000", mux))
}