package parentorder

import (
	"fmt"

	"vinet/market"
	"vinet/message"
	"vinet/order"
	"vinet/util"
)

func (r Response) String() string {
	return r.ParentOrderAcceptanceID
}

// Order makes a child order
func Order() error {
	method, err := order.SelectMethod()
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

	paramNum := getNumberOfParameters(method)
	parameters := make([]Parameter, paramNum, paramNum)
	for i := 0; i < paramNum; i++ {
		productCode, err := market.SelectProductCode()
		if err != nil {
			return err
		}

		conditionType, err := order.SelectConditionType()
		if err != nil {
			return err
		}

		side, err := order.SelectSide()
		if err != nil {
			return err
		}

		size, err := order.InputSize()
		if err != nil {
			return err
		}

		var price *float64
		if util.ContainsString([]string{order.TypeLimit, order.TypeStopLimit}, conditionType) {
			p, err := order.InputPrice()
			price = &p
			if err != nil {
				return err
			}
		}

		var trigerPrice *float64
		if util.ContainsString([]string{order.TypeStop, order.TypeStopLimit}, conditionType) {
			tp, err := order.InputTrigerPrice()
			trigerPrice = &tp
			if err != nil {
				return err
			}
		}

		var offset *int64
		if conditionType == order.TypeTrail {
			o, err := order.InputOffset()
			offset = &o
			if err != nil {
				return err
			}
		}

		parameters[i] = Parameter{
			ProductCode:   productCode,
			ConditionType: conditionType,
			Side:          side,
			Size:          size,
			Price:         price,
			TriggerPrice:  trigerPrice,
			Offset:        offset,
		}
	}

	b := Body{
		OrderMethod:    method,
		MinuteToExpire: miniteToExpire,
		TimeInForce:    timeInForce,
		Parameters:     parameters,
	}

	return send(b)
}

func send(b Body) error {
	s, err := sendParentOrder(b)
	if err != nil {
		return err
	}

	fmt.Println(message.GetParentOrderAcceptanceID())
	fmt.Println(s)
	return nil
}

func getNumberOfParameters(method string) int {
	switch method {
	case order.MethodSimple:
		return 1
	case order.MethodIFD:
		return 2
	case order.MethodOCO:
		return 2
	case order.MethodIFDOCO:
		return 3
	}
	return 0
}
