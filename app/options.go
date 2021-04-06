package app

import "time"

type Options struct {
	Production bool
	Interval   time.Duration
	Workers    int
	Tor        bool //experimental
}
