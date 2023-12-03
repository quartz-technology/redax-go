package relay

import (
	"errors"
	"fmt"
)

func NewClientConfigurationValidationError(err error) error {
	return fmt.Errorf("NewClient: failed to validate new service client configuration: %w", err)
}

// ErrEmptyClientRequest is raised when the Client's request is nil.
var ErrEmptyClientRequest = errors.New("do: client request must be non-nil")

// NewClientRequestInvalidTarget is raised when an error occurred while acquiring the request
// target.
func NewClientRequestInvalidTarget(err error) error {
	return fmt.Errorf("do: failed to acquire request target: %w", err)
}

// NewClientRequestCreationError is raised when the Client fails to instantiate the HTTP request
// object.
func NewClientRequestCreationError(err error) error {
	return fmt.Errorf("do: failed to create client request: %w", err)
}

// NewClientRequestExecutionError is raised when the Client fails to perform the HTTP request.
func NewClientRequestExecutionError(err error) error {
	return fmt.Errorf("do: failed to execute client request: %w", err)
}

// NewClientResponseReadError is raised when the Client fails to read the response data.
func NewClientResponseReadError(err error) error {
	return fmt.Errorf("do: failed to read response: %w", err)
}

// NewClientResponseCloseError is raised when the Client fails to close the response data buffer.
func NewClientResponseCloseError(err error) error {
	return fmt.Errorf("do: failed to close response body reader: %w", err)
}

// ClientResponseError is the error returned when the API sends back a response with a non-ok status
// code, indicating that something went wrong with the initial request.
type ClientResponseError struct {
	// The string-formatted error status code, for example "500 Internal Server Error"
	Status string

	// The optional error response, which might contain message indicating why the error occurred.
	Body []byte
}

// NewClientResponseError creates a ClientResponseError given the string-formatted error status code
// and an optional error message.
func NewClientResponseError(status string, body []byte) ClientResponseError {
	return ClientResponseError{
		Status: status,
		Body:   body,
	}
}

func (err ClientResponseError) Error() string {
	if len(err.Body) == 0 {
		return err.Status
	}

	return fmt.Sprintf("do: %v: %s", err.Status, err.Body)
}

// NewClientResponseDeserializationError is raised when the response data can not be deserialized
// into the placeholder passed by the Client's caller.
func NewClientResponseDeserializationError(err error) error {
	return fmt.Errorf("do: failed to unmarshal response data: %w", err)
}
