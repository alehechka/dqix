package utils

import (
	"golang.org/x/net/html"
)

func GetNodeText(n *html.Node) string {
	if n.Type == html.TextNode {
		return n.Data
	}

	// traverses the HTML of the webpage from the first child node
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if text := GetNodeText(c); text != "" {
			return text
		}
	}

	return ""
}
