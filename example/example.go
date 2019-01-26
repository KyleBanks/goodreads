package main

import (
	"fmt"
	"os"

	"github.com/KyleBanks/goodreads"
)

func main() {
	key := os.Getenv("API_KEY")
	if key == "" {
		fmt.Println("Missing required env var: API_KEY")
		os.Exit(1)
	}

	c := goodreads.NewClient(key)

	u, err := c.UserShow("38763538")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Loaded user details of %s:\n", u.Name)
	fmt.Println(u)
}
