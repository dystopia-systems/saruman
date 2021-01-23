package mysql

import (
	"github.com/dystopia-systems/alaskalog"
	"saruman/src/models"
)

func GetYahooQuote(symbol string) *models.PriceQuote {
	var res []models.PriceQuote

	context := GetDb()

	context.
		Where("symbol = ?", symbol).
		Order("date_added desc").
		First(&res)

	if context.Error != nil {
		alaskalog.Logger.Warnf("Failed to execute query %v", context.Error)
		return &models.PriceQuote{}
	}

	if len(res) > 0 {
		return &res[0]
	}

	return nil
}

func CreateYahooQuote(quote *models.PriceQuote) (*models.PriceQuote, error) {
	context := GetDb()

	context.Create(&quote)

	result := context.Save(&quote)

	if result.Error != nil {
		alaskalog.Logger.Warnf("Failed to execute query %v", context.Error)
		return nil, result.Error
	}

	return quote, nil
}
