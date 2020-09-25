package mysql

import (
	"github.com/vectorman1/alaskalog"
	"github.com/vectorman1/saruman/src/models"
)

func InitMigration() error {
	db := GetDb()

	alaskalog.Logger.Infoln("Auto-migrating...")

	err := db.AutoMigrate(&models.ApiKey{})

	if err != nil {
		return err
	}

	return nil
}