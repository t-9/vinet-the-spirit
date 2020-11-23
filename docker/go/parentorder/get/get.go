package get

import (
	"encoding/json"
	"log"

	"vinet/auth"
)

// ParentOrder represents a parent order.
type ParentOrder struct {
	ID                      int64   `json:"id"`
	ParentOrderID           string  `json:"parent_order_id"`
	ProductCode             string  `json:"product_code"`
	Side                    string  `json:"side"`
	ParentOrderType         string  `json:"parent_order_type"`
	Price                   float64 `json:"price"`
	AveragePrice            float64 `json:"average_price"`
	Size                    float64 `json:"size"`
	ParentOrderState        string  `json:"parent_order_state"`
	ExpireDate              string  `json:"expire_date"`
	ParentOrderDate         string  `json:"parent_order_date"`
	ParentOrderAcceptanceID string  `json:"parent_order_acceptance_id"`
	OutstandingSize         float64 `json:"outstanding_size"`
	CancelSize              float64 `json:"cancel_size"`
	ExecutedSize            float64 `json:"executed_size"`
	TotalCommision          float64 `json:"total_commission"`
}

func getParentOrders(params map[string]string) ([]ParentOrder, error) {
	var orders []ParentOrder

	body, err := auth.GetRequest("/v1/me/getparentorders", params)
	if err != nil {
		return orders, err
	}

	if err := json.Unmarshal(body, &orders); err != nil {
		log.Println(string(body))
		return orders, err
	}

	return orders, nil
}
