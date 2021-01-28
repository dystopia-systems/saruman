package service

import (
	"encoding/json"
	"github.com/dystopia-systems/alaskalog"
	"github.com/google/uuid"
	"github.com/piquette/finance-go"
	"github.com/piquette/finance-go/chart"
	"saruman/src/core/db/mysql"
)

func GetHistoricalYahoo(request *chart.Params) (*mysql.Historical, error) {
	requestBytes, _ := json.Marshal(request)
	requestUuid, _ := uuid.FromBytes(requestBytes)

	savedQuote := mysql.GetHistoricalYahooQuote(requestUuid.String())

	if savedQuote == nil {
		c := chart.Get(request)

		var bars []finance.ChartBar

		for c.Next() {
			bars = append(bars, *c.Bar())
		}

		if err := c.Err(); err != nil {
			alaskalog.Logger.Warnln(err)
		}

		return nil, nil
	}

	return savedQuote, nil
}
