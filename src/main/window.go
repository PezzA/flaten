package main

import (
	"github.com/pezza/advent-of-wasm/wasm"
	"github.com/pezza/flaten/src/Generated"
	"github.com/pezza/flaten/src/model"
)

// window holds the host runtime environment (wasm JS wrappers in this case) and all the assets used.
type window struct {
	wasm.JsDoc
	*wasm.JsCanvas
	windowSettings
	gfx assetMap
	sfx assetMap
}

// region defines the top left and bottom right of a rectangle, used to identity areas of the screen
type region struct {
	tl model.Point
	br model.Point
}

// windowSettings contains precalculated layout values
type windowSettings struct {
	// cellSize holds the pixel width of each cell
	cellSize int

	// gridCellWidth holds the width of the grid in cells
	gridCellWidth int

	// gridCellHeight holds the height of the grid in cells
	gridCellHeight int

	// gridPixelWidth holds the width of the grid in pixels
	gridPixelWidth int

	// gridPixelHeight holds the height of the grid in pixels
	gridPixelHeight int

	// canvasWidth is the pixel width of the canvas
	canvasWidth int

	// canvasHeight is the pixel height of the canvas
	canvasHeight int

	regions map[regionKey]region
}

func (win window) playSound(asset assetKey, vol float64) {
	wasm.SetVolume(win.sfx[asset], vol)
	wasm.PlaySound(win.sfx[asset])
}

// draw takes the current state and renders to the window
func (win window) draw(s *state) {
	win.Clear()
	outlineCanvas(win, s)
	if !s.running {
		drawStartMenu(win, s)
		return
	}
	switch s.game.State {
	case model.Running:
		drawGameGrid(win, s)
		drawIncoming(win, s)
		drawProgressPanel(win, s)
	case model.GameOver:
		drawGameGrid(win, s)
		drawGameOver(win, s)
	default:
	}

}

func (ws windowSettings) isClick(pos model.Point, reg regionKey) bool {
	return pos.X >= ws.regions[reg].tl.X &&
		pos.Y >= ws.regions[reg].tl.Y &&
		pos.X < ws.regions[reg].br.X &&
		pos.Y < ws.regions[reg].br.Y
}

func getWindowSettings(cellSizePix int, cellGridWidth int, cellGridHeight int) windowSettings {
	gridDisplayWidth, gridDisplayHeight := cellGridWidth*cellSizePix, cellGridHeight*cellSizePix

	canvasWidth, canvasHeight := gridDisplayWidth*2, gridDisplayHeight+(2*cellSizePix)

	var regions = map[regionKey]region{
		grid: {
			tl: model.Point{X: gridDisplayWidth / 2, Y: cellSizePix / 2},
			br: model.Point{X: (gridDisplayWidth / 2) + gridDisplayWidth, Y: (cellSizePix / 2) + gridDisplayHeight},
		},
		incoming: {
			tl: model.Point{X: gridDisplayWidth / 2, Y: gridDisplayHeight + cellSizePix},
			br: model.Point{X: gridDisplayHeight + gridDisplayWidth, Y: (gridDisplayHeight + cellSizePix) + cellSizePix},
		},
		progress: {
			tl: model.Point{X: gridDisplayWidth/2 + gridDisplayWidth, Y: 50},
		},
	}
	return windowSettings{
		cellSize:        cellSizePix,
		gridCellWidth:   cellGridWidth,
		gridCellHeight:  cellGridHeight,
		gridPixelWidth:  gridDisplayWidth,
		gridPixelHeight: gridDisplayHeight,
		canvasWidth:     canvasWidth,
		canvasHeight:    canvasHeight,
		regions:         regions,
	}
}

func getWindow(cellSizePix int, cellGridWidth int, cellGridHeight int) window {
	settings := getWindowSettings(cellSizePix, cellGridWidth, cellGridHeight)
	doc := wasm.NewJsDoc()
	canvas := doc.GetOrCreateCanvas("gameCanvas", settings.canvasWidth, settings.canvasHeight, true, false)

	gfxList := make(assetMap, 0)
	gfxList[gfxBlueSquare] = doc.GetCanvasImage(Generated.GfxResource_blueSquare.Data, Generated.GfxResource_blueSquare.Width, Generated.GfxResource_blueSquare.Height)
	gfxList[gfxGreenSquare] = doc.GetCanvasImage(Generated.GfxResource_greenSquare.Data, Generated.GfxResource_greenSquare.Width, Generated.GfxResource_greenSquare.Height)
	gfxList[gfxRedSquare] = doc.GetCanvasImage(Generated.GfxResource_redSquare.Data, Generated.GfxResource_redSquare.Width, Generated.GfxResource_redSquare.Height)
	gfxList[gfxPurpleSquare] = doc.GetCanvasImage(Generated.GfxResource_purpleSquare.Data, Generated.GfxResource_purpleSquare.Width, Generated.GfxResource_purpleSquare.Height)
	gfxList[gfxBlueClear] = doc.GetCanvasImage(Generated.GfxResource_blueClear.Data, Generated.GfxResource_blueClear.Width, Generated.GfxResource_blueClear.Height)
	gfxList[gfxGreenClear] = doc.GetCanvasImage(Generated.GfxResource_greenClear.Data, Generated.GfxResource_greenClear.Width, Generated.GfxResource_greenClear.Height)
	gfxList[gfxRedClear] = doc.GetCanvasImage(Generated.GfxResource_redClear.Data, Generated.GfxResource_redClear.Width, Generated.GfxResource_redClear.Height)
	gfxList[gfxPurpleClear] = doc.GetCanvasImage(Generated.GfxResource_purpleClear.Data, Generated.GfxResource_purpleClear.Width, Generated.GfxResource_purpleClear.Height)
	gfxList[gfxSlideLeft] = doc.GetCanvasImage(Generated.GfxResource_slideLeft.Data, Generated.GfxResource_slideLeft.Width, Generated.GfxResource_slideLeft.Height)
	gfxList[gfxSlideUp] = doc.GetCanvasImage(Generated.GfxResource_slideUp.Data, Generated.GfxResource_slideUp.Width, Generated.GfxResource_slideUp.Height)
	gfxList[gfxBomb] = doc.GetCanvasImage(Generated.GfxResource_bomb.Data, Generated.GfxResource_bomb.Width, Generated.GfxResource_bomb.Height)

	sfxList := make(assetMap, 0)
	sfxList[sfxMusic] = doc.GetElementByID("music")
	sfxList[sfxClick] = doc.GetElementByID("click")
	sfxList[sfxClear] = doc.GetElementByID("clear")
	sfxList[sfxBomb] = doc.GetElementByID("bombsfx")
	sfxList[sfxIncoming] = doc.GetElementByID("incoming")
	sfxList[sfxIncomingPip] = doc.GetElementByID("incomingpip")

	return window{
		windowSettings: settings,
		JsDoc:          doc,
		JsCanvas:       canvas,
		gfx:            gfxList,
		sfx:            sfxList,
	}
}
