package service

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	constants "github.com/sureshsolanki17/ant-golang/const"
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
	if app.GetAuthorization() == "" {
		return nil, fmt.Errorf("authorization token is not set")
	}

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
func (app *AntApp) GetFund() ([]GetRmsLimitsResponse, error) {
	if app.GetAuthorization() == "" {
		return nil, fmt.Errorf("authorization token is not set")
	}

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

	var data []GetRmsLimitsResponse
	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (app *AntApp) Holdings() (*HoldingsResponse, error) {
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

	var data *HoldingsResponse
	err = json.Unmarshal(body, &data)
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

type GetRmsLimitsResponse struct {
	Symbol                 string `json:"symbol"`
	CncMarginUsed          string `json:"cncMarginUsed"`
	Spanmargin             string `json:"spanmargin"`
	BranchAdhoc            string `json:"branchAdhoc"`
	AdhocMargin            string `json:"adhocMargin"`
	Payoutamount           string `json:"payoutamount"`
	CdsSpreadBenefit       string `json:"cdsSpreadBenefit"`
	Adhocscripmargin       string `json:"adhocscripmargin"`
	Exposuremargin         string `json:"exposuremargin"`
	Scripbasketmargin      string `json:"scripbasketmargin"`
	Credits                string `json:"credits"`
	Segment                string `json:"segment"`
	Net                    string `json:"net"`
	Turnover               string `json:"turnover"`
	Grossexposurevalue     string `json:"grossexposurevalue"`
	MfssAmountUsed         string `json:"mfssAmountUsed"`
	RealizedMtomPrsnt      string `json:"realizedMtomPrsnt"`
	Product                string `json:"product"`
	Stat                   string `json:"stat"`
	CncSellCrditPrsnt      string `json:"cncSellCrditPrsnt"`
	Debits                 string `json:"debits"`
	Varmargin              string `json:"varmargin"`
	Multiplier             string `json:"multiplier"`
	Elm                    string `json:"elm"`
	Mfamount               string `json:"mfamount"`
	Cashmarginavailable    string `json:"cashmarginavailable"`
	BrokeragePrsnt         string `json:"brokeragePrsnt"`
	CncRealizedMtomPrsnt   string `json:"cncRealizedMtomPrsnt"`
	NotionalCash           string `json:"notionalCash"`
	Directcollateralvalue  string `json:"directcollateralvalue"`
	CncBrokeragePrsnt      string `json:"cncBrokeragePrsnt"`
	Valueindelivery        string `json:"valueindelivery"`
	NfoSpreadBenefit       string `json:"nfoSpreadBenefit"`
	Losslimit              string `json:"losslimit"`
	Subtotal               string `json:"subtotal"`
	RmsPayInAmnt           string `json:"rmsPayInAmnt"`
	UnrealizedMtomPrsnt    string `json:"unrealizedMtomPrsnt"`
	CoverOrderMarginPrsnt  string `json:"coverOrderMarginPrsnt"`
	Exchange               string `json:"exchange"`
	Category               string `json:"category"`
	Collateralvalue        string `json:"collateralvalue"`
	RmsIpoAmnt             string `json:"rmsIpoAmnt"`
	CncUnrealizedMtomPrsnt string `json:"cncUnrealizedMtomPrsnt"`
	PremiumPrsnt           string `json:"premiumPrsnt"`
}
