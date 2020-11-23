package balance

import (
	"encoding/json"
	"log"

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

	body, err := auth.GetRequest("/v1/me/getbalance", map[string]string{})
	if err != nil {
		return balances, err
	}

	if err := json.Unmarshal(body, &balances); err != nil {
		log.Println(string(body))
		return balances, err
	}

	return balances, nil
}
