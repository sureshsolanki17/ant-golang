package antgolang

import "time"

const (
	name           string        = "antgolang"
	version        string        = "1.0"
	requestTimeout time.Duration = 7000 * time.Millisecond
	baseURI        string        = "https://ant.aliceblueonline.com/rest/AliceBlueAPIService/api/"
	contractURL    string        = "https://v2api.aliceblueonline.com/restpy/contract_master?exch="
)
