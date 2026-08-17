// +build linux
// +build !no_x11
// +build !gtk_3_6,!gtk_3_8

package gdk

// #include <gdk/gdk.h>
// #include <gdk/gdkx.h>
import "C"
import (
	"unsafe"

	"github.com/kozorork/gotk3/glib"
)

// GetDesktop is a wrapper around gdk_x11_window_get_desktop().
// It only works on GDK versions compiled with X11 support - its return value can't be used if WorkspaceControlSupported returns false
func (v *Window) GetDesktop() uint32 {
	return uint32(C.gdk_x11_window_get_desktop(v.native()))
}

// MoveToDesktop is a wrapper around gdk_x11_window_move_to_desktop().
// It only works on GDK versions compiled with X11 support - its return value can't be used if WorkspaceControlSupported returns false
func (v *Window) MoveToDesktop(d uint32) {
	C.gdk_x11_window_move_to_desktop(v.native(), C.guint32(d))
}

