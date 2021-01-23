package service

import (
	"encoding/json"
	"github.com/dystopia-systems/alaskalog"
	"github.com/google/uuid"
	"github.com/piquette/finance-go"
	"github.com/piquette/finance-go/chart"
	"saruman/src/core/db/mysql"
	"saruman/src/models"
)

func GetHistoricalYahoo(request *chart.Params) (*models.Historical, error) {
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

		res, err := CreateHistoricalYahoo(bars, requestUuid)

		if err != nil {
			return nil, err
		}

		return res, nil
	}

	return savedQuote, nil
}

func CreateHistoricalYahoo(bars []finance.ChartBar, requestUuid uuid.UUID) (*models.Historical, error) {

}
