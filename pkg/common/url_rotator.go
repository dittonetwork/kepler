package common

import (
	"os"
	"strings"
)

type ApiUrls struct {
	urls      []string
	currentId int
}

// NewApiUrls initializes ApiUrls with a list of URLs.
func NewApiUrls(urls []string) *ApiUrls {
	return &ApiUrls{
		urls:      urls,
		currentId: 0,
	}
}

// GetCurrentUrl returns the current URL.
func (au *ApiUrls) GetCurrentUrl() string {
	if len(au.urls) == 0 {
		return "" // or handle the error as appropriate
	}
	return au.urls[au.currentId]
}

// RotateUrl moves to the next URL in the list, cycling back to the start if necessary.
func (au *ApiUrls) RotateUrl() {
	au.currentId = (au.currentId + 1) % len(au.urls)
}

func GetUrlsFromEnv(envVar string, defaultUrls []string) []string {
	urlStr := os.Getenv(envVar)
	if urlStr == "" {
		return defaultUrls
	}
	return strings.Split(urlStr, ",")
}
