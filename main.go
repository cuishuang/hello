package main

import (
	"fmt"
	"os"
	"os/user"
)

func main() {
	currentUser, err := user.Current()
	if err != nil {
		fmt.Printf("Failed to get current user: %s\n", err)
		os.Exit(1)
	}

	fmt.Printf("Hello %s!\n", currentUser.Username)
}
