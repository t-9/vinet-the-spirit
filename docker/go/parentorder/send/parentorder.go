package send

import (
	"encoding/json"
	"log"

	"vinet/order"
)

// Body represents a request body.
type Body struct {
	OrderMethod    string      `json:"order_method"`
	MinuteToExpire int64       `json:"minute_to_expire"`
	TimeInForce    string      `json:"time_in_force"`
	Parameters     []Parameter `json:"parameters"`
}

// Parameter represents a parameter.
type Parameter struct {
	ProductCode   string   `json:"product_code"`
	ConditionType string   `json:"condition_type"`
	Side          string   `json:"side"`
	Size          float64  `json:"size"`
	Price         *float64 `json:"price"`
	TriggerPrice  *float64 `json:"trigger_price"`
	Offset        *int64   `json:"offset"`
}

// Response reprents a response body.
type Response struct {
	ParentOrderAcceptanceID string `json:"parent_order_acceptance_id"`
}

func sendParentOrder(b Body) (Response, error) {
	var resBody Response

	body, err := order.SendOrder(b, "/v1/me/sendparentorder")
	if err != nil {
		return resBody, err
	}

	if err := json.Unmarshal(body, &resBody); err != nil {
		log.Println(string(body))
		return resBody, err
	}

	return resBody, nil
}
