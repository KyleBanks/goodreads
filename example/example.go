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
	fmt.Printf("User [%s] %s:\n", u.ID, u.Name)
	fmt.Printf(" Link: %s\n ImageURL: %s\n LastActive: %s\n", u.Link, u.ImageURL, u.LastActive)

	fmt.Println("")

	reviews, err := c.ReviewList(u.ID, "read", "date_read", "", "d", 1, 5)
	if err != nil {
		panic(err)
	}
	fmt.Println("Reviews:")
	for i, rev := range reviews {
		fmt.Printf(" %d. [%d stars, %s] %s\n", i+1, rev.Rating, rev.ReadAt, rev.Book.Title)
	}

	fmt.Println("")

	a, err := c.AuthorShow("18541")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Author [%s] %s:\n", a.ID, a.Name)
	fmt.Printf(" Link: %s\n ImageURL: %s\n WorksCount: %d\n", a.Link, a.ImageURL, a.WorksCount)
}
