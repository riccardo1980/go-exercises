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

func getLink(n *html.Node) Link {
	log.Printf("Parsing anchor attributes")
	var ret Link
	for _, a := range n.Attr {
		if a.Key == "href" {
			log.Printf("Found href: %s", a.Val)
			ret.Href = a.Val
			break
		}
	}
	ret.Text = getText(n)
	return ret
}

func getText(n *html.Node) string {
	log.Printf("Parsing anchor for text")
	if n.Type == html.TextNode {
		log.Printf("Found text: %s", n.Data)
		return n.Data
	}
	if n.Type != html.ElementNode {
		return ""
	}
	var ret string
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		ret += getText(c)
	}
	return strings.Join(strings.Fields(ret), " ")
}

func (e *Extractor) Extract(bytes []byte) Collection {

	doc, err := html.Parse(strings.NewReader(string(bytes)))
	if err != nil {
		log.Fatalf("Error on parsing: %s", err.Error())
	}

	anchors := getAnchors(doc)
	var c Collection

	for _, anchor := range anchors {
		c = append(c, getLink(anchor))
	}
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
