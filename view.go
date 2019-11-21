package main

import (
	"github.com/jroimartin/gocui"
)

const (
	V_TREE         = " Tree "
	V_EDITOR       = " Editor "
	V_RESULTS      = " Results "
	V_ERROR        = " Error "
	V_HELP         = " Help "
	V_QueryList    = " Query "
	V_LINE_NUMBERS = "─"
)

var views *ViewManager

type ViewManager struct {
	Window      *View
	Tree        *View
	Editor      *View
	EditorLines *View
	Results     *View
	Error       *View
	Help        *View
	QueryList   *View

	CurrentView string
}

func (v *ViewManager) SelectView(g *gocui.Gui, view *View) {
	v.CurrentView = view.Title
	g.SetCurrentView(v.CurrentView)
}

type View struct {
	Title     string
	Width     int
	Height    int
	WidthMod  int
	HeightMod int
	X         int
	Y         int
	XMod      int
	YMod      int
	Visible   bool
}

func (v *View) Coords(g *gocui.Gui) (int, int, int, int) {
	maxX, maxY := g.Size()

	return percent(v.X, maxX) + v.XMod,
		percent(v.Y, maxY) + min(v.Y, 1) + v.YMod,
		max(1, percent(v.X+v.Width, maxX)-1) + v.WidthMod,
		max(1, percent(v.Y+v.Height, maxY)) + v.HeightMod
}

func (v *View) Toggle() {
	v.Visible = !v.Visible
}

func (v *View) View(g *gocui.Gui) *gocui.View {
	view, err := g.View(v.Title)
	if nil != err {
		return nil
	}

	return view
}

func (v *View) bindChangeView(g *gocui.Gui, view *View, key interface{}, cursor bool) error {
	return g.SetKeybinding(v.Title, key, gocui.ModNone, func(g *gocui.Gui, v *gocui.View) error {
		if !view.Visible {
			return nil
		}

		g.Cursor = cursor
		views.SelectView(g, view)
		return nil
	})
}

func initViewManager() {
	views = &ViewManager{
		Window: &View{
			Title: "",
		},
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
			Height:  30,
			Width:   70,
			Visible: true,
			X:       30,
			Y:       0,
			XMod:    2,
		},
		EditorLines: &View{
			Title:    V_LINE_NUMBERS,
			Height:   30,
			Width:    0,
			Visible:  true,
			X:        30,
			Y:        0,
			WidthMod: 3,
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
			Width:   0,
			Visible: false,
			X:       100,
			Y:       0,
			XMod:    -10,
		},
		CurrentView: V_TREE,
	}
}

func initView(g *gocui.Gui, view *View) (*gocui.View, error) {
	x, y, x2, y2 := view.Coords(g)
	v, err := g.SetView(view.Title, x, y, x2, y2)

	if nil != err && gocui.ErrUnknownView != err {
		return nil, err
	}

	v.SelFgColor = gocui.ColorBlack
	v.SelBgColor = gocui.ColorBlue
	v.Title = view.Title

	return v, nil
}
