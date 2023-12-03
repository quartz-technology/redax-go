package relay

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
)

// Request contains all the contents related to performing a request on the Service API.
type Request struct {
	Method  string
	Path    string
	Headers http.Header
	Query   url.Values
	Body    io.Reader
}

// getURL constructs a URL based on the base API URL and the request's path.
//
// Example: "library.com" + "/books".
func (req *Request) getURL(baseAPIURL string) (*url.URL, error) {
	parsedURL, err := url.Parse(baseAPIURL + req.Path)
	if err != nil {
		return nil, NewRequestURLParseError(err)
	}

	parsedURL.RawQuery = req.Query.Encode()

	return parsedURL, nil
}

// SetBody serializes the given body and write the json content type
// to the request.
func (req *Request) SetBody(body any) error {
	var (
		contentType string
		content     io.Reader
	)

	buf, err := json.Marshal(body)
	if err != nil {
		return NewRequestBodySerializationError(err)
	}

	contentType = "application/json"
	content = bytes.NewReader(buf)

	if req.Headers == nil {
		req.Headers = http.Header{}
	}

	req.Headers.Set("Content-Type", contentType)
	req.Body = content

	return nil
}
