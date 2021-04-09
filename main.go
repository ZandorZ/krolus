package main

import (
	_ "embed"
	"krolus/app"
	"time"
)

//go:embed front/frontend/dist/frontend/main.js
var js string

//go:embed front/frontend/dist/frontend/styles.css
var css string

func main() {

	new(app.KrolusApp).Start(app.Options{
		Js:         js,
		Css:        css,
		Production: true,
		Interval:   30 * time.Minute,
		Workers:    3,
	})
}
