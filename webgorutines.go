package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Page struct {
	URL  string
	Size int
}

func ResponseSize(url string, channel chan Page) {
	//fmt.Println("Getting", url)
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	//channel <- len(body)
	channel <- Page{URL: url, Size: len(body)}
}

func main() {
	pages := make(chan Page)
	//sizes := make(chan int)
	/*
		go ResponseSize("https://example.com", sizes)
		go ResponseSize("https://golang.org", sizes)
		go ResponseSize("https://golang.org/doc", sizes)
		//time.Sleep(5 * time.Second)
		fmt.Println(<-sizes)
		fmt.Println(<-sizes)
		fmt.Println(<-sizes)
	*/
	urls := []string{"https://example.com", "https://golang.org", "https://golang.org/doc"}
	for _, url := range urls {
		go ResponseSize(url, pages)
	}
	for i := 0; i < len(urls); i++ {
		page := <-pages
		fmt.Printf("%s: %d\n", page.URL, page.Size)
	}
}
