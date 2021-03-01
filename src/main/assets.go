package main

import (
	"syscall/js"

	"github.com/pezza/flaten/src/model"
)

type assetKey int
type assetMap map[assetKey]js.Value

const (
	sfxMusic       assetKey = iota
	sfxClick       assetKey = iota
	sfxClear       assetKey = iota
	sfxBomb        assetKey = iota
	sfxIncoming    assetKey = iota
	sfxIncomingPip assetKey = iota
)

const (
	gfxBlueSquare   = assetKey(model.Blue)
	gfxRedSquare    = assetKey(model.Red)
	gfxGreenSquare  = assetKey(model.Green)
	gfxPurpleSquare = assetKey(model.Purple)
	gfxBlueClear    = assetKey(model.BlueClear)
	gfxRedClear     = assetKey(model.RedClear)
	gfxGreenClear   = assetKey(model.GreenClear)
	gfxPurpleClear  = assetKey(model.PurpleClear)
	gfxBomb         = assetKey(model.Bomb)
	gfxSlideLeft    = assetKey(model.SlideLeft)
	gfxSlideUp      = assetKey(model.SlideUp)

	gfxMeadowSprite = 1000
)
