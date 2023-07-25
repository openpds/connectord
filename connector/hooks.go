package connector

import "context"

type PreTransferCreation interface {
	PreCreateTransfer(ctx context.Context, input *TransferInput) error
}

type PostTransferCreation interface {
	PostCreateTransfer(ctx context.Context, input *TransferOutput) error
}

type PreTransferCancellation interface {
	PreCancelTransfer(ctx context.Context, input *TransferInput) error
}

type PostTransferCancellation interface {
	PostCancelTransfer(ctx context.Context, input *TransferOutput) error
}

type PreTransferConfirmation interface {
	PreConfirmTransfer(ctx context.Context, input *TransferInput) error
}

type PostTransferConfirmation interface {
	PostConfirmTransfer(ctx context.Context, input *TransferOutput) error
}
