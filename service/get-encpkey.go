package service

import (
	"bytes"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	constants "github.com/sureshsolanki17/ant-golang/const"
)

type GetAPIEncKeyBody struct {
	UserId string `json:"userId"`
}

func (app *AntApp) getAPIEncpkey() (string, error) {
	u := constants.BaseURL + constants.URLGetAPIEncpkey
	body := GetAPIEncKeyBody{
		UserId: app.UserId,
	}

	jsonBody, err := json.Marshal(body)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest("POST", u, bytes.NewBuffer(jsonBody))
	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("unexpected status code: %d, body: %s", resp.StatusCode, string(respBody))
	}

	var result struct {
		EncKey string `json:"encKey"`
	}
	if err := json.Unmarshal(respBody, &result); err != nil {
		return "", err
	}

	return result.EncKey, nil
}

type GetUserSIDBody struct {
	UserId   string `json:"userId"`
	UserData string `json:"userData"`
}

func (app *AntApp) getUserSID() (string, error) {
	u := constants.BaseURL + constants.URLGetUserSID

	ud, err := app.generateUserData()
	if err != nil {
		return "", err
	}

	body := GetUserSIDBody{
		UserId:   app.UserId,
		UserData: ud,
	}

	jsonBody, err := json.Marshal(body)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest("POST", u, bytes.NewBuffer(jsonBody))
	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("unexpected status code: %d, body: %s", resp.StatusCode, string(respBody))
	}

	var result struct {
		SessionID string `json:"sessionID"`
	}
	if err := json.Unmarshal(respBody, &result); err != nil {
		return "", err
	}

	return result.SessionID, nil
}

func (app *AntApp) generateUserData() (string, error) {
	var err error

	dt := strings.Builder{}
	dt.WriteString(app.UserId)
	dt.WriteString(app.ApiKey)
	dt.WriteString(app.EncKey)

	t := dt.String()

	h := sha256.New()
	h.Write([]byte(t))

	b := fmt.Sprintf("%x", h.Sum(nil))

	return string(b), err
}
