package game

import (
	"math/rand"
	"time"
)

// State holds the current run state of the game
type State int

const (
	New State = iota
	Running
	Finished
)

// Game holds all the game state
type Game struct {
	blocks [][]Block
	Width  int
	Height int
	State  State
	Timer  float64
	Dist   int
}

// NewGame returns a new game
func NewGame(width int, height int) Game {
	rand.Seed(time.Now().UnixNano())
	g := Game{
		blocks: make([][]Block, height),
		Width:  width,
		Height: height,
		State:  New,
		Timer:  0,
		Dist:   0,
	}

	for index := range g.blocks {
		g.blocks[index] = make([]Block, width)
	}

	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			bt := rand.Intn(playTileMax-playTileMin+1) + playTileMin

			maxFrame := 0
			if bt == Red || bt == Green {
				maxFrame = 4
			}

			if bt == Blue || bt == Purple {
				maxFrame = 8
			}

			//currentframe := rand.Intn(maxFrame)
			g.blocks[y][x] = Block{
				Point:        Point{X: x, Y: y},
				Type:         bt,
				Moving:       false,
				Drop:         0,
				MaxFrame:     maxFrame,
				CurrentFrame: 0,
				FrameTimer:   0,
			}
		}
	}

	return g
}

// GetBlock returns a specified block from the grid
func (g *Game) GetBlock(x int, y int) Block {
	return g.blocks[y][x]
}

// ClickGrid handle a click on a cell
func (g *Game) ClickGrid(x int, y int) bool {
	if g.State != Running || x >= g.Width || y >= g.Height || x < 0 || y < 0 {
		return false
	}

	if g.blocks[y][x].Type == Empty {
		return false
	}

	blockGroup := g.getBlockGroup(g.blocks[y][x].Type, []Point{Point{X: x, Y: y}})

	if len(blockGroup) < 3 {
		return false
	}

	colList := make(map[int]bool, 0)
	for _, b := range blockGroup {
		g.blocks[b.Y][b.X].Type = Empty
		colList[b.X] = true
	}

	for k := range colList {
		g.shuntCol(k)
	}

	// squash empty cols
	g.shiftLeft()
	g.shiftRight()

	return true
}

func (g *Game) gameOver() bool {
	for x := 0; x < g.Width; x++ {
		for y := 0; y < g.Height; y++ {
			if g.blocks[y][x].Type == Empty || g.blocks[y][x].Type == Temp {
				continue
			}
			if g.blocks[y][x].Moving || g.blocks[y][x].Type == Empty || g.blocks[y][x].Type == Temp {
				return false
			}
			blockGroup := g.getBlockGroup(g.blocks[y][x].Type, []Point{Point{X: x, Y: y}})
			if len(blockGroup) > 2 {
				return false
			}
		}
	}
	return true
}

func (g *Game) isColEmpty(x int) bool {
	for index := range g.blocks {
		if g.blocks[index][x].Type != Empty {
			return false
		}
	}
	return true
}

func (g *Game) shiftLeft() {
	shunt := 0
	for i := (g.Width / 2) - 1; i >= 0; i-- {
		if g.isColEmpty(i) {
			shunt++
		} else {
			if shunt > 0 {
				for index := range g.blocks {
					g.blocks[index][i+shunt].Type = g.blocks[index][i].Type
					g.blocks[index][i+shunt].Moving = g.blocks[index][i].Moving
					g.blocks[index][i+shunt].Drop = g.blocks[index][i].Drop
					g.blocks[index][i+shunt].Dist = g.blocks[index][i].Dist
					g.blocks[index][i].Type = Empty
				}
			}
		}
	}
}

func (g *Game) shiftRight() {
	shunt := 0
	for i := (g.Width / 2); i <= (g.Width - 1); i++ {
		if g.isColEmpty(i) {
			shunt++
		} else {
			if shunt > 0 {
				for index := range g.blocks {
					g.blocks[index][i-shunt].Type = g.blocks[index][i].Type
					g.blocks[index][i-shunt].Moving = g.blocks[index][i].Moving
					g.blocks[index][i-shunt].Drop = g.blocks[index][i].Drop
					g.blocks[index][i-shunt].Dist = g.blocks[index][i].Dist
					g.blocks[index][i].Type = Empty
				}
			}
		}
	}
}

func (g *Game) shuntCol(x int) {
	shunt := 0
	for i := g.Height - 1; i >= 0; i-- {
		if g.blocks[i][x].Type == Empty {
			shunt++
		} else {
			if shunt > 0 {
				g.blocks[i+shunt][x].Type = g.blocks[i][x].Type
				g.blocks[i+shunt][x].Moving = true
				g.blocks[i+shunt][x].Dist = shunt
				g.blocks[i+shunt][x].TotalDrop = float64(shunt) * dropTime
				g.blocks[i][x].Type = Temp
			}
		}
	}

	for i := g.Height - 1; i >= 0; i-- {
		if g.blocks[i][x].Type == Temp {
			g.blocks[i][x].Type = Empty
		}
	}
}

func pointExists(p Point, pl []Point) bool {
	for index := range pl {
		if p.X == pl[index].X && p.Y == pl[index].Y {
			return true
		}
	}
	return false
}

func (g *Game) getBlockGroup(bt BlockType, points []Point) []Point {
	lp := points[len(points)-1]

	// look left
	if lp.X > 0 {
		newP := Point{X: lp.X - 1, Y: lp.Y}
		if g.blocks[newP.Y][newP.X].Type == bt && !pointExists(newP, points) && !g.blocks[newP.Y][newP.X].Moving {
			points = g.getBlockGroup(bt, append(points, newP))
		}
	}

	// look up
	if lp.Y > 0 {
		newP := Point{X: lp.X, Y: lp.Y - 1}
		if g.blocks[newP.Y][newP.X].Type == bt && !pointExists(newP, points) && !g.blocks[newP.Y][newP.X].Moving {
			points = g.getBlockGroup(bt, append(points, newP))
		}
	}

	// look right
	if lp.X < (g.Width - 1) {
		newP := Point{X: lp.X + 1, Y: lp.Y}
		if g.blocks[newP.Y][newP.X].Type == bt && !pointExists(newP, points) && !g.blocks[newP.Y][newP.X].Moving {
			points = g.getBlockGroup(bt, append(points, newP))
		}
	}

	// look down
	if lp.Y < (g.Height - 1) {
		newP := Point{X: lp.X, Y: lp.Y + 1}
		if g.blocks[newP.Y][newP.X].Type == bt && !pointExists(newP, points) && !g.blocks[newP.Y][newP.X].Moving {
			points = g.getBlockGroup(bt, append(points, newP))
		}
	}

	return points
}

// Update modifies the game model based on the delta
func (g *Game) Update(now float64) {
	// First off, lets see if the game has finished
	if g.gameOver() {
		g.State = Finished
	}

	if g.State == Running {
		delta := now - g.Timer
		g.Timer = now
		g.updateBlocks(delta)
	}
}

func (g *Game) updateBlocks(delta float64) {
	for x := 0; x < g.Width; x++ {
		for y := 0; y < g.Height; y++ {
			g.blocks[y][x].update(delta)
		}
	}
}

// Start gets the game going!
func (g *Game) Start() {
	g.State = Running
}
