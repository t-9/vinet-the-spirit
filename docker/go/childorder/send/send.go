package send

import (
	"encoding/json"
	"log"

	"vinet/auth"
)

// Body represents a request body.
type Body struct {
	ProductCode    string   `json:"product_code"`
	ChildOrderType string   `json:"child_order_type"`
	Side           string   `json:"side"`
	Price          *float64 `json:"price"`
	Size           float64  `json:"size"`
	MinuteToExpire int64    `json:"minute_to_expire"`
	TimeInForce    string   `json:"time_in_force"`
}

// Response reprents a response body.
type Response struct {
	ChildOrderAcceptanceID string `json:"child_order_acceptance_id"`
}

func sendChildOrder(b Body) (Response, error) {
	var resBody Response

	body, err := auth.PostRequest("/v1/me/sendchildorder", b)
	if err != nil {
		return resBody, err
	}

	if err := json.Unmarshal(body, &resBody); err != nil {
		log.Println(string(body))
		return resBody, err
	}

	return resBody, nil
}
