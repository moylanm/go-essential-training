// Calling GitHub API
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type User struct {
	Login    string `json:"login"`
	Name     string `json:"name"`
	NumRepos int	`json:"public_repos"`
}

// userInfo return information on github user
func userInfo(login string) (*User, error) {
	url := fmt.Sprintf("https://api.github.com/users/%s", login)
	resp, err := http.Get(url)
	if err != nil {
		return &User{}, err
	}
	defer resp.Body.Close()

	user := User{}

	parser := json.NewDecoder(resp.Body)
	if err = parser.Decode(&user); err != nil {
		return &User{}, err
	}

	return &user, nil
}

func main() {
	user, err := userInfo("tebeka")
	if err != nil {
		log.Fatalf("error: %s", err)
	}

	fmt.Printf("%#v\n", user)
}
