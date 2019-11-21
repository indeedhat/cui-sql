package main

import (
	"github.com/indeedhat/tree"
	"github.com/jroimartin/gocui"
)

var tre *tree.Tree

func bindTreeKeys(g *gocui.Gui) error {
	err := g.SetKeybinding(V_TREE, 'k', gocui.ModNone, func(g *gocui.Gui, v *gocui.View) error {
		v.MoveCursor(0, -1, false)
		return nil
	})
	if nil != err {
		return err
	}

	err = g.SetKeybinding(V_TREE, 'j', gocui.ModNone, func(g *gocui.Gui, v *gocui.View) error {
		v.MoveCursor(0, 1, false)
		return nil
	})
	if nil != err {
		return err
	}

	err = g.SetKeybinding(V_TREE, 'o', gocui.ModNone, func(g *gocui.Gui, v *gocui.View) error {
		_, offset := v.Origin()
		_, y := v.Cursor()

		limb := tre.FindByIndex(y+offset, true)
		switch branch := limb.(type) {
		case *tree.Branch:
			branch.Toggle()
			reDrawTree(v)
		}

		return nil
	})

	err = g.SetKeybinding(V_TREE, gocui.KeyEnter, gocui.ModNone, func(g *gocui.Gui, v *gocui.View) error {
		_, offset := v.Origin()
		_, y := v.Cursor()
		l := tre.FindByIndex(y+offset, true)
		switch branch := l.(type) {
		case *tree.Branch:
			database = branch.Key
			selectDatabase()
		}

		reDrawTree(v)
		return nil
	})

	if err := views.Tree.bindChangeView(g, views.Editor, gocui.KeyCtrlL, true); nil != err {
		return err
	}
	if err := views.Tree.bindChangeView(g, views.Editor, gocui.KeyCtrlK, true); nil != err {
		return err
	}
	if err := views.Tree.bindChangeView(g, views.Results, gocui.KeyCtrlJ, true); nil != err {
		return err
	}

	return nil
}

func renderTree(g *gocui.Gui) error {
	if !views.Tree.Visible {
		g.DeleteView(views.Tree.Title)
		return nil
	}

	v, err := initView(g, views.Tree)
	if nil != err {
		return err
	}

	v.Highlight = true
	v.Wrap = false

	reDrawTree(v)

	return nil
}

func reDrawTree(v *gocui.View) {
	v.Clear()
	data := tre.Render()
	v.Write([]byte(data[:len(data)-1]))
}

func plantTree() error {
	tre = tree.NewTree()

	dbs, err := fetchDatabases()
	if nil != err {
		return err
	}

	for _, db := range dbs {
		node := &tree.Branch{
			Key:  db,
			Text: db,
		}

		tables, err := fetchTables(db)
		if nil != err {
			return err
		}

		for _, table := range tables {
			node.Limbs = append(node.Limbs, &tree.Leaf{
				Key:  table,
				Text: table,
			})
		}

		tre.Root.Limbs = append(tre.Root.Limbs, node)
	}

	tre.Plant()

	return nil
}
