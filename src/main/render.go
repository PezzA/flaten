package main

import (
	"fmt"

	"github.com/pezza/flaten/src/model"
	"github.com/pezza/wasm"
)

type regionKey int

const (
	grid regionKey = iota
	incoming
	progress
)

func (win window) drawHits(s *state) {
	win.Save()
	win.Canvas.Set("lineWidth", 1)
	for _, click := range s.clickHistory {

		newA := click.age / 2000
		if newA < 0 {
			newA = 0
		}

		style := "#FFFFFF"

		switch click.bt {
		case model.Blue:
			style = "#0000FF"
		case model.Red:
			style = "#FF0000"
		case model.Green:
			style = "#00FF00"
		}

		win.SetGlobalAlpha(newA)

		font := fmt.Sprintf("%dpx goldbox", 20+click.cleared)

		text := fmt.Sprintf("+ %d", click.cleared)

		offset := 60 - (click.age / (1000 / 60))
		win.DrawText(text, click.X, click.Y-int(offset), true, font, style, wasm.TextAlignLeft, wasm.TextBaseLineMiddle)
		win.DrawText(text, click.X, click.Y-int(offset), false, font, "#FFFFFF", wasm.TextAlignLeft, wasm.TextBaseLineMiddle)

	}

	win.Restore()
}

func (win window) drawBackGround(s *state) {

	w, h := float64(480), float64(272)
	dw, dh := w*2, h*2

	win.DrawImageF(win.gfx[gfxMeadowSprite], 960, 640, w, h, s.parralax[0]-960, 0, dw, dh)
	win.DrawImageF(win.gfx[gfxMeadowSprite], 960, 640, w, h, s.parralax[0], 0, dw, dh)

	win.DrawImageF(win.gfx[gfxMeadowSprite], 0, 640, w, h, s.parralax[1]-960, 0, dw, dh)
	win.DrawImageF(win.gfx[gfxMeadowSprite], 0, 640, w, h, s.parralax[1], 0, dw, dh)

	win.DrawImageF(win.gfx[gfxMeadowSprite], 768, 0, w, h, s.parralax[2]-960, 0, dw, dh)
	win.DrawImageF(win.gfx[gfxMeadowSprite], 768, 0, w, h, s.parralax[2], 0, dw, dh)

	win.DrawImageF(win.gfx[gfxMeadowSprite], 0, 912, w, h, s.parralax[3]-960, 0, dw, dh)
	win.DrawImageF(win.gfx[gfxMeadowSprite], 0, 912, w, h, s.parralax[3], 0, dw, dh)

	win.DrawImageF(win.gfx[gfxMeadowSprite], 480, 912, w, h, s.parralax[4]-960, 0, dw, dh)
	win.DrawImageF(win.gfx[gfxMeadowSprite], 480, 912, w, h, s.parralax[4], 0, dw, dh)

	win.DrawImageF(win.gfx[gfxMeadowSprite], 480, 640, w, h, s.parralax[5]-960, 0, dw, dh)
	win.DrawImageF(win.gfx[gfxMeadowSprite], 480, 640, w, h, s.parralax[5], 0, dw, dh)

	win.DrawImageF(win.gfx[gfxMeadowSprite], 768, 272, w, h, s.parralax[6]-960, 0, dw, dh)
	win.DrawImageF(win.gfx[gfxMeadowSprite], 768, 272, w, h, s.parralax[6], 0, dw, dh)

}

func drawGameOver(win window, state *state) {
	wasm.SetVolume(win.sfx[sfxMusic], state.fade)
	win.SetGlobalAlpha(1)
	win.DrawText("Game Over", win.canvasWidth/2, win.canvasHeight/2, true, "78px goldbox", "black", wasm.TextAlignCenter, wasm.TextBaseLineMiddle)
}

func drawStartMenu(win window) {
	x, y := win.canvasWidth/2, win.canvasHeight/2
	win.DrawText("Gem POP!", x+3, y+3, true, "78px goldbox", "#ffffff", wasm.TextAlignCenter, wasm.TextBaseLineMiddle)
	win.DrawText("Gem POP!", x, y, true, "78px goldbox", "yellow", wasm.TextAlignCenter, wasm.TextBaseLineMiddle)
	win.DrawText("Click anywhere to start...", x+1, y+100+1, true, "39px goldbox", "#333333", wasm.TextAlignCenter, wasm.TextBaseLineMiddle)
	win.DrawText("Click anywhere to start...", x, y+100, true, "39px goldbox", "white", wasm.TextAlignCenter, wasm.TextBaseLineMiddle)
}

func outlineCanvas(win window) {
	win.SetStrokeStyle("#FFFFFF")
	win.DrawRectInt(0, 0, win.canvasWidth, win.canvasHeight, false)
}

func drawProgressPanel(win window, state *state) {

	font, style := "26px goldbox", "white"
	reg := win.getRegion(progress)

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

	if win.isInRegion(state.mousePos, grid) {
		gReg := win.getRegion(grid)
		transX, transY := (state.mousePos.X-gReg.tl.X)/win.cellSize, (state.mousePos.Y-gReg.tl.Y)/win.cellSize

		win.DrawText(fmt.Sprintf("[%d,%d]", transX, transY), xOffSet+cs+10, yOffset+(5*cs), true, font, style, wasm.TextAlignLeft, wasm.TextBaseLineTop)
	}
}

func drawIncoming(win window, s *state) {
	reg := win.getRegion(incoming)

	win.SetGlobalAlpha(0.6)
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
	reg := win.getRegion(grid)
	//win.canvas.SetGlobalAlpha(backGroundAlpha)
	//win.canvas.SetFillStyle("#222222")
	//win.canvas.DrawRect(float64(win.setting.canvasTL.X), float64(win.setting.canvasTL.Y), float64(win.setting.gridWidth), float64(win.setting.gridHeight), true)
	win.SetGlobalAlpha(0.6)
	win.SetFillStyle("#222222")
	win.DrawRectInt(reg.tl.X, reg.tl.Y, win.gridPixelWidth, win.cellSize*win.gridCellHeight, true)
	win.SetGlobalAlpha(1)

	for x := 0; x < win.gridCellWidth; x++ {
		for y := 0; y < win.gridCellHeight; y++ {
			bl := s.game.GetBlock(x, y)

			drawX, drawY := (x*win.cellSize)+reg.tl.X, y*win.cellSize+reg.tl.Y

			if bl.Type != model.Empty {
				win.DrawImage(win.gfx[assetKey(bl.Type)], 0, 0, win.cellSize, win.cellSize, drawX, drawY, win.cellSize, win.cellSize)
			}
		}
	}
}

func (win window) highlightCells(s *state) {
	reg := win.getRegion(grid)
	transX, transY := win.getGridCells(s)

	points := s.game.GetCellGroup(transX, transY)

	if len(points) < 3 {
		return
	}

	win.SetFillStyle("#FFFFFF")
	win.SetGlobalAlpha(0.3)
	for _, point := range points {
		win.DrawRectInt( (point.X*win.cellSize)+reg.tl.X,(point.Y*win.cellSize)+reg.tl.Y,win.cellSize,win.cellSize,true)
	}
	win.SetGlobalAlpha(1)
}
