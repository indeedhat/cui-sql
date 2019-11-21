package main

import (
	"github.com/jroimartin/gocui"
)

func renderEditor(g *gocui.Gui) error {
	if !views.Editor.Visible {
		return nil
	}

	v, err := initView(g, views.Editor)

	if nil != err {
		return err
	}

	v.Editable = true
	v.Wrap = true
	v.Autoscroll = true

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
			views.Results.Visible = false
			views.Error.Visible = true
			reDrawError(resView, err)
			return nil
		}

		data, err := fetchRows(rows)
		if nil != err {
			views.Results.Visible = false
			views.Error.Visible = true
			reDrawError(resView, err)
			return nil
		}

		views.Results.Visible = true
		views.Error.Visible = false
		reDrawResults(resView, data)

		return nil
	})

	if nil != err {
		return err
	}
	return nil
}
