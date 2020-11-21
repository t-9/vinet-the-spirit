package address

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"vinet/auth"
)

// Address represents the address where
// the virtual currency is deposited into a bitFlyer account.
type Address struct {
	Type         string `json:"type"`
	CurrencyCode string `json:"currency_code"`
	Address      string `json:"address"`
}

func getAddresses() ([]Address, error) {
	var addresses []Address

	path := "/v1/me/getaddresses"
	url := "https://api.bitflyer.com" + path
	method := http.MethodGet

	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return addresses, err
	}

	timestamp := auth.GenerateTimestamp()
	sign, err := auth.GenerateSign(timestamp, method, path, "")
	if err != nil {
		return addresses, err
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
		return addresses, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return addresses, err
	}

	if resp.StatusCode != 200 {
		return addresses, fmt.Errorf(string(body))
	}

	if err := json.Unmarshal(body, &addresses); err != nil {
		log.Println(string(body))
		return addresses, err
	}

	return addresses, nil
}
