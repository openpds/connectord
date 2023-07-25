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

package main

import (
	"encoding/json"
	"fmt"

	"github.com/openpds/connectord/config"
	connector "github.com/openpds/connectord/connector"
	dummyv1mz "github.com/openpds/connectord/connector/dummy-v1-mz"
	dummyv2mz "github.com/openpds/connectord/connector/dummy-v2-mz"
)

func main() {
	reg := connector.NewRegistry(
		dummyv1mz.New(),
		dummyv2mz.New(),
	)

	cfg, err := config.Init(reg)
	if err != nil {
		panic(err)
	}

	b, err := json.MarshalIndent(cfg, "", "  ")
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", string(b))

	reg.Walk(func(c connector.Connector) {
		print(c)
	})
}

func print(c connector.Connector) {
	fmt.Printf("ID: %s\nNAME: %s\nVERSION: %s\n\n", c.Manifest().ID, c.Manifest().Name, c.Manifest().Version)
}
