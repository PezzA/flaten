package view

import (
	"fmt"
	"sync"
	"syscall/js"
)

var BlueSprite, RedSprite, GreenSprite, PurpleSprite js.Value

var ClearBlue, ClearRed, ClearGreen, ClearPurple js.Value

var Bomb, SlideLeft, SlideUp js.Value

var Sushi js.Value

var SfxMusic, SfxClick, SfxClear, SfxBomb, SfxIncoming js.Value

var elements []string

func (d *JsDoc) addImageElement(id, src string) {
	pattern := fmt.Sprintf("<img id=\"%v\" onload=\"loadEvent('%v');\"/>", id, id)
	html := d.GetElementByID("assets").Get("innerHTML")

	d.GetElementByID("assets").Set("innerHTML", fmt.Sprintf("%v%v", html, pattern))
	d.GetElementByID(id).Set("src", src)
	elements = append(elements, id)
}

func (d *JsDoc) addAudioElement(id, src, audioType string) {
	pattern := fmt.Sprintf("<audio id=\"%v\" oncanplaythrough=\"loadEvent('%v');\"/><source id=\"%vwav\" type=\"%v\"/></audio>", id, id, id, audioType)
	html := d.GetElementByID("assets").Get("innerHTML")

	d.GetElementByID("assets").Set("innerHTML", fmt.Sprintf("%v%v", html, pattern))
	d.GetElementByID(id+"wav").Set("src", src)
	d.GetElementByID(id).Call("load")
	elements = append(elements, id)
}

func loadEvent(this js.Value, i []js.Value) interface{} {
	fmt.Println(i)
	return nil
}

func (d *JsDoc) initAssets() {

	var wg sync.WaitGroup
	wg.Add(1)

	js.Global().Set("loadEvent", js.FuncOf(loadEvent))

	// add the images to the dom
	d.addImageElement("bluesprite", "assets/gfx/gems/blue.png")
	d.addImageElement("redsprite", "assets/gfx/gems/red.png")
	d.addImageElement("greensprite", "assets/gfx/gems/green.png")
	d.addImageElement("purplesprite", "assets/gfx/gems/purple.png")

	d.addImageElement("redclear", "assets/gfx/gems/redclear.png")
	d.addImageElement("greenclear", "assets/gfx/gems/greenclear.png")
	d.addImageElement("blueclear", "assets/gfx/gems/blueclear.png")
	d.addImageElement("purpleclear", "assets/gfx/gems/purpleclear.png")

	d.addImageElement("bomb", "assets/gfx/gems/bomb.png")
	d.addImageElement("slideleft", "assets/gfx/gems/slideleft.png")
	d.addImageElement("slideup", "assets/gfx/gems/slideup.png")
	d.addImageElement("sushi", "assets/gfx/gems/sushi.png")

	d.addAudioElement("click", "assets/sfx/click.wav", "audio/wav")
	d.addAudioElement("clear", "assets/sfx/clear.wav", "audio/wav")
	d.addAudioElement("incoming", "assets/sfx/incoming.wav", "audio/wav")
	d.addAudioElement("bombsfx", "assets/sfx/bomb.wav", "audio/wav")

	wg.Wait()
	// load them when they are done
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
	d.GetElementByID("click").Set("oncanplaythrough", "")
	SfxClear = d.GetElementByID("clear")
	d.GetElementByID("clear").Set("oncanplaythrough", "")
	SfxBomb = d.GetElementByID("bombsfx")
	d.GetElementByID("bombsfx").Set("oncanplaythrough", "")
	SfxIncoming = d.GetElementByID("incoming")
	d.GetElementByID("incoming").Set("oncanplaythrough", "")
}
