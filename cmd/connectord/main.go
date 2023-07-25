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
	"fmt"
	"log"
	"time"

	"github.com/openpds/connectord/config"
	connector "github.com/openpds/connectord/connector"
	dummymz "github.com/openpds/connectord/connector/dummy-mz"
)

func main() {
	reg := connector.NewRegistry(dummymz.New())

	cfg, err := config.Init(reg)
	if err != nil {
		panic(err)
	}

	log.Printf("%q", cfg)

	reg.Walk(func(c connector.Connector) {
		fmt.Printf("ID: %s\nNAME: %s\nVERSION: %s\n", c.Manifest().ID, c.Manifest().Name, c.Manifest().Version)
	})

	ticker := time.NewTicker(10 * time.Second)

	for {
		select {
		case t := <-ticker.C:
			run(t)
		}
	}
}

func run(t time.Time) {
	fmt.Println(t)
}

/*func doConfigure(ctx context.Context, connctx *connector.Context, conn connector.Connector) error {
	if err := ctx.Err(); err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()

	errC := make(chan error)
	defer close(errC)

	go func() {
		if err := conn.Configure(connctx); err != nil {
			errC <- err
		}
	}()

	select {
	case <-ctx.Done():
		return context.Canceled
	case err := <-errC:
		return err
	}
}

func (Connectord) doCreateTransfer(ctx context.Context, c connector.Connector, input *connector.CreateTransferInput) (*connector.CreateTransferOutput, error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}

	errC := make(chan error)
	defer close(errC)

	respC := make(chan *connector.CreateTransferOutput)
	defer close(respC)

	go func() {
		creator, ok := c.(connector.TransferCreator)
		if !ok {
			errC <- connector.ErrMotImplemented
			return
		}

		resp, err := creator.CreateTransfer(ctx, input)
		if err != nil {
			errC <- err
			return
		}

		respC <- resp
	}()

	select {
	case <-ctx.Done():
		return nil, context.Canceled
	case err := <-errC:
		return nil, err
	case resp := <-respC:
		return resp, nil
	}
}

func (Connectord) doConfirmTransfer(ctx context.Context, c connector.Connector, input *connector.CreateTransferInput) (*connector.CreateTransferOutput, error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}

	errC := make(chan error)
	defer close(errC)

	respC := make(chan *connector.CreateTransferOutput)
	defer close(respC)

	go func() {
		creator, ok := c.(connector.TransferCreator)
		if !ok {
			errC <- connector.ErrMotImplemented
			return
		}

		resp, err := creator.CreateTransfer(ctx, input)
		if err != nil {
			errC <- err
			return
		}

		respC <- resp
	}()

	select {
	case <-ctx.Done():
		return nil, context.Canceled
	case err := <-errC:
		return nil, err
	case resp := <-respC:
		return resp, nil
	}
}

func (Connectord) doCancelTransfer(ctx context.Context, c connector.Connector, input *connector.CreateTransferInput, opts ...connector.Option) (*connector.CreateTransferOutput, error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}

	o := connector.Options{}

	for _, opt := range opts {
		opt(&o)
	}

	ctx, cancel := context.WithTimeout(ctx, o.Timeout)
	defer cancel()

	errC := make(chan error)
	defer close(errC)

	respC := make(chan *connector.CreateTransferOutput)
	defer close(respC)

	go func() {
		creator, ok := c.(connector.TransferCreator)
		if !ok {
			errC <- connector.ErrMotImplemented
			return
		}

		resp, err := creator.CreateTransfer(ctx, input)
		if err != nil {
			errC <- err
			return
		}

		respC <- resp
	}()

	select {
	case <-ctx.Done():
		return nil, context.Canceled
	case err := <-errC:
		return nil, err
	case resp := <-respC:
		return resp, nil
	}
}
*/
