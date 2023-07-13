package antgolang

import (
	"encoding/json"
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
	body, err := GinGet(URLAccountDetails, app.Authorization)
	if err != nil {
		return nil, err
	}

	var data *AccountDetailsResponse
	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, err
	}

	return data, nil
}
