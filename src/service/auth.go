package service

import (
	"github.com/vectorman1/saruman/src/core/apikey"
	"github.com/vectorman1/saruman/src/models"
)

func VerifyApiKey(key string) bool {
	apiKey := apikey.GetApiKey(key)

	if apiKey == nil {
		return false
	}

	return true
}

func CreateApiKey() (*models.ApiKey, bool) {
	apiKey, err := apikey.CreateApiKey()

	if err != nil {
		return nil, false
	}

	return apiKey, true
}