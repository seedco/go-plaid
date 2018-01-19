package plaid

import (
	"bytes"
	"encoding/json"
	"errors"
)

type AuthGetResponse struct {
	Accounts []Account `json:"accounts"`
	Numbers  []struct {
		Account     string `json:"account"`
		AccountID   string `json:"account_id"`
		Routing     string `json:"routing"`
		WireRouting string `json:"wireRouting"`
	} `json:"numbers"`
	Item      map[string]interface{} `json:"item"`
	RequestID string                 `json:"request_id"`
}

type Account struct {
	AccountID string `json:"account_id"`
	Balances  struct {
		Available float64 `json:"available"`
		Current   float64 `json:"current"`
		Limit     float64 `json:"limit"`
	} `json:"balances"`
	Mask         string `json:"mask"`
	Name         string `json:"name"`
	OfficialName string `json:"official_name"`
	Type         string `json:"type"`
	SubType      string `json:"sub_type"`
}

type AuthRequest struct {
	ClientID    string `json:"client_id"`
	Secret      string `json:"secret"`
	AccessToken string `json:"access_token"`
}

func (c *Client) AuthGet(accessToken string) (*AuthGetResponse, error) {
	if accessToken == "" {
		return nil, errors.New("access token cannot be empty")
	}

	req := AuthRequest{
		ClientID:    c.id,
		Secret:      c.secret,
		AccessToken: accessToken,
	}

	body := bytes.NewBuffer(nil)
	if err := json.NewEncoder(body).Encode(req); err != nil {
		return nil, err
	}

	var response AuthGetResponse
	if err := c.postAndUnmarshal("/auth/get", body, &response); err != nil {
		return nil, err
	}
	return &response, nil
}
