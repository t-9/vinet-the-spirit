package balance

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"vinet/auth"
)

// Balance represents the asset balance.
type Balance struct {
	CurrencyCode string  `json:"currency_code"`
	Amount       float64 `json:"amount"`
	Available    float64 `json:"available"`
}

func getBalances() ([]Balance, error) {
	var balances []Balance

	path := "/v1/me/getbalance"
	url := "https://api.bitflyer.com" + path
	method := http.MethodGet

	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return balances, err
	}

	timestamp := auth.GenerateTimestamp()
	sign, err := auth.GenerateSign(timestamp, method, path, "")
	if err != nil {
		return balances, err
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
		return balances, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return balances, err
	}

	if resp.StatusCode != 200 {
		return balances, fmt.Errorf(string(body))
	}

	if err := json.Unmarshal(body, &balances); err != nil {
		log.Println(string(body))
		return balances, err
	}

	return balances, nil
}
