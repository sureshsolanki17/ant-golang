package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"log"

	constants "github.com/sureshsolanki17/ant-golang/const"
	"github.com/sureshsolanki17/ant-golang/service"
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

func PlaceOrder(a service.AntApp, orderType, price, symbolID, tradingSymbol, transactionType string, quantity int) (*[]OrderResponse, error) {
	order := Order{
		Complexty:       "regular",
		DiscQty:         "0",
		Exchange:        "NSE",
		PCode:           "mis",
		PriceType:       orderType,
		Price:           price,
		Quantity:        quantity,
		Retention:       "DAY",
		SymbolID:        symbolID,
		TradingSymbol:   tradingSymbol,
		TransactionType: transactionType,
		TriggerPrice:    "",
	}

	jsonData, err := json.Marshal([]Order{order})
	if err != nil {
		return nil, fmt.Errorf("failed to marshal order: %v", err)
	}

	resp, err := a.MakeRequest("POST", constants.URLExecutePlaceOrder, string(jsonData))
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
