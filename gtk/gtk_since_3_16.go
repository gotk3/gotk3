// +build !gtk_3_6,!gtk_3_8,!gtk_3_10,!gtk_3_12,!gtk_3_14

// See: https://developer.gnome.org/gtk3/3.16/api-index-3-16.html

package gtk

// #cgo pkg-config: gtk+-3.0
// #include <gtk/gtk.h>
import "C"

// SetOverlayScrolling is a wrapper around gtk_scrolled_window_set_overlay_scrolling().
func (v *ScrolledWindow) SetOverlayScrolling(scrolling bool) {
	C.gtk_scrolled_window_set_overlay_scrolling(v.native(), gbool(scrolling))
}

// GetOverlayScrolling is a wrapper around gtk_scrolled_window_get_overlay_scrolling().
func (v *ScrolledWindow) GetOverlayScrolling() bool {
	return gobool(C.gtk_scrolled_window_get_overlay_scrolling(v.native()))
}

// SetWideHandle is a wrapper around gtk_paned_set_wide_handle().
func (v *Paned) SetWideHandle(wide bool) {
	C.gtk_paned_set_wide_handle(v.native(), gbool(wide))
}

// GetWideHandle is a wrapper around gtk_paned_get_wide_handle().
func (v *Paned) GetWideHandle() bool {
	return gobool(C.gtk_paned_get_wide_handle(v.native()))
}

// GetXAlign is a wrapper around gtk_label_get_xalign().
func (v *Label) GetXAlign() float64 {
	c := C.gtk_label_get_xalign(v.native())
	return float64(c)
}

// GetYAlign is a wrapper around gtk_label_get_yalign().
func (v *Label) GetYAlign() float64 {
	c := C.gtk_label_get_yalign(v.native())
	return float64(c)
}

// SetXAlign is a wrapper around gtk_label_set_xalign().
func (v *Label) SetXAlign(n float64) {
	C.gtk_label_set_xalign(v.native(), C.gfloat(n))
}

// SetYAlign is a wrapper around gtk_label_set_yalign().
func (v *Label) SetYAlign(n float64) {
	C.gtk_label_set_yalign(v.native(), C.gfloat(n))
}
