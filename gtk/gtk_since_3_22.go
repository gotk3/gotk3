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

// TODO:
// gtk_scrolled_window_get_max_content_width().
// gtk_scrolled_window_set_max_content_width().
// gtk_scrolled_window_get_max_content_height().
// gtk_scrolled_window_set_max_content_height().
// gtk_scrolled_window_get_propagate_natural_width().
// gtk_scrolled_window_set_propagate_natural_width().
// gtk_scrolled_window_get_propagate_natural_height().
// gtk_scrolled_window_set_propagate_natural_height().
