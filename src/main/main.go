package main

import (
	"syscall/js"

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
	window.AddEventListener(window.Document, "mousemove", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		e := args[0]
		state.handleMouseMove(window, model.Point{X: int(e.Get("clientX").Float()), Y: int(e.Get("clientY").Float())})
		return nil
	}))

	window.AddEventListener(window.Canvas, "click", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		state.handleClick(window)
		return nil
	}))

	// Start the game loop
	window.StartAnimLoop(func(now float64) {
		state.update(now)
		window.draw(state)
	})
}
