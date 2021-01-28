package mysql

import (
	"github.com/dystopia-systems/alaskalog"
	"github.com/google/uuid"
	"saruman/src/models"
	"time"
)

type ApiKeyRepository struct

func (*ApiKeyRepository) Get(key string) *models.ApiKey {
	var res []*models.ApiKey

	context := GetDb()

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
	context := GetDb()

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
