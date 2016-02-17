//+build gtk_3_6 gtk_3_8 gtk_3_10 gtk_3_12 gtk_3_14

package impl

// #cgo pkg-config: gtk+-3.0
// #include <gtk/gtk.h>
// #include <stdlib.h>
import "C"

import (
	"unsafe"

	gdk_impl "github.com/gotk3/gotk3/gdk/impl"
	"github.com/gotk3/gotk3/gtk"
)

// OverrideColor is a wrapper around gtk_widget_override_color().
func (v *widget) OverrideColor(state gtk.StateFlags, color *gdk_impl.RGBA) {
	var cColor *C.GdkRGBA
	if color != nil {
		cColor = (*C.GdkRGBA)(unsafe.Pointer(color.Native()))
	}
	C.gtk_widget_override_color(v.native(), C.GtkStateFlags(state), cColor)
}

// OverrideFont is a wrapper around gtk_widget_override_font().
func (v *widget) OverrideFont(description string) {
	cstr := C.CString(description)
	defer C.free(unsafe.Pointer(cstr))
	c := C.pango_font_description_from_string(cstr)
	C.gtk_widget_override_font(v.native(), c)
}