package goodreads

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"net/url"
	"reflect"
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
	c := newTestClient(mockDecoder{
		expectFn: "author/list/12345",
		expectQuery: url.Values(map[string][]string{
			"key":  []string{testApiKey},
			"page": []string{"1"},
		}),
		response: `<response><author><id>AuthorID</id><name>AuthorName</name></author></response>`,
	})

	a, err := c.AuthorBooks("12345", 1)
	assert.Nil(t, err)
	assert.Equal(t, Author{
		ID:   "AuthorID",
		Name: "AuthorName",
	}, *a)
}

func TestClient_AuthorShow(t *testing.T) {
	c := newTestClient(mockDecoder{
		expectFn: "author/show/12345",
		expectQuery: url.Values(map[string][]string{
			"key": []string{testApiKey},
		}),
		response: `<response><author><id>AuthorID</id><name>AuthorName</name></author></response>`,
	})

	a, err := c.AuthorShow("12345")
	assert.Nil(t, err)
	assert.Equal(t, Author{
		ID:   "AuthorID",
		Name: "AuthorName",
	}, *a)
}

func TestClient_ReviewList(t *testing.T) {
	c := newTestClient(mockDecoder{
		expectFn: "review/list",
		expectQuery: url.Values(map[string][]string{
			"key":      []string{testApiKey},
			"id":       []string{"user-id"},
			"v":        []string{"2"},
			"shelf":    []string{"read"},
			"sort":     []string{"date_read"},
			"search":   []string{"search"},
			"order":    []string{"d"},
			"page":     []string{"1"},
			"per_page": []string{"200"},
		}),
		response: `<response>
			<reviews>
				<review><id>Review1</id><rating>1</rating></review>
				<review><id>Review2</id><rating>2</rating></review>
				<review><id>Review3</id><rating>3</rating></review>
			</reviews>
		</response>`,
	})

	r, err := c.ReviewList("user-id", "read", "date_read", "search", "d", 1, 200)
	assert.Nil(t, err)
	assert.Equal(t, 3, len(r))
	assert.Equal(t, []Review{
		Review{ID: "Review1", Rating: 1},
		Review{ID: "Review2", Rating: 2},
		Review{ID: "Review3", Rating: 3},
	}, r)
}

type mockDecoder struct {
	expectFn    string
	expectQuery url.Values
	response    string
}

func (m mockDecoder) Decode(fn string, q url.Values, v interface{}) error {
	if fn != m.expectFn {
		return fmt.Errorf("Unexpected function; expected=%s, got=%s", m.expectFn, fn)
	}
	if !reflect.DeepEqual(q, m.expectQuery) {
		return fmt.Errorf("Unexpected query; expected=%s, got=%s", m.expectQuery, q)
	}
	return xml.NewDecoder(bytes.NewBufferString(m.response)).Decode(v)
}

func newTestClient(m mockDecoder) *Client {
	return &Client{
		ApiKey:  testApiKey,
		Decoder: m,
	}
}
