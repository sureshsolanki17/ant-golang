package service

import (
	"encoding/json"
	"fmt"
	"io"
	"log"

	constants "github.com/sureshsolanki17/ant-golang/const"
)

type OrderResponse struct {
	Stat   string `json:"stat"`
	NOrdNo string `json:"NordNo"`
}

type Order struct {
	Complexty       string `json:"complexty"`
	DiscQty         string `json:"discqty"`
	Exchange        string `json:"exch"`
	PCode           string `json:"pCode"`
	PriceType       string `json:"prctyp"`
	Price           string `json:"price"`
	Quantity        int    `json:"qty"`
	Retention       string `json:"ret"`
	SymbolID        string `json:"symbol_id"`
	TradingSymbol   string `json:"trading_symbol"`
	TransactionType string `json:"transtype"`
	TriggerPrice    string `json:"trigPrice"`
}

func (app *AntApp) Buy(symbolID, tradingSymbol, price, transactionType string, quantity int) (*[]OrderResponse, error) {
	order := Order{
		Complexty:       "regular",
		DiscQty:         "0",
		Exchange:        app.Exchange,
		PCode:           "mis",
		PriceType:       transactionType,
		Price:           price,
		Quantity:        quantity,
		Retention:       "DAY",
		SymbolID:        symbolID,
		TradingSymbol:   tradingSymbol,
		TransactionType: constants.TransactionTypeBuy,
		TriggerPrice:    "",
	}

	jsonData, err := json.Marshal([]Order{order})
	if err != nil {
		return nil, fmt.Errorf("failed to marshal order: %v", err)
	}

	resp, err := app.MakeRequest("POST", constants.URLExecutePlaceOrder, string(jsonData))
	if err != nil {
		log.Fatalf("Error placing order: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return nil, err
	}

	var data *[]OrderResponse
	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (app *AntApp) BuyWithLimit(symbolID, tradingSymbol, price string, quantity int) (*[]OrderResponse, error) {
	return app.Buy(symbolID, tradingSymbol, price, constants.OrderTypeLimit, quantity)
}

func (app *AntApp) BuyWithMarket(symbolID, tradingSymbol, price string, quantity int) (*[]OrderResponse, error) {
	return app.Buy(symbolID, tradingSymbol, price, constants.OrderTypeMarket, quantity)
}

type PositionBookBody struct {
	Ret string `json:"ret"`
}

func (app *AntApp) PositionBook(DAY string) (*string, error) {
	u := constants.BaseURL + constants.URLPositionBook

	requestBody := PositionBookBody{
		Ret: DAY,
	}

	jsonBody, err := json.Marshal(requestBody)
	if err != nil {
		fmt.Println("Error marshaling request body:", err)
		return nil, err
	}

	req, err := app.MakeRequest("POST", u, string(jsonBody))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", app.Authorization)

	return nil, nil
}
