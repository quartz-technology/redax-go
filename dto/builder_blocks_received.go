package dto

import (
	"fmt"
	"math/big"
	"strings"

	"github.com/attestantio/go-eth2-client/spec/phase0"
	"github.com/ethereum/go-ethereum/common"
)

type BuilderBlocksReceived struct {
	Slot             phase0.Slot
	BlockHash        *common.Hash
	BlockNumber      *big.Int
	BuilderPublicKey *phase0.BLSPubKey
	Limit            uint
}

func (o *BuilderBlocksReceived) EncodeToQueryParams() string {
	var args []string

	if o.Slot > 0 {
		args = append(args, fmt.Sprintf("slot=%d", o.Slot))
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

	if o.Limit > 0 {
		args = append(args, fmt.Sprintf("limit=%d", o.Limit))
	}

	params := strings.Join(args, "&")
	if len(params) > 0 {
		params = "?" + params
	}

	return params
}
