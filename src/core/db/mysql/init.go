package mysql

import (
	"github.com/dystopia-systems/alaskalog"
	"saruman/src/models"
)

func InitMigration() error {
	db := GetDb()

	alaskalog.Logger.Infoln("Auto-migrating...")

	_ = db.AutoMigrate(&models.ApiKey{})
	_ = db.AutoMigrate(&models.PriceQuote{})

	return nil
}