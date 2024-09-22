package httpmockmonkeypatching

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func GetUser(userID int) (*User, error) {
	response, err := http.Get(fmt.Sprintf("https://api.example.com/users/%d", userID))
	if err != nil {
			return nil, err
	}

	defer response.Body.Close()

	var user User
	if err := json.NewDecoder(response.Body).Decode(&user); err != nil {
			return nil, err
	}

	return &user, nil
}