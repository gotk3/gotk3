// Same copyright and license as the rest of the files in this project

// +build !gtk_3_6,!gtk_3_8,!gtk_3_10,!gtk_3_12

package gtk

// #include <gtk/gtk.h>
import "C"

// SetInteractiveDebugging is a wrapper around gtk_window_set_interactive_debugging().
func SetInteractiveDebugging(enable bool) {
	C.gtk_window_set_interactive_debugging(gbool(enable))
}
