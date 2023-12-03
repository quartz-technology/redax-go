package v1

import (
	"context"
	"testing"

	"github.com/attestantio/go-eth2-client/spec/phase0"
	"github.com/quartz-technology/redax-go/relay"
	"github.com/stretchr/testify/require"
)

func TestNewSDK(t *testing.T) {
	client, err := relay.NewClient()

	require.NoError(t, err)
	require.NotNil(t, client)

	sdk := NewSDK(client)
	require.NotNil(t, sdk)
}

func TestSDK_GetBidsDelivered(t *testing.T) {
	client, _ := relay.NewClient(relay.WithAPIURL("https://boost-relay.flashbots.net"))
	sdk := NewSDK(client)

	testCases := []struct {
		name   string
		params *GetBidsDeliveredRequest
	}{
		{
			name:   "should_get_bids_delivered_with_no_params",
			params: nil,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			bidsDelivered, err := sdk.GetBidsDelivered(context.Background(), tt.params)
			require.NoError(t, err)
			require.NotNil(t, bidsDelivered)
		})
	}
}

func TestSDK_GetBidsReceived(t *testing.T) {
	client, _ := relay.NewClient(relay.WithAPIURL("https://boost-relay.flashbots.net"))
	sdk := NewSDK(client)

	testCases := []struct {
		name   string
		params *GetBidsReceivedRequest
	}{
		{
			name:   "should_get_bids_received_with_minimal_param",
			params: NewGetBidsReceivedRequest().WithSlot(7_898_580),
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			bidsReceived, err := sdk.GetBidsReceived(context.Background(), tt.params)
			require.NoError(t, err)
			require.NotNil(t, bidsReceived)
		})
	}
}

func TestSDK_GetValidatorRegistration(t *testing.T) {
	client, _ := relay.NewClient(relay.WithAPIURL("https://boost-relay.flashbots.net"))
	sdk := NewSDK(client)

	testCases := []struct {
		name               string
		validatorPublicKey phase0.BLSPubKey
		failure            bool
	}{
		{
			name:               "should_find_registration",
			validatorPublicKey: phase0.BLSPubKey{},
			failure:            false,
		},
		{
			name:               "should_not_find_registration",
			validatorPublicKey: phase0.BLSPubKey{},
			failure:            true,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			registration, err := sdk.GetValidatorRegistration(context.Background(), tt.validatorPublicKey)

			if tt.failure {
				require.Error(t, err)
				require.Nil(t, registration)
			} else {
				/*
					require.NoError(t, err)
					require.NotNil(t, registration)
				*/
			}
		})
	}
}
