package relay

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

// Client is the Service client which performs API requests.
//
// This client should be passed in the `NewApi` functions whenever an API instance is created.
type Client struct {
	apiURL    string
	requester Requester
}

// NewClient instantiate a new Client object.
//
// Zero or more ClientOption object can be passed as a parameter.
// These options will then be applied to the client.
func NewClient(opts ...ClientOption) (*Client, error) {
	conf := newClientConfiguration()

	// Apply the options to the configuration.
	conf.apply(append(defaultClientOptions(), opts...))

	// Validate the applied options.
	if err := conf.validate(); err != nil {
		return nil, NewClientConfigurationValidationError(err)
	}

	return &Client{
		apiURL:    conf.apiURL,
		requester: conf.requester,
	}, nil
}

// Do performs HTTP request based on the Request object.
func (c *Client) Do(ctx context.Context, req *Request, res any) error {
	return c.do(ctx, req, res)
}

//nolint:cyclop
func (c *Client) do(ctx context.Context, req *Request, res any) (err error) {
	if req == nil {
		return ErrEmptyClientRequest
	}

	// Acquiring target URL.
	reqURL, err := req.getURL(c.apiURL)
	if err != nil {
		return NewClientRequestInvalidTarget(err)
	}

	// Instantiating HTTP request.
	httpRequest, err := http.NewRequest(req.Method, reqURL.String(), req.Body)
	if err != nil {
		return NewClientRequestCreationError(err)
	}

	httpRequest = httpRequest.WithContext(ctx)

	// Performing HTTP request.
	httpResponse, err := c.requester.Do(httpRequest)
	if err != nil {
		return NewClientRequestExecutionError(err)
	}

	// Reading response data.
	data, err := io.ReadAll(httpResponse.Body)
	if err != nil {
		return NewClientResponseReadError(err)
	}

	// Closing response data.
	defer func() {
		closeErr := httpResponse.Body.Close()

		if closeErr != nil {
			err = errors.Join(err, NewClientResponseCloseError(closeErr))
		}
	}()

	// Handling response data errors.
	if httpResponse.StatusCode < 200 || httpResponse.StatusCode >= 300 {
		return NewClientResponseError(httpResponse.Status, data)
	}

	// Deserializing response data.
	if res != nil {
		if err := json.Unmarshal(data, res); err != nil {
			return NewClientResponseDeserializationError(err)
		}
	}

	return nil
}
