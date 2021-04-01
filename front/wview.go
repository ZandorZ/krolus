package main

import "github.com/webview/webview"

func main() {
	debug := true
	w := webview.New(debug)
	defer w.Destroy()
	w.SetTitle("krolus")
	w.SetSize(1024, 768, webview.HintNone)
	w.Navigate("http://localhost:4200")
	w.Run()
}
