package app

import "time"

type Options struct {
	Js           string
	Css          string
	Production   bool
	Interval     time.Duration
	Workers      int
	Tor          bool //experimental
	CheckAtStart bool
}
