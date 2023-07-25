package connector

import "time"

type Option func(*Options)

type Options struct {
	Timeout time.Duration
}
