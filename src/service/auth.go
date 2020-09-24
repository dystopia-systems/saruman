package service

import (
	"github.com/vectorman1/saruman/src/core/auth"
)

func VerifyApiKey(key string) bool {
	apiKey := auth.GetApiKey(key)

	if apiKey == nil {
		return false
	}

	return true
}

func CreateApiKey(key string) bool {
	err := auth.CreateApiKey(key)

	if err != nil {
		return false
	}

	return true
}