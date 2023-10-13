package dto

import (
	"fmt"
	"math/big"
	"strings"

	"github.com/attestantio/go-eth2-client/spec/phase0"
	"github.com/ethereum/go-ethereum/common"
)

type ResultsOrder string

const (
	IncreasingValue ResultsOrder = "value"
	DecreasingValue ResultsOrder = "-value"
)

type ProposerPayloadDelivered struct {
	Slot              phase0.Slot
	Cursor            uint
	BlockHash         *common.Hash
	BlockNumber       *big.Int
	BuilderPublicKey  *phase0.BLSPubKey
	ProposerPublicKey *phase0.BLSPubKey
	Limit             uint
	OrderBy           ResultsOrder
}

func (o *ProposerPayloadDelivered) EncodeToQueryParams() string {
	var args []string

	if o.Slot > 0 {
		args = append(args, fmt.Sprintf("slot=%d", o.Slot))
	}

	if o.Cursor > 0 {
		args = append(args, fmt.Sprintf("cursor=%d", o.Cursor))
	}

	if o.BlockHash != nil {
		args = append(args, fmt.Sprintf("block_hash=%s", o.BlockHash))
	}

	if o.BlockNumber != nil {
		args = append(args, fmt.Sprintf("block_number=%d", o.BlockNumber))
	}

	if o.BuilderPublicKey != nil {
		args = append(args, fmt.Sprintf("builder_pubkey=%s", o.BuilderPublicKey))
	}

	if o.ProposerPublicKey != nil {
		args = append(args, fmt.Sprintf("proposer_pubkey=%s", o.ProposerPublicKey))
	}

	if o.Limit > 0 {
		args = append(args, fmt.Sprintf("limit=%d", o.Limit))
	}

	orderBy := o.OrderBy

	if o.OrderBy == "" {
		orderBy = IncreasingValue
	}

	args = append(args, fmt.Sprintf("order_by=%s", orderBy))

	params := strings.Join(args, "&")
	if len(params) > 0 {
		params = "?" + params
	}

	return params
}
