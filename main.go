package main

import (
	"github.com/vectorman1/alaskalog"
	"github.com/vectorman1/saruman/src/core/db"
	"github.com/vectorman1/saruman/src/web/routes"
	"github.com/vectorman1/saruman/src/web/serve"
	"net/http"
)

func main(){
	routes.InitializeMap()

	r := serve.SetupRoutes()
	dbErr := db.InitDb()

	if dbErr != nil {
		alaskalog.Logger.Fatalf("Failed to init db. %v", dbErr)
		return
	}

	migrationErr := db.InitMigration()

	if migrationErr != nil {
		alaskalog.Logger.Fatalf("Failed to migrate db. %v", dbErr)
		return
	}


	alaskalog.Logger.Infoln("Saruman is now running. Listening on port :3000...")

	alaskalog.Logger.Fatal(http.ListenAndServe(":3000", r))
}