package service

import (
	"github.com/piquette/finance-go"
	"github.com/piquette/finance-go/quote"
	"saruman/src/core/db/mysql"
	"saruman/src/models"
	"time"
)

func CreateYahooQuote(quote *finance.Quote) (*models.PriceQuote, error) {
	yahooQuote := models.PriceQuote{
		Quote:     *quote,
		DateAdded: time.Now(),
	}

	res, err := mysql.CreateYahooQuote(&yahooQuote)

	if err != nil {
		return nil, err
	}

	return res, err
}

func GetYahooQuote(symbol string) (*models.PriceQuote, error) {
	savedQuote := mysql.GetYahooQuote(symbol)

	if savedQuote == nil || time.Now().Sub(savedQuote.DateAdded) > time.Minute * time.Duration(5) {
		q, err := quote.Get(symbol)

		if err != nil {
			return nil, err
		}

		res, err := CreateYahooQuote(q)

		if err != nil {
			return nil, err
		}

		return res, nil
	}

	return savedQuote, nil
}