package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	constants "github.com/sureshsolanki17/ant-golang/const"
	"github.com/sureshsolanki17/ant-golang/model"
)

// History gets stock price history for given token, timeframe and dates.
// token: stock symbol, resolution: "1"(1min)/"D"(daily), fromDate/toDate: date range
// Returns OHLCV data in ChartHistoryResponse or error if request fails
func (app *AntApp) History(token, resolution, fromDate, toDate string) (*model.ChartHistoryResponse, error) {
	url := constants.BaseURL + constants.URLHistory

	requestBody := HistoryBody{
		Token:      token,
		Resolution: resolution,
		From:       fromDate,
		To:         toDate,
		Exchange:   app.Exchange,
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
	var data model.ChartHistoryResponse
	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

type HistoryBody struct {
	Token      string `json:"token"`
	Resolution string `json:"resolution"`
	From       string `json:"from"`
	To         string `json:"to"`
	Exchange   string `json:"exchange"`
}
