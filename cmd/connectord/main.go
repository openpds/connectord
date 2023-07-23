package main

import (
	"fmt"

	connectorsdk "github.com/openpds/connector-sdk"
	dummymz "github.com/openpds/connector-sdk/example/dummy-mz"
)

var (
	r = connectorsdk.NewRegistry(
		dummymz.New(),
	)
)

func main() {
	r.For(func(c connectorsdk.Connector) {
		fmt.Printf("ID: %s\nNAME: %s\nVERSION: %s", c.ID(), c.Name(), c.Version())
	})
}
