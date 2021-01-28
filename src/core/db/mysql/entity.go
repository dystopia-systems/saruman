package mysql

import (
	"github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
	"github.com/piquette/finance-go"
	"gorm.io/gorm"
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

type Historical struct {
	gorm.Model

	Candles []*finance.ChartBar

	RequestUuid uuid.UUID
}


type PriceQuote struct {
	ID uint64 `gorm:"primaryKey" json:"id"`

	finance.Quote

	DateAdded time.Time `json:"dateAdded"`
}

