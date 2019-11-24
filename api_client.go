package goodreads

import (
	"bytes"
	"fmt"
	"net/http"
	"net/url"
)

// DefaultAPIRoot specifies a root for the client, which we point at goodreads.com.
// It's their API, after all.
const defaultAPIRoot = "https://www.goodreads.com"

// The default client, which we configure to work with the Goodreads public API.
var defaultAPIClient APIClient = &httpClient{
	Client:  http.DefaultClient,
	APIRoot: defaultAPIRoot,
}

// APIClient defines a client that can perform an action
// against a Goodreads API function, with parameters,
// and decode the response to a local struct.
type APIClient interface {
	Get(string, func([]byte, interface{}) error, url.Values, interface{}) error
}

type httpClient struct {
	Client  *http.Client
	APIRoot string
	Verbose bool
}

func (h *httpClient) Get(endpoint string, decoder func([]byte, interface{}) error, q url.Values, v interface{}) error {
	url := fmt.Sprintf("%s/%s?%s", h.APIRoot, endpoint, q.Encode())
	if h.Verbose {
		fmt.Printf("GET %s\n", url)
	}

	res, err := h.Client.Get(url)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	buf := new(bytes.Buffer)
	_, err = buf.ReadFrom(res.Body)
	if err != nil {
		return err
	}

	if res.StatusCode < 200 || res.StatusCode > 299 {
		return fmt.Errorf("unexpected response code: %d", res.StatusCode)
	}

	return decoder(buf.Bytes(), v)
}
