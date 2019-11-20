package main

import (
	"github.com/cheynewallace/tabby"
	"github.com/jroimartin/gocui"
	"log"
	"text/tabwriter"
)

func renderResults(g *gocui.Gui) error {
	x, y := g.Size()

	v, err := g.SetView(V_RESULTS, x/3+1, y/6*4+1, x-1, y-1)

	if nil != err && gocui.ErrUnknownView != err {
		log.Println("render")
		log.Fatal(err)
	}

	v.SelFgColor = gocui.ColorBlack
	v.SelBgColor = gocui.ColorBlue
	v.Highlight = true
	v.Wrap = false

	v.Title = V_RESULTS

	return nil
}

func reDrawResultsError(v *gocui.View, err error) {
	v.Clear()

	v.Write([]byte(err.Error()))
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
