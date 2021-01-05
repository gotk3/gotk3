// Same copyright and license as the rest of the files in this project

// +build !gtk_3_6,!gtk_3_8,!gtk_3_10,!gtk_3_12,!gtk_3_14,!gtk_3_16

package gtk

// #include <gtk/gtk.h>
// #include "gtk.go.h"
import "C"
import (
	"unsafe"

	"github.com/gotk3/gotk3/gdk"
)

// FullscreenOnMonitor is a wrapper around gtk_window_fullscreen_on_monitor().
func (v *Window) FullscreenOnMonitor(screen *gdk.Screen, monitor int) {
	C.gtk_window_fullscreen_on_monitor(v.native(), C.toGdkScreen(unsafe.Pointer(screen.Native())), C.gint(monitor))
}
