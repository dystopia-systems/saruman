package service

import (
	"github.com/vectorman1/saruman/src/core/db/mysql"
	"github.com/vectorman1/saruman/src/models"
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