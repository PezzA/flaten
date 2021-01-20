package main

import (
	"fmt"

	"github.com/pezza/advent-of-wasm/wasm"
	"github.com/pezza/flaten/src/model"
)

type regionKey int

const (
	grid regionKey = iota
	incoming
	progress
)

func drawGameOver(win window, state *state) {
	wasm.SetVolume(win.sfx[sfxMusic], state.fade)
	win.SetGlobalAlpha(1)
	win.DrawText("Game Over", win.canvasWidth/2, win.canvasHeight/2, true, "78px goldbox", "black", wasm.TextAlignCenter, wasm.TextBaseLineMiddle)
}

func drawStartMenu(win window, state *state) {
	x, y := win.canvasWidth/2, win.canvasHeight/2
	win.DrawText("Gem POP!", x+3, y+3, true, "78px goldbox", "#ffffff", wasm.TextAlignCenter, wasm.TextBaseLineMiddle)
	win.DrawText("Gem POP!", x, y, true, "78px goldbox", "yellow", wasm.TextAlignCenter, wasm.TextBaseLineMiddle)
	win.DrawText("Click anywhere to start...", x+1, y+100+1, true, "39px goldbox", "#333333", wasm.TextAlignCenter, wasm.TextBaseLineMiddle)
	win.DrawText("Click anywhere to start...", x, y+100, true, "39px goldbox", "white", wasm.TextAlignCenter, wasm.TextBaseLineMiddle)
}

func outlineCanvas(win window, state *state) {
	win.SetStrokeStyle("#FFFFFF")
	win.DrawRectInt(0, 0, win.canvasWidth, win.canvasHeight, false)
}

func drawProgressPanel(win window, state *state) {

	font, style := "26px goldbox", "white"
	reg := win.regions[progress]

	xOffSet, yOffset := reg.tl.X+5, reg.tl.Y

	cs := win.cellSize

	res := state.game.GetResults()
	win.DrawImage(win.gfx[assetKey(model.Red)], 0, 0, cs, cs, xOffSet, yOffset+(1*cs), cs, cs)
	win.DrawText(fmt.Sprintf("x %v", res.BlockClears[model.Red]), xOffSet+cs+10, yOffset+(1*cs), true, font, style, wasm.TextAlignLeft, wasm.TextBaseLineTop)

	win.DrawImage(win.gfx[assetKey(model.Blue)], 0, 0, cs, cs, xOffSet, yOffset+(2*cs), cs, cs)
	win.DrawText(fmt.Sprintf("x %v", res.BlockClears[model.Blue]), xOffSet+cs+10, yOffset+(2*cs), true, font, style, wasm.TextAlignLeft, wasm.TextBaseLineTop)

	win.DrawImage(win.gfx[assetKey(model.Green)], 0, 0, cs, cs, xOffSet, yOffset+(3*cs), cs, cs)
	win.DrawText(fmt.Sprintf("x %v", res.BlockClears[model.Green]), xOffSet+cs+10, yOffset+(3*cs), true, font, style, wasm.TextAlignLeft, wasm.TextBaseLineTop)

	win.DrawImage(win.gfx[assetKey(model.Purple)], 0, 0, cs, cs, xOffSet, yOffset+(4*cs), cs, cs)
	win.DrawText(fmt.Sprintf("x %v", res.BlockClears[model.Purple]), xOffSet+cs+10, yOffset+(4*cs), true, font, style, wasm.TextAlignLeft, wasm.TextBaseLineTop)
}

func drawIncoming(win window, s *state) {
	reg := win.regions[incoming]

	win.SetFillStyle("#555555")
	win.DrawRectInt(reg.tl.X, reg.tl.Y, win.gridPixelWidth, win.cellSize, true)

	for i := range s.event {
		if s.event[i] == "incoming" {
			wasm.SetVolume(win.sfx[sfxIncoming], 0.4)
			wasm.PlaySound(win.sfx[sfxIncoming])
		}
	}

	win.SetGlobalAlpha(0.7)

	for index, bl := range s.game.GetIncomingRow() {
		drawX, drawY := (index*win.cellSize)+reg.tl.X, reg.tl.Y

		if bl.Type != model.Empty {
			win.DrawImage(win.gfx[assetKey(bl.Type)], 0, 0, win.cellSize, win.cellSize, drawX, drawY, win.cellSize, win.cellSize)
		}
	}
	win.SetGlobalAlpha(1)
}

func drawGameGrid(win window, s *state) {
	reg := win.regions[grid]
	//win.canvas.SetGlobalAlpha(backGroundAlpha)
	//win.canvas.SetFillStyle("#222222")
	//win.canvas.DrawRect(float64(win.setting.canvasTL.X), float64(win.setting.canvasTL.Y), float64(win.setting.gridWidth), float64(win.setting.gridHeight), true)
	win.SetGlobalAlpha(1)

	win.SetFillStyle("#222222")
	win.DrawRectInt(reg.tl.X, reg.tl.Y, win.gridPixelWidth, win.cellSize*win.gridCellHeight, true)

	for x := 0; x < win.gridCellWidth; x++ {
		for y := 0; y < win.gridCellHeight; y++ {
			fmt.Println(x, y)
			bl := s.game.GetBlock(x, y)

			drawX, drawY := (x*win.cellSize)+reg.tl.X, y*win.cellSize+reg.tl.Y

			if bl.Type != model.Empty {
				win.DrawImage(win.gfx[assetKey(bl.Type)], 0, 0, win.cellSize, win.cellSize, drawX, drawY, win.cellSize, win.cellSize)
			}
		}
	}
}
