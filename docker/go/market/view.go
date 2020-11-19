package market

import "fmt"

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

func PrintChoices() ([]string, error) {
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
