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
	s.ScrapeFrom("/system:list-all-pages")
	s.ScrapeFrom("/system:page-tags-list")

	return s.WriteFile("data/wikidot.json")
}

func (s WikidotScraper) WriteFile(path string) (err error) {
	file, err := json.MarshalIndent(s.pages, "", " ")
	if err != nil {
		return err
	}

	return os.WriteFile(path, file, 0644)
}

func (s WikidotScraper) ScrapeFrom(path string) (err error) {
	page, err := s.ScrapePage(path)
	if err != nil {
		return err
	}

	s.pages[page.Path] = page

	if err := s.ScrapePageLinks(page); err != nil {
		return err
	}

	return
}

func (s WikidotScraper) ScrapePageLinks(page scraper.PageContent) (err error) {
	fmt.Println("Scraping: ", page.Path)

	for path := range page.Links {
		if _, ok := s.pages[path]; !ok {
			pageContent, err := s.ScrapePage(path)
			if err != nil {
				return err
			}
			s.pages[path] = pageContent
			s.ScrapePageLinks(pageContent)
		}
	}
	return
}

func (s WikidotScraper) ScrapePage(path string) (page scraper.PageContent, err error) {
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
	var link func(*html.Node, bool)
	link = func(n *html.Node, isMainContent bool) {
		if n.Type == html.ElementNode && n.Data == "script" {
			return
		}

		if !isMainContent && n.Type == html.ElementNode && n.Data == "div" {
			for _, a := range n.Attr {
				if a.Key == "id" {
					if a.Val == "main-content" {
						isMainContent = true
					}
					break
				}
			}
		}

		if isMainContent {

			// the tag pages have a floating box with all the tags listed and it just adds a bunch of fluff, so this skips the node if found
			if n.Type == html.ElementNode && n.Data == "div" {
				for _, a := range n.Attr {
					if a.Key == "class" && a.Val == "pages-tag-cloud-box" {
						return
					}
				}
			}

			if n.Type == html.ElementNode && n.Data == "a" {
				for _, a := range n.Attr {
					if a.Key == "href" {
						hrefPath := strings.Split(a.Val, "#")[0]
						if strings.HasPrefix(hrefPath, "/") {
							links[hrefPath] = utils.GetNodeText(n)
						}
					}
				}
			}

			if n.Type == html.TextNode {
				t := strings.TrimSpace(n.Data)
				if t != "" {
					text = append(text, t)
				}
			}
		}

		// traverses the HTML of the webpage from the first child node
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			link(c, isMainContent)
		}
	}
	link(doc, false)

	return scraper.PageContent{
		Path:  trimmedPath,
		Text:  text,
		Links: links,
	}, nil
}
