package balance

import (
	"fmt"
)

func (b Balance) String() string {
	return fmt.Sprintf("%s, %f, %f", b.CurrencyCode, b.Amount, b.Available)
}

func PrintList() error {
	balances, err := getBalances()
	if err != nil {
		return err
	}

	fmt.Println("CurrencyCode, Amount, Available")
	for _, b := range balances {
		fmt.Println(b)
	}
	return nil
}
