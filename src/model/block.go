package model

import "math/rand"

type BlockType int

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
	SlideLeft   BlockType = 10
	SlideUp     BlockType = 11
)

const (
	playTileMin    = 1
	playTileMax    = 4
	clearTileMin   = 5
	clearTileMax   = 8
	specialTileMin = 9
	specialTileMax = 11
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
	0, 0, 0, 0, 0, 0, 0, 1, 1, 2}

func newBlock() Block {
	return Block{
		Type: BlockType(getRandomPlayTile()),
	}
}

func IsBomb(b BlockType) bool {
	return b == Bomb
}

func IsClearTile(b BlockType) bool {
	return b == RedClear || b == BlueClear || b == GreenClear || b == PurpleClear
}

func IsBlock(b BlockType) bool {
	return b == Red || b == Blue || b == Green || b == Purple
}

func IsSlide(b BlockType) bool {
	return b == SlideLeft || b == SlideUp
}

func getRandomPlayTile() int {
	class := rand.Intn(len(tileDist))

	if tileDist[class] == 0 {
		return rand.Intn(playTileMax-playTileMin+1) + playTileMin
	}

	if tileDist[class] == 2 {
		return rand.Intn(specialTileMax-specialTileMin+1) + specialTileMin
	}

	if tileDist[class] == 1 {
		return rand.Intn(clearTileMax-clearTileMin+1) + clearTileMin
	}

	return rand.Intn(playTileMax-playTileMin+1) + playTileMin
}
