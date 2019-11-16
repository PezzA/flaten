package main

import (
	"fmt"
	"syscall/js"

	"github.com/pezza/flaten/model"
	"github.com/pezza/flaten/view"
)

var d view.JsDoc
var g model.Game

var gridSize = 32
var gameWidth, gameHeight = 12, 10

var gridDisplayWidth, gridDisplayHeight = gameWidth * gridSize, gameHeight * gridSize
var canvasWidth, canvasHeight = gridDisplayWidth * 2, (gridDisplayHeight + (2 * gridSize))

// top left of grid
var gridXOffSet = gridDisplayWidth / 2
var gridYOffSet = gridSize / 2

// top left of incoming row
var incomingXOffSet = gridXOffSet
var incomingYOffSet = gridDisplayHeight + gridSize

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

	if x >= incomingXOffSet && y >= incomingYOffSet && x < incomingXOffSet+gridDisplayWidth && y < incomingYOffSet+gridSize {
		g.ClickIncoming()
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
		drawIncoming()
		drawProgressPanel(g.GetResults())
	} else if g.State == model.GameOver {
		d.ClearFrame(0, 0, canvasWidth, canvasHeight)
		drawGameGrid(0.2, 0.5)
		d.SetGlobalAlpha(1)
		d.DrawText("Game Over", "45px Comic Sans MS", "black", "center", "middle", canvasWidth/2, canvasHeight/2)
	}
}

func drawProgressPanel(res model.Results) {
	xOffSet, yOffset := gridXOffSet+gridDisplayWidth+(gridSize/2), (gridSize / 2)
	d.DrawImage(view.RedSprite, 0, 0, gridSize, gridSize, xOffSet, yOffset+(1*gridSize), gridSize, gridSize)
	d.DrawText(fmt.Sprintf("x %v", res.BlockClears[model.Red]), "24px Consolas", "white", "left", "top", xOffSet+gridSize+10, yOffset+(1*gridSize)+24)
	d.DrawImage(view.BlueSprite, 0, 0, gridSize, gridSize, xOffSet, yOffset+(2*gridSize), gridSize, gridSize)
	d.DrawText(fmt.Sprintf("x %v", res.BlockClears[model.Blue]), "24px Consolas", "white", "left", "top", xOffSet+gridSize+10, yOffset+(2*gridSize)+24)
	d.DrawImage(view.GreenSprite, 0, 0, gridSize, gridSize, xOffSet, yOffset+(3*gridSize), gridSize, gridSize)
	d.DrawText(fmt.Sprintf("x %v", res.BlockClears[model.Green]), "24px Consolas", "white", "left", "top", xOffSet+gridSize+10, yOffset+(3*gridSize)+24)
	d.DrawImage(view.PurpleSprite, 0, 0, gridSize, gridSize, xOffSet, yOffset+(4*gridSize), gridSize, gridSize)
	d.DrawText(fmt.Sprintf("x %v", res.BlockClears[model.Purple]), "24px Consolas", "white", "left", "top", xOffSet+gridSize+10, yOffset+(4*gridSize)+24)
}

func drawIncoming() {
	d.SetGlobalAlpha(0.7)
	for index, bl := range g.GetIncomingRow() {
		drawX, drawY := (index*gridSize)+incomingXOffSet, incomingYOffSet

		if bl.Type != model.Empty {
			d.DrawImage(getSprite(bl.Type), 0, 0, gridSize, gridSize, drawX, drawY, gridSize, gridSize)
		}
	}
	d.SetGlobalAlpha(1)
	for i := 0; i < gameWidth; i++ {
		drawX, drawY := (i*gridSize)+incomingXOffSet, incomingYOffSet
		d.StrokeRect(drawX, drawY, gridSize, gridSize, "#dddddd")
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

			if bl.Type != model.Empty {
				d.DrawImage(getSprite(bl.Type), 0, 0, gridSize, gridSize, drawX, drawY, gridSize, gridSize)
			}
		}
	}
}

func getSprite(t model.BlockType) js.Value {
	switch t {
	case model.Red:
		return view.RedSprite
	case model.Blue:
		return view.BlueSprite
	case model.Green:
		return view.GreenSprite
	case model.Purple:
		return view.PurpleSprite
	case model.Bomb:
		return view.Bomb
	case model.RedClear:
		return view.ClearRed
	case model.BlueClear:
		return view.ClearBlue
	case model.GreenClear:
		return view.ClearGreen
	case model.PurpleClear:
		return view.ClearPurple
	case model.SlideLeft:
		return view.SlideLeft
	case model.SlideUp:
		return view.SlideUp
	}

	return view.Sushi
}
