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
	"fmt"
	"github.com/conformal/gotk3/glib"
	"github.com/conformal/gotk3/gtk"
	"log"
	"time"
)

var (
	topLabel    *gtk.Label
	bottomLabel *gtk.Label
	nSets       = 1
)

func main() {
	gtk.Init(nil)

	win, err := gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
	if err != nil {
		log.Fatal("Unable to create window:", err)
	}
	win.Connect("destroy", func() {
		gtk.MainQuit()
	})

	win.Add(windowWidget())

	// Native GTK is not thread safe, and thus, gotk3's GTK bindings may not
	// be used from other goroutines.  Instead, glib.IdleAdd() must be used
	// to add a function to run in the GTK main loop when it is in an idle
	// state.
	//
	// Two examples of using glib.IdleAdd() are shown below.  The first runs
	// a user created function, LabelSetTextIdle, and passes it two
	// arguments for a label and the text to set it with.  The second calls
	// (*gtk.Label).SetText directly, passing in only the text as an
	// argument.
	//
	// If the function passed to glib.IdleAdd() returns one argument, and
	// that argument is a bool, this return value will be used in the same
	// manner as a native g_idle_add() call.  If this return value is false,
	// the function will be removed from executing in the GTK main loop's
	// idle state.  If the return value is true, the function will continue
	// to execute when the GTK main loop is in this state.
	go func() {
		for {
			time.Sleep(time.Second)
			s := fmt.Sprintf("Set a label %d time(s)!", nSets)
			_, err := glib.IdleAdd(LabelSetTextIdle, topLabel, s)
			if err != nil {
				log.Fatal("IdleAdd() failed:", err)
			}
			nSets++
			s = fmt.Sprintf("Set a label %d time(s)!", nSets)
			_, err = glib.IdleAdd(bottomLabel.SetText, s)
			if err != nil {
				log.Fatal("IdleAdd() failed:", err)
			}
			nSets++
		}
	}()

	win.ShowAll()
	gtk.Main()
}

func windowWidget() *gtk.Widget {
	grid, err := gtk.GridNew()
	if err != nil {
		log.Fatal("Unable to create grid:", err)
	}
	grid.SetOrientation(gtk.ORIENTATION_VERTICAL)

	topLabel, err = gtk.LabelNew("Text set by initializer")
	if err != nil {
		log.Fatal("Unable to create label:", err)
	}
	bottomLabel, err = gtk.LabelNew("Text set by initializer")
	if err != nil {
		log.Fatal("Unable to create label:", err)
	}

	grid.Add(topLabel)
	grid.Add(bottomLabel)
	topLabel.SetHExpand(true)
	topLabel.SetVExpand(true)
	bottomLabel.SetHExpand(true)
	bottomLabel.SetVExpand(true)

	return &grid.Container.Widget
}

func LabelSetTextIdle(label *gtk.Label, text string) bool {
	label.SetText(text)

	// Returning false here is unnecessary, as anything but returning true
	// will remove the function from being called by the GTK main loop.
	return false
}
