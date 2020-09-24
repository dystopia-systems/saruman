package models

import "gorm.io/gorm"

type ApiKey struct {
	gorm.Model
	Key              string
	ConfigPermission bool
}
