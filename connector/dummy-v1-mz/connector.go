package dummyv1mz

import (
	"context"
	"errors"

	"github.com/openpds/connectord/connector"
	"github.com/openpds/connectord/hooks"
)

func New() *dummymz {
	return &dummymz{}
}

type dummymz struct {
	config config
}

type config struct {
	ip    string
	port  string
	path  string
	token string
}

func (d *dummymz) Configure(opts *connector.ConfigureOptions) error {
	var err error

	d.config.ip, err = opts.Variables.GetString("ip")
	if err != nil {
		return err
	}

	d.config.port, err = opts.Variables.GetString("port")
	if err != nil {
		return err
	}

	return nil
}

func (d dummymz) Manifest() connector.Manifest {
	return connector.Manifest{
		ID:      "dummy-v1-mz",
		Name:    "Dummy",
		Version: "0.1.0",
		Variables: []connector.Variable{
			{
				Name:    "ip",
				Type:    "string",
				Desc:    "IP represents an IP address",
				Default: "127.0.0.1",
			},
			{
				Name: "port",
				Type: "string",
				Desc: "IP represents an IP address",
			},
			{
				Name: "path",
				Type: "string",
				Desc: "IP represents an IP address",
			},
		},
		Secrets: []connector.Secret{
			{
				Name:        "token",
				Type:        "string",
				Description: "IP represents an IP address",
			},
		},
		Hooks: connector.Hooks{
			PreCreateTransfer: []connector.TransferCreatorMiddleware{
				hooks.ValidateReceivingCurrency("MZN"),
				hooks.DenySendingCountry("CUB"),
			},
		},
	}
}

func (d *dummymz) CreateTransfer(ctx context.Context, input *connector.TransferInput, opts ...connector.Option) (*connector.TransferOutput, error) {
	return nil, errors.New("not implemented")
}

func (d *dummymz) CancelTransfer(ctx context.Context, input *connector.TransferInput, opts ...connector.Option) (*connector.TransferOutput, error) {
	return nil, errors.New("not implemented")
}

func (d *dummymz) ConfirmTransfer(ctx context.Context, input *connector.TransferInput, opts ...connector.Option) (*connector.TransferOutput, error) {
	return nil, errors.New("not implemented")
}

func (d *dummymz) CheckTransfer(ctx context.Context, input *connector.TransferInput, opts ...connector.Option) (*connector.TransferOutput, error) {
	return nil, errors.New("not implemented")
}
