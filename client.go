package api

import (
	"context"
	"fmt"
	"io"
	_ "log"
	"net/http"
	"net/url"
	"path/filepath"

	"github.com/whosonfirst/go-ioutil"
)

// V1 is the root URL endpoint for version one of the FlySFO API.
const V1 string = "https://api.flysfo.com/sfo/v1.0"

// type Client is a struct for performing requests to the FlySFO API.
type Client struct {
	// http_client is the underlying `net/http.Client` instance to perform HTTP requests.
	http_client *http.Client
	// apikey is the FlySFO API key used to execute API requests.
	apikey string
	// endpoint is the root URL endpoint for API requests.
	endpoint string
}

func NewClient(ctx context.Context, apikey string) (*Client, error) {
	return NewClientWithEndpoint(ctx, V1, apikey)
}

func NewClientWithEndpoint(ctx context.Context, endpoint string, apikey string) (*Client, error) {

	_, err := url.Parse(endpoint)

	if err != nil {
		return nil, fmt.Errorf("Failed to parse endpoint, %w", err)
	}

	http_client := &http.Client{}

	cl := &Client{
		http_client: http_client,
		endpoint:    endpoint,
		apikey:      apikey,
	}

	return cl, nil
}

func (cl *Client) ExecuteMethod(ctx context.Context, uri string, params *url.Values) (io.ReadSeekCloser, error) {

	u, _ := url.Parse(cl.endpoint)

	u.Path = filepath.Join(u.Path, uri)

	if params != nil {
		u.RawQuery = params.Encode()
	}

	req, err := http.NewRequestWithContext(ctx, "GET", u.String(), nil)

	if err != nil {
		return nil, fmt.Errorf("Failed to create new request, %w", err)
	}

	req.Header.Set("apikey", cl.apikey)
	req.Header.Set("accept", "application/json")

	rsp, err := cl.http_client.Do(req)

	if err != nil {
		return nil, fmt.Errorf("Failed to execute request, %w", err)
	}

	if rsp.StatusCode != 200 {
		return nil, fmt.Errorf("Request failed with response %s", rsp.Status)
	}

	rcs, err := ioutil.NewReadSeekCloser(rsp.Body)

	if err != nil {
		return nil, fmt.Errorf("Failed to create ReadSeekCloser for response, %w", err)
	}

	return rcs, nil
}
