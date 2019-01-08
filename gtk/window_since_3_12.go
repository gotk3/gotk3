// +build !gtk_3_6,!gtk_3_8,!gtk_3_10

package gtk

// #include <gtk/gtk.h>
import "C"

// IsMaximized is a wrapper around gtk_window_is_maximized().
func (v *Window) IsMaximized() bool {
	c := C.gtk_window_is_maximized(v.native())
	return gobool(c)
}
