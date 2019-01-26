// Package goodreads provides a REST client for the public goodreads.com API.
//
// https://www.goodreads.com/api
package goodreads

import (
	"encoding/xml"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

const (
	defaultApiRoot = "https://goodreads.com"
	defaultTimeout = time.Second * 5
)

type Client struct {
	ApiKey string

	ApiRoot string
	Http    *http.Client
}

// NewClient initializes a Client with default parameters.
func NewClient(key string) *Client {
	return &Client{
		ApiKey: key,

		ApiRoot: defaultApiRoot,
		Http: &http.Client{
			Timeout: defaultTimeout,
		},
	}
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

	res, err := c.Http.Get(fmt.Sprintf("%s/%s?%s", c.ApiRoot, endpoint, q.Encode()))
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
