package main

import (
	"fmt"
	"github.com/NickPresta/go-wall/gowall"
)

func main() {
	user, err := gowall.FetchUser("NickPresta")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Name: %s (%s)\n", user.Name, user.Username)
	fmt.Println("Location:", user.Location)
	fmt.Println("Endorsements:", user.Endorsements)
	fmt.Println("Team:", user.Team)

	fmt.Println("\nAccounts:")
	for key, val := range user.Accounts {
		fmt.Printf("\t%s: %s\n", key, val)
	}

	fmt.Printf("\nBadges (%d):\n", len(user.Badges))
	for _, badge := range user.Badges {
		fmt.Printf("\t%s on %s\n", badge.Name, badge.Created)
	}
}
