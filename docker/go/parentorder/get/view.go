package get

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"vinet/market"
	"vinet/message"
)

func (p ParentOrder) String() string {
	return fmt.Sprintf(
		"%d, %s, %s, %s, %s, %f, %f, %f, %s, %s, %s, %s, %f, %f, %f, %f",
		p.ID,
		p.ParentOrderID,
		p.ProductCode,
		p.Side,
		p.ParentOrderType,
		p.Price,
		p.AveragePrice,
		p.Size,
		p.ParentOrderState,
		p.ExpireDate,
		p.ParentOrderDate,
		p.ParentOrderAcceptanceID,
		p.OutstandingSize,
		p.CancelSize,
		p.ExecutedSize,
		p.TotalCommision,
	)
}

// PrintList displays parent orders.
func PrintList() error {
	productCode, err := market.SelectProductCode()
	if err != nil {
		return err
	}

	params := map[string]string{"product_code": productCode}
	orders, err := getParentOrders(params)
	if err != nil {
		return err
	}

	fmt.Println("ID, ParentOrderID, ProductCode, Side, ParentOrderType, Price, AveragePrice," +
		" Size, ParentOrderState, ExpireDate, ParentOrderDate, ParentOrderAcceptanceID," +
		" OutstandingSize, CancelSize, ExecutedSize, TotalCommision")
	for _, d := range orders {
		fmt.Println(d)
	}
	return nil
}

func printChoices(productCode string) ([]string, error) {
	var choices []string

	params := map[string]string{"product_code": productCode}
	orders, err := getParentOrders(params)
	if err != nil {
		return choices, err
	}

	choices = make([]string, len(orders), len(orders))

	fmt.Println("Number, ID, ParentOrderID, ProductCode, Side, ParentOrderType, Price, AveragePrice," +
		" Size, ParentOrderState, ExpireDate, ParentOrderDate, ParentOrderAcceptanceID," +
		" OutstandingSize, CancelSize, ExecutedSize, TotalCommision")
	for i, o := range orders {
		choices[i] = o.ParentOrderID
		fmt.Printf("%d. %s\n", i, o)
	}
	return choices, nil
}

// SelectParentOrderID makes you select the parent order id.
func SelectParentOrderID(productCode string) (string, error) {
	fmt.Println(message.GetWhichParentOrderID())
	choices, err := printChoices(productCode)
	if err != nil {
		return "", err
	}
	fmt.Print(message.GetInputLine())
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	fmt.Println("")

	c, cerr := strconv.Atoi(scanner.Text())
	if cerr != nil || c >= len(choices) || c < 0 {
		return "", fmt.Errorf(message.GetWrongChoice())
	}
	return choices[c], nil
}
