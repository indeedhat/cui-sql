package main

import (
	"fmt"
	"strconv"

	"github.com/jroimartin/gocui"
)

func renderEditor(g *gocui.Gui) error {
	if !views.Editor.Visible {
		g.DeleteView(views.Editor.Title)
		return nil
	}

	renderLineNumbers(g)
	v, err := initView(g, views.Editor)

	if nil != err {
		return err
	}

	v.Editable = true
	v.Wrap = true
	v.Autoscroll = true

	return nil
}

func renderLineNumbers(g *gocui.Gui) {
	if !views.EditorLines.Visible {
		g.DeleteView(views.EditorLines.Title)
		return
	}

	scrollPos := 0
	_, _, _, height := views.EditorLines.Coords(g)

	if editor := views.Editor.View(g); nil != editor {
		_, scrollPos = editor.Origin()
		height += scrollPos
	}

	viewWidth := len(strconv.Itoa(height)) + 2

	views.Editor.XMod = viewWidth - 1
	views.EditorLines.WidthMod = viewWidth

	v, err := initView(g, views.EditorLines)
	if nil != err {
		return
	}

	v.Clear()

	for i := 0; i < height; i++ {
		v.Write([]byte(fmt.Sprintf("%d\n", i+1)))
	}

	v.SetOrigin(0, scrollPos)
}

func bindEditor(g *gocui.Gui) error {
	err := g.SetKeybinding(V_EDITOR, gocui.KeyCtrlSpace, gocui.ModNone, func(g *gocui.Gui, v *gocui.View) error {
		resView := views.Results.View(g)
		if nil == resView {
			return nil
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

	err = g.SetKeybinding(V_EDITOR, gocui.KeyCtrlSlash, gocui.ModNone, func(g *gocui.Gui, v *gocui.View) error {
		views.QueryList.Toggle()
		views.SelectView(g, views.QueryList)
		return nil
	})
	if nil != err {
		return err
	}

	if err := views.Editor.bindChangeView(g, views.Tree, gocui.KeyCtrlH, true); nil != err {
		return err
	}
	if err := views.Editor.bindChangeView(g, views.Results, gocui.KeyCtrlJ, true); nil != err {
		return err
	}

	return nil
}
