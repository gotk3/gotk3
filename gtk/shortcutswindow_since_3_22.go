// +build !gtk_3_6,!gtk_3_8,!gtk_3_10,!gtk_3_12,!gtk_3_14,!gtk_3_16,!gtk_3_18,!gtk_3_20

package gtk

// #include <gtk/gtk.h>
// #include "gtk.go.h"
// #include "shortcutswindow_since_3_22.go.h"
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
