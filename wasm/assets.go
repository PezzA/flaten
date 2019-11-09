package main

import "syscall/js"

var blue, red, green, purple js.Value
var blueSprite, redSprite, greenSprite, purpleSprite js.Value

var music, click js.Value

func initAssets() {
	blue = page.getElementByID("bluesquare")
	red = page.getElementByID("redsquare")
	green = page.getElementByID("greensquare")
	purple = page.getElementByID("purplesquare")

	blueSprite = page.getElementByID("bluesprite")
	redSprite = page.getElementByID("redsprite")
	greenSprite = page.getElementByID("greensprite")
	purpleSprite = page.getElementByID("purplesprite")

	music = page.getElementByID("music")
	click = page.getElementByID("click")
}
