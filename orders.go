package antgolang

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (app *AntApp) ExecutePlaceOrder() (*[]ExecutePlaceOrderResponse, error) {
	url := baseURI + URLExecutePlaceOrder

	requestBody := ExecutePlaceOrderBody{
		Discqty:       "",
		TradingSymbol: "",
		Exch:          app.Exchange,
		Transtype:     "",
		Ret:           "",
		Prctyp:        "",
		Qty:           "",
		SymbolID:      "",
		Price:         "",
		TrigPrice:     "",
		PCode:         "",
		Complexty:     "",
		OrderTag:      "",
	}

	jsonBody, err := json.Marshal(requestBody)
	if err != nil {
		fmt.Println("Error marshaling request body:", err)
		return nil, err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonBody))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", app.Authorization)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var data *[]ExecutePlaceOrderResponse
	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

type ExecutePlaceOrderBody struct {
	Discqty       string `json:"discqty"`
	TradingSymbol string `json:"trading_symbol"`
	Exch          string `json:"exch"`
	Transtype     string `json:"transtype"`
	Ret           string `json:"ret"`
	Prctyp        string `json:"prctyp"`
	Qty           string `json:"qty"`
	SymbolID      string `json:"symbol_id"`
	Price         string `json:"price"`
	TrigPrice     string `json:"trigPrice"`
	PCode         string `json:"pCode"`
	Complexty     string `json:"complexty"`
	OrderTag      string `json:"orderTag"`
}

type ExecutePlaceOrderResponse struct {
	Stat   string `json:"stat"`
	NOrdNo string `json:"NOrdNo"`
}
