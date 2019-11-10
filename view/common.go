package view

import (
	"fmt"
	"syscall/js"
)

func (d *JsDoc) Log(text string) {
	fmt.Println(text)
}

func (d *JsDoc) GetElementInnerHTML(elementID string, html string) {
	d.document.Call("getElementById", elementID).Set("innerHTML", html)
}

func (d *JsDoc) GetElementByID(elementID string) js.Value {
	return d.document.Call("getElementById", elementID)
}

func (d *JsDoc) SetGlobalAlpha(alpha float64) {
	d.TwoDCtx.Set("globalAlpha", alpha)
}

func (d *JsDoc) SetCanvasSize(x, y int) {
	d.canvasElem.Set("width", x)
	d.canvasElem.Set("height", y)

	canvasWidth = float64(x)
	canvasHeight = float64(y)
}

func (d *JsDoc) StartAnimLoop() {
	js.Global().Call("requestAnimationFrame", renderFrameEvt)
}

func (d *JsDoc) DrawImage(img js.Value, sX, sY, sW, sH, dX, dY, dW, dH int) {
	d.TwoDCtx.Call("drawImage", img, sX, sY, sW, sH, dX, dY, dW, dH)
}

func (d *JsDoc) ClearFrame(x, y, w, h int) {
	d.TwoDCtx.Call("clearRect", x, y, w, h)
}

// DrawText draws text to the canvas
func (d *JsDoc) DrawText(text, font, fillStyle, textAlign string, x, y int) {
	d.TwoDCtx.Set("font", font)
	d.TwoDCtx.Set("fillStyle", fillStyle)
	d.TwoDCtx.Set("textAlign", textAlign)
	d.TwoDCtx.Call("fillText", text, x, y)
}

// DrawRect draws a rectangle to the canvas
func (d *JsDoc) DrawRect(x, y, w, h int, fillStyle string) {
	d.TwoDCtx.Set("fillStyle", fillStyle)
	d.TwoDCtx.Call("fillRect", x, y, w, h)
}
