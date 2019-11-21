package main

import (
	"fmt"
	"github.com/jroimartin/gocui"
)

func renderQueryList(g *gocui.Gui) error {
	if !views.QueryList.Visible {
		g.DeleteView(views.QueryList.Title)
		return nil
	}

	v, err := initView(g, views.QueryList)

	if nil != err {
		return err
	}

	v.Editable = false
	v.Wrap = false
	v.Autoscroll = true
	v.Highlight = true
	v.BgColor = gocui.ColorCyan
	v.FgColor = gocui.ColorBlack

	reDrawQueryList(v)
	return nil
}

func reDrawQueryList(v *gocui.View) {
	if "" != v.Buffer() {
		return
	}

	var text string
	for k := range queryList() {
		text += fmt.Sprintln(k)
	}

	v.Write([]byte(text[:len(text)-1]))
}

func queryList() map[string]string {
	return map[string]string{
		" INSERT ": "INSERT INTO `%s` (\n\n) VALUES (\n\n)",
		" SELECT ": "SELECT * FROM `%s`",
		" UPDATE ": "UPDATE `%s` SET ",
		" DELETE ": "DELETE FROM `%s` WHERE ",
	}
}

func writeQueryToEditor(g *gocui.Gui, key string) {
	queries := queryList()
	if _, ok := queries[key]; !ok {
		return
	}

	editor := views.Editor.View(g)
	if nil == editor {
		return
	}

	editor.Clear()
	editor.Write([]byte(fmt.Sprintf(queries[key], database)))
}

func bindQueryList(g *gocui.Gui) error {
	err := g.SetKeybinding(V_QueryList, gocui.KeyEnter, gocui.ModNone, func(g *gocui.Gui, v *gocui.View) error {
		_, y := v.Cursor()
		key := v.ViewBufferLines()[y]

		views.QueryList.Toggle()
		views.SelectView(g, views.Editor)
		writeQueryToEditor(g, key)

		return nil
	})
	if nil != err {
		return err
	}

	err = g.SetKeybinding(V_QueryList, gocui.KeyEsc, gocui.ModNone, func(g *gocui.Gui, v *gocui.View) error {
		views.QueryList.Toggle()
		views.SelectView(g, views.Editor)

		return nil
	})
	if nil != err {
		return err
	}

	err = g.SetKeybinding(V_QueryList, 'k', gocui.ModNone, func(g *gocui.Gui, v *gocui.View) error {
		v.MoveCursor(0, -1, false)
		return nil
	})
	if nil != err {
		return err
	}

	err = g.SetKeybinding(V_QueryList, 'j', gocui.ModNone, func(g *gocui.Gui, v *gocui.View) error {
		v.MoveCursor(0, 1, false)
		return nil
	})
	if nil != err {
		return err
	}

	return nil
}
