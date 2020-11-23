package deposit

import (
	"encoding/json"
	"log"

	"vinet/auth"
)

// Deposit represents a deposit history.
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

	body, err := auth.GetRequest("/v1/me/getdeposits", map[string]string{})
	if err != nil {
		return deposits, err
	}

	if err := json.Unmarshal(body, &deposits); err != nil {
		log.Println(string(body))
		return deposits, err
	}

	return deposits, nil
}
