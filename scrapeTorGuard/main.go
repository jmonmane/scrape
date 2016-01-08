package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/yhat/scrape"
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

func main() {
	// request and parse the front page
	resp, err := http.Get("https://torguard.net/downloads.php")
	if err != nil {
		panic(err)
	}
	root, err := html.Parse(resp.Body)
	if err != nil {
		panic(err)
	}

	// define a matcher
	matcher := func(n *html.Node) bool {
		// must check for nil values
		// if n.DataAtom == atom.A && n.Parent != nil && n.Parent.Parent != nil {
		if n.DataAtom == atom.Tr {
			return true
		}
		return false
	}
	// grab all articles and print them
	articles := scrape.FindAll(root, matcher)
	for _, article := range articles {
		if strings.Contains(scrape.Text(article), "DEBIAN x64Bit") {
			fmt.Printf("%s\n", scrape.Text(article))
		}
		//fmt.Printf("%2d %s (%s)\n", i, scrape.Text(article), scrape.Attr(article, "href"))
	}
}
