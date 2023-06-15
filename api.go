package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func apiGET() {
	url := "https://fqdn.com"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}

	// reg.Header.Add("header-value", "API_KEY")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}
	fmt.Println(string(body))
}

func main() {
	apiGET()
}
