// Same copyright and license as the rest of the files in this project
// This file contains accelerator related functions and structures

// +build !gtk_3_6,!gtk_3_8
// not use this: go build -tags gtk_3_8'. Otherwise, if no build tags are used, GTK 3.10

package gtkf

// #cgo pkg-config: gtk+-3.0
// #include <stdlib.h>
// #include <gtk/gtk.h>
// #include "gtk_since_3_10.go.h"
import "C"
import "github.com/gotk3/gotk3/gtk"

/*
 * GtkWindow
 */

// SetTitlebar is a wrapper around gtk_window_set_titlebar().
func (v *window) SetTitlebar(titlebar gtk.Widget) {
	C.gtk_window_set_titlebar(v.native(), castToWidget(titlebar))
}

// Close is a wrapper around gtk_window_close().
func (v *window) Close() {
	C.gtk_window_close(v.native())
}
