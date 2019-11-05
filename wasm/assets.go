package main

import "syscall/js"

var blue, red, green, purple js.Value

func initAssets() {
	blue = page.getElementByID("bluesquare")
	red = page.getElementByID("redsquare")
	green = page.getElementByID("greensquare")
	purple = page.getElementByID("purplesquare")
}
