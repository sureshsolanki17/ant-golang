package service

import (
	"fmt"
	"net/http"
	"strings"
)

type AntApp struct {
	BaseURL       string
	Authorization string
	HTTPClient    *http.Client
	Exchange      string
}

// New creates a new instance of AntApp
func NewConfig(cf Config) *AntApp {
	return &AntApp{
		BaseURL:    cf.BaseURL,
		HTTPClient: &http.Client{},
		Exchange:   strings.ToUpper(cf.Exchange),
	}
}

// SetToken sets the authorization token for the app
func (app *AntApp) SetToken(userID, token string) {
	app.Authorization = "Bearer " + userID + " " + token
}

// GetAuthorization retrieves the current authorization header
func (app *AntApp) GetAuthorization() string {
	return app.Authorization
}

// MakeRequest is a reusable function to send HTTP requests
func (app *AntApp) MakeRequest(method, endpoint string, body string) (*http.Response, error) {
	url := fmt.Sprintf("%s/%s", app.BaseURL, strings.TrimLeft(endpoint, "/"))

	req, err := http.NewRequest(method, url, strings.NewReader(body))
	if err != nil {
		return nil, err
	}

	// Add Authorization header
	if auth := app.GetAuthorization(); auth != "" {
		req.Header.Set("Authorization", auth)
		req.Header.Set("Content-Type", "application/json")
	}

	// Execute the HTTP request
	return app.HTTPClient.Do(req)
}
