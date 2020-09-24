package models

import (
	"github.com/go-sql-driver/mysql"
	"time"
)

type ApiKey struct {
	ID               uint
	CreatedAt        time.Time
	UpdatedAt        time.Time
	DeletedAt        mysql.NullTime
	Key              string
	ConfigPermission bool
	Name             string
}
