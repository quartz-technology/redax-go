package relay

import (
	"net"
	"net/http"
	"time"
)

// Requester wraps the standard HTTP method used to perform a request.
type Requester interface {
	Do(*http.Request) (*http.Response, error)
}

// newDefaultRequester creates a Requester using the standard http.Client.
func newDefaultRequester() Requester {
	//nolint:gomnd,exhaustruct
	return &http.Client{
		Timeout: 30 * time.Second,
		Transport: &http.Transport{
			DialContext:           (&net.Dialer{Timeout: 5 * time.Second}).DialContext,
			TLSHandshakeTimeout:   5 * time.Second,
			ResponseHeaderTimeout: 30 * time.Second,
			MaxIdleConnsPerHost:   20,
		},
	}
}
