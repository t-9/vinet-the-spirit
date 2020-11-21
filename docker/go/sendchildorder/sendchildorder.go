package sendchildorder

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"vinet/auth"
)

const (
	// TypeLimit represents that the order type is a limit.
	TypeLimit = "LIMIT"
	// TypeMarket represents that the order type is a market.
	TypeMarket = "MARKET"

	// SideBuy represents that the side is a buy.
	SideBuy = "BUY"
	// SideSell represents that the side is a sell.
	SideSell = "SELL"

	// TimeInForceGTC represents that the time in force is good till close.
	TimeInForceGTC = "GTC"
	// TimeInForceIOC represents that the time in force is immediate or close.
	TimeInForceIOC = "IOC"
	// TimeInForceFOK represents that the time in force is fill or kill.
	TimeInForceFOK = "FOK"
)

// Body represents a request body.
type Body struct {
	ProductCode    string  `json:"product_code"`
	ChildOrderType string  `json:"child_order_type"`
	Side           string  `json:"side"`
	Price          float64 `json:"price"`
	Size           float64 `json:"size"`
	MinuteToExpire int64   `json:"minute_to_expire"`
	TimeInForce    string  `json:"time_in_force"`
}

// Response reprents a response body.
type Response struct {
	ChildOrderAcceptanceID string `json:"child_order_acceptance_id"`
}

func sendChildOrder(b Body) (Response, error) {
	var resBody Response

	path := "/v1/me/sendchildorder"
	url := "https://api.bitflyer.com" + path
	method := http.MethodPost

	reqBody, err := json.Marshal(b)
	if err != nil {
		return resBody, err
	}

	req, err := http.NewRequest(method, url, bytes.NewReader(reqBody))
	if err != nil {
		return resBody, err
	}

	timestamp := auth.GenerateTimestamp()
	sign, err := auth.GenerateSign(timestamp, method, path, string(reqBody))
	if err != nil {
		return resBody, err
	}

	req.Header.Set("ACCESS-KEY", auth.GetAccessKey())
	req.Header.Set("ACCESS-TIMESTAMP", timestamp)
	req.Header.Set("ACCESS-SIGN", sign)
	req.Header.Set("Content-Type", "application/json")

	timeout := time.Duration(5 * time.Second)
	client := &http.Client{
		Timeout: timeout,
	}

	resp, err := client.Do(req)
	if err != nil {
		return resBody, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return resBody, err
	}

	if resp.StatusCode != 200 {
		return resBody, fmt.Errorf(string(body))
	}

	if err := json.Unmarshal(body, &resBody); err != nil {
		log.Println(string(body))
		return resBody, err
	}

	return resBody, nil
}
