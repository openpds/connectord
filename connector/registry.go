package connector

import (
	"errors"
	"fmt"
	"sync"
)

var (
	r    *Registry
	once sync.Once
)

func Walk(f func(c Connector)) {
	r.Walk(f)
}

func NewRegistry(c ...Connector) *Registry {
	r := &Registry{
		connector: make(map[string]Connector, 0),
	}

	for _, conn := range c {
		r.Register(conn)
	}

	return r
}

type Registry struct {
	mu        sync.Mutex
	connector map[string]Connector
}

func (r *Registry) Register(conn Connector) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	m := conn.Manifest()

	if _, ok := r.connector[m.ID]; ok {
		return fmt.Errorf("connector with ID %s already exists", m.ID)
	}

	r.connector[m.ID] = conn

	return nil
}

func (r *Registry) Walk(f func(Connector)) {
	for _, c := range r.connector {
		r.mu.Lock()
		f(c)
		r.mu.Unlock()
	}
}

func (r *Registry) Find(s string) (Connector, error) {
	conn, ok := r.connector[s]
	if !ok {
		return nil, errors.New("not found")
	}

	return conn, nil
}

type Introspection struct{}

func Introspect(conn Connector) Introspection {
	return Introspection{}
}
