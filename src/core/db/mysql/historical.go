package mysql

import "saruman/src/models"

func GetHistoricalYahooQuote(uuid string) *models.Historical {
	context := GetDb()

	var res []models.Historical

	context.
		Where("request_id = ?", uuid).
		First(&res)

	if len(res) > 1 {
		return &res[0]
	}

	return nil
}
