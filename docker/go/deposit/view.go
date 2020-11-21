package deposit

import (
	"fmt"
)

func (d Deposit) String() string {
	return fmt.Sprintf(
		"%d, %s, %s, %f, %s, %s",
		d.ID,
		d.OrderID,
		d.CurrencyCode,
		d.Amount,
		d.Status,
		d.EventDate,
	)
}

// PrintList displays deposit histories.
func PrintList() error {
	deposits, err := getDeposits()
	if err != nil {
		return err
	}

	fmt.Println("ID, OrderID, CurrencyCode, Amount, Status, EventDate")
	for _, d := range deposits {
		fmt.Println(d)
	}
	return nil
}
