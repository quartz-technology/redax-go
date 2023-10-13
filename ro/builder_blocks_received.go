package ro

import (
	"math/big"

	v1 "github.com/attestantio/go-builder-client/api/v1"
)

type BuilderBlocksReceivedRO struct {
	v1.BidTrace

	BlockNumber          *big.Int `json:"block_number"`
	NumTx                uint     `json:"num_tx"`
	Timestamp            uint64   `json:"timestamp"`
	TimestampMs          uint64   `json:"timestamp_ms"`
	OptimisticSubmission bool     `json:"optimistic_submission"`
}
