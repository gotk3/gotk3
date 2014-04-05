package main

import (
	"github.com/conformal/gotk3/gtk"
)

// Setup the Window.
func setupWindow() *gtk.Window {
	w, _ := gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
	w.Connect("destroy", gtk.MainQuit)
	w.SetDefaultSize(500, 300)
	w.SetPosition(gtk.WIN_POS_CENTER)
	w.SetTitle("TextView properties example")

	return w
}

// Setup the TextView, put it in a ScrolledWindow, and add both to box.
func setupTextView(box *gtk.Box) *gtk.TextView {
	sw, _ := gtk.ScrolledWindowNew(nil, nil)
	tv, _ := gtk.TextViewNew()
	sw.Add(tv)
	box.PackStart(sw, true, true, 0)
	return tv
}

type BoolProperty struct {
	Name string
	Get  func() bool
	Set  func(bool)
}

func setupPropertyCheckboxes(tv *gtk.TextView, outer *gtk.Box, props []*BoolProperty) {
	box, _ := gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, 0)
	for _, prop := range props {
		chk, _ := gtk.CheckButtonNewWithLabel(prop.Name)
		// initialize the checkbox with the property's current value
		chk.SetActive(prop.Get())
		p := prop // w/o this all the checkboxes will toggle the last property in props
		chk.Connect("toggled", func() {
			p.Set(chk.GetActive())
		})
		box.PackStart(chk, true, true, 0)
	}
	outer.PackStart(box, false, false, 0)
}

func main() {
	gtk.Init(nil)

	win := setupWindow()
	box, _ := gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 0)
	win.Add(box)

	tv := setupTextView(box)

	props := []*BoolProperty{
		&BoolProperty{"cursor visible", (*tv).GetCursorVisible, (*tv).SetCursorVisible},
		&BoolProperty{"editable", (*tv).GetEditable, (*tv).SetEditable},
		&BoolProperty{"overwrite", (*tv).GetOverwrite, (*tv).SetOverwrite},
		&BoolProperty{"accepts tab", (*tv).GetAcceptsTab, (*tv).SetAcceptsTab},
	}

	setupPropertyCheckboxes(tv, box, props)

	win.ShowAll()

	gtk.Main()
}
