package mysql

import (
	"github.com/dystopia-systems/alaskalog"
	"saruman/src/models"
	"time"
)

func GetApiKey(key string) *models.ApiKey {
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

func CreateApiKey(key string) (*models.ApiKey, error) {
	context := GetDb()

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
