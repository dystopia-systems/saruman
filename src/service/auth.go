package service

import (
	"github.com/vectorman1/saruman/src/core/auth"
	"github.com/vectorman1/saruman/src/models"
)

func VerifyApiKey(key string) bool {
	apiKey := auth.GetApiKey(key)

	if apiKey == nil {
		return false
	}

	return true
}

func CreateApiKey(key string) (*models.ApiKey, bool) {
	apiKey, err := auth.CreateApiKey(key)

	if err != nil {
		return nil, false
	}

	return apiKey, true
}