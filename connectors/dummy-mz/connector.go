package dummymz

import (
	"context"
	"errors"

	connectorsdk "github.com/openpds/connector-sdk"
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

func (d *dummymz) Configure(opts *connectorsdk.ConfigureOptions) error {
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

func (d dummymz) Manifest() connectorsdk.Manifest {
	return connectorsdk.Manifest{
		ID:      "dummy-mz",
		Name:    "Dummy",
		Version: "0.1.0",
		Variables: []connectorsdk.Variable{
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
		Secrets: []connectorsdk.Secret{
			{
				Name:        "token",
				Type:        "string",
				Description: "IP represents an IP address",
			},
		},
	}
}

func (d *dummymz) CreateTransfer(ctx context.Context, input *connectorsdk.CreateTransferInput, opts ...connectorsdk.Option) (*connectorsdk.CreateTransferOutput, error) {
	return nil, errors.New("not implemented")
}
