package gowall

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"encoding/json"
)

const coderwallApiUrl = "http://coderwall.com/%s.json"

type Badge struct {
	badge       string
	created     string
	description string
	name        string
}

type User struct {
	endorsements string
	location     string
	name         string
	team         string
	username     string
	accounts     []string
	badges       []Badge
}

func fetchUser(username string) (body []byte, err error) {
	response, err := http.Get(fmt.Sprintf(coderwallApiUrl, username))
	defer response.Body.Close()

	if err != nil {
		fmt.Errorf("Could not fetch user, %s, from Coderwall.", username)
		return
	}

	body, err = ioutil.ReadAll(response.Body)

	return
}

func GetUser(username string) (user User, err error) {
	userResponse, err := fetchUser(username)
	if err != nil {
		fmt.Errorf("Could not fetch user, %s, from Coderwall.", username)
		return
	}

	err = json.Unmarshal(userResponse, &user)
	if err != nil {
		fmt.Errorf("Could not parse JSON response.")
		return
	}

	return
}
