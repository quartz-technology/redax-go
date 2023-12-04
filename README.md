# redax-go

Go SDK for the Relay Data Transparency API on Ethereum.

## Getting Started

To install the SDK in your own go project, run the following command:
```shell
go get github.com/quartz-technology/redax-go
```

### Examples

_More examples will come soon !_

```go
package test

import (
	"context"
	"fmt"
	"testing"

	"github.com/quartz-technology/redax-go/relay"
	"github.com/quartz-technology/redax-go/sdk"
	"github.com/stretchr/testify/require"
)

func TestRedaxSDK(t *testing.T) {
	client, err := relay.NewClient(relay.WithAPIURL("https://boost-relay.flashbots.net"))
	require.NoError(t, err)

	relaySDK := sdk.NewRelaySDK(client)

	// Get the bids delivered.
	bidsDelivered, err := relaySDK.Data().V1().GetBidsDelivered(context.Background(), nil)
	require.NoError(t, err)

	// Print the first bid's value.
	fmt.Println(bidsDelivered[0].Value.String())
}
```

### API

Below is a list of the supported API endpoints:

| Name       	                                            | Status 	 |
|:--------------------------------------------------------|:--------:|
| `/relay/v1/data/bidtraces/proposer_payload_delivered` 	 |  ✅   	   |
| `/relay/v1/data/bidtraces/builder_blocks_received` 	    |  ✅   	   |
| `/relay/v1/data/validator_registration` 	               |  ✅   	   |

## Issues

This SDK is still under active development, if you find any bug or have a feature request please
submit an appropriate issue [here](https://github.com/quartz-technology/redax-go/issues/new/choose).

## Contributing

If you would like to contribute to this project, please refer to the instructions in the
dedicated document [here](./CONTRIBUTING.md).

## Authors

This project is a pure open-source contribution to the Ethereum ecosystem.
It is currently maintained by the 🤖 at [Quartz Technology](https://github.com/quartz-technology).