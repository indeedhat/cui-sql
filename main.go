package main

import (
	"github.com/jroimartin/gocui"
	"log"
)

func main() {
	mustNotError(
		mysqlConnect(),
	)
	defer connection.Close()

	mustNotError(
		plantTree(),
	)

	initViewManager()

	g, err := initGui()
	mustNotError(err)
	defer g.Close()

	selectDatabase()
	if err := g.MainLoop(); nil != err && gocui.ErrQuit != err {
		log.Println("main loop")
		mustNotError(err)
	}
}

func mustNotError(err error) {
	if nil != err {
		log.Fatal(err)
	}
}

func initGui() (*gocui.Gui, error) {
	g, err := gocui.NewGui(gocui.OutputNormal)
	if nil != err {
		return nil, err
	}

	g.SetManagerFunc(render)

	g.Highlight = true
	g.SelFgColor = gocui.ColorBlue

	if err := bindKeys(g); nil != err {
		return nil, err
	}

	return g, nil
}

func render(g *gocui.Gui) error {
	if err := renderTree(g); nil != err {
		return err
	}

	if err := renderEditor(g); nil != err {
		return err
	}

	if err := renderQueryList(g); nil != err {
		return err
	}

	if err := renderError(g); nil != err {
		return err
	}

	if err := renderResults(g); nil != err {
		return err
	}

	if err := renderHelp(g); nil != err {
		return err
	}

	g.SetCurrentView(views.CurrentView)
	return nil
}

func bindKeys(g *gocui.Gui) error {
	err := g.SetKeybinding("", gocui.KeyCtrlQ, gocui.ModNone, func(g *gocui.Gui, v *gocui.View) error {

		return gocui.ErrQuit
	})
	if nil != err {
		return err
	}

	if err := views.Window.bindChangeView(g, views.Tree, gocui.KeyCtrlT, true); nil != err {
		return err
	}

	if err := views.Window.bindChangeView(g, views.Editor, gocui.KeyCtrlE, true); nil != err {
		return err
	}

	if err := views.Window.bindChangeView(g, views.Results, gocui.KeyCtrlR, true); nil != err {
		return err
	}

	err = g.SetKeybinding("", '?', gocui.ModNone, func(g *gocui.Gui, v *gocui.View) error {
		views.Help.Toggle()
		if views.Help.Visible {
			views.SelectView(g, views.Help)
		} else {
			views.SelectView(g, views.Tree)
		}

		return nil
	})
	if nil != err {
		return err
	}

	if err = bindTreeKeys(g); nil != err {
		return err
	}

	if err = bindEditor(g); nil != err {
		return err
	}

	if err = bindResults(g); nil != err {
		return err
	}

	if err = bindQueryList(g); nil != err {
		return err
	}

	return nil
}
