package plaid

import (
	"bytes"
	"encoding/json"
	"errors"
	"time"
)

type ExchangeTokenResponse struct {
	AccessToken string `json:"access_token"`
	ItemID      string `json:"item_id"`
	RequestID   string `json:"request_id"`
}

type ExchangeTokenRequest struct {
	ClientID    string `json:"client_id"`
	Secret      string `json:"secret"`
	PublicToken string `json:"public_token"`
}

func (c *Client) ExchangeToken(publicToken string) (*ExchangeTokenResponse, error) {
	if publicToken == "" {
		return nil, errors.New("public token cannot be empty")
	}

	req := ExchangeTokenRequest{
		ClientID:    c.id,
		Secret:      c.secret,
		PublicToken: publicToken,
	}

	body := bytes.NewBuffer(nil)
	if err := json.NewEncoder(body).Encode(req); err != nil {
		return nil, err
	}

	var response ExchangeTokenResponse
	if err := c.postAndUnmarshal("/item/public_token/exchange", body, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

type CreatePublicTokenRequest struct {
	ClientID    string `json:"client_id"`
	Secret      string `json:"secret"`
	AccessToken string `json:"access_token"`
}

type CreatePublicTokenResponse struct {
	Expiration  time.Time `json:"expiration"`
	PublicToken string    `json:"public_token"`
	RequestID   string    `json:"request_id"`
}

func (c *Client) CreatePublicToken(accessToken string) (*CreatePublicTokenResponse, error) {
	if accessToken == "" {
		return nil, errors.New("public token cannot be empty")
	}

	req := CreatePublicTokenRequest{
		ClientID:    c.id,
		Secret:      c.secret,
		AccessToken: accessToken,
	}

	body := bytes.NewBuffer(nil)
	if err := json.NewEncoder(body).Encode(req); err != nil {
		return nil, err
	}

	var response CreatePublicTokenResponse
	if err := c.postAndUnmarshal("/item/public_token/create", body, &response); err != nil {
		return nil, err
	}
	return &response, nil
}
