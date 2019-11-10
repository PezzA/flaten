package main

import (
	"github.com/pezza/flaten/model"
	"github.com/pezza/flaten/view"
)

var d view.JsDoc
var g model.Game

var gridSize = 32
var gameWidth, gameHeight = 12, 10

var gridDisplayWidth, gridDisplayHeight = gameWidth * gridSize, gameHeight * gridSize
var canvasWidth, canvasHeight = gridDisplayWidth * 2, gridDisplayHeight + (2 * gridSize)
var gridXOffSet = gridDisplayWidth / 2
var gridYOffSet = gridSize

func init() {
	d = view.NewJsDoc(handleClick, initGame, doGameLoop)
	d.SetCanvasSize(canvasWidth, canvasHeight)
	d.StartAnimLoop()
}

func main() {
	//run forever
	c := make(chan struct{}, 0)
	<-c
}

func handleClick(x int, y int) {
	// check to see if the click is in the confines of the grid
	if x >= gridXOffSet && y >= gridYOffSet && x < gridXOffSet+gridDisplayWidth && y < gridYOffSet+gridDisplayHeight {
		transX, transY := (x-gridXOffSet)/gridSize, (y-gridYOffSet)/gridSize
		g.ClickGrid(transX, transY)
	}
}

func initGame() {
	g = model.NewGame(gameWidth, gameHeight)
	g.State = model.Running
}

func doGameLoop(now float64) {
	if g.State == model.Running {
		g.Update(now)

		d.ClearFrame(0, 0, canvasWidth, canvasHeight)
		drawGameGrid(0.3, 1.0)
	} else if g.State == model.GameOver {
		d.ClearFrame(0, 0, canvasWidth, canvasHeight)
		drawGameGrid(0.2, 0.5)
		d.SetGlobalAlpha(1)
		d.DrawText("Game Over", "45px Comic Sans MS", "black", "center", canvasWidth/2, canvasHeight/2)
	}
}

func drawGameGrid(backGroundAlpha float64, alpha float64) {
	d.SetGlobalAlpha(backGroundAlpha)
	d.DrawRect(gridXOffSet, gridYOffSet, gridDisplayWidth, gridDisplayHeight, "#666666")
	d.SetGlobalAlpha(alpha)

	for x := 0; x < g.Width; x++ {
		for y := 0; y < g.Height; y++ {
			bl := g.GetBlock(x, y)

			drawX, drawY := (x*gridSize)+gridXOffSet, y*gridSize+gridYOffSet

			if bl.Type == model.Red {
				d.DrawImage(view.RedSprite, 0, 0, gridSize, gridSize, drawX, drawY, gridSize, gridSize)
			} else if bl.Type == model.Blue {
				d.DrawImage(view.BlueSprite, 0, 0, gridSize, gridSize, drawX, drawY, gridSize, gridSize)
			} else if bl.Type == model.Green {
				d.DrawImage(view.GreenSprite, 0, 0, gridSize, gridSize, drawX, drawY, gridSize, gridSize)
			} else if bl.Type == model.Purple {
				d.DrawImage(view.PurpleSprite, 0, 0, gridSize, gridSize, drawX, drawY, gridSize, gridSize)
			}
		}
	}
}
