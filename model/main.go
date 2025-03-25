package model

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

type GetScripQuoteDetailsResponse struct {
	Optiontype       string `json:"optiontype"`
	SQty             int    `json:"SQty"`
	VwapAveragePrice string `json:"vwapAveragePrice"`
	LTQ              string `json:"LTQ"`
	DecimalPrecision int    `json:"DecimalPrecision"`
	OpenPrice        string `json:"openPrice"`
	LTP              string `json:"LTP"`
	Ltp              string `json:"Ltp"`
	BRate            string `json:"BRate"`
	Defmktproval     string `json:"defmktproval"`
	Symbolname       string `json:"symbolname"`
	NoMktPro         string `json:"noMktPro"`
	BQty             int    `json:"BQty"`
	Mktpro           string `json:"mktpro"`
	LTT              string `json:"LTT"`
	TickSize         string `json:"TickSize"`
	Multiplier       int    `json:"Multiplier"`
	Strikeprice      string `json:"strikeprice"`
	TotalSell        string `json:"TotalSell"`
	High             string `json:"High"`
	Stat             string `json:"stat"`
	YearlyLowPrice   string `json:"yearlyLowPrice"`
	YearlyHighPrice  string `json:"yearlyHighPrice"`
	ExchFeedTime     string `json:"exchFeedTime"`
	BodLotQty        int    `json:"BodLotQty"`
	PrvClose         string `json:"PrvClose"`
	Change           string `json:"Change"`
	SRate            string `json:"SRate"`
	Series           string `json:"Series"`
	TotalBuy         string `json:"TotalBuy"`
	Low              string `json:"Low"`
	UniqueKey        string `json:"UniqueKey"`
	PerChange        string `json:"PerChange"`
	Companyname      string `json:"companyname"`
	TradeVolume      string `json:"TradeVolume"`
	TSymbl           string `json:"TSymbl"`
	Exp              string `json:"Exp"`
	LTD              string `json:"LTD"`
}

type ChartHistoryResponse struct {
	Stat    string      `json:"stat"`
	Result  []result    `json:"result"`
	Message interface{} `json:"message"`
}

type result struct {
	Volume float64 `json:"volume"`
	High   float64 `json:"high"`
	Low    float64 `json:"low"`
	Time   string  `json:"time"`
	Close  float64 `json:"close"`
	Open   float64 `json:"open"`
}
