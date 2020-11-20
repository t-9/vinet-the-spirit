package address

import "fmt"

func (a Address) String() string {
	return fmt.Sprintf(
		"%s, %s, %s",
		a.Type,
		a.CurrencyCode,
		a.Address,
	)
}

func PrintList() error {
	addresses, err := getAddresses()
	if err != nil {
		return err
	}

	fmt.Println("Type, CurrencyCode, Address")
	for _, a := range addresses {
		fmt.Println(a)
	}
	return nil
}
