package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"github.com/gocolly/colly"
)

type Fact struct {
	Header string `json:"header"`
	Link   string `json:"description"`
}

func main() {
	allFacts := make([]Fact, 0)

	collector := colly.NewCollector(
		colly.AllowedDomains("armlur.am", "www.armlur.am"),
	)
	collector.OnHTML(".armlur-content", func(element *colly.HTMLElement) {
		factHeader := element.DOM.Find(".armlur-posts-list-header").Text()

		factLink := element.Attr("href")

		fact := Fact{
			Header: factHeader,
			Link:   factLink,
		}

		allFacts = append(allFacts, fact)

	})

	collector.OnRequest(func(request *colly.Request) {
		fmt.Print("Scrapping  ", request.URL.String())
	})

	collector.Visit("https://www.armlur.am/category/newsfeedcat/world/")

	//writeJson(allFacts)

	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", " ")
	enc.Encode(allFacts)
}

func writeJson(data []Fact) {
	file, err := json.MarshalIndent(data, "", " ")

	if err != nil {
		log.Println(err)
		return
	}

	_ = ioutil.WriteFile("news.json", file, 0644)
}
