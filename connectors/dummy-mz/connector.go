package dummymz

import (
	"context"
	"errors"

	connectorsdk "github.com/openpds/connector-sdk"
)

func New() *dummymz {
	return &dummymz{}
}

type RMSpec struct {
	minLength *int
	maxLength *int
	length    *int
	prefix    *string
	suffix    *string
}

func Int(i int) *int {
	return &i
}

type RMOption func(*RMSpec)

func WithMinLength(l int) RMSpec {
	return RMSpec{
		minLength: &l,
	}
}

func WithMaxLength(l int) RMSpec {
	return RMSpec{
		maxLength: &l,
	}
}

func WithPrefix(s string) RMSpec {
	return RMSpec{
		prefix: &s,
	}
}

type config struct {
	ip    string
	port  string
	path  string
	token string
}

type dummymz struct {
	config config
}

func (d dummymz) ID() string {
	return "dummy-mz"
}

func (d dummymz) Name() string {
	return "Dummy"
}

func (d dummymz) Version() string {
	return "0.1.0"
}

func (d dummymz) Requires() connectorsdk.Requirement {
	return connectorsdk.Requirement{
		Vars: []connectorsdk.VarSpec{
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
		Secrets: []connectorsdk.SecretSpec{
			{
				Name:        "token",
				Type:        "string",
				Description: "IP represents an IP address",
			},
		},
		ReceivingMethods: []RMSpec{
			WithMinLength(10),
			WithMaxLength(12),
			WithPrefix("258"),
		},
	}
}

func (d *dummymz) Configure(ctx *connectorsdk.Context) error {
	var err error

	d.config.ip, err = ctx.Vars.GetString("ip")
	if err != nil {
		return err
	}

	d.config.port, err = ctx.Vars.GetString("port")
	if err != nil {
		return err
	}

	return nil
}

func (d *dummymz) CreateTransfer(ctx context.Context, input *connectorsdk.CreateTransferInput, opts ...connectorsdk.Option) (*connectorsdk.CreateTransferOutput, error) {
	return nil, errors.New("not implemented")
}
