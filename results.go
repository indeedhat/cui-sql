package main

import (
	"github.com/cheynewallace/tabby"
	"github.com/jroimartin/gocui"
	"text/tabwriter"
)

func renderResults(g *gocui.Gui) error {
	if !views.Results.Visible {
		return nil
	}

	v, err := initView(g, views.Results)
	if nil != err {
		return err
	}

	v.Highlight = true
	v.Wrap = false

	return nil
}

func reDrawResults(v *gocui.View, data []map[string]string) {
	v.Clear()

	writer := tabwriter.NewWriter(v, 0, 0, 2, ' ', 0)
	table := tabby.NewCustom(writer)

	var fields []string
	var tmp []interface{}
	for field := range data[0] {
		fields = append(fields, field)
		tmp = append(tmp, field)
	}

	table.AddHeader(tmp...)

	for _, row := range data {
		tmp = []interface{}{}
		for _, field := range fields {
			tmp = append(tmp, row[field])
		}

		table.AddLine(tmp...)
	}

	table.Print()
}

func bindResults(g *gocui.Gui) error {
	err := g.SetKeybinding(V_RESULTS, 'k', gocui.ModNone, func(g *gocui.Gui, v *gocui.View) error {
		v.MoveCursor(0, -1, false)
		return nil
	})
	if nil != err {
		return err
	}

	err = g.SetKeybinding(V_RESULTS, 'j', gocui.ModNone, func(g *gocui.Gui, v *gocui.View) error {
		v.MoveCursor(0, 1, false)
		return nil
	})

	return err
}
