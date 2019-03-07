// Same copyright and license as the rest of the files in this project
// This file contains accelerator related functions and structures

package gtk

// #include <gtk/gtk.h>
// #include "gtk.go.h"
import "C"
import (
	"unsafe"

	"github.com/gotk3/gotk3/glib"
)

/*
 * GtkMenu
 */

// Menu is a representation of GTK's GtkMenu.
type Menu struct {
	MenuShell
}

// IMenu is an interface type implemented by all structs embedding
// a Menu.  It is meant to be used as an argument type for wrapper
// functions that wrap around a C GTK function taking a
// GtkMenu.
type IMenu interface {
	toMenu() *C.GtkMenu
	toWidget() *C.GtkWidget
}

// native() returns a pointer to the underlying GtkMenu.
func (v *Menu) native() *C.GtkMenu {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkMenu(p)
}

func (v *Menu) toMenu() *C.GtkMenu {
	if v == nil {
		return nil
	}
	return v.native()
}

func marshalMenu(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := glib.Take(unsafe.Pointer(c))
	return wrapMenu(obj), nil
}

func wrapMenu(obj *glib.Object) *Menu {
	return &Menu{MenuShell{Container{Widget{glib.InitiallyUnowned{obj}}}}}
}

// MenuNew() is a wrapper around gtk_menu_new().
func MenuNew() (*Menu, error) {
	c := C.gtk_menu_new()
	if c == nil {
		return nil, nilPtrErr
	}
	return wrapMenu(glib.Take(unsafe.Pointer(c))), nil
}

// Popdown() is a wrapper around gtk_menu_popdown().
func (v *Menu) Popdown() {
	C.gtk_menu_popdown(v.native())
}

// ReorderChild() is a wrapper around gtk_menu_reorder_child().
func (v *Menu) ReorderChild(child IWidget, position int) {
	C.gtk_menu_reorder_child(v.native(), child.toWidget(), C.gint(position))
}
