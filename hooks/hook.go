package hooks

import (
	"context"
	"errors"

	"github.com/openpds/connectord/connector"
)

func ValidateReceivingCurrency(s string) connector.TransferCreatorMiddleware {
	return func(tc connector.TransferCreator) connector.TransferCreator {
		return connector.TransferCreatorFunc(func(ctx context.Context, ti *connector.TransferInput, o ...connector.Option) (*connector.TransferOutput, error) {
			if ti.ReceivingCurrency == nil {
				return nil, errors.New("receiving currency should not be nil")
			}

			c := *ti.ReceivingCurrency
			if c != s {
				return nil, errors.New("invalid currency")
			}

			return tc.CreateTransfer(ctx, ti, o...)
		})
	}
}

func ValidateSendingCurrency(s string) connector.TransferCreatorMiddleware {
	return func(tc connector.TransferCreator) connector.TransferCreator {
		return connector.TransferCreatorFunc(func(ctx context.Context, ti *connector.TransferInput, o ...connector.Option) (*connector.TransferOutput, error) {
			if ti.ReceivingCurrency == nil {
				return nil, errors.New("receiving currency should not be nil")
			}

			c := *ti.ReceivingCurrency
			if c != s {
				return nil, errors.New("invalid currency")
			}

			return tc.CreateTransfer(ctx, ti, o...)
		})
	}
}

func DenySendingCountry(countries ...string) connector.TransferCreatorMiddleware {
	return func(tc connector.TransferCreator) connector.TransferCreator {
		return connector.TransferCreatorFunc(func(ctx context.Context, ti *connector.TransferInput, o ...connector.Option) (*connector.TransferOutput, error) {
			if ti.SendingCountry == nil {
				return nil, errors.New("sending country invalid")
			}

			for _, c := range countries {
				if *ti.SendingCountry == c {
					return nil, errors.New("Block it")
				}
			}

			return tc.CreateTransfer(ctx, ti, o...)
		})
	}
}
