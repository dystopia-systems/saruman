package models

import (
	"github.com/piquette/finance-go"
	"time"
)

type YahooQuote struct {
	ID uint64 `gorm:"primaryKey" json:"id"`

	finance.Quote

	DateAdded time.Time `json:"dateAdded"`
}
