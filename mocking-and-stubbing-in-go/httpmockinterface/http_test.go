package httpmockinterface

import (
	"bytes"
	"io"
	"testing"

	"net/http"
)

type MockHTTPClient struct {}

func (m *MockHTTPClient) Get(url string) (*http.Response, error) {
	mockResponse := &http.Response{
		StatusCode: http.StatusOK,
		Body:       io.NopCloser(bytes.NewBufferString(`{"id": 1, "name": "Alice", "age": 30}`)),
	}
	return mockResponse, nil
}

func TestGetUser(t *testing.T) {
	client := &MockHTTPClient{}
	user, err := GetUser(client, 1)
	if err != nil {
		t.Fatalf("Failed to get user: %v", err)
	}

	if user.Name != "Alice" || user.Age != 30 {
		t.Errorf("Expected Alice, 30; got %s, %d", user.Name, user.Age)
	}
}