package antgolang

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

var Authorization string
var CurrentTime time.Time = time.Now()

const (
	// Varieties

	// Products

	// Order types

	// Validities

	// Position Type

	// Transaction type
	TransactionTypeBuy  = "BUY"
	TransactionTypeSell = "SELL"

	// Exchanges
	ExchangeNSE = "NSE"
	ExchangeBSE = "BSE"
	ExchangeMCX = "MCX"
	ExchangeNFO = "NFO"
	ExchangeBFO = "BFO"
	ExchangeCDS = "CDS"
	ExchangeBCD = "BCD"

	// Margins segments

	// Order status
	OrderStatusComplete  = "COMPLETE"
	OrderStatusRejected  = "REJECTED"
	OrderStatusCancelled = "CANCELLED"
)

const (
	URLAccountDetails = "customer/accountDetails"
	URLGetRmsLimits   = "limits/getRmsLimits"
)

func GinGet(url, token string) ([]byte, error) {
	u := baseURI + url

	client := &http.Client{}
	req, err := http.NewRequest("GET", u, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		panic(err)
	}

	req.Header.Set("Authorization", token)

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return nil, err
	}

	return body, nil
}
