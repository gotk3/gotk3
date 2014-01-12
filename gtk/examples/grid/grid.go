/*
 * Copyright (c) 2013-2014 Conformal Systems <info@conformal.com>
 *
 * This file originated from: http://opensource.conformal.com/
 *
 * Permission to use, copy, modify, and distribute this software for any
 * purpose with or without fee is hereby granted, provided that the above
 * copyright notice and this permission notice appear in all copies.
 *
 * THE SOFTWARE IS PROVIDED "AS IS" AND THE AUTHOR DISCLAIMS ALL WARRANTIES
 * WITH REGARD TO THIS SOFTWARE INCLUDING ALL IMPLIED WARRANTIES OF
 * MERCHANTABILITY AND FITNESS. IN NO EVENT SHALL THE AUTHOR BE LIABLE FOR
 * ANY SPECIAL, DIRECT, INDIRECT, OR CONSEQUENTIAL DAMAGES OR ANY DAMAGES
 * WHATSOEVER RESULTING FROM LOSS OF USE, DATA OR PROFITS, WHETHER IN AN
 * ACTION OF CONTRACT, NEGLIGENCE OR OTHER TORTIOUS ACTION, ARISING OUT OF
 * OR IN CONNECTION WITH THE USE OR PERFORMANCE OF THIS SOFTWARE.
 */

package main

import (
	"github.com/conformal/gotk3/gtk"
	"log"
)

func main() {
	gtk.Init(nil)

	win, err := gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
	if err != nil {
		log.Fatal("Unable to create window:", err)
	}
	win.SetTitle("Grid Example")
	win.Connect("destroy", func() {
		gtk.MainQuit()
	})

	// Create a new grid widget to arrange child widgets
	grid, err := gtk.GridNew()
	if err != nil {
		log.Fatal("Unable to create grid:", err)
	}

	// gtk.Grid embeds an Orientable struct to simulate the GtkOrientable
	// GInterface.  Set the orientation from the default horizontal to
	// vertical.
	grid.SetOrientation(gtk.ORIENTATION_VERTICAL)

	// Create some widgets to put in the grid.
	lab, err := gtk.LabelNew("Just a label")
	if err != nil {
		log.Fatal("Unable to create label:", err)
	}

	btn, err := gtk.ButtonNewWithLabel("Button with label")
	if err != nil {
		log.Fatal("Unable to create button:", err)
	}

	entry, err := gtk.EntryNew()
	if err != nil {
		log.Fatal("Unable to create entry:", err)
	}

	spnBtn, err := gtk.SpinButtonNewWithRange(0.0, 1.0, 0.001)
	if err != nil {
		log.Fatal("Unable to create spin button:", err)
	}

	nb, err := gtk.NotebookNew()
	if err != nil {
		log.Fatal("Unable to create notebook:", err)
	}

	// Calling (*gtk.Container).Add() with a gtk.Grid will add widgets next
	// to each other, in the order they were added, to the right side of the
	// last added widget when the grid is in a horizontal orientation, and
	// at the bottom of the last added widget if the grid is in a vertial
	// orientation.  Using a grid in this manner works similar to a gtk.Box,
	// but unlike gtk.Box, a gtk.Grid will respect its child widget's expand
	// and margin properties.
	grid.Add(btn)
	grid.Add(lab)
	grid.Add(entry)
	grid.Add(spnBtn)

	// Widgets may also be added by calling (*gtk.Grid).Attach() to specify
	// where to place the widget in the grid, and optionally how many rows
	// and columns to span over.
	//
	// Additional rows and columns are automatically added to the grid as
	// necessary when new widgets are added with (*gtk.Container).Add(), or,
	// as shown in this case, using (*gtk.Grid).Attach().
	//
	// In this case, a notebook is added beside the widgets inserted above.
	// The notebook widget is inserted with a left position of 1, a top
	// position of 1 (starting at the same vertical position as the button),
	// a width of 1 column, and a height of 2 rows (spanning down to the
	// same vertical position as the entry).
	//
	// This example also demonstrates how not every area of the grid must
	// contain a widget.  In particular, the area to the right of the label
	// and the right of spin button have contain no widgets.
	grid.Attach(nb, 1, 1, 1, 2)
	nb.SetHExpand(true)
	nb.SetVExpand(true)

	// Add a child widget and tab label to the notebook so it renders.
	nbChild, err := gtk.LabelNew("Notebook content")
	if err != nil {
		log.Fatal("Unable to create button:", err)
	}
	nbTab, err := gtk.LabelNew("Tab label")
	if err != nil {
		log.Fatal("Unable to create label:", err)
	}
	nb.AppendPage(nbChild, nbTab)

	// Add the grid to the window, and show all widgets.
	win.Add(grid)
	win.ShowAll()

	gtk.Main()
}
