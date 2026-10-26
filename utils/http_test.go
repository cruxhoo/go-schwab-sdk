package utils

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
)

type MockStruct struct {
	Ben bool
}

func TestDoRequest_HandleReqError(t *testing.T) {
	expectedError := errors.New("something happened")
	mockHttpClient := MockHttpClient{}
	mockHttpClient.SetErrorResponse(expectedError)

	mockReq, _ := http.NewRequest(http.MethodGet, "http://localhost:8080", nil)

	shouldBeNil, err := DoRequest[MockStruct](&mockHttpClient, mockReq)

	if err == nil {
		t.Error("method call should had return an error")
	}

	if shouldBeNil != nil {
		t.Error("method call should had return error and nil value")
	}
}

func TestDoRequest_HandleHttpBadStatusCode(t *testing.T) {
	mockResponse := http.Response{
		StatusCode: http.StatusBadRequest,
		Body:       io.NopCloser(strings.NewReader("")),
	}
	mockHttpClient := MockHttpClient{}
	mockHttpClient.SetResponse(&mockResponse)

	mockReq, _ := http.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	shouldBeNil, err := DoRequest[MockStruct](&mockHttpClient, mockReq)

	if err == nil {
		t.Error("method call should had return an error")
	}

	if shouldBeNil != nil {
		t.Error("method call should had return error and nil value")
	}

	if !strings.EqualFold(err.Error(), "received not 2xx status code for request, status: 400") {
		t.Error("received error message doesn't match expected")
	}
}

func TestDoRequest_HandleBadJson(t *testing.T) {
	mockResponse := http.Response{
		StatusCode: http.StatusOK,
		Body:       io.NopCloser(strings.NewReader("hay sol bueno y mar de espuma")),
	}
	mockHttpClient := MockHttpClient{}
	mockHttpClient.SetResponse(&mockResponse)

	mockReq, _ := http.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	shouldBeNil, err := DoRequest[MockStruct](&mockHttpClient, mockReq)

	if err == nil {
		t.Error("method call should had return an error")
	}

	if shouldBeNil != nil {
		t.Error("method call should had return error and nil value")
	}

	if _, ok := errors.AsType[*json.SyntaxError](err); !ok {
		t.Fatalf("expected *json.SyntaxError, got %T (%v)", err, err)
	}
}

func TestDoRequest(t *testing.T) {
	mockStruct := MockStruct{
		Ben: false,
	}
	byteStruct, _ := json.Marshal(mockStruct)
	mockResponse := http.Response{
		StatusCode: http.StatusOK,
		Body:       io.NopCloser(bytes.NewReader(byteStruct)),
	}
	mockHttpClient := MockHttpClient{}
	mockHttpClient.SetResponse(&mockResponse)

	mockReq, _ := http.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	respStruct, err := DoRequest[MockStruct](&mockHttpClient, mockReq)
	if err != nil {
		t.Errorf("method call received unspected error %s", err)
	}

	if diff := cmp.Diff(*respStruct, mockStruct); diff != "" {
		t.Errorf("received object doesn't match %s", diff)
	}
}
