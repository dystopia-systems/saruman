package db

import "github.com/vectorman1/saruman/src/models"

func InitMigration() error {
	db := GetDb()

	err := db.AutoMigrate(&models.ApiKey{})

	if err != nil {
		return err
	}

	return nil
}