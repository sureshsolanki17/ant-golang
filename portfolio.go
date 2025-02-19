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
