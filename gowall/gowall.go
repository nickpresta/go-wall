// A Go implementation of the Coderwall API (http://coderwall.com/api)
package gowall

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Coderwall API URL.
const coderwallApiUrl = "http://coderwall.com/%s.json"

// Badge type representing user badges.
type Badge struct {
	Badge       string
	Created     string
	Description string
	Name        string
}

// User type representing the user profile.
type User struct {
	Endorsements int
	Location     string
	Name         string
	Team         string
	Username     string
	Accounts     map[string]string
	Badges       []Badge
}

// Performs the low-level HTTP GET request to the API to fetch the user profile data.
func fetch(username string) ([]byte, error) {
	response, err := http.Get(fmt.Sprintf(coderwallApiUrl, username))
	if err != nil {
		return []byte{}, err
	}
	defer response.Body.Close()

	if response.StatusCode == http.StatusNotFound {
		return []byte{}, fmt.Errorf("404 -- Could not find user")
	}

	body, err := ioutil.ReadAll(response.Body)

	return body, err
}

// Fetches the user data and parses the JSON into the User type.
func FetchUser(username string) (User, error) {
	userResponse, err := fetch(username)
	if err != nil {
		return User{}, fmt.Errorf("Could not fetch user, %s, from Coderwall: %s", username, err)
	}

	user := User{}
	err = json.Unmarshal(userResponse, &user)
	if err != nil {
		return user, fmt.Errorf("Could not parse JSON: %s", err)
	}

	return user, err
}
