# go-plaid
Go client for Plaid. This is not a comprehensive client. It includes the following endpoints

## Endpoints

- POST /auth/get
- POST /item/public_token/exchange
- POST /item/public_token/create

## Usage
### Instantiate client
```
client := plaid.NewClient("some_id","some_secret",plaid.Sandbox)
```
### Call methods
```
resp, err := client.AuthGet("access token value here")
```

## Legacy Client
Included also is a legacy client, instantiated in the same way. It will automatically use the correct urls.

### Instantiate client
```
client := plaid.NewLegacyClient("some_id","some_secret",plaid.Sandbox)
```
