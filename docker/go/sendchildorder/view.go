package sendchildorder

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"vinet/market"
	"vinet/message"
)

func (r Response) String() string {
	return r.ChildOrderAcceptanceID
}

func Order() error {
	productCode, err := market.SelectProductCode()
	if err != nil {
		return err
	}

	childOrderType, err := selectOrderType()
	if err != nil {
		return err
	}

	side, err := selectSide()
	if err != nil {
		return err
	}

	var price float64
	if childOrderType == TypeLimit {
		price, err = inputPrice()
		if err != nil {
			return err
		}
	}

	size, err := inputSize()
	if err != nil {
		return err
	}

	miniteToExpire, err := inputMinuteToExpire()
	if err != nil {
		return err
	}

	timeInForce, err := selectTimeInForce()
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

	fmt.Println("ChildOrderAcceptanceID")
	fmt.Println(s)
	return nil
}

func selectOrderType() (string, error) {
	orderTypeList := []string{
		TypeLimit,
		TypeMarket,
	}

	return selectItemString(orderTypeList, message.GetOrderType())
}

func selectSide() (string, error) {
	sideList := []string{
		SideBuy,
		SideSell,
	}

	return selectItemString(sideList, message.GetSide())
}

func selectTimeInForce() (string, error) {
	timeInForceList := []string{
		TimeInForceGTC,
		TimeInForceIOC,
		TimeInForceFOK,
	}

	return selectItemString(timeInForceList, message.GetTimeInForce())
}

func selectItemString(items []string, mes string) (string, error) {
	fmt.Println(mes)
	for i, t := range items {
		fmt.Printf("%d. %s\n", i, t)
	}
	fmt.Print(message.GetInputLine())
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	fmt.Println("")

	c, err := strconv.ParseInt(scanner.Text(), 10, 64)
	if err != nil || c < 0 || c >= int64(len(items)) {
		return "", fmt.Errorf(message.GetWrongChoice())
	}

	return items[c], nil
}

func inputPrice() (float64, error) {
	return inputPositiveFloat64(message.GetPrice())
}

func inputSize() (float64, error) {
	return inputPositiveFloat64(message.GetSize())
}

func inputPositiveFloat64(mes string) (float64, error) {
	fmt.Println(mes)
	fmt.Print(message.GetInputLine())
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	fmt.Println("")
	in, err := strconv.ParseFloat(scanner.Text(), 64)
	if err != nil || in < 0.0 {
		return in, fmt.Errorf(message.GetInvalidInputValue())
	}
	return in, nil
}

func inputMinuteToExpire() (int64, error) {
	fmt.Println(message.GetMinuteToExpire())
	fmt.Print(message.GetInputLine())
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	fmt.Println("")
	minuteToExpire, err := strconv.ParseInt(scanner.Text(), 10, 64)
	if err != nil || minuteToExpire < 0.0 {
		return minuteToExpire, fmt.Errorf(message.GetInvalidInputValue())
	}
	return minuteToExpire, nil
}
