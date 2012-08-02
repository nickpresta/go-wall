package gowall

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Coderwall API URL
const coderwallApiUrl = "http://coderwall.com/%s.json"

// Badge type representing user badges
type Badge struct {
	Badge       string
	Created     string
	Description string
	Name        string
}

// User type representing the user profile
type User struct {
	Endorsements int
	Location     string
	Name         string
	Team         string
	Username     string
	Accounts     map[string]interface{}
	Badges       []Badge
}

// Performs the low-level HTTP GET request to the API to fetch the user profile data.
func fetch(username string) (body []byte, err error) {
	response, err := http.Get(fmt.Sprintf(coderwallApiUrl, username))
	if err != nil {
		return
	}
	defer response.Body.Close()

	body, err = ioutil.ReadAll(response.Body)

	return
}

// Fetches the user data and parses the JSON into the User type.
func FetchUser(username string) (user User, err error) {
	userResponse, err := fetch(username)
	if err != nil {
		return user, fmt.Errorf("Could not fetch user, %s, from Coderwall: %s", username, err)
	}

	err = json.Unmarshal(userResponse, &user)
	if err != nil {
		return user, fmt.Errorf("Could not parse JSON: %s", err)
	}

	return
}
