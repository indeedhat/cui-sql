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

	defer g.Close()
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

	if err := renderError(g); nil != err {
		return err
	}

	if err := renderResults(g); nil != err {
		return err
	}

	if nil == g.CurrentView() {
		g.SetCurrentView(V_TREE)
	}

	return nil
}

func bindKeys(g *gocui.Gui) error {
	err := g.SetKeybinding("", gocui.KeyCtrlQ, gocui.ModNone, func(g *gocui.Gui, v *gocui.View) error {
		return gocui.ErrQuit
	})
	if nil != err {
		return err
	}

	err = g.SetKeybinding("", gocui.KeyCtrlT, gocui.ModNone, func(g *gocui.Gui, v *gocui.View) error {
		g.Cursor = false
		_, err := g.SetCurrentView(V_TREE)
		return err
	})
	if nil != err {
		return err
	}

	err = g.SetKeybinding("", gocui.KeyCtrlE, gocui.ModNone, func(g *gocui.Gui, v *gocui.View) error {
		g.Cursor = true
		_, err := g.SetCurrentView(V_EDITOR)
		return err
	})
	if nil != err {
		return err
	}

	err = g.SetKeybinding("", gocui.KeyCtrlR, gocui.ModNone, func(g *gocui.Gui, v *gocui.View) error {
		g.Cursor = false
		_, err := g.SetCurrentView(V_RESULTS)
		return err
	})
	if nil != err {
		return err
	}

	err = bindTreeKeys(g)
	if nil != err {
		return err
	}

	err = bindEditor(g)
	if nil != err {
		return err
	}

	err = bindResults(g)
	if nil != err {
		return err
	}

	return nil
}
