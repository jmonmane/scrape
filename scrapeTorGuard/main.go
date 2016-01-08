package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/yhat/scrape"
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

func getLink(r *html.Node) (s string) {
	buttons := scrape.FindAll(r, scrape.ByClass("downloadbtn"))
	for _, button := range buttons {
		windowLocation := scrape.Attr(button, "onclick")
		link := strings.Split(windowLocation, "=")[1]
		s := strings.Trim(link, "'")
		return s
	}
	return
}

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

	rows := scrape.FindAll(root, scrape.ByTag(atom.Tr))
	for _, row := range rows {
		if strings.Contains(scrape.Text(row), "DEBIAN x64") {
			l := getLink(row)
			fmt.Printf("%s \n %s \n", scrape.Text(row), l)
		}
	}
}

//fmt.Printf("%2d %s (%s)\n", i, scrape.Text(article), scrape.Attr(article, "href"))
