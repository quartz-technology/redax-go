package v1

import "errors"

//nolint:lll
var (
	ErrMissingMandatoryParam = errors.New("params validation: need to query for specific slot or block_hash or block_number or builder_pubkey")
	ErrConflictingParams     = errors.New("params validation: conflicting params, cannot specify both slot and cursor")
)
