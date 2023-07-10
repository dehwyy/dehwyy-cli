package utils

import (
	"fmt"
	"testing"
	"time"
)

func assert_fetch(url string, expectStatusCode int, t *testing.T) {
	fetchStartTime := time.Now()
	actualStatusCode := FetchUrl(url, struct{}{})
	if actualStatusCode != expectStatusCode {
		t.Errorf("Expected %d, got %d statusCode", expectStatusCode, actualStatusCode)
	} else {
		timePassed := time.Since(fetchStartTime).Seconds()
		fmt.Printf("Fetch: passed %s -> %d in %.2fs\n",  url, actualStatusCode, timePassed)
	}
}

func TestFetchGoogle(t *testing.T) {
	t.Parallel()
	assert_fetch("https://google.com", 200, t)
}

func TestFetchTwitch(t *testing.T) {
	t.Parallel()
	assert_fetch("https://twitch.tv", 200, t)
}

func TestFetchVK(t *testing.T) {
	t.Parallel()
	assert_fetch("https://vk.com", 200, t)
}
