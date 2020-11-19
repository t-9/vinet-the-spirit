package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	markets, err := getMarkets()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(markets))
}

func getMarkets() (string, error) {
	url := "https://api.bitflyer.com/v1/getmarkets"

	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	byteArray, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(byteArray), nil
}
