package main

import (
	"flag"
	"fmt"
	//"golang.org/x/net/html"
	"io/ioutil"
	"net/http"
)

type Page struct {
	Title string
	Description string
}

func Get(url string) {
	resp, err := http.Get(url)
	if err != nil{
		panic("error")
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil{
		panic("error")
	}
	fmt.Printf("%s", string(b))
}

func Head(url string) {
	//resp, err := http.Get(url)
	fmt.Print("Head_test\n")
}

func main() {
	flag.Parse()
	m := flag.Arg(0)
	u := flag.Arg(1)// URL

	switch m {
	case "GET":
		Get(u)

	case "I":
		Head(u)
	}
}