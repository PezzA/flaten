package main

import (
	"syscall/js"
)

var mouseMoveEvt, resetClickEvt, renderFrameEvt, canvasClickEvt, muteToggleEvt js.Func
var mousePos [2]float64

var musicVol = 1

func initEvents() {
	mouseMoveEvt = js.FuncOf(mouseMove)
	page.doc.Call("addEventListener", "mousemove", mouseMoveEvt)

	resetClickEvt = js.FuncOf(resetClick)
	page.doc.Call("getElementById", "reset").Call("addEventListener", "click", resetClickEvt)

	muteToggleEvt = js.FuncOf(muteToggleClick)
	page.doc.Call("getElementById", "muteToggle").Call("addEventListener", "click", muteToggleEvt)

	canvasClickEvt = js.FuncOf(canvasClick)
	page.canvas.Call("addEventListener", "click", canvasClickEvt)

	renderFrameEvt = js.FuncOf(renderFrame)
}

func releaseEvents() {
	mouseMoveEvt.Release()
	resetClickEvt.Release()
	renderFrameEvt.Release()
	canvasClickEvt.Release()
}

func canvasClick(this js.Value, args []js.Value) interface{} {
	if handleClick(int(mousePos[0]), int(mousePos[1])) {
		click.Set("currentTime", 0)
		click.Call("play")
	}
	return nil
}
func renderFrame(this js.Value, args []js.Value) interface{} {
	page.ctx.Call("clearRect", 0, 0, displayWidth, displayHeight)
	updateGame(args[0].Float())
	drawGameGrid()

	js.Global().Call("requestAnimationFrame", renderFrameEvt)
	return nil
}

func muteToggleClick(this js.Value, args []js.Value) interface{} {
	if musicVol == 0 {
		musicVol = 1
	} else {
		musicVol = 0
	}
	music.Set("volume", musicVol)
	return nil
}

func resetClick(this js.Value, args []js.Value) interface{} {
	initGame()
	music.Set("currentTime", 0)
	music.Call("play")
	return nil
}

func mouseMove(this js.Value, args []js.Value) interface{} {
	e := args[0]

	mousePos[0] = e.Get("clientX").Float() - page.canvas.Get("offsetLeft").Float()
	mousePos[1] = e.Get("clientY").Float() - page.canvas.Get("offsetTop").Float()

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
