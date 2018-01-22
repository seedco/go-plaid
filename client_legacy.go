package plaid

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

type LegacyClient struct {
	client
}

func NewLegacyClient(id, secret, environment string) (*LegacyClient, error) {
	httpClient := &http.Client{}

	var environmentUrl string
	switch environment {
	case Production:
		environmentUrl = legacyProductionUrl
	case Development, Sandbox:
		environmentUrl = legacyDevelopmentUrl
	default:
		return nil, errors.New("invalid environment")
	}

	client := &LegacyClient{
		client{
			id:             id,
			secret:         secret,
			environmentUrl: environmentUrl,
			httpClient:     httpClient,
		},
	}

	return client, nil
}

type errorMessageLegacy struct {
	Code    int    `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
	Resolve string `json:"resolve,omitempty"`
}

func (c *LegacyClient) postAndUnmarshal(endpoint string, body io.Reader, result interface{}) error {
	req, err := http.NewRequest("POST", string(c.environmentUrl)+endpoint, body)
	if err != nil {
		return fmt.Errorf("Error when creating plaid post request %v: %v", endpoint, err)
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("User-Agent", "seed-plaid-go")
	res, err := c.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("Error when executing plaid post request %v: %v", endpoint, err)
	}
	defer res.Body.Close()
	// throw an error on any non-200 response
	if res.StatusCode/100 != 2 {
		var errorMsg errorMessageLegacy
		if err := json.NewDecoder(res.Body).Decode(&errorMsg); err != nil {
			return fmt.Errorf("Error when decoding plaid error response %v: %v", endpoint, err)
		}
		return fmt.Errorf("Non-200 response when exeucting plaid post request %v: %d %v", endpoint, res.StatusCode, errorMsg.Message)
	}

	if err := json.NewDecoder(res.Body).Decode(result); err != nil {
		return fmt.Errorf("Error when decoding plaid post response %v: %v", endpoint, err)
	}
	return nil
}
