package sdk

// ServiceOption is used to customize / update the Service configuration.
type ServiceOption func(s *serviceConfiguration)

// serviceConfiguration holds all the options used to construct a customized Service.
type serviceConfiguration struct {
	requester Requester
}

// newEmptyServiceConfiguration creates a Service configuration with all the options being empty.
func newEmptyServiceConfiguration() *serviceConfiguration {
	return &serviceConfiguration{
		requester: nil,
	}
}

// applyOptions takes a list of options and apply each one of those to the configuration.
func (cfg *serviceConfiguration) applyOptions(opts []ServiceOption) {
	for _, opt := range opts {
		opt(cfg)
	}
}

// WithRequester is an option used to pass a custom HTTP client used by the Service.
func WithRequester(requester Requester) ServiceOption {
	return func(s *serviceConfiguration) {
		s.requester = requester
	}
}
