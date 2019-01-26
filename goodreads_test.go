package goodreads

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewClient(t *testing.T) {
	c := NewClient("api-key")
	assert.NotNil(t, c)
	assert.Equal(t, "api-key", c.ApiKey)
	assert.Equal(t, defaultApiRoot, c.ApiRoot)
	assert.Equal(t, http.DefaultClient, c.Http)
}
