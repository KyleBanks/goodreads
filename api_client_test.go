package goodreads

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHttpClient_Get(t *testing.T) {
	testCases := []struct {
		DecoderType string
		Decoder     func([]byte, interface{}) error
		Data        []byte
	}{
		{"xml", xml.Unmarshal, []byte(`<?xml version="1.0" encoding="UTF-8"?><response><id>SampleID</id></response>`)},
		{"json", json.Unmarshal, []byte(`{ "id": "SampleID" }`)},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("with %s decoder", tc.DecoderType), func(t *testing.T) {
			s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, "/foo/bar?p1=v1&p2=v2", r.URL.String())
				_, _ = w.Write(tc.Data)
			}))
			defer s.Close()

			v := url.Values{}
			v.Set("p1", "v1")
			v.Set("p2", "v2")
			var res struct {
				ID string `xml:"id" json:"id"`
			}
			h := httpClient{Client: http.DefaultClient, APIRoot: s.URL, Verbose: true}
			err := h.Get("foo/bar", tc.Decoder, v, &res)
			assert.Nil(t, err)
			assert.Equal(t, "SampleID", res.ID)
		})
	}
}
