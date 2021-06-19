package controllers

import (
	"github.com/revel/revel"
)

type stuff struct {
	Hello string
	Foo string ` json:"foo" xml:"foo" `
	Bar int ` json:"bar" xml:"bar" `
}

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	stuff := stuff{Foo: "Oke", Bar: 69, Hello: "Hellooo"}

	return c.RenderJSON(stuff)
}


