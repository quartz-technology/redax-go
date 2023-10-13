package sdk

import "fmt"

// ErrFailedToParseRelayAddress is raised when the relay address provided to create a requester
// fails to be parsed in a proper URL.
func ErrFailedToParseRelayAddress(err error) error {
	return fmt.Errorf("failed to parse relay address: %w", err)
}

// MethodNotSupportedError is the error returned by the SDK when the requester tries to
// send a request using a non-supported HTTP method unsupported by the API.
type MethodNotSupportedError struct {
	Method string
}

// NewMethodNotSupportedError creates a MethodNotSupportedError given
// the unsupported method and the target API version.
func NewMethodNotSupportedError(method string) MethodNotSupportedError {
	return MethodNotSupportedError{
		Method: method,
	}
}

func (err MethodNotSupportedError) Error() string {
	return fmt.Sprintf("HTTP method [%s] is not supported by the API", err.Method)
}

// ErrFailedToCreateRequestWithContext wraps the error raised when a request object can not be
// created. This can happen for multiple reasons: when parsing the URL,
// when serializing the request body, ...
func ErrFailedToCreateRequestWithContext(err error) error {
	return fmt.Errorf("failed to create new HTTP request with given context: %w", err)
}

// ErrFailedToPerformRequest wraps the error raised when a request can not be performed correctly.
// This can happen for multiple reasons: when reading a closed request body,
// when parsing the request headers, ...
func ErrFailedToPerformRequest(err error) error {
	return fmt.Errorf("failed to perform HTTP request: %w", err)
}

// ErrFailedToReadResponseBody wraps the error raised when the response body can not be read.
func ErrFailedToReadResponseBody(err error) error {
	return fmt.Errorf("failed to read HTTP response body: %w", err)
}

// ResponseError is the error returned when the API sends back a response with a non-ok status
// code, indicating that something went wrong with the initial request.
type ResponseError struct {
	// The string-formatted error status code, for example "500 Internal Server Error"
	Status string

	// The optional error response, which might contain message indicating why the error occurred.
	Body []byte
}

// NewResponseError creates a ResponseError given the string-formatted error status code and an
// optional error message.
func NewResponseError(status string, body []byte) ResponseError {
	return ResponseError{
		Status: status,
		Body:   body,
	}
}

func (err ResponseError) Error() string {
	if len(err.Body) == 0 {
		return err.Status
	}

	return fmt.Sprintf("%v: %s", err.Status, err.Body)
}

// ErrFailedToSendRequest is raised by the SDK Service when an error has occurred while the client
// was processing a request.
func ErrFailedToSendRequest(err error) error {
	return fmt.Errorf("failed to send request: %w", err)
}

// ErrFailedToParseResponseData is raised by the SDK Service when the response object can not be
// parsed.
func ErrFailedToParseResponseData(err error) error {
	return fmt.Errorf("failed to parse response data into RO: %w", err)
}
