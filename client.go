package plaid

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

const (
	// Environment constants
	Production  = "production"
	Development = "development"
	Sandbox     = "sandbox"

	// Environment Urls
	productionUrl  = "https://production.plaid.com"
	developmentUrl = "https://development.plaid.com"
	sandboxUrl     = "https://sandbox.plaid.com"

	// Legacy sandbox
	legacyProductionUrl  = "https://api.plaid.com"
	legacyDevelopmentUrl = "https://tartan.plaid.com"

	// Legacy test values
	legacyTestClientId = "test_id"
	legacyTestSecret   = "test_secret"
)

type client struct {
	id             string
	secret         string
	environmentUrl string
	httpClient     *http.Client
}

type Client struct {
	client
}

func NewClient(id, secret, environment string) (*Client, error) {
	httpClient := &http.Client{}

	var environmentUrl string
	switch environment {
	case Production:
		environmentUrl = productionUrl
	case Development:
		environmentUrl = developmentUrl
	case Sandbox:
		environmentUrl = sandboxUrl
	default:
		return nil, errors.New("invalid environment")
	}

	client := &Client{
		client{
			id:             id,
			secret:         secret,
			environmentUrl: environmentUrl,
			httpClient:     httpClient,
		},
	}

	return client, nil
}

func (c *client) postAndUnmarshal(endpoint string, body io.Reader, result interface{}) error {
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
	// throw an error on any non-200 response
	if res.StatusCode/100 != 2 {
		return fmt.Errorf("Non-200 response when exeucting plaid post request %v: %d %v", endpoint, res.StatusCode, res.Status)
	}

	defer res.Body.Close()
	if err := json.NewDecoder(res.Body).Decode(result); err != nil {
		return fmt.Errorf("Error when decoding plaid post response %v: %v", endpoint, err)
	}
	return nil
}
