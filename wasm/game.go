package main

import (
	"fmt"

	"github.com/pezza/flaten/game"
)

var g game.Game

var gridSize = 32
var gameWidth, gameHeight = 12, 15
var displayWidth, displayHeight = gameWidth * gridSize, gameHeight * gridSize

var gridXOffSet = (displayWidth - (gameWidth * gridSize)) / 2

func updateGame(now float64) {
	g.Update(now)

	page.setElementInnerHTML("gs", fmt.Sprintf("%v", g.State))
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
				page.ctx.Set("globalAlpha", 0.2)
			} else {
				if bl.Type == game.Empty || bl.Moving {
					page.ctx.Set("globalAlpha", 0.8)

					// max pixels distance it can be
					offset := bl.Dist * gridSize

					drawY -= int(float64(offset) * (1 - (bl.Drop / bl.TotalDrop)))
				} else {
					page.ctx.Set("globalAlpha", 1)
				}
			}

			if bl.Type == game.Red {
				page.ctx.Call("drawImage", redSprite, bl.CurrentFrame*gridSize, 0, gridSize, gridSize, drawX, drawY, gridSize, gridSize)
			} else if bl.Type == game.Blue {
				page.ctx.Call("drawImage", blueSprite, bl.CurrentFrame*gridSize, 0, gridSize, gridSize, drawX, drawY, gridSize, gridSize)
			} else if bl.Type == game.Green {
				page.ctx.Call("drawImage", greenSprite, bl.CurrentFrame*gridSize, 0, gridSize, gridSize, drawX, drawY, gridSize, gridSize)
			} else if bl.Type == game.Purple {
				page.ctx.Call("drawImage", purpleSprite, bl.CurrentFrame*gridSize, 0, gridSize, gridSize, drawX, drawY, gridSize, gridSize)
			}
		}
	}
}
