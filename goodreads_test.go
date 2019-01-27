package goodreads

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewClient(t *testing.T) {
	c := NewClient("api-key")
	assert.NotNil(t, c)
	assert.Equal(t, "api-key", c.ApiKey)
	assert.Equal(t, DefaultDecoder, c.Decoder)
}
