package v1

import (
	"context"
	"net/url"

	v1 "github.com/attestantio/go-builder-client/api/v1"
	"github.com/attestantio/go-eth2-client/spec/phase0"
	"github.com/quartz-technology/redax-go/relay"
)

type SDK struct {
	client *relay.Client
}

func NewSDK(client *relay.Client) *SDK {
	return &SDK{
		client: client,
	}
}

func (dataV1 *SDK) GetBidsDelivered(ctx context.Context, params *GetBidsDeliveredRequest) ([]*BidsDelivered, error) {
	if params == nil {
		params = new(GetBidsDeliveredRequest)
	}

	//nolint:exhaustruct
	relayReq := &relay.Request{
		Method: "GET",
		Path:   "/relay/v1/data/bidtraces/proposer_payload_delivered",
		Query:  params.queryParams,
	}

	resp := make([]*BidsDelivered, 0)

	if err := dataV1.client.Do(ctx, relayReq, &resp); err != nil {
		return nil, NewDataSDKError("GetBidsDelivered", err)
	}

	return resp, nil
}

func (dataV1 *SDK) GetBidsReceived(ctx context.Context, params *GetBidsReceivedRequest) ([]*BidsReceived, error) {
	if params == nil {
		return nil, NewDataSDKError("GetBidsReceived", ErrMissingMandatoryParam)
	}

	if err := params.validate(); err != nil {
		return nil, NewDataSDKError("GetBidsReceived", err)
	}

	//nolint:exhaustruct
	relayReq := &relay.Request{
		Method: "GET",
		Path:   "/relay/v1/data/bidtraces/builder_blocks_received",
		Query:  params.queryParams,
	}

	resp := make([]*BidsReceived, 0)

	if err := dataV1.client.Do(ctx, relayReq, &resp); err != nil {
		return nil, NewDataSDKError("GetBidsReceived", err)
	}

	return resp, nil
}

// GetValidatorRegistration TODO
//
//nolint:lll
func (dataV1 *SDK) GetValidatorRegistration(ctx context.Context, validatorPublicKey phase0.BLSPubKey) (*v1.SignedValidatorRegistration, error) {
	query := url.Values{}

	query.Set("pubkey", validatorPublicKey.String())

	//nolint:exhaustruct
	relayReq := &relay.Request{
		Method: "GET",
		Path:   "/relay/v1/data/validator_registration",
		Query:  query,
	}

	resp := new(v1.SignedValidatorRegistration)

	if err := dataV1.client.Do(ctx, relayReq, &resp); err != nil {
		return nil, NewDataSDKError("GetValidatorRegistration", err)
	}

	return resp, nil
}
