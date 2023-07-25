package connector

import (
	"context"
)

type TransferCanceller interface {
	CancelTransfer(context.Context, *TransferInput, ...Option) (*TransferOutput, error)
}

type TransferCancellerFunc func(context.Context, *TransferInput, ...Option) (*TransferOutput, error)

func (t TransferCancellerFunc) CancelTransfer(ctx context.Context, input *TransferInput, opt ...Option) (*TransferOutput, error) {
	return t(ctx, input, opt...)
}

type TransferConfirmer interface {
	ConfirmTransfer(context.Context, *TransferInput, ...Option) (*TransferOutput, error)
}

type TransferConfirmerFunc func(context.Context, *TransferInput, ...Option) (*TransferOutput, error)

func (t TransferConfirmerFunc) ConfirmTransfer(ctx context.Context, input *TransferInput, opt ...Option) (*TransferOutput, error) {
	return t(ctx, input, opt...)
}

type TransferCreator interface {
	CreateTransfer(context.Context, *TransferInput, ...Option) (*TransferOutput, error)
}

type TransferCreatorFunc func(context.Context, *TransferInput, ...Option) (*TransferOutput, error)

func (t TransferCreatorFunc) CreateTransfer(ctx context.Context, input *TransferInput, opt ...Option) (*TransferOutput, error) {
	return t(ctx, input, opt...)
}

type TransferChecker interface {
	CheckTransfer(context.Context, *TransferInput, ...Option) (*TransferOutput, error)
}

type TransferCheckerFunc func(context.Context, *TransferInput, ...Option) (*TransferOutput, error)

func (t TransferCheckerFunc) CheckTransfer(ctx context.Context, input *TransferInput, opt ...Option) (*TransferOutput, error) {
	return t(ctx, input, opt...)
}

type TransferInput struct {
	SendingAmount     *int64  `json:"sending_amount"`
	ReceivingAmount   *int64  `json:"receiving_amount"`
	SendingCurreny    *string `json:"sending_currency"`
	ReceivingCurrency *string `json:"receiving_urrency"`
	SendingCountry    *string `json:"sending_country"`
	ReceivingCountry  *string `json:"receiving_country"`
}

type TransferOutput struct {
	Status string
}
