package view

import (
	"syscall/js"
)

type JsDoc struct {
	document   js.Value
	canvasElem js.Value
	TwoDCtx    js.Value
}

var clickCallback func(x, y int)
var frameCallback func(now float64)

var offSetLeft, offSetTop float64
var canvasWidth, canvasHeight float64

var mouseMoveEvt, renderFrameEvt, canvasClickEvt js.Func
var mousePos [2]float64

// NewJsDoc returns a new JsDoc initted with assets and events
func NewJsDoc(click func(x, y int), frame func(now float64)) JsDoc {
	doc := js.Global().Get("document")
	canvas := doc.Call("getElementById", "canv")
	ctx := canvas.Call("getContext", "2d")

	offSetLeft = canvas.Get("offsetLeft").Float()
	offSetTop = canvas.Get("offsetTop").Float()

	clickCallback = click
	frameCallback = frame

	jsDoc := JsDoc{
		document:   doc,
		canvasElem: canvas,
		TwoDCtx:    ctx,
	}

	jsDoc.initEvents()
	jsDoc.initAssets()

	return jsDoc
}

func (d *JsDoc) initEvents() {
	mouseMoveEvt = js.FuncOf(mouseMove)
	d.document.Call("addEventListener", "mousemove", mouseMoveEvt)

	canvasClickEvt = js.FuncOf(canvasClick)
	d.canvasElem.Call("addEventListener", "click", canvasClickEvt)

	renderFrameEvt = js.FuncOf(renderFrame)
}

func releaseEvents() {
	mouseMoveEvt.Release()
	renderFrameEvt.Release()
	canvasClickEvt.Release()
}

func canvasClick(this js.Value, args []js.Value) interface{} {
	clickCallback(int(mousePos[0]), int(mousePos[1]))
	return nil
}

// MAIN GAME LOOP
func renderFrame(this js.Value, args []js.Value) interface{} {
	frameCallback(args[0].Float())
	js.Global().Call("requestAnimationFrame", renderFrameEvt)
	return nil
}

func mouseMove(this js.Value, args []js.Value) interface{} {
	e := args[0]

	mousePos[0] = e.Get("clientX").Float() - offSetLeft
	mousePos[1] = e.Get("clientY").Float() - offSetTop

	if mousePos[0] < 0 {
		mousePos[0] = 0
	}
	if mousePos[1] < 0 {
		mousePos[1] = 0
	}

	if mousePos[0] > canvasWidth {
		mousePos[0] = canvasWidth
	}

	if mousePos[1] > canvasHeight {
		mousePos[1] = canvasHeight
	}

	return nil
}
