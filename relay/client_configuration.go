package relay

import "net/url"

// clientConfiguration holds all the options used by the Client.
type clientConfiguration struct {
	apiURL    string
	requester Requester
}

// newClientConfiguration creates an empty clientConfiguration.
func newClientConfiguration() *clientConfiguration {
	return &clientConfiguration{
		apiURL:    "",
		requester: nil,
	}
}

// apply is used during the Client creation and stores the client options in the
// clientConfiguration.
func (cs *clientConfiguration) apply(opts []ClientOption) {
	for _, opt := range opts {
		opt(cs)
	}
}

// validate is used at the Client creation and ensures that the values used in the
// clientConfiguration are valid.
func (cs *clientConfiguration) validate() error {
	if _, err := url.Parse(cs.apiURL); err != nil {
		return NewClientConfigurationAPIUrlParseError(err)
	}

	if cs.apiURL[len(cs.apiURL)-1:] == "/" {
		return ErrAPIURLHasTrailingSlash
	}

	return nil
}
