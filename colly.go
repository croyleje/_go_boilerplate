package main

import (
	"fmt"
	"log"

	"github.com/gocolly/colly"
)

func main() {
	c := colly.NewCollector(
		colly.AllowedDomains("www.website.com", "www.fqdn.com"),
	)

	// On every a element which has href attribute call callback
	// c.OnHTML("div.white-balls", func(e *colly.HTMLElement) {
	c.OnHTML("css selector", func(e *colly.HTMLElement) {
		e.ForEach("css selector", func(_ int, l *colly.HTMLElement) {
			fmt.Println(l.Text)
		})

	})

	// Before making a request print "Visiting ..."
	c.OnRequest(func(e *colly.Request) {
		fmt.Println("Visiting", e.URL.String())
	})

	c.OnResponse(func(r *colly.Response) {
		fmt.Println(r.StatusCode)

		for key, value := range *r.Headers {
			fmt.Printf("%s: %s\n", key, value)
		}
	})

	url := "https://fqdn.com"
	err := c.Visit(url)
	if err != nil {
		log.Fatal(err)
	}

}
