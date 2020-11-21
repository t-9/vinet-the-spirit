package coinin

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"vinet/auth"
)

type CoinIn struct {
	ID           int64   `json:"id"`
	OrderID      string  `json:"order_id"`
	CurrencyCode string  `json:"currency_code"`
	Amount       float64 `json:"amount"`
	Address      string  `json:"address"`
	TxHash       string  `json:"tx_hash"`
	Status       string  `json:"status"`
	EventDate    string  `json:"event_date"`
}

func getCoinIns() ([]CoinIn, error) {
	var coinIns []CoinIn

	path := "/v1/me/getcoinins"
	url := "https://api.bitflyer.com" + path
	method := http.MethodGet

	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return coinIns, err
	}

	timestamp := auth.GenerateTimestamp()
	sign, err := auth.GenerateSign(timestamp, method, path, "")
	if err != nil {
		return coinIns, err
	}

	req.Header.Set("ACCESS-KEY", auth.GetAccessKey())
	req.Header.Set("ACCESS-TIMESTAMP", timestamp)
	req.Header.Set("ACCESS-SIGN", sign)
	req.Header.Set("Content-Type", "application/json")

	timeout := time.Duration(5 * time.Second)
	client := &http.Client{
		Timeout: timeout,
	}

	resp, err := client.Do(req)
	if err != nil {
		return coinIns, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return coinIns, err
	}

	if resp.StatusCode != 200 {
		return coinIns, fmt.Errorf(string(body))
	}

	if err := json.Unmarshal(body, &coinIns); err != nil {
		log.Println(string(body))
		return coinIns, err
	}

	return coinIns, nil
}
