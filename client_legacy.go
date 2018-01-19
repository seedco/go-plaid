package plaid

import (
	"errors"
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
		id = letacyTestClientId
		secret = legacyTestSecret
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
