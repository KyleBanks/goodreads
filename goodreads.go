// Package goodreads provides a REST client for the public goodreads.com API.
//
// https://www.goodreads.com/api
package goodreads

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"github.com/KyleBanks/goodreads/responses"
	"github.com/KyleBanks/goodreads/responses/work"
	"net/url"
	"strconv"
	"strings"
)

// Client wraps the public Goodreads API.
type Client struct {
	APIKey     string
	httpClient APIClient
}

// NewClient initializes a Client with default parameters.
func NewClient(key string) *Client {
	return &Client{
		APIKey:     key,
		httpClient: defaultAPIClient,
	}
}

// AuthorBooks returns a list of books by a particular author.
// https://www.goodreads.com/api/index#author.books
func (c *Client) AuthorBooks(authorID string, page int) (*responses.Author, error) {
	v := c.defaultValues()
	if page > 0 {
		v.Set("page", strconv.Itoa(page))
	}

	var r struct {
		Author responses.Author `xml:"author"`
	}
	err := c.httpClient.Get(fmt.Sprintf("author/list/%s", authorID), xml.Unmarshal, v, &r)
	if err != nil {
		return nil, err
	}
	return &r.Author, nil
}

// AuthorShow returns the full details of an author.
// https://www.goodreads.com/api/index#author.show
func (c *Client) AuthorShow(authorID string) (*responses.Author, error) {
	var r struct {
		Author responses.Author `xml:"author"`
	}
	err := c.httpClient.Get(fmt.Sprintf("author/show/%s", authorID), xml.Unmarshal, c.defaultValues(), &r)
	if err != nil {
		return nil, err
	}
	return &r.Author, nil
}

// BookReviewCounts returns the review statistics for a given list of ISBNs.
// https://www.goodreads.com/api/index#book.review_counts
func (c *Client) BookReviewCounts(isbns []string) ([]responses.ReviewCounts, error) {
	v := c.defaultValues()
	v.Set("isbns", strings.Join(isbns, ","))
	var r struct {
		ReviewCounts []responses.ReviewCounts `json:"books"`
	}
	err := c.httpClient.Get("book/review_counts.json", json.Unmarshal, v, &r)
	if err != nil {
		return nil, err
	}
	return r.ReviewCounts, nil
}

// ReviewList returns the books on a members shelf.
// https://www.goodreads.com/api/index#reviews.list
func (c *Client) ReviewList(userID, shelf, sort, search, order string, page, perPage int) ([]responses.Review, error) {
	v := c.defaultValues()
	v.Set("v", "2")
	if shelf != "" {
		v.Set("shelf", shelf)
	}
	if sort != "" {
		v.Set("sort", sort)
	}
	if search != "" {
		v.Set("search", search)
	}
	if order != "" {
		v.Set("order", order)
	}
	if page > 0 {
		v.Set("page", strconv.Itoa(page))
	}
	if perPage > 0 {
		v.Set("per_page", strconv.Itoa(perPage))
	}

	var r struct {
		Reviews []responses.Review `xml:"reviews>review"`
	}
	err := c.httpClient.Get(fmt.Sprintf("review/list/%s.xml", userID), xml.Unmarshal, v, &r)
	if err != nil {
		return nil, err
	}
	return r.Reviews, nil
}

// SearchBooks returns a list of books based on a query string
// by title, author, or ISBN.
// https://www.goodreads.com/api/index#search.books
func (c *Client) SearchBooks(query string, page int, field SearchField) ([]work.Work, error) {
	v := c.defaultValues()
	v.Set("q", query)
	v.Set("search[field]", string(field))
	if page != 0 {
		v.Set("page", strconv.Itoa(page))
	}

	var r struct {
		Works []work.Work `xml:"search>results>work"`
	}

	err := c.httpClient.Get("search/index.xml", xml.Unmarshal, v, &r)
	if err != nil {
		return nil, err
	}

	return r.Works, nil
}

// ShelvesList returns the list of shelves belonging to a user.
// https://www.goodreads.com/api/index#shelves.list
func (c *Client) ShelvesList(userID string) ([]responses.UserShelf, error) {
	v := c.defaultValues()
	v.Set("user_id", userID)
	var r struct {
		Shelves []responses.UserShelf `xml:"shelves>user_shelf"`
	}
	err := c.httpClient.Get("shelf/list.xml", xml.Unmarshal, v, &r)
	if err != nil {
		return nil, err
	}
	return r.Shelves, nil
}

// UserShow returns the public information about a given Goodreads user.
// https://www.goodreads.com/api/index#user.show
func (c *Client) UserShow(id string) (*responses.User, error) {
	var r struct {
		User responses.User `xml:"user"`
	}
	err := c.httpClient.Get(fmt.Sprintf("user/show/%s.xml", id), xml.Unmarshal, c.defaultValues(), &r)
	if err != nil {
		return nil, err
	}
	return &r.User, nil
}

func (c *Client) defaultValues() url.Values {
	v := url.Values{}
	v.Set("key", c.APIKey)
	return v
}
