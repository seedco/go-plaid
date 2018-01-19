package plaid

import (
	"bytes"
	"encoding/json"
	"errors"
)

func (c *LegacyClient) ExchangeToken(publicToken string) (*ExchangeTokenResponse, error) {
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
	if err := c.postAndUnmarshal("/exchange_token", body, &response); err != nil {
		return nil, err
	}
	return &response, nil
}
