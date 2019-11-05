package main

import "syscall/js"

func (d *jsDoc) setElementInnerHTML(elementID string, html string) {
	d.doc.Call("getElementById", elementID).Set("innerHTML", html)
}

func (d *jsDoc) getElementByID(elementID string) js.Value {
	return d.doc.Call("getElementById", elementID)
}
