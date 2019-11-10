package model

import "math/rand"

// BlockType holds the type of the block
type BlockType = int

const (
	Empty  BlockType = 0
	Red    BlockType = 1
	Green  BlockType = 2
	Blue   BlockType = 3
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

func (b Block) GetTypeCode() string {
	if b.Type == Red {
		return "R"
	} else if b.Type == Green {
		return "G"
	} else if b.Type == Blue {
		return "B"
	} else if b.Type == Purple {
		return "P"
	} else if b.Type == Empty {
		return "E"
	}
	return string(b.Type)
}
