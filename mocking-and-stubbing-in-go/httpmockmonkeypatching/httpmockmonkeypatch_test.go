package httpmockmonkeypatching

import (
	"testing"

	"github.com/jarcoal/httpmock"
)


func TestGetUserWithHttpMock(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "https://api.example.com/users/1",
			httpmock.NewStringResponder(200, `{"id": 1, "name": "John Doe", "age": 30}`))

	user, err := GetUser(1)
	if err != nil {
			t.Fatalf("Expected no error, got %v", err)
	}

	if user.Name != "John Doe" {
			t.Errorf("Expected user name to be 'John Doe', got '%s'", user.Name)
	}
}