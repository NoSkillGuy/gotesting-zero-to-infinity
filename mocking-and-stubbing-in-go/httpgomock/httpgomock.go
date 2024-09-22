package httpgomock

import (
	"encoding/json"
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
	resp, err := client.Get("http://localhost:8080/users/1")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var user User
	if err := json.NewDecoder(resp.Body).Decode(&user); err != nil {
		return nil, err
	}

	return &user, nil
}