package goodreads

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

const testApiKey = "test-api-key"

func TestNewClient(t *testing.T) {
	c := NewClient("api-key")
	assert.NotNil(t, c)
	assert.Equal(t, "api-key", c.ApiKey)
	assert.Equal(t, DefaultDecoder, c.Decoder)
}

func TestClient_AuthorBooks(t *testing.T) {
	c, done := newTestClient(t, decodeTestCase{
		expectURL: fmt.Sprintf("/author/list/12345?key=%s&page=1", testApiKey),
		response:  `<response><author><id>AuthorID</id><name>AuthorName</name></author></response>`,
	})
	defer done()

	a, err := c.AuthorBooks("12345", 1)
	assert.Nil(t, err)
	assert.Equal(t, Author{
		ID:   "AuthorID",
		Name: "AuthorName",
	}, *a)
}

func TestClient_AuthorShow(t *testing.T) {
	c, done := newTestClient(t, decodeTestCase{
		expectURL: fmt.Sprintf("/author/show/12345?key=%s", testApiKey),
		response:  `<response><author><id>AuthorID</id><name>AuthorName</name></author></response>`,
	})
	defer done()

	a, err := c.AuthorShow("12345")
	assert.Nil(t, err)
	assert.Equal(t, Author{
		ID:   "AuthorID",
		Name: "AuthorName",
	}, *a)
}

func TestClient_ReviewList(t *testing.T) {
	c, done := newTestClient(t, decodeTestCase{
		expectURL: fmt.Sprintf("/review/list/user-id.xml?key=%s&order=d&page=1&per_page=200&search=search&shelf=read&sort=date_read&v=2", testApiKey),
		response: `<response>
			<reviews>
				<review><id>review1</id><rating>1</rating></review>
				<review><id>review2</id><rating>2</rating></review>
				<review><id>review3</id><rating>3</rating></review>
			</reviews>
		</response>`,
	})
	defer done()

	r, err := c.ReviewList("user-id", "read", "date_read", "search", "d", 1, 200)
	assert.Nil(t, err)
	assert.Equal(t, []Review{
		Review{ID: "review1", Rating: 1},
		Review{ID: "review2", Rating: 2},
		Review{ID: "review3", Rating: 3},
	}, r)
}

func TestClient_ShelvesList(t *testing.T) {
	c, done := newTestClient(t, decodeTestCase{
		expectURL: fmt.Sprintf("/shelves/list?key=%s&user_id=user-id", testApiKey),
		response: `<response>
			<shelves>
				<user_shelf><id>shelf1</id><name>Shelf 1</name></user_shelf>
				<user_shelf><id>shelf2</id><name>Shelf 2</name></user_shelf>
				<user_shelf><id>shelf3</id><name>Shelf 3</name></user_shelf>
			</shelves>
		</response>`,
	})
	defer done()

	s, err := c.ShelvesList("user-id")
	assert.Nil(t, err)
	assert.Equal(t, []UserShelf{
		UserShelf{ID: "shelf1", Name: "Shelf 1"},
		UserShelf{ID: "shelf2", Name: "Shelf 2"},
		UserShelf{ID: "shelf3", Name: "Shelf 3"},
	}, s)
}

func TestClient_UserShow(t *testing.T) {
	c, done := newTestClient(t, decodeTestCase{
		expectURL: fmt.Sprintf("/user/show/user-id.xml?key=%s", testApiKey),
		response: `<response>
			<user>
				<id>user-id</id>
				<name>User Name</name>
			</user>
		</response>`,
	})
	defer done()

	u, err := c.UserShow("user-id")
	assert.Nil(t, err)
	assert.Equal(t, User{
		ID:   "user-id",
		Name: "User Name",
	}, *u)
}

type decodeTestCase struct {
	expectURL string
	response  string
}

func newTestClient(t *testing.T, tc decodeTestCase) (*Client, func()) {
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, tc.expectURL, r.URL.String())
		w.Write([]byte(tc.response))
	}))

	return &Client{
		ApiKey: testApiKey,
		Decoder: &HttpDecoder{
			Client:  http.DefaultClient,
			ApiRoot: s.URL,
			Verbose: true,
		},
	}, s.Close
}
