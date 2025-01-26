package antgolang

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

var CurrentTime time.Time = time.Now()

const (
	// Varieties

	// Products

	// Order types
	OrderTypeMarket = "MARKET"
	OrderTypeLimit  = "LIMIT"
	OrderTypeSL     = "SL"
	OrderTypeSLM    = "SL-M"

	// Validities
	ValidityDay = "DAY"
	ValidityIOC = "IOC"
	ValidityTTL = "TTL"

	// Position Type
	PositionTypeDay       = "day"
	PositionTypeOvernight = "overnight"

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
	MarginsEquity    = "equity"
	MarginsCommodity = "commodity"

	// Order status
	OrderStatusComplete  = "COMPLETE"
	OrderStatusRejected  = "REJECTED"
	OrderStatusCancelled = "CANCELLED"
)

// API endpoints
const (
	URLAccountDetails    = "customer/accountDetails"
	URLGetRmsLimits      = "limits/getRmsLimits"
	URLPositionBook      = "positionAndHoldings/positionBook"
	URLHoldings          = "positionAndHoldings/holdings"
	URLExecutePlaceOrder = "placeOrder/executePlaceOrder"
	URLHistory           = "chart/history"
)

func GinGet(url, token string) ([]byte, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
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

// GinPost sends a POST request to the given URL with the given token and body