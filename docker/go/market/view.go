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
