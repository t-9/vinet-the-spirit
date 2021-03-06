package coinin

import (
	"fmt"
)

func (c CoinIn) String() string {
	return fmt.Sprintf(
		"%d, %s, %s, %f, %s, %s, %s, %s",
		c.ID,
		c.OrderID,
		c.CurrencyCode,
		c.Amount,
		c.Address,
		c.TxHash,
		c.Status,
		c.EventDate,
	)
}

// PrintList displays the history of virtual currency deposits.
func PrintList() error {
	coinins, err := getCoinIns()
	if err != nil {
		return err
	}

	fmt.Println("ID, OrderID, CurrencyCode, Amount, Address, TxHash, Status, EventDate")
	for _, c := range coinins {
		fmt.Println(c)
	}
	return nil
}
