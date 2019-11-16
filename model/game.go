package model

import (
	"math/rand"
	"time"
)

var startingTick = 500

// GameState returns an int that represents the game's current state
type GameState int

const (
	// New means the game is setup, but has not been started
	New GameState = iota
	// Running indicates that the game is currently running
	Running
	// GameOver indcates that the game is finished
	GameOver
)

// NewGame returns a new game
func NewGame(width int, height int) Game {
	rand.Seed(time.Now().UnixNano())
	g := Game{
		blocks:      make([][]Block, height),
		Width:       width,
		Height:      height,
		State:       New,
		Timer:       0,
		Tick:        startingTick,
		CurrentTick: 0,
		newRow:      make([]Block, width),
		scores:      Results{make(map[BlockType]int), 0, 0, 0, 0},
	}

	for index := range g.blocks {
		g.blocks[index] = make([]Block, width)
	}

	return g
}

// GetIncomingRow returns the list of incoming blocks
func (g *Game) GetIncomingRow() []Block {
	return g.newRow
}

// GetResults gets the current Results from the game
func (g *Game) GetResults() Results {
	return g.scores
}

// GetBlock returns a specified block from the grid
func (g *Game) GetBlock(x int, y int) Block {
	return g.blocks[y][x]
}

// ClickIncoming handles a click on the incoming row
func (g *Game) ClickIncoming() bool {
	// first check to see if there is a tile on the top row, if so, dont click
	for x := 0; x < g.Width; x++ {
		if g.blocks[0][x].Type != Empty {
			return false
		}
	}

	for i := 0; i < g.Width; i++ {
		if g.newRow[i].Type == Empty {
			g.newRow[i] = Block{getRandomPlayTile()}
		}
	}

	g.insertIncomingRow()

	return true
}

func (g *Game) clearCol(x int) []Point {
	points := []Point{}
	for y := 0; y < g.Height; y++ {
		if g.blocks[y][x].Type != Empty {
			points = append(points, Point{x, y})
		}
	}
	return points
}

func (g *Game) clearRow(y int) []Point {
	points := []Point{}
	for x := 0; x < g.Width; x++ {
		if g.blocks[y][x].Type != Empty {
			points = append(points, Point{x, y})
		}
	}
	return points
}

// ClickGrid handle a click on a cell
func (g *Game) ClickGrid(x int, y int) ClickResult {
	if g.State != Running || x >= g.Width || y >= g.Height || x < 0 || y < 0 {
		return ClickResult{false, 0, 0}
	}

	clickedType := g.blocks[y][x].Type

	var blockGroup []Point

	// perform the block action
	if clickedType == Empty {
		return ClickResult{false, 0, 0}
	} else if clickedType == Bomb {
		blockGroup = g.clearBomb(Point{x, y})
	} else if clickedType == RedClear {
		blockGroup = g.clearType([]BlockType{Red, RedClear}, Point{x, y})
	} else if clickedType == GreenClear {
		blockGroup = g.clearType([]BlockType{Green, GreenClear}, Point{x, y})
	} else if clickedType == PurpleClear {
		blockGroup = g.clearType([]BlockType{Purple, PurpleClear}, Point{x, y})
	} else if clickedType == BlueClear {
		blockGroup = g.clearType([]BlockType{Blue, BlueClear}, Point{x, y})
	} else if clickedType == SlideLeft {
		blockGroup = g.clearRow(y)
	} else if clickedType == SlideUp {
		blockGroup = g.clearCol(x)
	} else {
		blockGroup = g.getBlockGroup(clickedType, []Point{Point{X: x, Y: y}})

		if len(blockGroup) < 3 {
			return ClickResult{false, 0, 0}
		}
	}

	blocksCleared := len(blockGroup)

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

	// score based on the blocks cleard
	g.scores.OverallScore += blocksCleared * blocksCleared * 10
	g.scores.BlockClears[clickedType] += blocksCleared

	return ClickResult{true, len(blockGroup), 0}
}

// Update modifies the game model based on the delta
func (g *Game) Update(now float64) {
	if g.State == GameOver {
		return
	}

	delta := int(now - g.Timer)
	g.Timer = now

	g.CurrentTick += delta

	if g.CurrentTick > g.Tick {
		if g.addIncomingBlock() {
			g.insertIncomingRow()
		}

		g.CurrentTick = 0
	}

	return
}

func (g *Game) addIncomingBlock() bool {
	for i := 0; i < g.Width; i++ {
		if g.newRow[i].Type == Empty {
			g.newRow[i] = Block{getRandomPlayTile()}
			return i+1 == g.Width
		}
	}
	return false
}

func (g *Game) insertIncomingRow() {
	g.shuntGrid()
	g.newRow = make([]Block, g.Width)
	g.scores.Rows++
	g.Tick = startingTick - (g.scores.Rows * 5)
}

func (g *Game) fillGrid() {
	for x := 0; x < g.Width; x++ {
		for y := 0; y < g.Height; y++ {
			g.blocks[y][x] = newBlock()
		}
	}
}

func (g *Game) noMatchGameOver() bool {
	for x := 0; x < g.Width; x++ {
		for y := 0; y < g.Height; y++ {
			if g.blocks[y][x].Type == Empty {
				continue
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

func isPointInBound(x, y, minX, minY, maxX, maxY int) bool {
	return x >= minX && y >= minY && x < maxX && y < maxY
}

func (g *Game) clearBomb(p Point) []Point {
	points := []Point{p}

	testPoints := []Point{
		Point{p.X - 1, p.Y - 1},
		Point{p.X, p.Y - 1},
		Point{p.X + 1, p.Y - 1},
		Point{p.X + 1, p.Y},
		Point{p.X + 1, p.Y + 1},
		Point{p.X, p.Y + 1},
		Point{p.X - 1, p.Y + 1},
		Point{p.X - 1, p.Y},
	}

	for _, tp := range testPoints {
		if isPointInBound(tp.X, tp.Y, 0, 0, g.Width, g.Height) {
			points = append(points, tp)
		}
	}

	return points
}

func (g *Game) clearType(types []BlockType, p Point) []Point {
	points := []Point{p}
	for x := 0; x < g.Width; x++ {
		for y := 0; y < g.Height; y++ {
			for _, t := range types {
				if g.blocks[y][x].Type == t {
					points = append(points, Point{x, y})
				}
			}
		}
	}
	return points
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
				g.blocks[i][x].Type = Empty
			}
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
		if g.blocks[newP.Y][newP.X].Type == bt && !pointExists(newP, points) {
			points = g.getBlockGroup(bt, append(points, newP))
		}
	}

	// look up
	if lp.Y > 0 {
		newP := Point{X: lp.X, Y: lp.Y - 1}
		if g.blocks[newP.Y][newP.X].Type == bt && !pointExists(newP, points) {
			points = g.getBlockGroup(bt, append(points, newP))
		}
	}

	// look right
	if lp.X < (g.Width - 1) {
		newP := Point{X: lp.X + 1, Y: lp.Y}
		if g.blocks[newP.Y][newP.X].Type == bt && !pointExists(newP, points) {
			points = g.getBlockGroup(bt, append(points, newP))
		}
	}

	// look down
	if lp.Y < (g.Height - 1) {
		newP := Point{X: lp.X, Y: lp.Y + 1}
		if g.blocks[newP.Y][newP.X].Type == bt && !pointExists(newP, points) {
			points = g.getBlockGroup(bt, append(points, newP))
		}
	}

	return points
}

func (g *Game) shuntGrid() bool {
	for y := 0; y < g.Height; y++ {
		for x := 0; x < g.Width; x++ {
			// if top row, shunting non empty cell is GAME OVER
			if y == 0 && g.blocks[y][x].Type != Empty {
				g.State = GameOver
				return true
			}

			// if bottom row take from the new row
			if y == g.Height-1 {
				g.blocks[y][x].Type = g.newRow[x].Type
				continue
			}
			g.blocks[y][x].Type = g.blocks[y+1][x].Type
		}
	}
	return false
}
