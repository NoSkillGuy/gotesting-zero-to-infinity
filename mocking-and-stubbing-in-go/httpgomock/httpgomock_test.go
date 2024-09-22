package httpgomock

import (
	"bytes"
	"io"
	"testing"

	"github.com/NoSkillGuy/gotesting-zero-to-infinity/mocking-and-stubbing-in-go/httpgomock/mocks"
	"github.com/golang/mock/gomock"

	"net/http"
)

func TestGetUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClient := mocks.NewMockHTTPClient(ctrl)

	user := &User{
		ID:   1,
		Name: "Alice",
		Age:  30,
	}

	mockResponse := &http.Response{
		StatusCode: http.StatusOK,
		Body:       io.NopCloser(bytes.NewBufferString(`{"id": 1, "name": "Alice", "age": 30}`)),
	}

	mockClient.EXPECT().Get("http://localhost:8080/users/1").Return(mockResponse, nil)

	u, err := GetUser(mockClient, 1)
	if err != nil {
		t.Fatalf("Failed to get user: %v", err)
	}

	if u.Name != user.Name || u.Age != user.Age {
		t.Errorf("Expected %s, %d; got %s, %d", user.Name, user.Age, u.Name, u.Age)
	}

}
