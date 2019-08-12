package main

import (
	"fmt"
	"golang.org/x/net/html"
	"net/http"
)

		type Page struct {
			Title string
			Description string
		}


		func isDescription(attrs []html.Attribute) bool {
			for _, attr := range attrs {
		if attr.Key == "name" && attr.Val == "description"{
			return true
		}
	}
	return false
}

func Get(url string) (*Page, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return nil, err
	}
	page := Page{}
	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "title" {
			page.Title = n.FirstChild.Data
		}
		if n.Type == html.ElementNode && n.Data == "meta" {
			if isDescription(n.Attr) {
				for _, attr := range n.Attr {
					if attr.Key == "content" {
						page.Description = attr.Val
					}
				}
			}
		}
		if page.Title != "" && page.Description != "" {
			return
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)
	if page.Title == "" || page.Description == "" {
		return nil, fmt.Errorf("There is no find Title or Description")
	}
	return &page, nil
}

func main() {
	page, err := Get("http://voyagegroup.com")
	if err != nil {
		fmt.Errorf("Error occur in request.")
	}
	fmt.Printf("%#v", page)
}