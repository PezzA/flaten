package view

import "syscall/js"

var BlueSprite, RedSprite, GreenSprite, PurpleSprite js.Value

var ClearBlue, ClearRed, ClearGreen, ClearPurple js.Value

var Bomb, SlideLeft, SlideUp js.Value

var Sushi js.Value

var SfxMusic, SfxClick, SfxClear, SfxBomb, SfxIncoming, SfxIncomingPip js.Value

func (d *JsDoc) initAssets() {
	BlueSprite = d.GetElementByID("bluesprite")
	RedSprite = d.GetElementByID("redsprite")
	GreenSprite = d.GetElementByID("greensprite")
	PurpleSprite = d.GetElementByID("purplesprite")

	ClearBlue = d.GetElementByID("blueclear")
	ClearRed = d.GetElementByID("redclear")
	ClearGreen = d.GetElementByID("greenclear")
	ClearPurple = d.GetElementByID("purpleclear")

	Bomb = d.GetElementByID("bomb")
	SlideLeft = d.GetElementByID("slideleft")
	SlideUp = d.GetElementByID("slideup")
	Sushi = d.GetElementByID("sushi")

	SfxMusic = d.GetElementByID("music")
	SfxClick = d.GetElementByID("click")
	SfxClear = d.GetElementByID("clear")
	SfxBomb = d.GetElementByID("bombsfx")
	SfxIncoming = d.GetElementByID("incoming")
	SfxIncomingPip = d.GetElementByID("incomingpip")
}
