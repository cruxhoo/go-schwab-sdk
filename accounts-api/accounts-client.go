package accounts_api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"time"

	authapi "github.com/cruxhoo/go-schwab-sdk/auth-api"
	"github.com/cruxhoo/go-schwab-sdk/utils"
)

const (
	BaseApiUrl    = "https://api.schwabapi.com/trader/v1"
	UTCTimeFormat = "2006-01-02T15:04:05.000Z"
)

type AccountsClient struct {
	token      authapi.AccessToken
	httpClient utils.HttpClient
}

func NewTraderClient(token authapi.AccessToken, client utils.HttpClient) *AccountsClient {
	return &AccountsClient{
		token:      token,
		httpClient: client,
	}
}

func (c *AccountsClient) UpdateToken(token authapi.AccessToken) {
	c.token = token
}

func (c *AccountsClient) GetAccountNumbers() ([]AccountNumber, error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/accounts/accountNumbers", BaseApiUrl), nil)
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

func (c *AccountsClient) GetAccounts(accountHash, fields string) ([]Account, error) {
	url := fmt.Sprintf("%s/accounts", BaseApiUrl)
	if accountHash != "" {
		url = fmt.Sprintf("%s/accounts/%s", BaseApiUrl, accountHash)
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

func (c *AccountsClient) GetOrders(accountHash string, maxResults int64, fromTime, toTime time.Time) ([]Order, error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/accounts/%s/orders", BaseApiUrl, accountHash), nil)
	if err != nil {
		slog.Error("error creating get orders request", "error", err)
		return nil, err
	}
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.token.AccessToken))
	// set the query params
	queryParams := req.URL.Query()
	queryParams.Add("maxResults", fmt.Sprintf("%d", maxResults))
	queryParams.Add("fromEnteredTime", fromTime.UTC().Format(UTCTimeFormat))
	queryParams.Add("toEnteredTime", toTime.UTC().Format(UTCTimeFormat))
	req.URL.RawQuery = queryParams.Encode()

	orders, err := utils.DoRequest[[]Order](c.httpClient, req)

	if err != nil {
		slog.Error("error executing orders request", "error", err)
		return nil, err
	}
	return *orders, nil
}

func (c *AccountsClient) GetOrder(accountHash, orderId string) (*Order, error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/accounts/%s/orders/%s", BaseApiUrl, accountHash, orderId), nil)
	if err != nil {
		slog.Error("error creating get order request", "error", err)
		return nil, err
	}
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.token.AccessToken))
	order, err := utils.DoRequest[Order](c.httpClient, req)

	if err != nil {
		slog.Error("error executing get order request", "error", err)
		return nil, err
	}
	return order, nil
}

func (c *AccountsClient) PutOrder(accountHash string, order Order) error {
	byteOrder, err := json.Marshal(order)
	if err != nil {
		slog.Error("failed to encode order to json", "error", err)
		return err
	}
	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/accounts/%s/orders", BaseApiUrl, accountHash), bytes.NewReader(byteOrder))
	if err != nil {
		slog.Error("error creating account request", "error", err)
		return err
	}
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.token.AccessToken))
	req.Header.Add("Content-Type", "application/json")
	resp, err := c.httpClient.Do(req)
	if err != nil {
		slog.Error("failed to execute http request", "error", err)
		return err
	}

	if resp.StatusCode < http.StatusOK || resp.StatusCode > http.StatusMultipleChoices {
		slog.Error("received not 2xx status code for request", "responseStatus", resp.StatusCode)
		return fmt.Errorf("received not 2xx status code for request, status: %d", resp.StatusCode)
	}
	return nil
}

func (c *AccountsClient) DeleteOrder(accountHash, orderId string) error {
	req, err := http.NewRequest(http.MethodDelete, fmt.Sprintf("%s/accounts/%s/orders/%s", BaseApiUrl, accountHash, orderId), nil)
	if err != nil {
		slog.Error("error creating get order request", "error", err)
		return err
	}
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.token.AccessToken))
	resp, err := c.httpClient.Do(req)
	if err != nil {
		slog.Error("failed to execute http request", "error", err)
		return err
	}

	if resp.StatusCode < http.StatusOK || resp.StatusCode > http.StatusMultipleChoices {
		slog.Error("received not 2xx status code for request", "responseStatus", resp.StatusCode)
		return fmt.Errorf("received not 2xx status code for request, status: %d", resp.StatusCode)
	}
	return nil
}

func (c *AccountsClient) PreviewOrder(accountHash string, previewOrder OrderPreview) (*Order, error) {
	byteOrder, err := json.Marshal(previewOrder)
	if err != nil {
		slog.Error("failed to encode preview order to json", "error", err)
		return nil, err
	}
	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/accounts/%s/previewOrder", BaseApiUrl, accountHash), bytes.NewReader(byteOrder))
	if err != nil {
		slog.Error("error creating account request", "error", err)
		return nil, err
	}
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.token.AccessToken))
	req.Header.Add("Content-Type", "application/json")

	order, err := utils.DoRequest[Order](c.httpClient, req)
	if err != nil {
		slog.Error("error executing get preview order request", "error", err)
		return nil, err
	}
	return order, nil
}
