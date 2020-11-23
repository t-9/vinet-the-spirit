package get

import (
	"fmt"

	"vinet/market"
)

func (p ParentOrder) String() string {
	return fmt.Sprintf(
		"%d, %s, %s, %s, %s, %f, %f, %f, %s, %s, %s, %s, %f, %f, %f, %f",
		p.ID,
		p.ParentOrder,
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

	fmt.Println("ID, ParentOrder, ProductCode, Side, ParentOrderType, Price, AveragePrice," +
		" Size, ParentOrderState, ExpireDate, ParentOrderDate, ParentOrderAcceptanceID," +
		" OutstandingSize, CancelSize, ExecutedSize, TotalCommision")
	for _, d := range orders {
		fmt.Println(d)
	}
	return nil
}
