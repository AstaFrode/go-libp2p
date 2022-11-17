package libp2pquic

type Option func(opts *config) error

type config struct {
	disableReuseport bool
	disableDraft29   bool
	metrics          bool
}

func (cfg *config) apply(opts ...Option) error {
	for _, opt := range opts {
		if err := opt(cfg); err != nil {
			return err
		}
	}

	return nil
}

func DisableReuseport() Option {
	return func(cfg *config) error {
		cfg.disableReuseport = true
		return nil
	}
}

// DisableDraft29 disables support for QUIC draft-29.
// This option should be set, unless support for this legacy QUIC version is needed for backwards compatibility.
// Support for QUIC draft-29 is already deprecated and will be removed in the future, see https://github.com/libp2p/go-libp2p/issues/1841.
func DisableDraft29() Option {
	return func(cfg *config) error {
		cfg.disableDraft29 = true
		return nil
	}
}

// WithMetrics enables Prometheus metrics collection.
func WithMetrics() Option {
	return func(cfg *config) error {
		cfg.metrics = true
		return nil
	}
}
