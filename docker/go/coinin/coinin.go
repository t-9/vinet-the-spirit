package coinin

import (
	"encoding/json"
	"log"

	"vinet/auth"
)

// CoinIn represents the history of virtual currency deposits.
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

	body, err := auth.GetRequest("/v1/me/getcoinins", map[string]string{})
	if err != nil {
		return coinIns, err
	}

	if err := json.Unmarshal(body, &coinIns); err != nil {
		log.Println(string(body))
		return coinIns, err
	}

	return coinIns, nil
}
