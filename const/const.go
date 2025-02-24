package constants

import "time"

var CurrentTime time.Time = time.Now()

const (
	Name           string        = "antgolang"
	Version        string        = "1.0"
	RequestTimeout time.Duration = 7000 * time.Millisecond
	BaseURL        string        = "https://ant.aliceblueonline.com/rest/AliceBlueAPIService/api/"
	ContractURL    string        = "https://v2api.aliceblueonline.com/restpy/contract_master?exch="
)

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
	URLGetAPIEncpkey     = "customer/getAPIEncpkey"
	URLGetUserSID        = "customer/getUserSID"
)
