package handlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/piquette/finance-go/quote"
	"io/ioutil"
	"net/http"
	"saruman/src/core/db/mysql"
	"saruman/src/models"
	"saruman/src/service"
	"time"
)

func QuoteYahooGet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	pathParams := mux.Vars(r)

	symbol := pathParams["id"]

	if symbol == "" {
		w.WriteHeader(http.StatusBadRequest)

		res, _ := json.Marshal(models.Error{Code: http.StatusBadRequest, Message: "Provide a symbol"})
		_, _ = w.Write(res)

		return
	}

	requestBody, err := ioutil.ReadAll(r.Body)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		res, _ := json.Marshal(models.Error{Code: http.StatusBadRequest, Message: "Error while reading body"})
		_, _ = w.Write(res)
		return
	}

	var request models.HistoricalRequest

	err = json.Unmarshal(requestBody, &request)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		res, _ := json.Marshal(models.Error{Code: http.StatusBadRequest, Message: "Error while reading body"})
		_, _ = w.Write(res)
		return
	}

	res, err := service.GetHistoricalYahoo(symbol)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)

		res, _ := json.Marshal(models.Error{Code: http.StatusBadRequest, Message: err.Error()})
		_, _ = w.Write(res)

		return
	}

	quoteBytes, _ := json.Marshal(res)

	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(quoteBytes)

	return


	savedQuote := mysql.GetYahooQuote(symbol)

	if savedQuote == nil || time.Now().Sub(savedQuote.DateAdded) > time.Minute * time.Duration(5) {
		q, err := quote.Get(symbol)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)

			res, _ := json.Marshal(models.Error{Code: http.StatusInternalServerError, Message: "Failed getting a quote from Yahoo"})
			_, _ = w.Write(res)

			return
		}

		res, err := service.CreateYahooQuote(q)
		resBody, _ := json.Marshal(res)

		w.WriteHeader(http.StatusOK)
		_, _ = w.Write(resBody)

		return
	}
}
