package main

import (
	"fmt"

	"github.com/pezza/flaten/src/model"
)

// state holds everything that can change as part of the game
type state struct {
	running  bool
	game     model.Game
	event    []string
	fade     float64
	mousePos model.Point
}

// getNewState returns an empty state
func getNewState() *state {
	return &state{
		fade:  1,
		event: []string{},
	}
}

// startNew
func (s *state) startNew(win window) {
	fmt.Println("init")
	s.running = true
	s.fade = 1
	s.game = model.NewGame(win.gridCellWidth, win.gridCellHeight)
	win.playSound(sfxMusic, 1)
}

func (s *state) update(now float64) {
	if !s.running {
		return
	}

	s.event = s.game.Update(now)

	if s.game.State == model.GameOver {
		s.fade -= 0.005
		if s.fade < 0.005 {
			s.running = false
		}
	}
}

func (s *state) handleMouseMove(win window, newPos model.Point) {
	if newPos.X < 0 {
		newPos.X = 0
	}
	if newPos.Y < 0 {
		newPos.Y = 0
	}

	if newPos.X > win.canvasWidth {
		newPos.X = win.canvasWidth
	}

	if newPos.Y > win.canvasHeight {
		newPos.Y = win.canvasHeight
	}

	s.mousePos = newPos
}

func (s *state) handleClick(win window) {
	if !s.running {
		s.startNew(win)
	} else {
		if win.isClick(s.mousePos, grid) {
			reg := win.regions[grid]

			transX, transY := (s.mousePos.X-reg.tl.X)/win.cellSize, (s.mousePos.Y-reg.tl.Y)/win.cellSize

			result := s.game.ClickGrid(transX, transY)

			if model.IsBomb(result.ClickType) {
				win.playSound(sfxBomb, 1)
			} else if model.IsBlock(result.ClickType) {
				win.playSound(sfxClick, 1)
			} else if model.IsClearTile(result.ClickType) {
				win.playSound(sfxClear, 1)
			}
		}

		if win.isClick(s.mousePos, incoming) {
			if s.game.ClickIncoming() {
				win.playSound(sfxIncoming, 0.4)
			}
		}
	}
}
