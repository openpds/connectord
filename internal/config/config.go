package config

import "time"

type Config struct {
	Connectors map[string]Operation `mapstructure:"connectors"`
}

type Operations struct {
	CreateTransfer  Operation `mapstructure:"create_transfer"`
	CancelTransfer  Operation `mapstructure:"cancel_transfer"`
	ConfirmTransfer Operation `mapstructure:"cancel_transfer"`
}

type Operation struct {
	Timeout time.Duration `mapstructure:"timeout"`
}
