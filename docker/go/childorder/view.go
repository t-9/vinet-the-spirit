package childorder

import (
	"fmt"

	"vinet/market"
	"vinet/message"
	"vinet/order"
)

func (r Response) String() string {
	return r.ChildOrderAcceptanceID
}

// Order makes a child order
func Order() error {
	productCode, err := market.SelectProductCode()
	if err != nil {
		return err
	}

	childOrderType, err := order.SelectOrderType()
	if err != nil {
		return err
	}

	side, err := order.SelectSide()
	if err != nil {
		return err
	}

	var price *float64
	if childOrderType == order.TypeLimit {
		p, err := order.InputPrice()
		price = &p
		if err != nil {
			return err
		}
	}

	size, err := order.InputSize()
	if err != nil {
		return err
	}

	miniteToExpire, err := order.InputMinuteToExpire()
	if err != nil {
		return err
	}

	timeInForce, err := order.SelectTimeInForce()
	if err != nil {
		return err
	}

	b := Body{
		ProductCode:    productCode,
		ChildOrderType: childOrderType,
		Side:           side,
		Price:          price,
		Size:           size,
		MinuteToExpire: miniteToExpire,
		TimeInForce:    timeInForce,
	}

	return send(b)
}

func send(b Body) error {
	s, err := sendChildOrder(b)
	if err != nil {
		return err
	}

	fmt.Println(message.GetChildOrderAcceptanceID())
	fmt.Println(s)
	return nil
}
