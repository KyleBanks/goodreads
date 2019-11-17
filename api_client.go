package goodreads

import (
	"encoding/xml"
	"fmt"
	"net/http"
	"net/url"
)

const DefaultAPIRoot = "https://www.goodreads.com"

var DefaultAPIClient APIClient = &HTTPClient{
	Client:  http.DefaultClient,
	ApiRoot: DefaultAPIRoot,
}

// APIClient defines a client that can perform an action
// against a Goodreads API function, with parameters,
// and decode the response to a local struct.
type APIClient interface {
	Get(string, url.Values, interface{}) error
}

type HTTPClient struct {
	Client  *http.Client
	ApiRoot string
	Verbose bool
}

func (h *HTTPClient) Get(endpoint string, q url.Values, v interface{}) error {
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
		return fmt.Errorf("unexpected response code: %d", res.StatusCode)
	}
	return xml.NewDecoder(res.Body).Decode(v)
}
