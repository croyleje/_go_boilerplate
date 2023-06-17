package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func api_GET_request() {
	url := "https://fqdn.com"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}
	// Appending API_KEY to request.
	// reg.Header.Add("header-value", "value")

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
	api_GET_request()
}
