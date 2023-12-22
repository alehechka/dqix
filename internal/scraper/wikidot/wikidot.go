package wikidot

import (
	"dqix/internal/scraper"
	"dqix/internal/scraper/utils"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/html"
)

type WikidotScraper struct {
	config *scraper.Config
	pages  scraper.Pages
}

func Init(config *scraper.Config) scraper.Scraper {
	return &WikidotScraper{
		config: config,
		pages:  make(scraper.Pages),
	}
}

func (s WikidotScraper) Scrape() (err error) {
	page, err := s.ScrapePage("", false)
	if err != nil {
		return err
	}

	s.pages[page.Path] = page

	if err := s.ScrapePageLinks(page); err != nil {
		return err
	}

	file, err := json.MarshalIndent(s.pages, "", " ")
	if err != nil {
		return err
	}

	return os.WriteFile("data/wikidot.json", file, 0644)
}

func (s WikidotScraper) ScrapePageLinks(page scraper.PageContent) (err error) {
	fmt.Println("Scraping: ", page.Path)

	for path := range page.Links {
		if _, ok := s.pages[path]; !ok {
			pageContent, err := s.ScrapePage(path, true)
			if err != nil {
				return err
			}
			s.pages[path] = pageContent
			s.ScrapePageLinks(pageContent)
		}
	}
	return
}

func (s WikidotScraper) ScrapePage(path string, mainContentOnly bool) (page scraper.PageContent, err error) {
	trimmedPath := strings.TrimPrefix(path, "/")
	res, err := http.Get(fmt.Sprintf("%s/%s", s.config.WikiURL, trimmedPath))
	if err != nil {
		return scraper.PageContent{}, err
	}
	defer res.Body.Close()

	doc, err := html.Parse(res.Body)
	if err != nil {
		return scraper.PageContent{}, err
	}

	// Find and print all links on the web page
	links := make(map[string]string)
	text := make([]string, 0)
	var link func(*html.Node, bool, bool)
	link = func(n *html.Node, shouldCaptureLinks, shouldCaptureText bool) {
		if n.Type == html.ElementNode && n.Data == "script" {
			return
		}

		var isMainContent bool
		if !shouldCaptureText && n.Type == html.ElementNode && n.Data == "div" {
			for _, a := range n.Attr {
				if a.Key == "id" {
					if a.Val == "main-content" {
						isMainContent = true
					}
					break
				}
			}
		}

		if shouldCaptureLinks && n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key == "href" {
					hrefPath := strings.Split(a.Val, "#")[0]
					if strings.HasPrefix(hrefPath, "/") {
						links[hrefPath] = utils.GetNodeText(n)
					}
				}
			}
		}

		if shouldCaptureText && n.Type == html.TextNode {
			t := strings.TrimSpace(n.Data)
			if t != "" {
				text = append(text, t)
			}
		}

		// traverses the HTML of the webpage from the first child node
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			link(c, (shouldCaptureLinks || isMainContent), (shouldCaptureText || isMainContent))
		}
	}
	link(doc, !mainContentOnly, false)

	return scraper.PageContent{
		Path:  trimmedPath,
		Text:  text,
		Links: links,
	}, nil
}
