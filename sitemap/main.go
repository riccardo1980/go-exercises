package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	sitemap "sitemap/crawler"
)

type options struct {
	domain string
	depth  int
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	fs := flag.NewFlagSet("link", flag.ExitOnError)

	var opts options
	fs.StringVar(&opts.domain, "domain", "", "Domain to parse")
	fs.IntVar(&opts.depth, "depth", 5, "Parse depth")

	fs.Parse(os.Args[1:])

	c := sitemap.New(sitemap.WithDepth(5))

	locs := c.Parse(opts.domain)
	fmt.Printf("Items :%d", len(locs))
}
