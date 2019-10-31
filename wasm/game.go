package main

import (
	"fmt"

	"github.com/pezza/wasm/game"
)

var g game.Game

var gameWidth, gameHeight, displayWidth, displayHeight = 12, 20, 360, 600
var gridSize = 30
var gridXOffSet = (displayWidth - (gameWidth * gridSize)) / 2

func updateGame(now float64) {
	g.Update(now)

	setElementInnerHTML("gs", fmt.Sprintf("%v", g.State))
}

func handleClick(x int, y int) {
	g.ClickGrid(x/gridSize, y/gridSize)
}

func initGame() {
	g = game.NewGame(gameWidth, gameHeight)
	g.Start()
}

func drawGameGrid() {
	for x := 0; x < g.Width; x++ {
		for y := 0; y < g.Height; y++ {

			bl := g.GetBlock(x, y)

			drawX, drawY := (x*gridSize)+gridXOffSet, y*gridSize

			if g.State != game.Running {
				ctx.Set("globalAlpha", 0.2)
			} else {
				if bl.Type == game.Empty || bl.Moving {
					ctx.Set("globalAlpha", 0.8)
					offset := bl.Dist * 30

					drawY -= int(offset - bl.Drop/float64(33))
				} else {
					ctx.Set("globalAlpha", 1)
				}
			}

			if bl.Type == game.Red {
				ctx.Call("drawImage", red, drawX, drawY)
			} else if bl.Type == game.Blue {
				ctx.Call("drawImage", blue, drawX, drawY)
			} else if bl.Type == game.Green {
				ctx.Call("drawImage", green, drawX, drawY)
			} else if bl.Type == game.White {
				ctx.Call("drawImage", purple, drawX, drawY)
			}
		}
	}
}
