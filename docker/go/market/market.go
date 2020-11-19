package market

import (
	"io/ioutil"
	"net/http"
)

func GetMarkets() (string, error) {
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
