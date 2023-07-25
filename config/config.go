// Copyright 2023 The Openpds Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// SPDX-License-Identifier: Apache-2.0

package config

import (
	"time"

	"github.com/openpds/connectord/connector"
)

var (
	DefaultOperation = &Operation{
		Timeout: time.Second,
	}
)

type configOptions struct{}

type Option func(*configOptions)

func Init(r *connector.Registry, opt ...Option) (*Config, error) {
	cfg := &Config{
		Connectors: map[string]*Operations{},
	}

	r.Walk(func(c connector.Connector) {
		cfg.Connectors[c.Manifest().ID] = &Operations{}
		if _, ok := c.(connector.TransferCreator); ok {
			cfg.Connectors[c.Manifest().ID].CreateTransfer = DefaultOperation
		}
	})

	return cfg, nil
}

type Config struct {
	Connectors map[string]*Operations `mapstructure:"connector"`
}

type Operations struct {
	CreateTransfer  *Operation `mapstructure:"create_transfer,omitempty"`
	CancelTransfer  *Operation `mapstructure:"cancel_transfer,omitempty"`
	ConfirmTransfer *Operation `mapstructure:"cancel_transfer,omitempty"`
}

type Operation struct {
	Timeout time.Duration `mapstructure:"timeout"`
}
