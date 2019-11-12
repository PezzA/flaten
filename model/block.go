package model

import "math/rand"

// BlockType holds the type of the block
type BlockType = int

const (
	// Empty - No Block Type
	Empty BlockType = 0
	// Red BlockType
	Red BlockType = 1
	// Green BlockType
	Green BlockType = 2
	// Blue BlockType
	Blue BlockType = 3
	// Purple BlockType
	Purple BlockType = 4
)

const (
	playTileMin = 1
	playTileMax = 4
)

func newBlock() Block {
	return Block{
		Type: getRandomPlayTile(),
	}
}

func getRandomPlayTile() int {
	return rand.Intn(playTileMax-playTileMin+1) + playTileMin
}
