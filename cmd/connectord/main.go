package main

import (
	"context"
	"fmt"
	"time"

	connectorsdk "github.com/openpds/connector-sdk"
	dummymz "github.com/openpds/connector-sdk/example/dummy-mz"
)

var (
	r = connectorsdk.NewRegistry(
		dummymz.New(),
	)
)

func main() {
	r.Walk(func(c connectorsdk.Connector) {
		fmt.Printf("ID: %s\nNAME: %s\nVERSION: %s", c.ID(), c.Name(), c.Version())
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

func doConfigure(ctx context.Context, connctx *connectorsdk.Context, conn connectorsdk.Connector) error {
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

type Connectord struct {
	config *Config
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
