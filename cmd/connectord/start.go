package main

import (
	"context"
	"time"

	"github.com/openpds/connectord/connector"
)

type ConnectorWrapperOption func(*ConnectorWrapper)

func NewConnectorWrapper(reg *connector.Registry, name string) (*ConnectorWrapper, error) {
	conn, err := reg.Find(name)
	if err != nil {
		return nil, err
	}

	// process metrics
	// process traces

	return &ConnectorWrapper{conn: conn}, nil
}

func WithBeforCreateTransfer(mw connector.TransferCreatorMiddleware) ConnectorWrapperOption {
	return func(cw *ConnectorWrapper) {
		if cw.beforCreateTransfer == nil {
			cw.beforCreateTransfer = make([]connector.TransferCreatorMiddleware, 0)
		}

		cw.beforCreateTransfer = append(cw.beforCreateTransfer, mw)
	}
}

type ConnectorWrapper struct {
	beforCreateTransfer []connector.TransferCreatorMiddleware
	conn                connector.Connector
}

func (c ConnectorWrapper) Configure(ctx context.Context, cfg *connector.ConfigureOptions) error {
	// process metrics
	// process traces

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

func (c ConnectorWrapper) CreateTransfer(ctx context.Context, input *connector.TransferInput) (*connector.TransferOutput, error) {
	// process metrics
	// process traces

	if err := ctx.Err(); err != nil {
		return nil, err
	}

	tc, ok := c.conn.(connector.TransferCreator)
	if !ok {
		return nil, connector.ErrNotImplemented
	}

	for _, mw := range c.beforCreateTransfer {
		tc = mw(tc)
	}

	if hook, ok := tc.(connector.PreTransferCreation); ok {
		if err := hook.PreCreateTransfer(ctx, input); err != nil {
			return nil, err
		}
	}

	errC := make(chan error)
	defer close(errC)

	respC := make(chan *connector.TransferOutput)
	defer close(respC)

	go func() {
		var (
			attempts    = 0
			maxAttempts = 10
		)

		for attempts = 0; attempts < maxAttempts; attempts++ {
			resp, err := tc.CreateTransfer(ctx, input)
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
		if hook, ok := tc.(connector.PostTransferCreation); ok {
			if err := hook.PostCreateTransfer(ctx, resp); err != nil {
				return nil, err
			}
		}

		return resp, nil
	}
}

func (c ConnectorWrapper) ConfirmTransfer(ctx context.Context, input *connector.TransferInput) (*connector.TransferOutput, error) {
	// process metrics
	// process traces

	if err := ctx.Err(); err != nil {
		return nil, err
	}

	cc, ok := c.conn.(connector.TransferCreator)
	if !ok {
		return nil, connector.ErrNotImplemented
	}

	errC := make(chan error)
	defer close(errC)

	respC := make(chan *connector.TransferOutput)
	defer close(respC)

	go func() {
		resp, err := cc.CreateTransfer(ctx, input)
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

func (c ConnectorWrapper) CancelTransfer(ctx context.Context, input *connector.TransferInput, opts ...connector.Option) (*connector.TransferOutput, error) {
	// process metrics
	// process traces

	if err := ctx.Err(); err != nil {
		return nil, err
	}

	cc, ok := c.conn.(connector.TransferCreator)
	if !ok {
		return nil, connector.ErrNotImplemented
	}

	o := connector.Options{}

	for _, opt := range opts {
		opt(&o)
	}

	ctx, cancel := context.WithTimeout(ctx, o.Timeout)
	defer cancel()

	errC := make(chan error)
	defer close(errC)

	respC := make(chan *connector.TransferOutput)
	defer close(respC)

	go func() {
		resp, err := cc.CreateTransfer(ctx, input)
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
