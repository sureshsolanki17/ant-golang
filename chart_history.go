package antgolang

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

var baseURI = ""

func (app *AntApp) ChartHistory(token, resolution, fromDate, toDate string) (*ChartHistoryResponse, error) {
	url := baseURI + URLHistory

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
	var data ChartHistoryResponse
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

type ChartHistoryResponse struct {
	Stat   string `json:"stat"`
	Result []struct {
		Volume float64 `json:"volume"`
		High   float64 `json:"high"`
		Low    float64 `json:"low"`
		Time   string  `json:"time"`
		Close  float64 `json:"close"`
		Open   float64 `json:"open"`
	} `json:"result"`
	Message interface{} `json:"message"`
}
