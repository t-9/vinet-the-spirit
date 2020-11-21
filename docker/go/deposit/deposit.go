package deposit

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"vinet/auth"
)

type Deposit struct {
	ID           int64  `json:"id"`
	OrderID      string `json:"order_id"`
	CurrencyCode string `json:"currency_code"`
	Amount       int64  `json:"amount"`
	Status       string `json:"status"`
	EventDate    string `json:"event_date"`
}

func getDeposits() ([]Deposit, error) {
	var deposits []Deposit

	path := "/v1/me/getdeposits"
	url := "https://api.bitflyer.com" + path
	method := http.MethodGet

	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return deposits, err
	}

	timestamp := auth.GenerateTimestamp()
	sign, err := auth.GenerateSign(timestamp, method, path, "")
	if err != nil {
		return deposits, err
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
		return deposits, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return deposits, err
	}

	if resp.StatusCode != 200 {
		return deposits, fmt.Errorf(string(body))
	}

	if err := json.Unmarshal(body, &deposits); err != nil {
		log.Println(string(body))
		return deposits, err
	}

	return deposits, nil
}
