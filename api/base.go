package api

import "net/url"

const (
	baseScheme = "https"
	baseURL    = "twitter.com"
)

var (
	BaseURL url.URL = url.URL{}
)

func init() {
	BaseURL.Host = baseURL
	BaseURL.Scheme = baseScheme
}
