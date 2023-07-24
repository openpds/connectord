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

package connectors

import (
	"sync"

	connectorsdk "github.com/openpds/connector-sdk"
	dummymz "github.com/openpds/connector-sdk/example/dummy-mz"
)

var (
	r    *connectorsdk.Registry
	once sync.Once
)

func init() {
	once.Do(func() {
		r = connectorsdk.NewRegistry(
			dummymz.New(),
		)
	})
}

func Walk(f func(c connectorsdk.Connector)) {
	r.Walk(f)
}
