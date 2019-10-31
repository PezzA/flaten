package main

import (
	"syscall/js"
)

type jsDoc struct {
	doc    js.Value
	canvas js.Value
	ctx    js.Value
}

var page jsDoc

func init() {

	doc := js.Global().Get("document")
	canvas := doc.Call("getElementById", "canv")
	ctx := canvas.Call("getContext", "2d")

	page = jsDoc{
		doc:    doc,
		canvas: canvas,
		ctx:    ctx,
	}

	initEvents()
	initAssets()
	initGame()
}

func main() {
	c := make(chan struct{}, 0)

	// clear down bound events
	defer releaseEvents()

	// start up the animation loop
	js.Global().Call("requestAnimationFrame", renderFrameEvt)

	// block forever
	<-c
}
