package scraper

import "net/url"

type Config struct {
	WikiURL *url.URL
}

type Scraper interface {
	Scrape() error
}

type PageContent struct {
	Path  string
	Text  []string
	Links map[string]string
}

type Pages map[string]PageContent
