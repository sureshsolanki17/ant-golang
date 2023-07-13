package antgolang

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type PositionBookBody struct {
	Ret string `json:"ret"`
}

func (app *AntApp) PositionBook(DAY string) (*string, error) {
	u := baseURI + URLPositionBook

	requestBody := PositionBookBody{
		Ret: DAY,
	}

	jsonBody, err := json.Marshal(requestBody)
	if err != nil {
		fmt.Println("Error marshaling request body:", err)
		return nil, err
	}

	req, err := http.NewRequest("POST", u, bytes.NewBuffer(jsonBody))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", app.Authorization)

	return nil, nil
}

func (app *AntApp) Holdings() (*HoldingsResponse, error) {
	u := baseURI + URLHoldings
	b, err := GinGet(u, app.Authorization)
	if err != nil {
		panic(err)
	}

	var data *HoldingsResponse
	err = json.Unmarshal(b, &data)
	if err != nil {
		return nil, err
	}
	
	return data, nil
}

type HoldingsResponse struct {
	Stat       string `json:"stat"`
	Clientid   string `json:"clientid"`
	HoldingVal []struct {
		WCqty           string      `json:"WCqty"`
		BSEHOldingValue string      `json:"BSEHOldingValue"`
		Hsflag          string      `json:"hsflag"`
		Series1         string      `json:"Series1"`
		HUqty           string      `json:"HUqty"`
		YSXHOldingValue string      `json:"YSXHOldingValue"`
		CSEHOldingValue string      `json:"CSEHOldingValue"`
		Ttrind          string      `json:"Ttrind"`
		DaysMTM         string      `json:"DaysMTM"`
		Csflag          string      `json:"csflag"`
		WHqty           string      `json:"WHqty"`
		Pcode           string      `json:"Pcode"`
		Price           string      `json:"Price"`
		BuyQty          string      `json:"BuyQty"`
		Exch4           string      `json:"Exch4"`
		Bsetsym         string      `json:"Bsetsym"`
		Exch5           string      `json:"Exch5"`
		LTcse           string      `json:"LTcse"`
		MCXHOldingValue string      `json:"MCXHOldingValue"`
		Holdqty         string      `json:"Holdqty"`
		Exch1           string      `json:"Exch1"`
		Exch2           string      `json:"Exch2"`
		Exch3           string      `json:"Exch3"`
		LTysx           string      `json:"LTysx"`
		Haircut         string      `json:"Haircut"`
		Scripcode       string      `json:"Scripcode"`
		LTPValuation    string      `json:"LTPValuation"`
		NSEHOldingValue string      `json:"NSEHOldingValue"`
		Ysxtsym         string      `json:"Ysxtsym"`
		Ltp             string      `json:"Ltp"`
		Coltype         string      `json:"Coltype"`
		Btst            string      `json:"Btst"`
		LTmcxsxcm       string      `json:"LTmcxsxcm"`
		Usedqty         string      `json:"Usedqty"`
		Token5          string      `json:"Token5"`
		Nsetsym         string      `json:"Nsetsym"`
		CUqty           string      `json:"CUqty"`
		Token2          string      `json:"Token2"`
		Token1          string      `json:"Token1"`
		Token4          string      `json:"Token4"`
		SellableQty     string      `json:"SellableQty"`
		Token3          string      `json:"Token3"`
		Mcxsxcmsym      string      `json:"Mcxsxcmsym"`
		Csetsym         string      `json:"Csetsym"`
		LTnse           string      `json:"LTnse"`
		Series          string      `json:"Series"`
		Colqty          string      `json:"Colqty"`
		ExchSeg5        interface{} `json:"ExchSeg5"`
		ExchSeg2        string      `json:"ExchSeg2"`
		ExchSeg1        string      `json:"ExchSeg1"`
		ExchSeg4        interface{} `json:"ExchSeg4"`
		LTbse           string      `json:"LTbse"`
		ExchSeg3        interface{} `json:"ExchSeg3"`
		Isin            string      `json:"isin"`
		Tprod           string      `json:"Tprod"`
	} `json:"HoldingVal"`
	Totalval struct {
		TotalMCXHoldingValue string `json:"TotalMCXHoldingValue"`
		TotalCSEHoldingValue string `json:"TotalCSEHoldingValue"`
		TotalNSEHoldingValue string `json:"TotalNSEHoldingValue"`
		TotalYSXHoldingValue string `json:"TotalYSXHoldingValue"`
		TotalBSEHoldingValue string `json:"TotalBSEHoldingValue"`
	} `json:"Totalval"`
}
