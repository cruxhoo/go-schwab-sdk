package accounts_api

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"regexp"
	"strings"
	"testing"
	"time"
	"uuid"

	authapi "github.com/cruxhoo/go-schwab-sdk/auth-api"
	"github.com/cruxhoo/go-schwab-sdk/utils"

	"github.com/google/go-cmp/cmp"
)

func getMockAccount() Account {
	return Account{
		SecuritiesAccount: SecuritiesAccount{
			AccountNumber: uuid.NewV7().String(),
			IsDayTrader:   false,
		},
		AggregatedBalance: AggregatedBalance{
			CurrentLiquidationValue: rand.Float64(),
			LiquidationValue:        rand.Float64(),
		},
	}
}

func getMockAccounts(count int64) []Account {
	accounts := make([]Account, count)
	for i := range count {
		accounts[i] = getMockAccount()
	}
	return accounts
}

func getMockOrder() Order {
	return Order{
		Session:   uuid.NewV7().String(),
		Duration:  uuid.NewV7().String(),
		OrderType: uuid.NewV7().String(),
		OrderId:   rand.Int63(),
	}
}

func getMockOrders(count int64) []Order {
	orders := make([]Order, count)
	for i := range count {
		orders[i] = getMockOrder()
	}
	return orders
}

func TestAccountsClient_GetAccountNumbers(t *testing.T) {
	accountNumbers := []AccountNumber{
		{
			AccountNumber: uuid.NewV7().String(),
			HashValue:     uuid.NewV7().String(),
		},
		{
			AccountNumber: uuid.NewV7().String(),
			HashValue:     uuid.NewV7().String(),
		},
		{
			AccountNumber: uuid.NewV7().String(),
			HashValue:     uuid.NewV7().String(),
		},
	}
	strAccounts, _ := json.Marshal(accountNumbers)
	mockHttpResp := http.Response{
		StatusCode: http.StatusOK,
		Body:       io.NopCloser(bytes.NewReader(strAccounts)),
	}

	mockHttp := utils.MockHttpClient{}
	mockHttp.SetResponse(&mockHttpResp)

	mockAccessToken := authapi.AccessToken{
		ExpiresIn:    3200,
		TokenType:    "random",
		Scope:        "random",
		RefreshToken: uuid.NewV7().String(),
		AccessToken:  uuid.NewV7().String(),
		IdToken:      uuid.NewV7().String(),
	}

	traderClient := NewTraderClient(mockAccessToken, &mockHttp)
	acc, err := traderClient.GetAccountNumbers()
	if err != nil {
		t.Errorf("unexpected error %s", err)
	}

	if acc == nil {
		t.Error("accountNumbers response was nil")
	}

	if len(acc) != len(accountNumbers) {
		t.Errorf("numbers of accounts do not match, expected: %d got: %d", len(accountNumbers), len(acc))
	}

	if diff := cmp.Diff(acc, accountNumbers); diff != "" {
		t.Errorf("results don't match %s", diff)
	}

	req := mockHttp.ReceivedRequest

	expectedUrl := "https://api.schwabapi.com/trader/v1/accounts/accountNumbers"

	if !strings.EqualFold(req.URL.String(), expectedUrl) {
		t.Errorf("request url doesn't match expected: %s got %s", expectedUrl, req.URL.String())
	}

	if auth := req.Header.Get("Authorization"); auth == "" {
		t.Error("missing auth header")
	} else {
		re := regexp.MustCompile(`Bearer\s[a-zA-Z0-9\-]*`)
		if !re.Match([]byte(auth)) {
			t.Error("header value doesn't match expected pattern")
		}
	}
}

func TestAccountsClient_GetAccountNumbers_HandleErrors(t *testing.T) {
	mockHttp := utils.MockHttpClient{}
	mockHttp.SetErrorResponse(errors.New("something happened here"))
	mockAccessToken := authapi.AccessToken{
		ExpiresIn:    3200,
		TokenType:    "random",
		Scope:        "random",
		RefreshToken: uuid.NewV7().String(),
		AccessToken:  uuid.NewV7().String(),
		IdToken:      uuid.NewV7().String(),
	}
	traderClient := NewTraderClient(mockAccessToken, &mockHttp)
	if _, err := traderClient.GetAccountNumbers(); err == nil {
		t.Error("expected function to return error")
	}
}

func TestAccountsClient_GetAccounts_All(t *testing.T) {
	mockAccounts := getMockAccounts(3)
	strMockAccounts, _ := json.Marshal(mockAccounts)
	mockResponse := http.Response{
		StatusCode: http.StatusOK,
		Body:       io.NopCloser(bytes.NewReader(strMockAccounts)),
	}

	mockClient := utils.MockHttpClient{}
	mockClient.SetResponse(&mockResponse)

	mockAccessToken := authapi.AccessToken{
		ExpiresIn:    3200,
		TokenType:    "random",
		Scope:        "random",
		RefreshToken: uuid.NewV7().String(),
		AccessToken:  uuid.NewV7().String(),
		IdToken:      uuid.NewV7().String(),
	}

	traderCLient := NewTraderClient(mockAccessToken, &mockClient)

	accounts, err := traderCLient.GetAccounts("", "positions")

	if err != nil {
		t.Errorf("unexpected error %s", err)
	}

	if diff := cmp.Diff(accounts, mockAccounts); diff != "" {
		t.Errorf("accounts don't match %s", diff)
	}

	req := mockClient.ReceivedRequest

	expectedUrl := "https://api.schwabapi.com/trader/v1/accounts?fields=positions"

	if !strings.EqualFold(req.URL.String(), expectedUrl) {
		t.Errorf("request url doesn't match expected: %s got %s", expectedUrl, req.URL.String())
	}

	if auth := req.Header.Get("Authorization"); auth == "" {
		t.Error("missing auth header")
	} else {
		re := regexp.MustCompile(`Bearer\s[a-zA-Z0-9\-]*`)
		if !re.Match([]byte(auth)) {
			t.Error("header value doesn't match expected pattern")
		}
	}

}

func TestAccountsClient_GetAccounts_One(t *testing.T) {
	mockAccounts := getMockAccounts(1)
	strMockAccounts, _ := json.Marshal(mockAccounts[0])
	mockResponse := http.Response{
		StatusCode: http.StatusOK,
		Body:       io.NopCloser(bytes.NewReader(strMockAccounts)),
	}

	mockClient := utils.MockHttpClient{}
	mockClient.SetResponse(&mockResponse)

	mockAccessToken := authapi.AccessToken{
		ExpiresIn:    3200,
		TokenType:    "random",
		Scope:        "random",
		RefreshToken: uuid.NewV7().String(),
		AccessToken:  uuid.NewV7().String(),
		IdToken:      uuid.NewV7().String(),
	}

	traderCLient := NewTraderClient(mockAccessToken, &mockClient)
	accountHash := uuid.NewV7().String()

	accounts, err := traderCLient.GetAccounts(accountHash, "positions")

	if err != nil {
		t.Errorf("unexpected error %s", err)
	}

	if diff := cmp.Diff(accounts, mockAccounts); diff != "" {
		t.Errorf("accounts don't match %s", diff)
	}

	req := mockClient.ReceivedRequest

	expectedUrl := fmt.Sprintf("https://api.schwabapi.com/trader/v1/accounts/%s?fields=positions", accountHash)

	if !strings.EqualFold(req.URL.String(), expectedUrl) {
		t.Errorf("request url doesn't match expected: %s got %s", expectedUrl, req.URL.String())
	}

	if auth := req.Header.Get("Authorization"); auth == "" {
		t.Error("missing auth header")
	} else {
		re := regexp.MustCompile(`Bearer\s[a-zA-Z0-9\-]*`)
		if !re.Match([]byte(auth)) {
			t.Error("header value doesn't match expected pattern")
		}
	}
}

func TestGetOrders_HandleHttpError(t *testing.T) {
	expectedError := errors.New("something happened in downtown")
	mockClient := utils.MockHttpClient{}
	mockClient.SetErrorResponse(expectedError)
	mockAccessToken := authapi.AccessToken{
		ExpiresIn:    3200,
		TokenType:    "random",
		Scope:        "random",
		RefreshToken: uuid.NewV7().String(),
		AccessToken:  uuid.NewV7().String(),
		IdToken:      uuid.NewV7().String(),
	}

	traderClient := NewTraderClient(mockAccessToken, &mockClient)
	accountHash := uuid.NewV7().String()
	shouldBeNil, err := traderClient.GetOrders(accountHash, 20, time.Now().Add(-10*time.Hour), time.Now())

	if err == nil {
		t.Error("expected error response")
	}

	if shouldBeNil != nil {
		t.Error("return values should be nil")
	}

}

func TestAccountsClient_GetOrders(t *testing.T) {
	mockOrder := getMockOrders(5)
	strMockAccounts, _ := json.Marshal(mockOrder)
	mockResponse := http.Response{
		StatusCode: http.StatusOK,
		Body:       io.NopCloser(bytes.NewReader(strMockAccounts)),
	}

	mockClient := utils.MockHttpClient{}
	mockClient.SetResponse(&mockResponse)
	mockAccessToken := authapi.AccessToken{
		ExpiresIn:    3200,
		TokenType:    "random",
		Scope:        "random",
		RefreshToken: uuid.NewV7().String(),
		AccessToken:  uuid.NewV7().String(),
		IdToken:      uuid.NewV7().String(),
	}

	traderClient := NewTraderClient(mockAccessToken, &mockClient)
	accountHash := uuid.NewV7().String()
	orders, err := traderClient.GetOrders(accountHash, 20, time.Now().Add(-10*time.Hour), time.Now())
	if err != nil {
		t.Errorf("unexpected error %s", err)
	}
	if orders == nil {
		t.Error("getOrders returned nil value")
	}
	if diff := cmp.Diff(orders, mockOrder); diff != "" {
		t.Errorf("getOrders return value mismatch %s", diff)
	}

	req := mockClient.ReceivedRequest
	expectedPath := fmt.Sprintf("/trader/v1/accounts/%s/orders", accountHash)
	if req.Method != http.MethodGet {
		t.Errorf("http method mismatch expected: %s got: %s", http.MethodGet, req.Method)
	}
	if !strings.EqualFold(expectedPath, req.URL.Path) {
		t.Errorf("request path mismatch expected: %s got: %s", expectedPath, req.URL.Path)
	}
	// validate query params
	queryParams := req.URL.Query()
	// re for matching time in format 2026-10-26T07:28:35.733Z
	dateRe := regexp.MustCompile(`[0-9]{4}-[0-9]{2}-[0-9]{2}T[0-9]{2}:[0-9]{2}:[0-9]{2}\.[0-9]{3}Z`)
	if !dateRe.Match([]byte(queryParams.Get("fromEnteredTime"))) {
		t.Errorf("fromEnteredTime doesnt match expected format")
	}
	if !dateRe.Match([]byte(queryParams.Get("toEnteredTime"))) {
		t.Errorf("toEnteredTime doesnt match expected format")
	}
}

func TestAccountsClient_PutOrder(t *testing.T) {
	mockOrder := getMockOrder()
	strMockOrder, _ := json.Marshal(mockOrder)
	mockResponse := http.Response{
		StatusCode: http.StatusOK,
		Body:       io.NopCloser(bytes.NewReader(strMockOrder)),
	}

	mockNewOrder := getMockOrder()

	mockClient := utils.MockHttpClient{}
	mockClient.SetResponse(&mockResponse)
	mockAccessToken := authapi.AccessToken{
		ExpiresIn:    3200,
		TokenType:    "random",
		Scope:        "random",
		RefreshToken: uuid.NewV7().String(),
		AccessToken:  uuid.NewV7().String(),
		IdToken:      uuid.NewV7().String(),
	}

	traderClient := NewTraderClient(mockAccessToken, &mockClient)
	accountHash := uuid.NewV7().String()
	err := traderClient.PutOrder(accountHash, mockNewOrder)
	if err != nil {
		t.Errorf("unexpected error %s", err)
	}

	req := mockClient.ReceivedRequest
	expectedPath := fmt.Sprintf("/trader/v1/accounts/%s/orders", accountHash)
	if req.Method != http.MethodPost {
		t.Errorf("http method mismatch expected: %s got: %s", http.MethodGet, req.Method)
	}
	if !strings.EqualFold(expectedPath, req.URL.Path) {
		t.Errorf("request path mismatch expected: %s got: %s", expectedPath, req.URL.Path)
	}
}

func TestAccountsClient_DeleteOrder(t *testing.T) {
	mockResponse := http.Response{
		StatusCode: http.StatusAccepted,
		Body:       io.NopCloser(strings.NewReader("")),
	}

	mockClient := utils.MockHttpClient{}
	mockClient.SetResponse(&mockResponse)
	mockAccessToken := authapi.AccessToken{
		ExpiresIn:    3200,
		TokenType:    "random",
		Scope:        "random",
		RefreshToken: uuid.NewV7().String(),
		AccessToken:  uuid.NewV7().String(),
		IdToken:      uuid.NewV7().String(),
	}

	traderClient := NewTraderClient(mockAccessToken, &mockClient)
	accountHash := uuid.NewV7().String()
	orderId := uuid.NewV7().String()
	err := traderClient.DeleteOrder(accountHash, orderId)
	if err != nil {
		t.Errorf("unexpected error %s", err)
	}

	req := mockClient.ReceivedRequest
	expectedPath := fmt.Sprintf("/trader/v1/accounts/%s/orders/%s", accountHash, orderId)
	if req.Method != http.MethodDelete {
		t.Errorf("http method mismatch expected: %s got: %s", http.MethodGet, req.Method)
	}
	if !strings.EqualFold(expectedPath, req.URL.Path) {
		t.Errorf("request path mismatch expected: %s got: %s", expectedPath, req.URL.Path)
	}
}

func TestAccountsClient_PreviewOrder(t *testing.T) {
	mockOrder := getMockOrder()
	strMockOrder, _ := json.Marshal(mockOrder)
	mockResponse := http.Response{
		StatusCode: http.StatusOK,
		Body:       io.NopCloser(bytes.NewReader(strMockOrder)),
	}

	mockClient := utils.MockHttpClient{}
	mockClient.SetResponse(&mockResponse)
	mockAccessToken := authapi.AccessToken{
		ExpiresIn:    3200,
		TokenType:    "random",
		Scope:        "random",
		RefreshToken: uuid.NewV7().String(),
		AccessToken:  uuid.NewV7().String(),
		IdToken:      uuid.NewV7().String(),
	}

	traderClient := NewTraderClient(mockAccessToken, &mockClient)
	accountHash := uuid.NewV7().String()
	previewOrder := OrderPreview{
		OrderId: rand.Int(),
	}
	order, err := traderClient.PreviewOrder(accountHash, previewOrder)
	if err != nil {
		t.Errorf("unexpected error %s", err)
	}

	req := mockClient.ReceivedRequest
	expectedPath := fmt.Sprintf("/trader/v1/accounts/%s/previewOrder", accountHash)
	if req.Method != http.MethodPost {
		t.Errorf("http method mismatch expected: %s got: %s", http.MethodGet, req.Method)
	}
	if !strings.EqualFold(expectedPath, req.URL.Path) {
		t.Errorf("request path mismatch expected: %s got: %s", expectedPath, req.URL.Path)
	}

	if order == nil {
		t.Error("order cannot be nil")
	}
}
