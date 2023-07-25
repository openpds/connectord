package connector

import (
	"errors"
)

type ConfigureOptions struct {
	Variables *Variables
	Secrets   *Secrets
}

type Connector interface {
	Manifest() Manifest
	Configure(*ConfigureOptions) error
}

type Manifest struct {
	ID         string     `json:"id"`
	Version    string     `json:"version"`
	Name       string     `json:"name"`
	Variables  []Variable `json:"variables"`
	Secrets    []Secret   `json:"secrets"`
	Deprecated bool       `json:"deprecated"`
	Enabled    bool       `json:"enabled"`
	Tags       []string   `json:"tags"`
}

func (m Manifest) Validate() error {
	if m.ID == "" {
		return errors.New("id should not be empty")
	}

	if m.Name == "" {
		return errors.New("name should not be empty")
	}

	if m.Version == "" {
		return errors.New("id should not be empty")
	}

	return nil
}

type Variables struct {
	entries map[string]Variables
}

type Variable struct {
	Name     string
	Type     string
	Desc     string
	Default  interface{}
	Example  string
	Required bool
	Value    interface{}
}

func (v Variables) Get(key string) (interface{}, error) {
	v, ok := v.entries[key]
	if !ok {
		return nil, errors.New("not found")
	}

	return v, nil
}

func (v Variables) GetString(key string) (string, error) {
	return "", errors.New("not defined")
}

type Secrets struct {
	entries map[string]Secret
}

type Secret struct {
	Name        string
	Type        string
	Description string
	Default     string
	Required    bool
	Value       interface{}
}

func (v Secrets) Get(key string) (interface{}, error) {
	return nil, errors.New("not defined")
}
