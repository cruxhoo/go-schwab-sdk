package utils

import "net/http"

type MockHttpClient struct {
	errorResponse    error
	expectedResponse *http.Response
	ReceivedRequest  *http.Request
}

func (m *MockHttpClient) Do(req *http.Request) (*http.Response, error) {
	m.ReceivedRequest = req
	if m.errorResponse != nil {
		return nil, m.errorResponse
	}
	return m.expectedResponse, nil
}

func (m *MockHttpClient) SetErrorResponse(expectedError error) {
	m.errorResponse = expectedError
}

func (m *MockHttpClient) SetResponse(expectedResponse *http.Response) {
	m.expectedResponse = expectedResponse
}

func (m *MockHttpClient) ClearMocks() {
	m.expectedResponse = nil
	m.errorResponse = nil
}
