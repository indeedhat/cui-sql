package main

import (
	"github.com/jroimartin/gocui"
)

func renderHelp(g *gocui.Gui) error {
	if !views.Help.Visible {
		g.DeleteView(views.Help.Title)
		return nil
	}

	v, err := initView(g, views.Help)
	if nil != err {
		return err
	}

	v.Highlight = false
	if "" != v.Buffer() {
		return nil
	}

	v.Write([]byte(`
GLOBAL
?      : toggle help
Ctrl+t : focus tree
Ctrl+e : focus editor
Ctrl+r : focus results
Ctrl+q : Qiot

TREE
j      : down
k      : up
o      : show/hide children
Enter  : select database (only if db is highlighted)
Ctrl+j : focus Results
Ctrl+k : focus Editor
Ctrl+l : focus Editor

RESULTS
j      : down
k      : up
Ctrl+h : focus tree
Ctrl+k : focus Editor

EDITOR
Ctrl+Space : Run Query
Ctrl+/     : Open Query List
Ctrl+h     : Focus Tree
Crtl+j     : Focus Editor

QUERY LIST
Enter : Open query in editor
Esc   : Close list
`))
	return nil
}
