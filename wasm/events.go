package main

import (
	"syscall/js"
)

var mouseMoveEvt, resetClickEvt, renderFrameEvt, canvasClickEvt js.Func
var mousePos [2]float64

func initEvents() {
	mouseMoveEvt = js.FuncOf(mouseMove)
	doc.Call("addEventListener", "mousemove", mouseMoveEvt)

	resetClickEvt = js.FuncOf(resetClick)
	doc.Call("getElementById", "reset").Call("addEventListener", "click", resetClickEvt)

	canvasClickEvt = js.FuncOf(canvasClick)
	canvas.Call("addEventListener", "click", canvasClickEvt)

	renderFrameEvt = js.FuncOf(renderFrame)
}

func releaseEvents() {
	mouseMoveEvt.Release()
	resetClickEvt.Release()
	renderFrameEvt.Release()
	canvasClickEvt.Release()
}

func canvasClick(this js.Value, args []js.Value) interface{} {
	handleClick(int(mousePos[0]), int(mousePos[1]))
	return nil
}
func renderFrame(this js.Value, args []js.Value) interface{} {
	ctx.Call("clearRect", 0, 0, displayWidth, displayHeight)
	updateGame(args[0].Float())
	drawGameGrid()

	js.Global().Call("requestAnimationFrame", renderFrameEvt)
	return nil
}

func resetClick(this js.Value, args []js.Value) interface{} {
	initGame()
	return nil
}

func mouseMove(this js.Value, args []js.Value) interface{} {
	e := args[0]

	mousePos[0] = e.Get("clientX").Float() - canvas.Get("offsetLeft").Float()
	mousePos[1] = e.Get("clientY").Float() - canvas.Get("offsetTop").Float()

	if mousePos[0] < 0 {
		mousePos[0] = 0
	}
	if mousePos[1] < 0 {
		mousePos[1] = 0
	}

	if mousePos[0] > float64(displayWidth) {
		mousePos[0] = float64(displayWidth)
	}

	if mousePos[1] > float64(displayHeight) {
		mousePos[1] = float64(displayHeight)
	}

	return nil
}
