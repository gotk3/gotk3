//+build gtk_3_6 gtk_3_8 gtk_3_10 gtk_3_12 gtk_3_14

package gtk

// #cgo pkg-config: gtk+-3.0
// #include <gtk/gtk.h>
// #include <stdlib.h>
import "C"

import (
	"github.com/gotk3/gotk3/gdk"
	"unsafe"
)

// OverrideColor is a wrapper around gtk_widget_override_color().
func (v *Widget) OverrideColor(state StateFlags, color *gdk.RGBA) {
	var cColor *C.GdkRGBA
	if color != nil {
		cColor = (*C.GdkRGBA)(unsafe.Pointer((&color.RGBA)))
	}
	C.gtk_widget_override_color(v.native(), C.GtkStateFlags(state), cColor)
}

// OverrideFont is a wrapper around gtk_widget_override_font().
func (v *Widget) OverrideFont(description string) {
	cstr := C.CString(description)
	defer C.free(unsafe.Pointer(cstr))
	c := C.pango_font_description_from_string(cstr)
	C.gtk_widget_override_font(v.native(), c)
}

// SupportsComposite() is a wrapper around gdk_display_supports_composite().
func (v *Display) SupportsComposite() bool {
	c := C.gdk_display_supports_composite(v.native())
	return gobool(c)
}
