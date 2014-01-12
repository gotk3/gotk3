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
	"container/list"
	"fmt"
	"github.com/conformal/gotk3/gtk"
	"log"
)

var labelList = list.New()

func main() {
	gtk.Init(nil)

	win, err := gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
	if err != nil {
		log.Fatal("Unable to create window:", err)
	}
	win.SetTitle("Add/Remove Widgets Example")
	win.Connect("destroy", func() {
		gtk.MainQuit()
	})

	win.Add(windowWidget())
	win.ShowAll()

	gtk.Main()
}

func windowWidget() *gtk.Widget {
	grid, err := gtk.GridNew()
	if err != nil {
		log.Fatal("Unable to create grid:", err)
	}
	grid.SetOrientation(gtk.ORIENTATION_VERTICAL)

	// Just as a demonstration, we create and destroy a Label without ever
	// adding it to a container.  In native GTK, this would result in a
	// memory leak, since gtk_widget_destroy() will not deallocate any
	// memory when passed a GtkWidget with a floating reference.
	//
	// gotk3 handles this situation by always sinking floating references
	// of any struct type embedding a glib.InitiallyUnowned, and by setting
	// a finalizer to unreference the object when Go has lost scope of the
	// variable.  Due to this design, widgets may be allocated freely
	// without worrying about handling memory incorrectly.
	//
	// The following code is not entirely useful (except to demonstrate
	// this point), but it is also not "incorrect" as the C equivalent
	// would be.
	unused, err := gtk.LabelNew("This label is never used")
	if err != nil {
		// Calling Destroy() is also unnecessary in this case.  The
		// memory will still be freed with or without calling it.
		unused.Destroy()
	}

	sw, err := gtk.ScrolledWindowNew(nil, nil)
	if err != nil {
		log.Fatal("Unable to create scrolled window:", err)
	}

	grid.Attach(sw, 0, 0, 2, 1)
	sw.SetHExpand(true)
	sw.SetVExpand(true)

	labelsGrid, err := gtk.GridNew()
	if err != nil {
		log.Fatal("Unable to create grid:", err)
	}
	labelsGrid.SetOrientation(gtk.ORIENTATION_VERTICAL)

	sw.Add(labelsGrid)
	labelsGrid.SetHExpand(true)

	insertBtn, err := gtk.ButtonNewWithLabel("Add a label")
	if err != nil {
		log.Fatal("Unable to create button:", err)
	}
	removeBtn, err := gtk.ButtonNewWithLabel("Remove a label")
	if err != nil {
		log.Fatal("Unable to create button:", err)
	}

	nLabels := 1
	insertBtn.Connect("clicked", func() {
		var s string
		if nLabels == 1 {
			s = fmt.Sprintf("Inserted %d label.", nLabels)
		} else {
			s = fmt.Sprintf("Inserted %d labels.", nLabels)
		}
		label, err := gtk.LabelNew(s)
		if err != nil {
			log.Print("Unable to create label:", err)
			return
		}

		labelList.PushBack(label)
		labelsGrid.Add(label)
		label.SetHExpand(true)
		labelsGrid.ShowAll()

		nLabels++
	})

	removeBtn.Connect("clicked", func() {
		e := labelList.Front()
		if e == nil {
			log.Print("Nothing to remove")
			return
		}
		lab, ok := labelList.Remove(e).(*gtk.Label)
		if !ok {
			log.Print("Element to remove is not a *gtk.Label")
			return
		}
		// (*Widget).Destroy() breaks this label's reference with all
		// other objects (in this case, the Grid container it was added
		// to).
		lab.Destroy()

		// At this point, only Go retains a reference to the GtkLabel.
		// When the lab variable goes out of scope when this function
		// returns, at the next garbage collector run, a finalizer will
		// be run to perform the final unreference and free the widget.
	})

	grid.Attach(insertBtn, 0, 1, 1, 1)
	grid.Attach(removeBtn, 1, 1, 1, 1)

	return &grid.Container.Widget
}
