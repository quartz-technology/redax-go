package dto

import (
	"math/big"
	"testing"

	"github.com/attestantio/go-eth2-client/spec/phase0"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"
)

func TestProposerPayloadDelivered_EncodeToQueryParams(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		input    *ProposerPayloadDelivered
		expected string
	}{
		{
			name: "must encode valid complete input",
			input: &ProposerPayloadDelivered{
				Slot:              42,
				Cursor:            42,
				BlockHash:         &common.Hash{},
				BlockNumber:       big.NewInt(42),
				BuilderPublicKey:  &phase0.BLSPubKey{},
				ProposerPublicKey: &phase0.BLSPubKey{},
				Limit:             42,
				OrderBy:           DecreasingValue,
			},
			expected: "?slot=42" +
				"&cursor=42" +
				"&block_hash=0x0000000000000000000000000000000000000000000000000000000000000000" +
				"&block_number=42" +
				"&builder_pubkey=0x000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000" +
				"&proposer_pubkey=0x000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000" +
				"&limit=42" +
				"&order_by=-value",
		},
		{
			name: "must encode valid partial input",
			input: &ProposerPayloadDelivered{
				BlockHash:   &common.Hash{},
				BlockNumber: big.NewInt(42),
				Limit:       42,
			},
			expected: "?block_hash=0x0000000000000000000000000000000000000000000000000000000000000000" +
				"&block_number=42" +
				"&limit=42" +
				"&order_by=value",
		},
		{
			name:     "must encode valid empty input",
			input:    &ProposerPayloadDelivered{},
			expected: "?order_by=value",
		},
	}

	for i := range testCases {
		//nolint:varnamelen
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			result := tc.input.EncodeToQueryParams()
			require.Equal(t, tc.expected, result)
		})
	}
}
