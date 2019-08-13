package main

import (
	"bytes"
	"flag"
	"fmt"
	"strings"

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

func POST(url string, content string, contenttype string){
	//resp, err := http.Get(url)
	fmt.Print("POST_test\n")
	req, err := http.NewRequest(
		"POST",
		url,
		bytes.NewBuffer([]byte(content)))
	if err != nil{
		panic("error")
	}
	splits := strings.Split(contenttype, ":")
	req.Header.Set(strings.ReplaceAll(strings.TrimLeft(splits[0], "{"), " ", ""),strings.ReplaceAll(strings.TrimRight(splits[1], "}"), " ", ""))
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic("error")
	}
	defer resp.Body.Close()
	//dumpResp, _ := httputil.DumpResponse(resp, true)
	//fmt.Printf("%s\n", dumpResp)
	fmt.Print("done.\n")
}

func main() {
	var (
		u = flag.String("u", "", "string flag")
		m = flag.String("X", "GET", "string flag")
		h = flag.String("H", "", "string flag")
		d = flag.String("d", "", "string flag")
	)
	flag.Parse()

	switch *m {
	case "GET":
		Get(*u)
	case "POST":
		POST(*u,*d,*h)
	default:
		panic("ERROR: There is no such method.")
	}

}