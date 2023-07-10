package utils

import (
	"encoding/json"
	"log"
	"net/http"
)

func FetchUrl(url string, v interface{}) int {
	response, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}

	defer response.Body.Close()
	json.NewDecoder(response.Body).Decode(v)

	return response.StatusCode
}
