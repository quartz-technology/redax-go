package v1

import "fmt"

func NewDataSDKError(method string, err error) error {
	return fmt.Errorf("data sdk: error occurred in method %s, reason: %w", method, err)
}
