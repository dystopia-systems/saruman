package service

import (
	"github.com/dystopia-systems/alaskalog"
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
	apiKey, err := mysql.CreateApiKey()

	if err != nil {
		alaskalog.Logger.Warnf("Failed to create Api-Key %+v", err)

		return nil, false
	}

	return apiKey, true
}