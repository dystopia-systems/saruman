package apikey

import (
	"github.com/google/uuid"
	"github.com/vectorman1/alaskalog"
	"github.com/vectorman1/saruman/src/core/db"
	"github.com/vectorman1/saruman/src/models"
	"time"
)

func GetApiKey(key string) *models.ApiKey {
	var res []*models.ApiKey

	context := db.GetDb()

	context.Find(&res)

	for _, apiKey := range res {
		if apiKey.Key == key {
			return apiKey
		}
	}

	if context.Error != nil {
		alaskalog.Logger.Warnf("Failed to execute query %v", context.Error)
		return &models.ApiKey{}
	}

	return nil
}

func CreateApiKey() (*models.ApiKey, error) {
	context := db.GetDb()

	apiKey := models.ApiKey{
		Key:              uuid.New().String(),
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
