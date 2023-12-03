package relay

// ClientOption is a function used to store the value of one option in the clientConfiguration.
type ClientOption func(*clientConfiguration)

// defaultClientOptions creates a list of ClientOption with valid default values.
func defaultClientOptions() []ClientOption {
	return []ClientOption{
		WithAPIURL("http://localhost:18550"),
		WithRequester(newDefaultRequester()),
	}
}

// WithAPIURL client option overrides the API URL of the Service API to the given URL.
func WithAPIURL(apiURL string) ClientOption {
	return func(s *clientConfiguration) {
		s.apiURL = apiURL
	}
}

// WithRequester client option overrides the Requester used to perform HTTP requests.
func WithRequester(requester Requester) ClientOption {
	return func(s *clientConfiguration) {
		s.requester = requester
	}
}
