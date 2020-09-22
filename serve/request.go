package serve

import (
	"net/http"
)

func HandleRequests() error {
	err := http.ListenAndServe(":3000", nil)

	if err != nil {
		return err
	}

	return nil
}