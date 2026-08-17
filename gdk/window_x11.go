// +build linux
// +build !no_x11

package gdk

// #include <gdk/gdk.h>
// #include <gdk/gdkx.h>
import "C"
import (
	"unsafe"

	"github.com/kozorork/gotk3/glib"
)

// MoveToCurrentDesktop is a wrapper around gdk_x11_window_move_to_current_desktop().
// It only works on GDK versions compiled with X11 support - its return value can't be used if WorkspaceControlSupported returns false
func (v *Window) MoveToCurrentDesktop() {
	C.gdk_x11_window_move_to_current_desktop(v.native())
}

// GetXID is a wrapper around gdk_x11_window_get_xid().
// It only works on GDK versions compiled with X11 support - its return value can't be used if WorkspaceControlSupported returns false
func (v *Window) GetXID() uint32 {
	return uint32(C.gdk_x11_window_get_xid(v.native()))
}

//ForeignNewForDisplay is a wrapper around gdk_x11_window_foreign_new_for_display()
// It only works on GDK versions compiled with X11 support - its return value can't be used if WorkspaceControlSupported returns false
func (v *Display) ForeignNewForDisplay(xid uint32) (*Window, error) {
	c := C.gdk_x11_window_foreign_new_for_display(v.native(), C.Window(xid))
	if c == nil {
		return nil, nilPtrErr
	}
	return &Window{glib.Take(unsafe.Pointer(c))}, nil
}
