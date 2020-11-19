package market

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Market struct {
	ProductCode string `json:"product_code"`
	MarketType  string `json:"market_type"`
	Alias       string `json:"alias"`
}

func (m Market) String() string {
	return fmt.Sprintf("%s, %s, %s", m.ProductCode, m.MarketType, m.Alias)
}

func GetMarkets() ([]Market, error) {
	url := "https://api.bitflyer.com/v1/getmarkets"

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var markets []Market
	if err := json.Unmarshal(body, &markets); err != nil {
		return nil, err
	}

	return markets, nil
}
