package main

import (
	"fmt"

	"github.com/pezza/flaten/src/model"
)

// state holds everything that can change as part of the game
type state struct {
	running      bool
	game         model.Game
	event        []string
	fade         float64
	mousePos     model.Point
	parralax     [7]float64
	delta        float64
	current      float64
	clickHistory []click
}

type click struct {
	model.Point
	bt      model.BlockType
	age     float64
	cleared int
}

func (s *state) updateParralax() {

	for i := range s.parralax {
		s.parralax[i] -= s.delta * float64(i) / 50

		if s.parralax[i] <= 0 {
			s.parralax[i] = 960
		}
	}

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
	s.delta = now - s.current
	s.current = now

	s.updateParralax()

	if !s.running {
		return
	}

	if s.game.State == model.GameOver {
		s.fade -= 0.005
		if s.fade < 0.005 {
			s.running = false
		}

		return
	}

	s.event = s.game.Update(now)

	clicks := make([]click, 0)

	for i := range s.clickHistory {
		if s.clickHistory[i].age >= 0 {
			s.clickHistory[i].age -= s.delta * 2
			clicks = append(clicks, s.clickHistory[i])
		}
	}

	s.clickHistory = clicks
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
		if win.isInRegion(s.mousePos, grid) {
			transX, transY := win.getGridCells(s)

			result := s.game.ClickGrid(transX, transY)

			if model.IsBomb(result.ClickType) {
				win.playSound(sfxBomb, 1)
			} else if model.IsBlock(result.ClickType) {

				s.clickHistory = append(s.clickHistory, click{
					cleared: result.BlocksCleared,
					bt:      result.ClickType,
					age:     2000,
					Point:   model.Point{X: s.mousePos.X, Y: s.mousePos.Y - 20},
				})

				win.playSound(sfxClick, 1)
			} else if model.IsClearTile(result.ClickType) {

				s.clickHistory = append(s.clickHistory, click{
					cleared: result.BlocksCleared,
					bt:      result.ClickType,
					age:     2000,
					Point:   model.Point{X: s.mousePos.X, Y: s.mousePos.Y - 20},
				})

				win.playSound(sfxClear, 1)
			}
		}

		if win.isInRegion(s.mousePos, incoming) {
			if s.game.ClickIncoming() {
				win.playSound(sfxIncoming, 0.4)
			}
		}
	}
}
