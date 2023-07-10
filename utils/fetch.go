package utils

import (
	"encoding/json"
	"net/http"

	e "github.com/dehwyy/dehwyy-cli/error-handler"
)

func FetchUrl(url string, v interface{}) int {
	response := e.WithFatal(http.Get(url))("Error fetching")

	defer response.Body.Close()
	json.NewDecoder(response.Body).Decode(v)

	return response.StatusCode
}
