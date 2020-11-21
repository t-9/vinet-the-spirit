package market

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// Market represents a market.
type Market struct {
	ProductCode string `json:"product_code"`
	MarketType  string `json:"market_type"`
	Alias       string `json:"alias"`
}

func getMarkets() ([]Market, error) {
	var markets []Market

	url := "https://api.bitflyer.com/v1/getmarkets"

	timeout := time.Duration(5 * time.Second)
	client := &http.Client{
		Timeout: timeout,
	}

	resp, err := client.Get(url)
	if err != nil {
		return markets, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return markets, err
	}

	if resp.StatusCode != 200 {
		return markets, fmt.Errorf(string(body))
	}

	if err := json.Unmarshal(body, &markets); err != nil {
		return markets, err
	}

	return markets, nil
}
