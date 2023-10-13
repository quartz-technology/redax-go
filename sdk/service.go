package sdk

import (
	"context"
	"encoding/json"
	"net/http"

	v1 "github.com/attestantio/go-builder-client/api/v1"
	"github.com/attestantio/go-eth2-client/spec/phase0"
	"github.com/quartz-technology/redax-go/dto"
	"github.com/quartz-technology/redax-go/ro"
)

type Service struct {
	requester Requester
}

// NewService creates a Service using a custom relay address to initialize the inner Requester.
func NewService(relayBaseURL string) (*Service, error) {
	requester, err := newDefaultRequester(relayBaseURL)
	if err != nil {
		return nil, err
	}

	return NewServiceWithOptions(WithRequester(requester))
}

// NewServiceWithOptions creates a Service using the given options.
func NewServiceWithOptions(opts ...ServiceOption) (*Service, error) {
	sdk := &Service{}
	cfg := newEmptyServiceConfiguration()

	cfg.applyOptions(opts)

	sdk.requester = cfg.requester

	return sdk, nil
}

func (s *Service) GetProposerPayloadDelivered(
	ctx context.Context, params *dto.ProposerPayloadDelivered,
) ([]*ro.ProposerPayloadDeliveredRO, error) {
	path := "/bidtraces/proposer_payload_delivered" + params.EncodeToQueryParams()

	data, err := s.requester.Send(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, ErrFailedToSendRequest(err)
	}

	res := make([]*ro.ProposerPayloadDeliveredRO, 0)

	if err = json.Unmarshal(data, &res); err != nil {
		return nil, ErrFailedToParseResponseData(err)
	}

	return res, nil
}

func (s *Service) GetBuilderBlocksReceived(
	ctx context.Context,
	params *dto.BuilderBlocksReceived,
) ([]*ro.BuilderBlocksReceivedRO, error) {
	path := "/bidtraces/builder_blocks_received" + params.EncodeToQueryParams()

	data, err := s.requester.Send(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, ErrFailedToSendRequest(err)
	}

	res := make([]*ro.BuilderBlocksReceivedRO, 0)

	if err = json.Unmarshal(data, &res); err != nil {
		return nil, ErrFailedToParseResponseData(err)
	}

	return res, nil
}

func (s *Service) GetValidatorRegistration(
	ctx context.Context,
	validatorPublicKey phase0.BLSPubKey,
) (*v1.SignedValidatorRegistration, error) {
	path := "/validator_registration?pubkey=" + validatorPublicKey.String()

	data, err := s.requester.Send(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, ErrFailedToSendRequest(err)
	}

	res := new(v1.SignedValidatorRegistration)

	if err = json.Unmarshal(data, res); err != nil {
		return nil, ErrFailedToParseResponseData(err)
	}

	return res, nil
}
