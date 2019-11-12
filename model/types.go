package model

// Point represents a position on a 2d grid
type Point struct {
	X int
	Y int
}

// ClickResult returns data that describes the outcome of the click
type ClickResult struct {
	HadEffect     bool
	BlocksCleared int
	ScoreDelta    int
}

// Game holds all the game state
type Game struct {
	blocks        [][]Block
	newRow        []Block
	Width         int
	Height        int
	State         GameState
	Timer         float64
	Score         int
	Tick          int
	CurrentTick   int
	BlocksCleared int
}

// Block is a area on a grid
type Block struct {
	Type BlockType
}
