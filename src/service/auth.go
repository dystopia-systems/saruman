package service

import (
	"github.com/dystopia-systems/alaskalog"
	"github.com/google/uuid"
	"os"
	"saruman/src/core/db/mysql"
	"saruman/src/models"
)

func VerifyApiKey(key string) bool {
	apiKey := mysql.GetApiKey(key)

	if apiKey == nil {
		return false
	}

	return true
}

func CreateApiKey() (*models.ApiKey, bool) {
	apiKey, err := mysql.CreateApiKey(uuid.New().String())

	if err != nil {
		alaskalog.Logger.Warnf("Failed to create Api-Key %+v", err)

		return nil, false
	}

	return apiKey, true
}

func CreateInitialKey() (*models.ApiKey, bool) {
	apiKey, err := mysql.CreateApiKey(os.Getenv("INITIAL_API_KEY"))

	if err != nil {
		alaskalog.Logger.Warnf("Failed to create Api-Key %+v", err)

		return nil, false
	}

	return apiKey, true
}