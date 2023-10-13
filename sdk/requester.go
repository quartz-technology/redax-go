package sdk

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type Requester interface {
	Send(ctx context.Context, method string, path string, bodyContent []byte) ([]byte, error)
}

type defaultRequester struct {
	client  *http.Client
	baseURL string
}

func newDefaultRequester(baseURL string) (*defaultRequester, error) {
	baseRelayURL, err := url.Parse(baseURL)
	if err != nil {
		return nil, ErrFailedToParseRelayAddress(err)
	}

	return &defaultRequester{
		client:  http.DefaultClient,
		baseURL: fmt.Sprintf("%s://%s%s", baseRelayURL.Scheme, baseRelayURL.Host, APIBasePath),
	}, nil
}

func (r *defaultRequester) Send(ctx context.Context, method string, path string, bodyContent []byte) ([]byte, error) {
	var body io.Reader

	switch method {
	case http.MethodGet:
		body = http.NoBody
	default:
		return nil, NewMethodNotSupportedError(method)
	}

	req, err := http.NewRequestWithContext(ctx, method, r.baseURL+path, body)
	if err != nil {
		return nil, ErrFailedToCreateRequestWithContext(err)
	}

	res, err := r.client.Do(req)
	if err != nil {
		return nil, ErrFailedToPerformRequest(err)
	}

	defer res.Body.Close()

	var buffer bytes.Buffer

	_, err = buffer.ReadFrom(res.Body)
	if err != nil {
		return nil, ErrFailedToReadResponseBody(err)
	}

	data := buffer.Bytes()

	if res.StatusCode < 200 || res.StatusCode >= 300 {
		return nil, NewResponseError(res.Status, data)
	}

	return data, nil
}
