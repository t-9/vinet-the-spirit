package cancel

import "vinet/auth"

// Body represents a request body.
type Body struct {
	ProductCode             string  `json:"product_code"`
	ParentOrderID           *string `json:"parent_order_id"`
	ParentOrderAcceptanceID *string `json:"parent_order_acceptance_id"`
}

func cancelParentOrder(b Body) error {
	_, err := auth.PostRequest("/v1/me/cancelparentorder", b)
	return err
}
