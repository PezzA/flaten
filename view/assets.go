package view

import "syscall/js"

var BlueSprite, RedSprite, GreenSprite, PurpleSprite js.Value

func (d *JsDoc) initAssets() {
	BlueSprite = d.GetElementByID("bluesprite")
	RedSprite = d.GetElementByID("redsprite")
	GreenSprite = d.GetElementByID("greensprite")
	PurpleSprite = d.GetElementByID("purplesprite")
}
