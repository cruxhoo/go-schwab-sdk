package auth_api

import (
	"encoding/base64"
	"errors"
	"fmt"
	"go-schwab-client/utils"
	"log/slog"
	"net/http"
	"net/url"
	"strings"
)

type AccessToken struct {
	ExpiresIn    int    `json:"expires_in"`
	TokenType    string `json:"token_type"`
	Scope        string `json:"scope"`
	RefreshToken string `json:"refresh_token"`
	AccessToken  string `json:"access_token"`
	IdToken      string `json:"id_token"`
}

type GrantType string

const (
	GrantAuthCode     GrantType = "authorization_code"
	GrantRefreshToken GrantType = "refresh_token"
)

type SchwabAuthClient struct {
	apiKey      string
	apiSecret   string
	callBackUrl string
	httpClient  utils.HttpClient
}

func (c *SchwabAuthClient) GenerateSignUrl() string {
	return fmt.Sprintf("https://api.schwabapi.com/v1/oauth/authorize?client_id=%s&redirect_uri=%s", c.apiKey, c.callBackUrl)
}

// GetOrRefreshAccessToken Gets or refresh the auth token for Schwab api
func (c *SchwabAuthClient) GetOrRefreshAccessToken(authCode, redirectUrl, refreshToken string, grantType GrantType) (*AccessToken, error) {
	/*
		Get access token
		curl -X POST https://api.schwabapi.com/v1/oauth/token \
		-H 'Authorization: Basic {BASE64_ENCODED_Client_ID:Client_Secret} \
		-H 'Content-Type: application/x-www-form-urlencoded' \
		-d 'grant_type=authorization_code&code={AUTHORIZATION_CODE_VALUE}&redirect_uri=https://example_url.com/callback_example'

		Refresh acces token
		curl -X POST https://api.schwabapi.com/v1/oauth/token \
		-H 'Authorization: Basic {BASE64_ENCODED_Client_ID:Client_Secret} \
		-H 'Content-Type: application/x-www-form-urlencoded' \
		-d 'grant_type=refresh_token&refresh_token={REFRESH_TOKEN_GENERATED_FROM_PRIOR_STEP}
	*/

	// request form
	reqForm := url.Values{}

	switch grantType {
	case GrantAuthCode:
		reqForm.Add("grant_type", "authorization_code")
		reqForm.Add("code", authCode)
		reqForm.Add("redirect_uri", redirectUrl)
	case GrantRefreshToken:
		reqForm.Add("grant_type", "refresh_token")
		reqForm.Add("refresh_token", refreshToken)
	default:
		return nil, errors.New("invalid grant type")
	}

	req, err := http.NewRequest(http.MethodPost, "https://api.schwabapi.com/v1/oauth/token", strings.NewReader(reqForm.Encode()))
	if err != nil {
		slog.Error("failed to create http req", "error", err)
		return nil, err
	}
	// set the auth header & content type
	clientIdEnc := base64.StdEncoding.EncodeToString([]byte(c.apiKey))
	req.Header.Add("Authorization", fmt.Sprintf("%s:%s", clientIdEnc, c.apiSecret))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	accessToken, err := utils.DoRequest[AccessToken](c.httpClient, req)
	if err != nil {
		slog.Error(" get access token http request failed", "error", err)
		return nil, err
	}

	return accessToken, nil
}

func NewClient(apiKey, apiSecret, callbackUrl string, httpClient utils.HttpClient) *SchwabAuthClient {
	return &SchwabAuthClient{
		apiKey:      apiKey,
		apiSecret:   apiSecret,
		callBackUrl: callbackUrl,
		httpClient:  httpClient,
	}
}
