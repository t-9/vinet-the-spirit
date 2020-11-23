package cancel

import (
	"fmt"

	"vinet/market"
	"vinet/message"
	getparent "vinet/parentorder/get"
)

// Cancel cancels a parent order.
func Cancel() error {
	productCode, err := market.SelectProductCode()
	if err != nil {
		return err
	}

	parentOrderID, err := getparent.SelectParentOrderID(productCode)
	if err != nil {
		return err
	}

	b := Body{
		ProductCode:   productCode,
		ParentOrderID: &parentOrderID,
	}

	return cancel(b)
}

func cancel(b Body) error {
	if err := cancelParentOrder(b); err != nil {
		return err
	}

	fmt.Println(message.GetParentOrderCancelSuccess())
	return nil
}
