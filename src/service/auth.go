package service

import (
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
		return nil, false
	}

	return apiKey, true
}