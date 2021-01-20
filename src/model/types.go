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
	ClickType     BlockType
}

// Results holds stuff
type Results struct {
	BlockClears  map[BlockType]int
	OverallScore int
	TenPlus      int
	TwentyPlus   int
	Rows         int
}

// Game holds all the game state
type Game struct {
	blocks      [][]Block
	newRow      []Block
	Width       int
	Height      int
	State       GameState
	Timer       float64
	Tick        int
	CurrentTick int
	scores      Results
}

// Block is a area on a grid
type Block struct {
	Type BlockType
}
