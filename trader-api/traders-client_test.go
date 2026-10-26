package trader_api

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	authapi "go-schwab-client/auth-api"
	"go-schwab-client/utils"
	"io"
	"math/rand"
	"net/http"
	"regexp"
	"strings"
	"testing"
	"uuid"

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

func TestTraderClient_GetAccountNumbers(t *testing.T) {
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

func TestTraderClient_GetAccountNumbers_HandleErrors(t *testing.T) {
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

func TestTraderClient_GetAccounts_All(t *testing.T) {
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

func TestTraderClient_GetAccounts_One(t *testing.T) {
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
