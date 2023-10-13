package sdk

import (
	"context"
	"math/big"
	"testing"

	"github.com/attestantio/go-eth2-client/spec/phase0"
	"github.com/quartz-technology/redax-go/dto"
	"github.com/stretchr/testify/require"
)

func TestService(t *testing.T) {
	t.Parallel()

	service, err := NewService("https://boost-relay.flashbots.net")
	require.NoError(t, err)

	validatorPublicKey := new(phase0.BLSPubKey)
	err = validatorPublicKey.UnmarshalJSON([]byte(
		"\"0x874699cc4a92d97c5cc44d30e6ea570d4f76242ce9321f6f134a4d9fa25a5219c0af11c25163b37960f570bc99a35d17\""))
	require.NoError(t, err)

	registrations, err := service.GetValidatorRegistration(context.Background(), *validatorPublicKey)
	require.NoError(t, err)
	require.NotNil(t, registrations)

	params1 := new(dto.BuilderBlocksReceived)
	params1.BlockNumber = big.NewInt(18343016)

	builderBlocksReceived, err := service.GetBuilderBlocksReceived(context.Background(), params1)
	require.NoError(t, err)
	require.NotNil(t, builderBlocksReceived)

	params2 := new(dto.ProposerPayloadDelivered)
	params2.BlockNumber = big.NewInt(18343016)

	proposerPayloadDelivered, err := service.GetProposerPayloadDelivered(context.Background(), params2)
	require.NoError(t, err)
	require.NotNil(t, proposerPayloadDelivered)
}
