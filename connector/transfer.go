package connector

import (
	"context"
)

type TransferCreator interface {
	CreateTransfer(context.Context, *CreateTransferInput, ...Option) (*CreateTransferOutput, error)
}

type TransferCreatorFunc func(context.Context, *CreateTransferInput, ...Option) (*CreateTransferOutput, error)

func (t TransferCreatorFunc) CreateTransfer(ctx context.Context, input *CreateTransferInput, opt ...Option) (*CreateTransferOutput, error) {
	return t(ctx, input, opt...)
}

type CreateTransferInput struct {
	SendingAmount     *int64  `json:"sending_amount"`
	ReceivingAmount   *int64  `json:"receiving_amount"`
	SendingCurreny    *string `json:"sending_currency"`
	ReceivingCurrency *string `json:"receiving_urrency"`
}

type CreateTransferOutput struct{}
