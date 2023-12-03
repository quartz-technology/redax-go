package sdk

import (
	"github.com/quartz-technology/redax-go/relay"
	"github.com/quartz-technology/redax-go/sdk/data"
)

type RelaySDK struct {
	dataSDK *data.SDK
	client  *relay.Client
}

func NewRelaySDK(client *relay.Client) *RelaySDK {
	return &RelaySDK{
		dataSDK: data.NewSDK(client),
		client:  client,
	}
}

func (s *RelaySDK) Data() *data.SDK {
	return s.dataSDK
}
