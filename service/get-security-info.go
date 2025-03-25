package service

import (
	"encoding/json"
	"fmt"
	"io"

	constants "github.com/sureshsolanki17/ant-golang/const"
	"github.com/sureshsolanki17/ant-golang/model"
)

type getScripBody struct {
	Exch   string `json:"exch"`
	Symbol string `json:"symbol"`
}

func (app *AntApp) GetScripQuoteDetails(symbol string) (*model.GetScripQuoteDetailsResponse, error) {
	reqBody := getScripBody{
		Exch:   app.Exchange,
		Symbol: symbol,
	}

	jsonData, err := json.Marshal(reqBody)
	if err != nil {
		return nil, fmt.Errorf("error marshaling request body: %v", err)
	}

	resp, err := app.MakeRequest("POST", constants.URLGetScripQuoteDetails, string(jsonData))
	if err != nil {
		return nil, fmt.Errorf("error making request: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return nil, err
	}

	var data *model.GetScripQuoteDetailsResponse
	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, err
	}

	return data, nil
}
