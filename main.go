package main

import (
	"github.com/jroimartin/gocui"
	"log"
)

const (
	V_TREE    = " Tree "
	V_EDITOR  = " Editor "
	V_RESULTS = " Results "
)

func main() {
	if err := mysqlConnect(); nil != err {
		log.Fatal("Failed to connect to db")
	}
	defer connection.Close()

	if err := plantTree(); nil != err {
		log.Fatalf("Failed to build tree: %s", err)
	}

	g, err := gocui.NewGui(gocui.OutputNormal)
	if nil != err {
		log.Fatal(err)
	}

	defer g.Close()
	g.SetManagerFunc(render)

	g.Highlight = true
	g.SelFgColor = gocui.ColorBlue

	if err := bindKeys(g); nil != err {
		log.Println("bind keys")
		log.Fatal(err)
	}

	selectDatabase()
	if err := g.MainLoop(); nil != err && gocui.ErrQuit != err {
		log.Println("main loop")
		log.Fatal(err)
	}
}

func render(g *gocui.Gui) error {
	err := renderTree(g)
	if nil != err {
		log.Println(err)
		return err
	}

	err = renderEditor(g)
	if nil != err {
		log.Println(err)
		return err
	}

	err = renderResults(g)
	if nil != err {
		log.Println(err)
		return err
	}

	if nil == g.CurrentView() {
		g.SetCurrentView(V_TREE)
	}
	return err
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
