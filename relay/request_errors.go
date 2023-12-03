package relay

import "fmt"

// NewRequestURLParseError is raised when the getURL method fails to parse the association of the
// base API URL and the request's path.
func NewRequestURLParseError(err error) error {
	return fmt.Errorf("getURL: failed to parse request url: %w", err)
}

// NewRequestBodySerializationError is raised when the SetBody method fails to serialize the body.
func NewRequestBodySerializationError(err error) error {
	return fmt.Errorf("SetBody: failed to marshal body content: %w", err)
}
