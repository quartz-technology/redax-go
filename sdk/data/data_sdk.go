package data

//nolint:goimports
import (
	"github.com/quartz-technology/redax-go/relay"
	"github.com/quartz-technology/redax-go/sdk/data/v1"
)

type SDK struct {
	sdkV1  *v1.SDK
	client *relay.Client
}

func NewSDK(client *relay.Client) *SDK {
	return &SDK{
		sdkV1:  v1.NewSDK(client),
		client: client,
	}
}

func (resource *SDK) V1() *v1.SDK {
	return resource.sdkV1
}
