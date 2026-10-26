package trader_api

import (
	"fmt"
	authapi "go-schwab-client/auth-api"
	"go-schwab-client/utils"
	"log/slog"
	"net/http"
)

const TraderApi = "https://api.schwabapi.com/trader/v1"

type TraderClient struct {
	token      authapi.AccessToken
	httpClient utils.HttpClient
}

func NewTraderClient(token authapi.AccessToken, client utils.HttpClient) *TraderClient {
	return &TraderClient{
		token:      token,
		httpClient: client,
	}
}

func (c *TraderClient) UpdateToken(token authapi.AccessToken) {
	c.token = token
}

func (c *TraderClient) GetAccountNumbers() ([]AccountNumber, error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/accounts/accountNumbers", TraderApi), nil)
	if err != nil {
		slog.Error("error creating account request", "error", err)
		return nil, err
	}
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.token.AccessToken))
	resp, err := utils.DoRequest[[]AccountNumber](c.httpClient, req)
	if err != nil {
		slog.Error("error executing account request", "error", err)
		return nil, err
	}
	return *resp, nil
}

func (c *TraderClient) GetAccounts(accountHash, fields string) ([]Account, error) {
	url := fmt.Sprintf("%s/accounts", TraderApi)
	if accountHash != "" {
		url = fmt.Sprintf("%s/accounts/%s", TraderApi, accountHash)
	}
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		slog.Error("error creating account request", "error", err)
		return nil, err
	}
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.token.AccessToken))
	if fields != "" {
		queryParams := req.URL.Query()
		queryParams.Add("fields", fields)
		req.URL.RawQuery = queryParams.Encode()
	}
	if accountHash != "" {
		resp, err := utils.DoRequest[Account](c.httpClient, req)
		if err != nil {
			slog.Error("error executing account request", "error", err)
			return nil, err
		}
		items := make([]Account, 1)
		items[0] = *resp
		return items, nil
	}
	resp, err := utils.DoRequest[[]Account](c.httpClient, req)
	if err != nil {
		slog.Error("error executing account request", "error", err)
		return nil, err
	}
	return *resp, nil
}
