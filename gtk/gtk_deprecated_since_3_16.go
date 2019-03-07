//+build gtk_3_6 gtk_3_8 gtk_3_10 gtk_3_12 gtk_3_14

package gtk

// #include <gtk/gtk.h>
// #include <stdlib.h>
import "C"

import (
	"unsafe"

	"github.com/gotk3/gotk3/gdk"
)

// OverrideColor is a wrapper around gtk_widget_override_color().
func (v *Widget) OverrideColor(state StateFlags, color *gdk.RGBA) {
	var cColor *C.GdkRGBA
	if color != nil {
		cColor = (*C.GdkRGBA)(unsafe.Pointer(color.Native()))
	}
	C.gtk_widget_override_color(v.native(), C.GtkStateFlags(state), cColor)
}

func (v *Widget) OverrideBackgroundColor(state StateFlags, color *gdk.RGBA) {
	var cColor *C.GdkRGBA
	if color != nil {
		cColor = (*C.GdkRGBA)(unsafe.Pointer(color.Native()))
	}
	C.gtk_widget_override_background_color(v.native(), C.GtkStateFlags(state), cColor)
}

func (v *Button) SetColor(color string) {
	rgba := C.GdkRGBA{}
	C.gdk_rgba_parse(&rgba, (*C.gchar)(C.CString(color)))
	C.gtk_widget_override_background_color(v.toWidget(), C.GTK_STATE_FLAG_NORMAL, &rgba)
}

// OverrideFont is a wrapper around gtk_widget_override_font().
func (v *Widget) OverrideFont(description string) {
	cstr := C.CString(description)
	defer C.free(unsafe.Pointer(cstr))
	c := C.pango_font_description_from_string(cstr)
	C.gtk_widget_override_font(v.native(), c)
}

func (v *Label) SetFont(font string) {
	v.OverrideFont(font)
}
