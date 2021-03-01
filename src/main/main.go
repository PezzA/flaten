package main

import (
	"syscall/js"

	"github.com/pezza/wasm"

	"github.com/pezza/flaten/src/model"
)

func main() {
	c := make(chan struct{}, 0)
	runGame()
	<-c
}

func runGame() {

	const cellSize = 32
	const cellWidth, cellHeight = 12, 10

	// setup all the window element, and get the initial state
	window, state := getWindow(cellSize, cellWidth, cellHeight), getNewState()

	// setup the closure bound event handlers
	window.AddEventListener(window.Canvas, "mousemove", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		e := args[0]
		mx := int(e.Get("clientX").Float()) - window.offSetX
		my := int(e.Get("clientY").Float()) - window.offSetY
		state.handleMouseMove(window, model.Point{X: mx, Y: my})
		return nil
	}))

	window.AddEventListener(window.Canvas, "click", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		state.handleClick(window)
		return nil
	}))

	window.AddEventListener(wasm.ParentWindow(), "resize", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		window.offSetX = window.JsCanvas.OffSetLeft()
		window.offSetY = window.JsCanvas.OffSetTop()
		return nil
	}))

	// Start the game loop
	window.StartAnimLoop(func(now float64) {
		state.update(now)
		window.draw(state)
	})
}
