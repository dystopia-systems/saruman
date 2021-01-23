package models

import (
	"github.com/google/uuid"
	"github.com/piquette/finance-go"
	"gorm.io/gorm"
)

type Historical struct {
	gorm.Model

	Candles []*finance.ChartBar

	RequestUuid uuid.UUID
}
