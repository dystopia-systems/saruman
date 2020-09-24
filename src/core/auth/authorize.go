package auth

import (
	"github.com/vectorman1/alaskalog"
	"github.com/vectorman1/saruman/src/core/db"
	"github.com/vectorman1/saruman/src/models"
	"gorm.io/gorm"
	"time"
)

func GetApiKey(key string) *models.ApiKey {
	var res *models.ApiKey

	context := db.GetDb()

	context.Where("key = ?", key).First(res)

	if context.Error != nil {
		alaskalog.Logger.Warnf("Failed to execute query %v", context.Error)
		return &models.ApiKey{}
	}

	return res
}

func CreateApiKey(key string) error {
	context := db.GetDb()

	apiKey := models.ApiKey{
		Model: gorm.Model{
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		Key:              key,
		ConfigPermission: false,
	}

	result := context.Create(&apiKey)

	if result.Error != nil {
		alaskalog.Logger.Warnf("Failed to execute query %v", context.Error)
		return result.Error
	}

	return nil
}
