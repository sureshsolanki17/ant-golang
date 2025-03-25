package service

import (
	"fmt"
	"net/http"
	"strings"
)

type AntApp struct {
	BaseURL       string
	UserId        string
	EncKey        string
	UserData      string
	ApiKey        string
	Authorization string
	HTTPClient    *http.Client
	Exchange      string
	SessionID     string
}

// NewConfig creates a new AntApp instance with the provided configuration.
// It validates the configuration and initializes the client with default HTTP client.
// Returns an error if the configuration is invalid.
func NewConfig(cf Config) (*AntApp, error) {
	if err := cf.Validate(); err != nil {
		return nil, fmt.Errorf("invalid configuration: %w", err)
	}

	return &AntApp{
		BaseURL:    cf.BaseURL,
		HTTPClient: &http.Client{},
		Exchange:   strings.ToUpper(cf.Exchange),
		ApiKey:     cf.APIKey,
		UserId:     cf.UserId,
	}, nil
}

// Connect establishes a connection by obtaining encryption key and session ID.
// It validates required credentials (ApiKey and UserId) and sets up authorization.
// Returns an error if authentication fails or required credentials are missing.
func (app *AntApp) Connect() error {
	if app.ApiKey == "" || app.UserId == "" {
		return fmt.Errorf("apikey or client id not provided")
	}

	enc, err := app.getAPIEncpkey()
	if err != nil {
		return fmt.Errorf("api encp key not generat")
	}

	app.EncKey = enc

	session, err := app.getUserSID()
	if err != nil {
		return err
	}

	app.SessionID = session
	app.Authorization = "Bearer " + app.UserId + " " + session

	return err
}

// GetAuthorization returns the current authorization token.
// The token is in format "Bearer {UserId} {SessionID}".
func (app *AntApp) GetAuthorization() string {
	return app.Authorization
}

// MakeRequest performs an HTTP request to the specified endpoint.
// Parameters:
//   - method: HTTP method (GET, POST, etc.)
//   - endpoint: API endpoint path
//   - body: Request body as string
//
// Returns the HTTP response and any error encountered.
// Automatically adds authorization headers if available.
func (app *AntApp) MakeRequest(method, endpoint string, body string) (*http.Response, error) {
	if app.GetAuthorization() == "" {
		return nil, fmt.Errorf("authorization token is not set")
	}
	url := fmt.Sprintf("%s/%s", app.BaseURL, strings.TrimLeft(endpoint, "/"))

	req, err := http.NewRequest(method, url, strings.NewReader(body))
	if err != nil {
		return nil, err
	}

	if auth := app.GetAuthorization(); auth != "" {
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", auth)
	}

	return app.HTTPClient.Do(req)
}
