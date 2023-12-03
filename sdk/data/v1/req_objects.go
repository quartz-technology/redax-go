package v1

import (
	"math/big"
	"net/url"
	"strconv"

	"github.com/attestantio/go-eth2-client/spec/phase0"
	"github.com/ethereum/go-ethereum/common"
)

type ResultsOrder string

const (
	IncreasingValue ResultsOrder = "value"
	DecreasingValue ResultsOrder = "-value"
)

type GetBidsDeliveredRequest struct {
	queryParams url.Values
}

func NewGetBidsDeliveredRequest() *GetBidsDeliveredRequest {
	return &GetBidsDeliveredRequest{
		queryParams: make(url.Values),
	}
}

func (r *GetBidsDeliveredRequest) WithSlot(slot phase0.Slot) *GetBidsDeliveredRequest {
	r.queryParams.Set("slot", strconv.FormatUint(uint64(slot), 10))

	return r
}

func (r *GetBidsDeliveredRequest) WithCursor(cursor uint) *GetBidsDeliveredRequest {
	r.queryParams.Set("cursor", strconv.FormatUint(uint64(cursor), 10))

	return r
}

func (r *GetBidsDeliveredRequest) WithLimit(limit uint) *GetBidsDeliveredRequest {
	r.queryParams.Set("limit", strconv.FormatUint(uint64(limit), 10))

	return r
}

func (r *GetBidsDeliveredRequest) WithBlockHash(blockHash *common.Hash) *GetBidsDeliveredRequest {
	if blockHash != nil {
		r.queryParams.Set("block_hash", blockHash.String())
	}

	return r
}

func (r *GetBidsDeliveredRequest) WithBlockNumber(blockNumber *big.Int) *GetBidsDeliveredRequest {
	if blockNumber != nil {
		r.queryParams.Set("block_number", blockNumber.String())
	}

	return r
}

func (r *GetBidsDeliveredRequest) WithProposerPublicKey(ppk *phase0.BLSPubKey) *GetBidsDeliveredRequest {
	if ppk != nil {
		r.queryParams.Set("proposer_pubkey", ppk.String())
	}

	return r
}

func (r *GetBidsDeliveredRequest) WithBuilderPublicKey(bpk *phase0.BLSPubKey) *GetBidsDeliveredRequest {
	if bpk != nil {
		r.queryParams.Set("builder_pubkey", bpk.String())
	}

	return r
}

func (r *GetBidsDeliveredRequest) WithOrder(order ResultsOrder) *GetBidsDeliveredRequest {
	r.queryParams.Set("order_by", string(order))

	return r
}

type GetBidsReceivedRequest struct {
	queryParams url.Values
}

func NewGetBidsReceivedRequest() *GetBidsReceivedRequest {
	return &GetBidsReceivedRequest{
		queryParams: make(url.Values),
	}
}

func (r *GetBidsReceivedRequest) WithSlot(slot phase0.Slot) *GetBidsReceivedRequest {
	r.queryParams.Set("slot", strconv.FormatUint(uint64(slot), 10))

	return r
}

func (r *GetBidsReceivedRequest) WithBlockHash(blockHash *common.Hash) *GetBidsReceivedRequest {
	if blockHash != nil {
		r.queryParams.Set("block_hash", blockHash.String())
	}

	return r
}

func (r *GetBidsReceivedRequest) WithBlockNumber(blockNumber *big.Int) *GetBidsReceivedRequest {
	if blockNumber != nil {
		r.queryParams.Set("block_number", blockNumber.String())
	}

	return r
}

func (r *GetBidsReceivedRequest) WithBuilderPublicKey(bpk *phase0.BLSPubKey) *GetBidsReceivedRequest {
	if bpk != nil {
		r.queryParams.Set("builder_pubkey", bpk.String())
	}

	return r
}

func (r *GetBidsReceivedRequest) WithLimit(limit uint) *GetBidsReceivedRequest {
	r.queryParams.Set("limit", strconv.FormatUint(uint64(limit), 10))

	return r
}

func (r *GetBidsReceivedRequest) validate() error {
	if !r.queryParams.Has("slot") && !r.queryParams.Has("block_hash") && !r.queryParams.Has(
		"block_number") && !r.queryParams.Has("builder_pubkey") {
		return ErrMissingMandatoryParam
	}

	if r.queryParams.Has("slot") && r.queryParams.Has("cursor") {
		return ErrConflictingParams
	}

	return nil
}
