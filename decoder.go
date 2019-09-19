package goodreads

import (
	"encoding/xml"
	"fmt"
	"net/http"
	"net/url"
)

const DefaultApiRoot = "https://www.goodreads.com"

var DefaultDecoder Decoder = &HttpDecoder{
	Client:  http.DefaultClient,
	ApiRoot: DefaultApiRoot,
}

// Decode defines a type that can decode a Goodreads API function,
// with parameters, to a local struct.
type Decoder interface {
	Decode(string, url.Values, interface{}) error
}

type HttpDecoder struct {
	Client  *http.Client
	ApiRoot string
	Verbose bool
}

func (h *HttpDecoder) Decode(endpoint string, q url.Values, v interface{}) error {
	url := fmt.Sprintf("%s/%s?%s", h.ApiRoot, endpoint, q.Encode())
	if h.Verbose {
		fmt.Printf("GET %s\n", url)
	}

	res, err := h.Client.Get(url)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode < 200 || res.StatusCode > 299 {
		return fmt.Errorf("Unexpected response code: %d", res.StatusCode)
	}
	return xml.NewDecoder(res.Body).Decode(v)
}
