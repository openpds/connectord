package main

import (
	"context"
	"time"

	"github.com/openpds/connectord/connector"
)

func NewConnectorWrapper(reg *connector.Registry, name string) (*ConnectorWrapper, error) {
	conn, err := reg.Find(name)
	if err != nil {
		return nil, err
	}

	return &ConnectorWrapper{conn}, nil
}

type ConnectorWrapper struct {
	conn connector.Connector
}

func (c ConnectorWrapper) Configure(ctx context.Context, cfg *connector.ConfigureOptions) error {
	if err := ctx.Err(); err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()

	errC := make(chan error)
	defer close(errC)

	go func() {

		if err := c.conn.Configure(cfg); err != nil {
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

func shouldRetry(err error) bool {
	return false
}

func (c ConnectorWrapper) CreateTransfer(ctx context.Context, input *connector.CreateTransferInput) (*connector.CreateTransferOutput, error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}

	creator, ok := c.conn.(connector.TransferCreator)
	if !ok {
		return nil, connector.ErrNotImplemented
	}

	errC := make(chan error)
	defer close(errC)

	respC := make(chan *connector.CreateTransferOutput)
	defer close(respC)

	go func() {
		var (
			attempts    = 0
			maxAttempts = 10
		)

		for attempts = 0; attempts < maxAttempts; attempts++ {
			resp, err := creator.CreateTransfer(ctx, input)
			if err != nil {
				if shouldRetry(err) && attempts < maxAttempts-1 {
					continue
				}

				errC <- err
				return
			}

			respC <- resp
			return
		}
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

func (c ConnectorWrapper) ConfirmTransfer(ctx context.Context, input *connector.CreateTransferInput) (*connector.CreateTransferOutput, error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}

	errC := make(chan error)
	defer close(errC)

	respC := make(chan *connector.CreateTransferOutput)
	defer close(respC)

	go func() {
		creator, ok := c.conn.(connector.TransferCreator)
		if !ok {
			errC <- connector.ErrNotImplemented
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

func (c ConnectorWrapper) CancelTransfer(ctx context.Context, input *connector.CreateTransferInput, opts ...connector.Option) (*connector.CreateTransferOutput, error) {
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
		creator, ok := c.conn.(connector.TransferCreator)
		if !ok {
			errC <- connector.ErrNotImplemented
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
