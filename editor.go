package main

import (
	"log"

	"github.com/jroimartin/gocui"
)

func renderEditor(g *gocui.Gui) error {
	x, y := g.Size()

	v, err := g.SetView(V_EDITOR, x/3+1, 0, x-1, y/6*4)

	if nil != err && gocui.ErrUnknownView != err {
		log.Println("render")
		log.Fatal(err)
	}

	v.SelFgColor = gocui.ColorBlack
	v.SelBgColor = gocui.ColorBlue
	v.Editable = true
	v.Wrap = true
	v.Autoscroll = true

	v.Title = V_EDITOR

	return nil
}

func bindEditor(g *gocui.Gui) error {
	err := g.SetKeybinding(V_EDITOR, gocui.KeyCtrlSpace, gocui.ModNone, func(g *gocui.Gui, v *gocui.View) error {
		resView, err := g.View(V_RESULTS)
		if nil != err {
			return err
		}

		rows, err := query(database, v.Buffer())
		if nil != err {
			reDrawResultsError(resView, err)
			return nil
		}

		data, err := fetchRows(rows)
		if nil != err {
			reDrawResultsError(resView, err)
			return nil
		}

		reDrawResults(resView, data)

		return nil
	})

	if nil != err {
		return err
	}
	return nil
}
