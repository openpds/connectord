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

	connectorsdk "github.com/openpds/connector-sdk"
	"github.com/openpds/connectord/config"
	"github.com/openpds/connectord/connectors"
)

func main() {
	cfg, err := config.Init()
	if err != nil {
		panic(err)
	}

	log.Printf("%q", cfg)

	connectors.Walk(func(c connectorsdk.Connector) {
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

/*func doConfigure(ctx context.Context, connctx *connectorsdk.Context, conn connectorsdk.Connector) error {
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

func (Connectord) doCreateTransfer(ctx context.Context, c connectorsdk.Connector, input *connectorsdk.CreateTransferInput) (*connectorsdk.CreateTransferOutput, error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}

	errC := make(chan error)
	defer close(errC)

	respC := make(chan *connectorsdk.CreateTransferOutput)
	defer close(respC)

	go func() {
		creator, ok := c.(connectorsdk.TransferCreator)
		if !ok {
			errC <- connectorsdk.ErrMotImplemented
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

func (Connectord) doConfirmTransfer(ctx context.Context, c connectorsdk.Connector, input *connectorsdk.CreateTransferInput) (*connectorsdk.CreateTransferOutput, error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}

	errC := make(chan error)
	defer close(errC)

	respC := make(chan *connectorsdk.CreateTransferOutput)
	defer close(respC)

	go func() {
		creator, ok := c.(connectorsdk.TransferCreator)
		if !ok {
			errC <- connectorsdk.ErrMotImplemented
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

func (Connectord) doCancelTransfer(ctx context.Context, c connectorsdk.Connector, input *connectorsdk.CreateTransferInput, opts ...connectorsdk.Option) (*connectorsdk.CreateTransferOutput, error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}

	o := connectorsdk.Options{}

	for _, opt := range opts {
		opt(&o)
	}

	ctx, cancel := context.WithTimeout(ctx, o.Timeout)
	defer cancel()

	errC := make(chan error)
	defer close(errC)

	respC := make(chan *connectorsdk.CreateTransferOutput)
	defer close(respC)

	go func() {
		creator, ok := c.(connectorsdk.TransferCreator)
		if !ok {
			errC <- connectorsdk.ErrMotImplemented
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
