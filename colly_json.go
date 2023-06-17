package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/gocolly/colly"
)

type Data struct {
	Date  string `json:"date"`
	N     int    `json:"n"`
	Value string `json:"value"`
}

func main() {
	c := colly.NewCollector(
		colly.AllowedDomains("domain.com"),
	)

	batch := make([]Data, 0)

	// On every a element which has href attribute call callback
	// c.OnHTML("div.white-balls", func(e *colly.HTMLElement) {
	c.OnHTML("div#selector", func(e *colly.HTMLElement) {
		d := Data{}
		e.ForEach("a.selector", func(_ int, ee *colly.HTMLElement) {
			d.Date = ee.ChildText("h1")
			d.N, _ = strconv.Atoi(ee.ChildText("div:nth-child(1)"))
			d.Value = ee.ChildText("span")
			batch = append(batch, d)
		})

	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.OnResponse(func(r *colly.Response) {
		fmt.Println("Response from", r.Request.URL)
	})

	c.OnError(func(r *colly.Response, e error) {
		fmt.Println("Error:", e)
	})

	c.OnScraped(func(r *colly.Response) {
		fmt.Println("Finished:", r.Request.URL)
		js, err := json.MarshalIndent(batch, "", "    ")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Writing data to file.")
		if err := os.WriteFile("batch.json", js, 0664); err == nil {
			fmt.Println("Data written successfully.")
		}

	})

	var err error
	err = nil
	for i := 1; err == nil; i++ {
		url := "https://domain.com?pg=" + strconv.Itoa(i)
		err = c.Visit(url)
		if err != nil {
			log.Fatal(err)
		}
	}

}
