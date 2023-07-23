package main

import (
	"fmt"

	connectorsdk "github.com/openpds/connector-sdk"
	dummymz "github.com/openpds/connector-sdk/example/dummy-mz"
)

var (
	r = connectorsdk.NewRegistry()
)

func main() {
	register()

	r.For(func(c connectorsdk.Connector) {
		fmt.Printf("ID: %s\nNAME: %s\nVERSION: %s", c.ID(), c.Name(), c.Version())
	})
}

func register() {
	r.Register(dummymz.New())
}
