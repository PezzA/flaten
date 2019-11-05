package main

import (
	"fmt"

	"github.com/pezza/flaten/game"
)

func debugTable() {
	html := ""
	for y := 0; y < g.Height; y++ {
		line := ""
		for x := 0; x < g.Width; x++ {

			output := ""

			if g.GetBlock(x, y).Type != game.Empty {
				output = fmt.Sprintf("%v", g.GetBlock(x, y))
			}

			line = fmt.Sprintf("%v<td>%v</td>", line, output)
		}
		html = fmt.Sprintf("%v<tr>%v</tr>", html, line)
	}
	page.setElementInnerHTML("dbg", fmt.Sprintf(`<table>%v</table>`, html))
}
