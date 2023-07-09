package utils

import (
	"encoding/json"
	"log"
	"net/http"
)

func FetchUrl(url string, v interface{}) {
	res, err := http.Get(url)
	if err != nil {
		log.Fatalf("Error fetching Jisho: %v", err)
	}

	defer res.Body.Close()

	json.NewDecoder(res.Body).Decode(v)
}
