package model

import "math/rand"

// BlockType holds the type of the block
type BlockType = int

const (
	Empty       BlockType = 0
	Red         BlockType = 1
	Green       BlockType = 2
	Blue        BlockType = 3
	Purple      BlockType = 4
	RedClear    BlockType = 5
	GreenClear  BlockType = 6
	BlueClear   BlockType = 7
	PurpleClear BlockType = 8
	Bomb        BlockType = 9
	LeftSlide   BlockType = 10
	UpSlide     BlockType = 11
	RightSlide  BlockType = 12
	DownSlide   BlockType = 13
)

const (
	playTileMin  = 1
	playTileMax  = 4
	clearTileMin = 5
	clearTileMax = 8
	slideTileMin = 10
	slideTileMax = 13
)

var tileDist []int = []int{
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 1, 1, 2, 2, 3, 3}

func newBlock() Block {
	return Block{
		Type: getRandomPlayTile(),
	}
}

func getRandomPlayTile() int {
	class := rand.Intn(len(tileDist) - 1)

	if tileDist[class] == 0 {
		return rand.Intn(playTileMax-playTileMin+1) + playTileMin
	}

	if tileDist[class] == 3 {
		return Bomb
	}

	if tileDist[class] == 1 {
		return rand.Intn(clearTileMax-clearTileMin+1) + clearTileMin
	}

	return rand.Intn(playTileMax-playTileMin+1) + playTileMin
	//if tileDist[class] == 2 {
	//	return rand.Intn(sli-clearTileMin+1) + clearTileMin
	//}

	// 1 in 50 chance for
	// 1 in 30 chance
	// 1 in 30 chance for bomb

}
