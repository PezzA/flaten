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

var startMenu = true
var gameOver = false

func isPointInBound(x, y, minX, minY, maxX, maxY int) bool {
	return x >= minX && y >= minY && x < maxX && y < maxY
}

func init() {
	d = view.NewJsDoc(handleClick, doGameLoop)
	d.SetCanvasSize(canvasWidth, canvasHeight)
	d.StartAnimLoop()
}

func main() {
	//run forever
	c := make(chan struct{}, 0)
	<-c
}

func handleClick(x int, y int) {

	if startMenu || gameOver {
		initGame()
	} else {
		if isPointInBound(x, y, gridXOffSet, gridYOffSet, gridXOffSet+gridDisplayWidth, gridYOffSet+gridDisplayHeight) {
			transX, transY := (x-gridXOffSet)/gridSize, (y-gridYOffSet)/gridSize
			result := g.ClickGrid(transX, transY)

			if model.IsBomb(result.ClickType) {
				d.PlaySound(view.SfxBomb)
			} else if model.IsBlock(result.ClickType) {
				d.PlaySound(view.SfxClick)
			} else if model.IsClearTile(result.ClickType) {
				d.PlaySound(view.SfxClear)
			}
		}

		if isPointInBound(x, y, incomingXOffSet, incomingYOffSet, incomingXOffSet+gridDisplayWidth, incomingYOffSet+gridSize) {
			if g.ClickIncoming() {
				d.SetVolume(view.SfxIncoming, 0.4)
				d.PlaySound(view.SfxIncoming)
			}
		}
	}
}

func initGame() {
	fadeOut = 1
	startMenu = false
	gameOver = false
	g = model.NewGame(gameWidth, gameHeight)
	g.State = model.Running
	d.SetVolume(view.SfxMusic, 1)
	d.PlaySound(view.SfxMusic)
}

var fadeOut float64 = 1

func drawStartMenu() {
	x, y := canvasWidth/2, canvasHeight/2

	d.DrawText("Gem POP!", "78px goldbox", "#333333", "center", "middle", x+3, y+3)
	d.DrawText("Gem POP!", "78px goldbox", "yellow", "center", "middle", x, y)

	d.DrawText("Click anywhere to start...", "39px goldbox", "#333333", "center", "middle", x+1, y+100+1)
	d.DrawText("Click anywhere to start...", "39px goldbox", "white", "center", "middle", x, y+100)
}

func doGameLoop(now float64) {

	if startMenu {
		d.ClearFrame(0, 0, canvasWidth, canvasHeight)
		drawStartMenu()
		return
	}

	if g.State == model.Running {
		events := g.Update(now)

		for i := range events {
			if events[i] == "incoming" {
				d.SetVolume(view.SfxIncoming, 0.4)
				d.PlaySound(view.SfxIncoming)
			}
		}

		d.ClearFrame(0, 0, canvasWidth, canvasHeight)
		drawGameGrid(0.3, 1.0)
		drawIncoming()
		drawProgressPanel(g.GetResults())
	} else if g.State == model.GameOver {
		d.ClearFrame(0, 0, canvasWidth, canvasHeight)
		drawGameGrid(0.3, fadeOut)
		d.SetVolume(view.SfxMusic, fadeOut)
		d.SetGlobalAlpha(1)
		d.DrawText("Game Over", "78px goldbox", "black", "center", "middle", canvasWidth/2, canvasHeight/2)
		fadeOut -= 0.005
		if fadeOut < 0.005 {

			fadeOut = 0
			gameOver = true
		}
	}
}

func drawProgressPanel(res model.Results) {
	xOffSet, yOffset := gridXOffSet+gridDisplayWidth+(gridSize/2), (gridSize / 2)
	d.DrawImage(view.RedSprite, 0, 0, gridSize, gridSize, xOffSet, yOffset+(1*gridSize), gridSize, gridSize)
	d.DrawText(fmt.Sprintf("x %v", res.BlockClears[model.Red]), "26px goldbox", "white", "left", "top", xOffSet+gridSize+10, yOffset+(1*gridSize)+24)
	d.DrawImage(view.BlueSprite, 0, 0, gridSize, gridSize, xOffSet, yOffset+(2*gridSize), gridSize, gridSize)
	d.DrawText(fmt.Sprintf("x %v", res.BlockClears[model.Blue]), "26px goldbox", "white", "left", "top", xOffSet+gridSize+10, yOffset+(2*gridSize)+24)
	d.DrawImage(view.GreenSprite, 0, 0, gridSize, gridSize, xOffSet, yOffset+(3*gridSize), gridSize, gridSize)
	d.DrawText(fmt.Sprintf("x %v", res.BlockClears[model.Green]), "26px goldbox", "white", "left", "top", xOffSet+gridSize+10, yOffset+(3*gridSize)+24)
	d.DrawImage(view.PurpleSprite, 0, 0, gridSize, gridSize, xOffSet, yOffset+(4*gridSize), gridSize, gridSize)
	d.DrawText(fmt.Sprintf("x %v", res.BlockClears[model.Purple]), "26px goldbox", "white", "left", "top", xOffSet+gridSize+10, yOffset+(4*gridSize)+24)
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
}

func drawGameGrid(backGroundAlpha float64, alpha float64) {
	if alpha < backGroundAlpha {
		backGroundAlpha = alpha
	}
	d.SetGlobalAlpha(backGroundAlpha)
	d.DrawRect(gridXOffSet, gridYOffSet, gridDisplayWidth, gridDisplayHeight, "#222222")
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
