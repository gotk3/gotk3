// +build !gtk_3_6,!gtk_3_8,!gtk_3_10,!gtk_3_12,!gtk_3_14,!gtk_3_16,!gtk_3_18,!gtk_3_20

// See: https://developer.gnome.org/gtk3/3.22/api-index-3-22.html

package gtk

// #include <gtk/gtk.h>
import "C"

// Popup is a wrapper around gtk_popover_popup().
func (v *Popover) Popup() {
	C.gtk_popover_popup(v.native())
}

// Popdown is a wrapper around gtk_popover_popdown().
func (v *Popover) Popdown() {
	C.gtk_popover_popdown(v.native())
}

/*
 * GtkScrolledWindow
 */

// GetMaxContentWidth is a wrapper around gtk_scrolled_window_get_max_content_width().
func (v *ScrolledWindow) GetMaxContentWidth() int {
	c := C.gtk_scrolled_window_get_max_content_width(v.native())
	return int(c)
}

// SetMaxContentWidth is a wrapper around gtk_scrolled_window_set_max_content_width().
func (v *ScrolledWindow) SetMaxContentWidth(width int) {
	C.gtk_scrolled_window_set_max_content_width(v.native(), C.gint(width))
}

// GetMaxContentHeight is a wrapper around gtk_scrolled_window_get_max_content_height().
func (v *ScrolledWindow) GetMaxContentHeight() int {
	c := C.gtk_scrolled_window_get_max_content_height(v.native())
	return int(c)
}

// SetMaxContentHeight is a wrapper around gtk_scrolled_window_set_max_content_height().
func (v *ScrolledWindow) SetMaxContentHeight(width int) {
	C.gtk_scrolled_window_set_max_content_height(v.native(), C.gint(width))
}

// GetPropagateNaturalWidth is a wrapper around gtk_scrolled_window_get_propagate_natural_width().
func (v *ScrolledWindow) GetPropagateNaturalWidth() bool {
	c := C.gtk_scrolled_window_get_propagate_natural_width(v.native())
	return gobool(c)
}

// SetPropagateNaturalWidth is a wrapper around gtk_scrolled_window_set_propagate_natural_width().
func (v *ScrolledWindow) SetPropagateNaturalWidth(propagate bool) {
	C.gtk_scrolled_window_set_propagate_natural_width(v.native(), gbool(propagate))
}

// GetPropagateNaturalHeight is a wrapper around gtk_scrolled_window_get_propagate_natural_height().
func (v *ScrolledWindow) GetPropagateNaturalHeight() bool {
	c := C.gtk_scrolled_window_get_propagate_natural_height(v.native())
	return gobool(c)
}

// SetPropagateNaturalHeight is a wrapper around gtk_scrolled_window_set_propagate_natural_height().
func (v *ScrolledWindow) SetPropagateNaturalHeight(propagate bool) {
	C.gtk_scrolled_window_set_propagate_natural_height(v.native(), gbool(propagate))
}
