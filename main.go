package main

import (
	"krolus/app"
	"time"
)

func main() {
	new(app.KrolusApp).Start(app.Options{
		Production: false,
		Interval:   30 * time.Minute,
		Workers:    3,
	})
}
