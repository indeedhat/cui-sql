package main

import (
	"github.com/jroimartin/gocui"
)

func renderError(g *gocui.Gui) error {
	if !views.Error.Visible {
		return nil
	}

	v, err := initView(g, views.Error)
	if nil != err {
		return err
	}

	v.Highlight = false
	v.Wrap = false

	return nil
}

func reDrawError(v *gocui.View, err error) {
	v.Clear()

	v.Write([]byte(err.Error()))
}
