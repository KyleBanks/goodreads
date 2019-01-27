// Package goodreads provides a REST client for the public goodreads.com API.
//
// https://www.goodreads.com/api
package goodreads

import (
	"encoding/xml"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

const (
	defaultApiRoot = "https://goodreads.com"
)

type Client struct {
	ApiKey string

	ApiRoot string
	Http    *http.Client
	Verbose bool
}

// NewClient initializes a Client with default parameters.
func NewClient(key string) *Client {
	return &Client{
		ApiKey: key,

		ApiRoot: defaultApiRoot,
		Http:    http.DefaultClient,
	}
}

// AuthorBooks returns a list of books by a particular author.
// https://www.goodreads.com/api/index#author.books
func (c *Client) AuthorBooks(authorID string, page int) (*Author, error) {
	q := url.Values{}
	if page > 0 {
		q.Set("page", strconv.Itoa(page))
	}

	var r struct {
		Author Author `xml:"author"`
	}
	if err := c.get(fmt.Sprintf("author/list/%s", authorID), &r, q); err != nil {
		return nil, err
	}
	return &r.Author, nil
}

// AuthorShow returns the full details of an author.
// https://www.goodreads.com/api/index#author.show
func (c *Client) AuthorShow(authorID string) (*Author, error) {
	var r struct {
		Author Author `xml:"author"`
	}
	if err := c.get(fmt.Sprintf("author/show/%s", authorID), &r, nil); err != nil {
		return nil, err
	}
	return &r.Author, nil
}

// ReviewList returns the books on a members shelf.
// https://www.goodreads.com/api/index#reviews.list
func (c *Client) ReviewList(userID, shelf, sort, search, order string, page, perPage int) ([]Review, error) {
	q := url.Values{}
	q.Set("id", userID)
	q.Set("v", "2")
	if shelf != "" {
		q.Set("shelf", shelf)
	}
	if sort != "" {
		q.Set("sort", sort)
	}
	if search != "" {
		q.Set("search", search)
	}
	if order != "" {
		q.Set("order", order)
	}
	if page > 0 {
		q.Set("page", strconv.Itoa(page))
	}
	if perPage > 0 {
		q.Set("per_page", strconv.Itoa(perPage))
	}

	var r struct {
		Reviews []Review `xml:"reviews>review"`
	}
	if err := c.get("review/list", &r, q); err != nil {
		return nil, err
	}
	return r.Reviews, nil
}

// ShelvesList returns the list of shelves belonging to a user.
// https://www.goodreads.com/api/index#shelves.list
func (c *Client) ShelvesList(userID string) ([]UserShelf, error) {
	q := url.Values{}
	q.Set("user_id", userID)
	var r struct {
		Shelves []UserShelf `xml:"shelves>user_shelf"`
	}
	if err := c.get("shelves/list", &r, q); err != nil {
		return nil, err
	}
	return r.Shelves, nil
}

// UserShow returns the public information about a given Goodreads user.
// https://www.goodreads.com/api/index#user.show
func (c *Client) UserShow(id string) (*User, error) {
	var r struct {
		User User `xml:"user"`
	}
	if err := c.get(fmt.Sprintf("user/show/%s.xml", id), &r, nil); err != nil {
		return nil, err
	}
	return &r.User, nil
}

func (c *Client) get(endpoint string, v interface{}, q url.Values) error {
	if q == nil {
		q = url.Values{}
	}
	q.Set("key", c.ApiKey)

	url := fmt.Sprintf("%s/%s?%s", c.ApiRoot, endpoint, q.Encode())
	if c.Verbose {
		fmt.Printf("GET %s\n", url)
	}
	res, err := c.Http.Get(url)
	if err != nil {
		return err
	} else if res.StatusCode < 200 || res.StatusCode > 299 {
		return fmt.Errorf("Unexpected response code: %d", res.StatusCode)
	}

	if err := xml.NewDecoder(res.Body).Decode(v); err != nil {
		return err
	}

	return nil
}
