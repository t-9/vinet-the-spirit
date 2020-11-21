package board

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// Board represents a board information.
type Board struct {
	MidPrice float64 `json:"mid_price"`
	Bids     []Order `json:"bids"`
	Asks     []Order `json:"asks"`
}

// Order represents an order information.
type Order struct {
	Price float64 `json:"price"`
	Size  float64 `json:"size"`
}

func getBoard(code string) (Board, error) {
	var board Board

	url := "https://api.bitflyer.com/v1/getboard"

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return board, err
	}

	params := req.URL.Query()
	params.Add("product_code", code)
	req.URL.RawQuery = params.Encode()

	timeout := time.Duration(5 * time.Second)
	client := &http.Client{
		Timeout: timeout,
	}

	resp, err := client.Do(req)
	if err != nil {
		return board, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return board, err
	}

	if resp.StatusCode != 200 {
		return board, fmt.Errorf(string(body))
	}

	if err := json.Unmarshal(body, &board); err != nil {
		return board, err
	}

	return board, nil
}
