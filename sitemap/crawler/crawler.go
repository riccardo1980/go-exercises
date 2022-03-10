package sitemap

import (
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"

	"golang.org/x/net/html"
)

func getAnchors(n *html.Node) []*html.Node {
	log.Printf("Parsing for anchors")
	if n.Type == html.ElementNode && n.Data == "a" {
		log.Printf("Found an anchor: %s", n.Data)
		return []*html.Node{n}
	}
	var ret []*html.Node
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		ret = append(ret, getAnchors(c)...)
	}
	return ret
}

type Loc struct {
	Link string
}

type Crawler struct {
	depth int
}

func New(opts ...CrawlerOptions) Crawler {
	c := Crawler{}
	for _, op := range opts {
		op(&c)
	}
	return c
}

type CrawlerOptions func(*Crawler)

func WithDepth(depth int) CrawlerOptions {
	return func(c *Crawler) {
		c.depth = depth
	}
}

func getDomain(URL string) string {
	u, _ := url.Parse(URL)
	return u.Host
}

func (c *Crawler) Parse(URL string) []Loc {
	var locs []Loc

	// retrieve domain
	domain := getDomain(URL)

	// split URL and get domain
	parse(URL, domain, c.depth, &locs)

	return locs
}

func parse(URL string, domain string, depth int, locs *[]Loc) {
	// get page
	resp, err := http.Get(URL)
	if err != nil {
		log.Printf("Error on getting page: %s", err.Error())
		return
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)

	// parse page:
	doc, err := html.Parse(strings.NewReader(string(body)))

	// 1 - build anchor list
	anchors := getAnchors(doc)

	// 2 - get link for each anchor
	// 3 - retain pages on same domain
	// call parser for each link (if not over depth)
	log.Println(string(body))
}
