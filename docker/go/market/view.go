package market

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"vinet/message"
)

func (m Market) String() string {
	return fmt.Sprintf("%s, %s, %s", m.ProductCode, m.MarketType, m.Alias)
}

func PrintList() error {
	markets, err := getMarkets()
	if err != nil {
		return err
	}

	fmt.Println("ProductCode, MarketType, Alias")
	for _, m := range markets {
		fmt.Println(m)
	}
	return nil
}

func printChoices() ([]string, error) {
	var choices []string
	markets, err := getMarkets()
	if err != nil {
		return choices, err
	}

	choices = make([]string, len(markets), len(markets))

	fmt.Println("Number, ProductCode, MarketType, Alias")
	for i, m := range markets {
		choices[i] = m.ProductCode
		fmt.Printf("%d. %s\n", i, m)
	}
	return choices, nil
}

// SelectProductCode makes you select the product code.
func SelectProductCode() (string, error) {
	fmt.Println(message.GetWhichBoard())
	choices, err := printChoices()
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
