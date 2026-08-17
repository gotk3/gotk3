// +build linux
// +build !no_x11

package gdk

// #include <gdk/gdk.h>
// #include <gdk/gdkx.h>
import "C"

func WorkspaceControlSupported() bool {
	return true
}

// GetScreenNumber is a wrapper around gdk_x11_screen_get_screen_number().
// It only works on GDK versions compiled with X11 support - its return value can't be used if WorkspaceControlSupported returns false
func (v *Screen) GetScreenNumber() int {
	return int(C.gdk_x11_screen_get_screen_number(v.native()))
}

