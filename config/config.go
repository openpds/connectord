package config

import (
	"time"

	connectorsdk "github.com/openpds/connector-sdk"
	"github.com/openpds/connectord/connectors"
)

var (
	DefaultOperation = &Operation{
		Timeout: time.Second,
	}
)

type configOptions struct{}

type Option func(*configOptions)

func Init(opt ...Option) (*Config, error) {
	cfg := &Config{
		Connectors: map[string]*Operations{},
	}

	connectors.Walk(func(c connectorsdk.Connector) {
		cfg.Connectors[c.ID()] = &Operations{}
		if _, ok := c.(connectorsdk.TransferCreator); ok {
			cfg.Connectors[c.ID()].CreateTransfer = DefaultOperation
		}
	})

	return cfg, nil
}

type Config struct {
	Connectors map[string]*Operations `mapstructure:"connectors"`
}

type Operations struct {
	CreateTransfer  *Operation `mapstructure:"create_transfer"`
	CancelTransfer  *Operation `mapstructure:"cancel_transfer"`
	ConfirmTransfer *Operation `mapstructure:"cancel_transfer"`
}

type Operation struct {
	Timeout time.Duration `mapstructure:"timeout"`
}
