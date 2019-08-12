package main

import (
	"encoding/json"
	"fmt"
	"golang.org/x/net/html"
	"log"
	"net/http"
)

type Page struct {
	URL			string `json:"url"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

func isDescription(attrs []html.Attribute) bool {
	for _, attr := range attrs {
		if attr.Key == "name" && attr.Val == "description"{
			return true
		}
	}
	return false
}

func lclserve(inputurl []string){
	pages := []Page{}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		for _, url := range inputurl {
			//fmt.Print(url)
			page, err := Get(url)
			if err != nil {
				fmt.Errorf("Error occur in request.")
			}
			pages = append(pages, *page)
		}
		//page, err := Get(inputurl[0])


		fmt.Printf("%#v", pages)

		w.Header().Set("Content-Type", "application/json")
		// json.NewEncoder はJSON用のエンコーダを初期化しています。
		// ここで初期化したエンコーダを使って `enc.Encode` をすることで、
		// 初期化時に指定した io.Writer に出力を書き込むことができます。
		// w はここでは http.ResponseWriter なので、HTTPの出力、
		// つまりHTTPレスポンスとしてJSONを返すことができます。便利ですね。
		enc := json.NewEncoder(w)
			if err := enc.Encode(pages); err != nil {
				http.Error(w, "encoding failed", http.StatusInternalServerError)
				return
			}

	})
	log.Fatal(http.ListenAndServe(":8080", nil))
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
	page.URL = url
	return &page, nil
}

func main() {
	urls := []string{"http://voyagegroup.com","https://smartplus-sec.com"}
	lclserve(urls)
}