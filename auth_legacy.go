package plaid

import (
	"bytes"
	"encoding/json"
	"errors"
)

type Account_legacy struct {
	ID      string `json:"_id"`
	ItemID  string `json:"_item"`
	UserID  string `json:"_user"`
	Balance struct {
		Available float64 `json:"available"`
		Current   float64 `json:"current"`
	} `json:"balance"`
	Meta struct {
		Number string `json:"number"`
		Name   string `json:"name"`
	} `json:"meta"`
	Numbers struct {
		Account     string `json:"account"`
		Routing     string `json:"routing"`
		WireRouting string `json:"wireRouting"`
	} `json:"numbers"`
	Type            string `json:"type"`
	InstitutionType string `json:"institution_type"`
}

type AuthGetResponse_legacy struct {
	Accounts    []Account_legacy `json:"accounts"`
	AccessToken string           `json:"access_token"`
}

func (c *LegacyClient) AuthGet(accessToken string) (*AuthGetResponse_legacy, error) {
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

	var response AuthGetResponse_legacy
	if err := c.postAndUnmarshal("/auth/get", body, &response); err != nil {
		return nil, err
	}
	return &response, nil
}
