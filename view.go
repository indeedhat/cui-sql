package main

import (
	"fmt"
	"github.com/jroimartin/gocui"
)

const (
	V_TREE      = " Tree "
	V_EDITOR    = " Editor "
	V_RESULTS   = " Results "
	V_ERROR     = " Error "
	V_HELP      = " Help "
	V_QueryList = " Query "
)

var views *ViewManager

type ViewManager struct {
	Tree      *View
	Editor    *View
	Results   *View
	Error     *View
	Help      *View
	QueryList *View
}

type View struct {
	Title   string
	Width   int
	Height  int
	X       int
	Y       int
	Visible bool
}

func (v *View) Coords(g *gocui.Gui) (int, int, int, int) {
	maxX, maxY := g.Size()

	return percent(v.X, maxX),
		percent(v.Y, maxY),
		max(1, percent(v.X+v.Width, maxX)-1),
		max(1, percent(v.Y+v.Height, maxY)-1)
}

func (v *View) Toggle() {
	v.Visible = !v.Visible
}

func initViewManager() {
	views = &ViewManager{
		Tree: &View{
			Title:   V_TREE,
			Height:  100,
			Width:   30,
			Visible: true,
			X:       0,
			Y:       0,
		},
		Editor: &View{
			Title:   V_EDITOR,
			Width:   70,
			Height:  30,
			Visible: true,
			X:       30,
			Y:       0,
		},
		Results: &View{
			Title:   V_RESULTS,
			Height:  70,
			Width:   70,
			Visible: true,
			X:       30,
			Y:       30,
		},
		Error: &View{
			Title:   V_ERROR,
			Height:  70,
			Width:   70,
			Visible: false,
			X:       30,
			Y:       30,
		},
		Help: &View{
			Title:   V_HELP,
			Height:  80,
			Width:   80,
			Visible: false,
			X:       10,
			Y:       10,
		},
		QueryList: &View{
			Title:   V_QueryList,
			Height:  30,
			Width:   20,
			Visible: false,
			X:       80,
			Y:       20,
		},
	}
}

func initView(g *gocui.Gui, view *View) (*gocui.View, error) {

	x, y, x2, y2 := view.Coords(g)
	fmt.Printf("init %s\n", view.Title)
	fmt.Println([]int{x, y, x2, y2})

	v, err := g.SetView(view.Title, x, y, x2, y2)

	if nil != err && gocui.ErrUnknownView != err {
		return nil, err
	}

	v.SelFgColor = gocui.ColorBlack
	v.SelBgColor = gocui.ColorBlue
	v.Title = V_TREE

	return v, nil
}
