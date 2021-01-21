package mysql

import (
	"github.com/dystopia-systems/alaskalog"
	"saruman/src/models"
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