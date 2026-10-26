package auth_api

import (
	"encoding/json"
	"errors"
	"go-schwab-client/utils"
	"io"
	"net/http"
	"strings"
	"testing"
	"uuid"

	"github.com/google/go-cmp/cmp"
)

func Test_GetOrRefreshAccessToken_HandleBadGrantParameter(t *testing.T) {
	mockClient := utils.MockHttpClient{}
	const badGrantParam GrantType = "badStuff"
	client := NewClient("apikey", "apiSecret", "https://example.com", &mockClient)
	_, err := client.GetOrRefreshAccessToken("some", "stuff", "", badGrantParam)
	if err == nil {
		t.Fatalf("the call must return an error")
	}
	if !strings.EqualFold(err.Error(), "invalid grant type") {
		t.Fatal("error message doesn't match expected error")
	}
}

func Test_GetOrRefreshAccessToken_HandleHttpError(t *testing.T) {
	mockClient := utils.MockHttpClient{}
	mockClient.SetErrorResponse(errors.New("something wrong with call"))
	client := NewClient("apikey", "apiSecret", "https://example.com", &mockClient)
	_, err := client.GetOrRefreshAccessToken("some", "stuff", "", GrantAuthCode)
	if err == nil {
		t.Fatalf("the call must return an error")
	}
	if !strings.EqualFold(err.Error(), "something wrong with call") {
		t.Fatal("error message doesn't match expected error")
	}
}

func Test_GetOrRefreshAccessToken_HandlesBadHttpStatus(t *testing.T) {
	response := http.Response{
		StatusCode: http.StatusBadRequest,
		Body:       io.NopCloser(strings.NewReader("")),
	}
	mockClient := utils.MockHttpClient{}
	mockClient.SetResponse(&response)
	client := NewClient("apikey", "apiSecret", "https://example.com", &mockClient)
	_, err := client.GetOrRefreshAccessToken("some", "stuff", "", GrantAuthCode)
	if err == nil {
		t.Errorf("the call must return an error")
	}
	if !strings.EqualFold(err.Error(), "received not 2xx status code for request, status: 400") {
		t.Errorf("error message doesn't match expected error")
	}
}

func Test_GetOrRefreshAccessToken_HandlesBadJsonData(t *testing.T) {
	response := http.Response{
		StatusCode: http.StatusOK,
		Body:       io.NopCloser(strings.NewReader("broken stuff")),
	}
	mockClient := utils.MockHttpClient{}
	mockClient.SetResponse(&response)
	client := NewClient("apikey", "apiSecret", "https://example.com", &mockClient)
	_, err := client.GetOrRefreshAccessToken("some", "stuff", "", GrantAuthCode)
	if err == nil {
		t.Errorf("the call must return an error")
	}
	if _, ok := errors.AsType[*json.SyntaxError](err); !ok {
		t.Fatalf("expected *json.SyntaxError, got %T (%v)", err, err)
	}
}

func Test_GetOrRefreshAccessToken(t *testing.T) {
	token := AccessToken{
		ExpiresIn:    3600,
		TokenType:    "token_type",
		Scope:        "scope",
		RefreshToken: uuid.NewV7().String(),
		AccessToken:  uuid.NewV7().String(),
		IdToken:      uuid.NewV7().String(),
	}
	bytesT, _ := json.Marshal(token)
	response := http.Response{
		StatusCode: http.StatusOK,
		Body:       io.NopCloser(strings.NewReader(string(bytesT))),
	}
	mockClient := utils.MockHttpClient{}
	mockClient.SetResponse(&response)
	client := NewClient("apikey", "apiSecret", "https://example.com", &mockClient)
	resp, err := client.GetOrRefreshAccessToken("some", "stuff", "", GrantAuthCode)
	if err != nil {
		t.Errorf("the call returned an error")
	}
	if resp == nil {
		t.Errorf("received nil response")
	}

	if diff := cmp.Diff(resp, &token); diff != "" {
		t.Errorf("AccessToken mismatch (-want +got):\n%s", diff)
	}
}
