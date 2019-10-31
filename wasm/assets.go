package main

import "syscall/js"

var blue, red, green, purple js.Value

func initAssets() {
	blue = d.doc.getElementById("bluesquare")
	red = d.doc.getElementById("redsquare")
	green = d.doc.getElementById("greensquare")
	purple = d.doc.getElementById("purplesquare")
}
