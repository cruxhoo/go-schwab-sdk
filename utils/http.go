package utils

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
)

type HttpClient interface {
	Do(req *http.Request) (*http.Response, error)
}

func DoRequest[T any](client HttpClient, req *http.Request) (*T, error) {
	resp, err := client.Do(req)
	if err != nil {
		slog.Error(" get access token http request failed", "error", err)
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode < http.StatusOK || resp.StatusCode > http.StatusMultipleChoices {
		slog.Error("received not 2xx status code for request", "responseStatus", resp.StatusCode)
		return nil, fmt.Errorf("received not 2xx status code for request, status: %d", resp.StatusCode)
	}
	var payload T
	if err := json.NewDecoder(resp.Body).Decode(&payload); err != nil {
		slog.Error("failed to decode response body", "err", err)
		return nil, err
	}
	return &payload, nil
}
