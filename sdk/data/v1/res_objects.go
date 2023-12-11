package v1

import (
	"encoding/json"
	"math/big"
	"time"

	v1 "github.com/attestantio/go-builder-client/api/v1"
)

type BidDeliveredMeta struct {
	BlockNumber *big.Int `json:"block_number"`
	NumTx       uint     `json:"num_tx"`
}

type bidDeliveredMetaJSON struct {
	BlockNumber string `json:"block_number"`
	NumTx       uint   `json:"num_tx,string"`
}

type BidDelivered struct {
	v1.BidTrace
	BidDeliveredMeta
}

func (b *BidDelivered) UnmarshalJSON(input []byte) error {
	bt := new(v1.BidTrace)

	if err := json.Unmarshal(input, bt); err != nil {
		return err
	}

	br := new(bidDeliveredMetaJSON)

	if err := json.Unmarshal(input, br); err != nil {
		return err
	}

	b.Slot = bt.Slot
	b.ParentHash = bt.ParentHash
	b.BlockHash = bt.BlockHash
	b.BuilderPubkey = bt.BuilderPubkey
	b.ProposerPubkey = bt.ProposerPubkey
	b.ProposerFeeRecipient = bt.ProposerFeeRecipient
	b.GasLimit = bt.GasLimit
	b.GasUsed = bt.GasUsed
	b.Value = bt.Value

	b.BlockNumber = big.NewInt(0).SetBytes([]byte(br.BlockNumber))
	b.NumTx = br.NumTx

	return nil
}

type BidReceivedMeta struct {
	BlockNumber          *big.Int
	NumTx                uint
	Timestamp            time.Time
	TimestampMs          time.Time
	OptimisticSubmission bool
}

type bidReceivedMetaJSON struct {
	BlockNumber          string `json:"block_number"`
	NumTx                uint   `json:"num_tx,string"`
	Timestamp            int64  `json:"timestamp,string"`
	TimestampMs          int64  `json:"timestamp_ms,string"`
	OptimisticSubmission bool   `json:"optimistic_submission"`
}

type BidReceived struct {
	v1.BidTrace
	BidReceivedMeta
}

func (b *BidReceived) UnmarshalJSON(input []byte) error {
	bt := new(v1.BidTrace)

	if err := json.Unmarshal(input, bt); err != nil {
		return err
	}

	br := new(bidReceivedMetaJSON)

	if err := json.Unmarshal(input, br); err != nil {
		return err
	}

	b.Slot = bt.Slot
	b.ParentHash = bt.ParentHash
	b.BlockHash = bt.BlockHash
	b.BuilderPubkey = bt.BuilderPubkey
	b.ProposerPubkey = bt.ProposerPubkey
	b.ProposerFeeRecipient = bt.ProposerFeeRecipient
	b.GasLimit = bt.GasLimit
	b.GasUsed = bt.GasUsed
	b.Value = bt.Value

	b.BlockNumber = big.NewInt(0).SetBytes([]byte(br.BlockNumber))
	b.NumTx = br.NumTx
	b.Timestamp = time.Unix(br.Timestamp, 0)
	b.TimestampMs = time.Unix(br.Timestamp, (br.Timestamp%1_000)*1_000_000)
	b.OptimisticSubmission = br.OptimisticSubmission

	return nil
}
