package gtk

// #include <gtk/gtk.h>
// #include "gtk.go.h"
import "C"
import (
	"unsafe"

	"github.com/gotk3/gotk3/glib"
)

func init() {
	tm := []glib.TypeMarshaler{
		{glib.Type(C.gtk_shortcuts_window_get_type()), marshalShortcutsWindow},
	}

	glib.RegisterGValueMarshalers(tm)

	WrapMap["GtkShortcutsWindow"] = wrapShortcutsWindow
}

/*
 * GtkShortcutsWindow
 */

// ShortcutsWindow is a representation of GTK's GtkShortcutsWindow.
type ShortcutsWindow struct {
	Window
}

func (v *ShortcutsWindow) native() *C.GtkShortcutsWindow {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGtkShortcutsWindow(p)
}

func marshalShortcutsWindow(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := glib.Take(unsafe.Pointer(c))
	return wrapShortcutsWindow(obj), nil
}

func wrapShortcutsWindow(obj *glib.Object) *ShortcutsWindow {
	return &ShortcutsWindow{Window{Bin{Container{Widget{glib.InitiallyUnowned{obj}}}}}}
}
