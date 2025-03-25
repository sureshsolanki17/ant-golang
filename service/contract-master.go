package service

import (
	"encoding/json"
	"io"
	"net/http"

	constants "github.com/sureshsolanki17/ant-golang/const"
	"github.com/sureshsolanki17/ant-golang/model"
)

func (app *AntApp) ContractMaster() (*[]model.StockData, error) {
	url := constants.ContractURL + app.Exchange

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if app.Exchange == "NSE" {
		var data *model.ContractMasterNSEResponse
		err = json.Unmarshal(body, &data)
		if err != nil {
			return nil, err
		}

		return &data.NSE, nil
	}

	if app.Exchange == "BSE" {
		var data *model.ContractMasterBSEResponse
		err = json.Unmarshal(body, &data)
		if err != nil {
			return nil, err
		}

		return &data.BSE, nil
	}

	return nil, nil
}

func (app *AntApp) ContractINDICES(ExchangeType string) (*[]model.Stock, error) {
	url := constants.BaseURL + "INDICES"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	defer req.Body.Close()

	body, err := io.ReadAll(req.Body)
	if err != nil {
		return nil, err
	}

	var data *model.INDICES

	err = json.Unmarshal(body, data)
	if err != nil {
		return nil, err
	}

	if ExchangeType == "BSE" {
		return &data.BSE, nil
	}
	if ExchangeType == "MCX" {
		return &data.MCX, nil
	}
	if ExchangeType == "NSE" {
		return &data.NSE, nil
	}

	return nil, nil
}
