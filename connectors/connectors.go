package connectors

import (
	"sync"

	connectorsdk "github.com/openpds/connector-sdk"
	dummymz "github.com/openpds/connector-sdk/example/dummy-mz"
)

var (
	r    *connectorsdk.Registry
	once sync.Once
)

func init() {
	once.Do(func() {
		r = connectorsdk.NewRegistry(
			dummymz.New(),
		)
	})
}

func Walk(f func(c connectorsdk.Connector)) {
	r.Walk(f)
}
