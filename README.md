# go-schwab-sdk

A lightweight Go SDK for the [Charles Schwab API](https://developer.schwab.com/), covering OAuth authentication and trader (accounts/orders) endpoints.

## Requirements

- Go 1.27+
- Schwab developer app credentials (API key + secret) and a registered callback URL

## Install

```bash
go get go-schwab-client
```

## Packages

| Package     | Purpose                                                              |
|-------------|----------------------------------------------------------------------|
| `auth-api`  | OAuth sign-in URL generation and access/refresh token exchange.      |
| `trader-api`| Trader endpoints (accounts, orders) and their response models.       |
| `utils`     | Shared `HttpClient` interface and the generic `DoRequest[T]` helper.  |

## Authentication

Schwab uses an OAuth authorization-code flow. First generate the sign-in URL, have the user authorize, then exchange the returned code for an access token.

```go
package main

import (
	"net/http"

	authapi "github.com/cruxhoo/go-schwab-sdk/auth-api"
)

func main() {
	client := authapi.NewClient(
		"API_KEY",
		"API_SECRET",
		"https://127.0.0.1/callback",
		&http.Client{},
	)

	// 1. Send the user to the sign-in URL to authorize.
	signInURL := client.GenerateSignUrl()
	_ = signInURL

	// 2. Exchange the authorization code for a token.
	token, err := client.GetOrRefreshAccessToken(
		"AUTH_CODE_FROM_CALLBACK",
		"https://127.0.0.1/callback",
		"",
		authapi.GrantAuthCode,
	)
	if err != nil {
		panic(err)
	}

	// 3. Later, refresh the token when it expires.
	refreshed, err := client.GetOrRefreshAccessToken(
		"",
		"",
		token.RefreshToken,
		authapi.GrantRefreshToken,
	)
	if err != nil {
		panic(err)
	}
	_ = refreshed
}
```

## Trader API

Construct a `TraderClient` with a valid `AccessToken` and an `HttpClient`:

```go
import (
	"net/http"

    authapi "github.com/cruxhoo/go-schwab-sdk/auth-api"
	accountapi "github.com/cruxhoo/go-schwab-sdk/trader-api"
)

trader := accountapi.NewTraderClient(*token, &http.Client{})
trader.UpdateToken(*refreshed) // swap in a refreshed token when needed
```

> Note: trader endpoints such as `GetAccounts` are still being implemented.

## Testing

```bash
go test ./...
```

Both `auth-api` and `utils` ship with an in-memory `MockHttpClient` implementing the `utils.HttpClient` interface, so requests can be tested without network calls.

## License

See repository for license details.
