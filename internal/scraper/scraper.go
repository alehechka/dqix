package scraper

import "net/url"

type Config struct {
	WikiURL *url.URL
}

type Scraper interface {
	Scrape() error
}
