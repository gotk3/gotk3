// Same copyright and license as the rest of the files in this project
// This file contains style related functions and structures

package impl

// #include <gtk/gtk.h>
// #include "gtk.go.h"
import "C"
import (
	"unsafe"

	glib_impl "github.com/gotk3/gotk3/glib/impl"
)

/*
 * GtkApplicationWindow
 */

// ApplicationWindow is a representation of GTK's GtkApplicationWindow.
type ApplicationWindow struct {
	Window
}

// native returns a pointer to the underlying GtkApplicationWindow.
func (v *ApplicationWindow) native() *C.GtkApplicationWindow {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkApplicationWindow(p)
}

func marshalApplicationWindow(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := wrapObject(unsafe.Pointer(c))
	return wrapApplicationWindow(obj), nil
}

func wrapApplicationWindow(obj *glib_impl.Object) *ApplicationWindow {
	return &ApplicationWindow{Window{Bin{Container{Widget{glib_impl.InitiallyUnowned{obj}}}}}}
}

// ApplicationWindowNew is a wrapper around gtk_application_window_new().
func ApplicationWindowNew(app *Application) (*ApplicationWindow, error) {
	c := C.gtk_application_window_new(app.native())
	if c == nil {
		return nil, nilPtrErr
	}
	return wrapApplicationWindow(wrapObject(unsafe.Pointer(c))), nil
}

// SetShowMenubar is a wrapper around gtk_application_window_set_show_menubar().
func (v *ApplicationWindow) SetShowMenubar(b bool) {
	C.gtk_application_window_set_show_menubar(v.native(), gbool(b))
}

// GetShowMenubar is a wrapper around gtk_application_window_get_show_menubar().
func (v *ApplicationWindow) GetShowMenubar() bool {
	return gobool(C.gtk_application_window_get_show_menubar(v.native()))
}

// GetID is a wrapper around gtk_application_window_get_id().
func (v *ApplicationWindow) GetID() uint {
	return uint(C.gtk_application_window_get_id(v.native()))
}
