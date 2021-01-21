package service

import (
	"github.com/piquette/finance-go"
	"saruman/src/core/db/mysql"
	"saruman/src/models"
	"time"
)

func CreateYahooQuote(quote *finance.Quote) (*models.YahooQuote, error) {
	yahooQuote := models.YahooQuote{
		Quote:     *quote,
		DateAdded: time.Now(),
	}

	res, err := mysql.CreateYahooQuote(&yahooQuote)

	if err != nil {
		return nil, err
	}

	return res, err
}
