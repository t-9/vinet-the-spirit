package address

import (
	"encoding/json"
	"log"

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

	body, err := auth.GetRequest("/v1/me/getaddresses", map[string]string{})
	if err != nil {
		return addresses, err
	}

	if err := json.Unmarshal(body, &addresses); err != nil {
		log.Println(string(body))
		return addresses, err
	}

	return addresses, nil
}
