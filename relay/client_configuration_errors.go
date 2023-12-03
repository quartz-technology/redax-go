package relay

import (
	"errors"
	"fmt"
)

func NewClientConfigurationAPIUrlParseError(err error) error {
	return fmt.Errorf("validate: failed to parse API URL: %w", err)
}

var ErrAPIURLHasTrailingSlash = errors.New("validate: API URL must not have a trailing slash")
