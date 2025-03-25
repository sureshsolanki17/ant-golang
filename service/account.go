package service

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	constants "github.com/sureshsolanki17/ant-golang/const"
	"github.com/sureshsolanki17/ant-golang/model"
)

type AccountDetailsResponse struct {
	AccountStatus string   `json:"accountStatus"`
	DpType        string   `json:"dpType"`
	AccountID     string   `json:"accountId"`
	SBrokerName   string   `json:"sBrokerName"`
	Product       []string `json:"product"`
	AccountName   string   `json:"accountName"`
	CellAddr      string   `json:"cellAddr"`
	EmailAddr     string   `json:"emailAddr"`
	ExchEnabled   string   `json:"exchEnabled"`
}

func (app *AntApp) GetProfile() (*AccountDetailsResponse, error) {
	resp, err := app.MakeRequest("GET", constants.URLAccountDetails, "")
	if err != nil {
		log.Fatalf("Error placing order: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return nil, err
	}

	var data *AccountDetailsResponse
	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// ant connecton as app
func (app *AntApp) GetFund() ([]model.GetRmsLimitsResponse, error) {
	u := constants.URLGetRmsLimits

	resp, err := app.MakeRequest("GET", u, "")
	if err != nil {
		log.Fatalf("Error placing order: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return nil, err
	}

	var data []model.GetRmsLimitsResponse
	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (app *AntApp) Holdings() (*model.HoldingsResponse, error) {
	u := constants.BaseURL + constants.URLHoldings
	req, err := http.NewRequest("GET", u, strings.NewReader(""))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", app.Authorization)

	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		log.Fatalf("Error making request: %v", err)
	}
	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		log.Fatalf("Request failed with status code: %d", response.StatusCode)
	}

	// Read the response body
	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatalf("Error reading response body: %v", err)
	}

	var data *model.HoldingsResponse
	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, err
	}

	return data, nil
}
