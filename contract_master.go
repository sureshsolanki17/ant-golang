package antgolang

import (
	"encoding/json"
	"io"
	"net/http"
)

func (app *AntApp) ContractMaster() (*[]StockData, error) {
	url := contractURL + app.Exchange

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
		var data *ContractMasterNSEResponse
		err = json.Unmarshal(body, &data)
		if err != nil {
			return nil, err
		}

		return &data.NSE, nil
	} else {
		var data *ContractMasterBSEResponse
		err = json.Unmarshal(body, &data)
		if err != nil {
			return nil, err
		}

		return &data.BSE, nil
	}
}

func (app *AntApp) ContractINDICES(ExchangeType string) (*[]Stock, error) {
	url := contractURL + "INDICES"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	defer req.Body.Close()

	body, err := io.ReadAll(req.Body)
	if err != nil {
		return nil, err
	}

	var data *INDICES

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

type ContractMasterNSEResponse struct {
	NSE          []StockData `json:"NSE"`
	ContractDate string      `json:"contract_date"`
}

type StockData struct {
	Exch             string `json:"exch"`
	ExchangeSegment  string `json:"exchange_segment"`
	FormattedInsName string `json:"formatted_ins_name"`
	GroupName        string `json:"group_name"`
	InstrumentType   string `json:"instrument_type"`
	LotSize          string `json:"lot_size"`
	Pdc              string `json:"pdc"`
	Symbol           string `json:"symbol"`
	TickSize         string `json:"tick_size"`
	Token            string `json:"token"`
	TradingSymbol    string `json:"trading_symbol"`
}

type ContractMasterBSEResponse struct {
	BSE          []StockData `json:"BSE"`
	ContractDate string      `json:"contract_date"`
}

type INDICES struct {
	BSE          []Stock `json:"BSE"`
	MCX          []Stock `json:"MCX"`
	NSE          []Stock `json:"NSE"`
	ContractDate string  `json:"contract_date"`
}

type Stock struct {
	Symbol string `json:"symbol"`
	Token  int    `json:"token"`
}
