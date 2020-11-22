package order

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"vinet/auth"
)

const (
	// TypeLimit represents that the order type is a limit.
	TypeLimit = "LIMIT"
	// TypeMarket represents that the order type is a market.
	TypeMarket = "MARKET"
	// TypeStop represents that the order type is a stop.
	TypeStop = "STOP"
	// TypeStopLimit represents that the order type is a stop limit.
	TypeStopLimit = "STOP_LIMIT"
	// TypeTrail represents that the order type is a trail.
	TypeTrail = "TRAIL"

	// SideBuy represents that the side is a buy.
	SideBuy = "BUY"
	// SideSell represents that the side is a sell.
	SideSell = "SELL"

	// TimeInForceGTC represents that the time in force is good till canceled.
	TimeInForceGTC = "GTC"
	// TimeInForceIOC represents that the time in force is immediate or cancel.
	TimeInForceIOC = "IOC"
	// TimeInForceFOK represents that the time in force is fill or kill.
	TimeInForceFOK = "FOK"

	// MethodSimple represents a special order that issues a single order.
	MethodSimple = "SIMPLE"
	// MethodIFD represents an IFD order.
	MethodIFD = "IFD"
	// MethodOCO represents an OCO order.
	MethodOCO = "OCO"
	// MethodIFDOCO represents an IFD-OCO order.
	MethodIFDOCO = "IFDOCO"
)

func SendOrder(b interface{}, path string) ([]byte, error) {
	url := "https://api.bitflyer.com" + path
	method := http.MethodPost

	reqBody, err := json.Marshal(b)
	if err != nil {
		return []byte{}, err
	}

	req, err := http.NewRequest(method, url, bytes.NewReader(reqBody))
	if err != nil {
		return []byte{}, err
	}

	timestamp := auth.GenerateTimestamp()
	sign, err := auth.GenerateSign(timestamp, method, path, string(reqBody))
	if err != nil {
		return []byte{}, err
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
		return []byte{}, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, err
	}

	if resp.StatusCode != 200 {
		return body, fmt.Errorf(string(body))
	}

	return body, nil
}
