package linklib

import (
	"encoding/json"
	"log"
	"strings"

	"golang.org/x/net/html"
)

type Extractor struct {
}

func NewExtractor() Extractor {
	return Extractor{}
}

type Link struct {
	Href string
	Text string
}

type Collection []Link

func (e *Extractor) Extract(bytes []byte) Collection {
	var c Collection

	doc, err := html.Parse(strings.NewReader(string(bytes)))
	if err != nil {
		log.Fatalf("Error on parsiong: %s", err.Error())
	}

	var t func(n *html.Node, out string) string
	t = func(n *html.Node, out string) string {
		if n.Type == html.TextNode {
			next := strings.Trim(n.Data, "\n")
			if len(next) > 0 {
				out = out + next
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			out = t(c, out)
		}
		return out
	}

	var f func(n *html.Node, out *Collection)
	f = func(n *html.Node, out *Collection) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key == "href" {

					*out = append(*out, Link{
						Href: a.Val,
						Text: strings.TrimSpace(t(n, "")),
					})
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c, out)
		}
	}

	f(doc, &c)

	return c
}

func (ll *Collection) ToJSON() ([]byte, error) {
	log.Printf("JSON encoding")

	var bytes []byte = nil
	var err error = nil

	if len(*ll) > 0 {
		bytes, err = json.Marshal(ll)
	}

	return bytes, err
}
