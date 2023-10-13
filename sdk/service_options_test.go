package sdk

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

type mockRequesterWithNoBehavior struct{}

func (m *mockRequesterWithNoBehavior) Send(_ context.Context, _ string, _ string, _ []byte) ([]byte, error) {
	return []byte{}, nil
}

func TestServiceOptions(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name                  string
		options               []ServiceOption
		expectedConfiguration *serviceConfiguration
	}{
		{
			name:    "Creates a configuration with no options",
			options: nil,
			expectedConfiguration: &serviceConfiguration{
				requester: nil,
			},
		},
		{
			name:    "Creates a configuration with one option",
			options: []ServiceOption{WithRequester(nil)},
			expectedConfiguration: &serviceConfiguration{
				requester: nil,
			},
		},
		{
			name:    "Creates a configuration with two options that overlap",
			options: []ServiceOption{WithRequester(nil), WithRequester(&mockRequesterWithNoBehavior{})},
			expectedConfiguration: &serviceConfiguration{
				requester: &mockRequesterWithNoBehavior{},
			},
		},
	}

	for i := range testCases {
		//nolint:varnamelen
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			configuration := newEmptyServiceConfiguration()
			configuration.applyOptions(tc.options)

			require.Equal(t, tc.expectedConfiguration, configuration)
		})
	}
}
