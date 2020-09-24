package auth

import (
	"github.com/vectorman1/alaskalog"
	"github.com/vectorman1/saruman/src/core/db"
	"github.com/vectorman1/saruman/src/models"
	"time"
)

func GetApiKey(key string) *models.ApiKey {
	var res *models.ApiKey

	context := db.GetDb()

	context.First(&models.ApiKey{Key: key})

	if context.Error != nil {
		alaskalog.Logger.Warnf("Failed to execute query %v", context.Error)
		return &models.ApiKey{}
	}

	return res
}

func CreateApiKey(key string) (*models.ApiKey, error) {
	context := db.GetDb()

	apiKey := models.ApiKey{
		Key:              key,
		ConfigPermission: false,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	result := context.Create(&apiKey)

	if result.Error != nil {
		alaskalog.Logger.Warnf("Failed to execute query %v", context.Error)
		return nil, result.Error
	}

	return &apiKey, nil
}
