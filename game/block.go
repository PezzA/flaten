package game

import "fmt"

// Point represents a position on a 2d grid
type Point struct {
	X int
	Y int
}

// BlockType holds the type of the block
type BlockType = int

const (
	Empty BlockType = 0
	Red   BlockType = 1
	Green BlockType = 2
	Blue  BlockType = 3
	White BlockType = 4
	Temp  BlockType = 5
)

const (
	playTileMin = 1
	playTileMax = 4
)

// Block is a area on a grid
type Block struct {
	Point
	Type   BlockType
	Moving bool
	Drop   float64
	Dist   float64
}

func (b Block) GetTypeCode() string {

	if b.Type == Red {
		return "R"
	} else if b.Type == Green {
		return "G"
	} else if b.Type == Blue {
		return "B"
	} else if b.Type == White {
		return "P"
	} else if b.Type == Temp {
		return "T"
	} else if b.Type == Empty {
		return "E"
	}
	return string(b.Type)
}

func (b Block) GetMovingCode() string {
	mc := "-"
	if b.Moving {
		mc = fmt.Sprintf("(%v,%v)", b.Drop, b.Dist)
	}
	return mc
}

func (b Block) String() string {
	return fmt.Sprintf("[%v,%v]%v%v", b.X, b.Y, b.GetTypeCode(), b.GetMovingCode())
}
