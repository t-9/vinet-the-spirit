package order

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"vinet/message"
)

// SelectOrderType lets you select an order type.
func SelectOrderType() (string, error) {
	orderTypeList := []string{
		TypeLimit,
		TypeMarket,
	}

	return selectItemString(orderTypeList, message.GetOrderType())
}

// SelectSide lets you select a side.
func SelectSide() (string, error) {
	sideList := []string{
		SideBuy,
		SideSell,
	}

	return selectItemString(sideList, message.GetSide())
}

// SelectTimeInForce lets you select a time in force.
func SelectTimeInForce() (string, error) {
	timeInForceList := []string{
		TimeInForceGTC,
		TimeInForceIOC,
		TimeInForceFOK,
	}

	return selectItemString(timeInForceList, message.GetTimeInForce())
}

// InputPrice lets you input a price.
func InputPrice() (float64, error) {
	return inputPositiveFloat64(message.GetPrice())
}

// InputSize lets you input a size.
func InputSize() (float64, error) {
	return inputPositiveFloat64(message.GetSize())
}

// InputMinuteToExpire lets you input a minute to expire.
func InputMinuteToExpire() (int64, error) {
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
