package httpmockinterface

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type HTTPClient interface {
	Get(url string) (*http.Response, error)
}

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func GetUser(client HTTPClient, id int) (*User, error) {
	resp, err := client.Get(fmt.Sprintf("http://localhost:8080/users/%d", id))
	if err != nil {
		return nil, fmt.Errorf("error sending request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var user User
	err = json.NewDecoder(resp.Body).Decode(&user)
	if err != nil {
		return nil, fmt.Errorf("error decoding response: %v", err)
	}

	return &user, nil
}