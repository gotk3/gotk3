// +build linux
// +build !no_x11
// +build !gtk_3_6,!gtk_3_8

package gdk

// #include <gdk/gdk.h>
// #include <gdk/gdkx.h>
import "C"

// GetNumberOfDesktops is a wrapper around gdk_x11_screen_get_number_of_desktops().
// It only works on GDK versions compiled with X11 support - its return value can't be used if WorkspaceControlSupported returns false
func (v *Screen) GetNumberOfDesktops() uint32 {
	return uint32(C.gdk_x11_screen_get_number_of_desktops(v.native()))
}

// GetCurrentDesktop is a wrapper around gdk_x11_screen_get_current_desktop().
// It only works on GDK versions compiled with X11 support - its return value can't be used if WorkspaceControlSupported returns false
func (v *Screen) GetCurrentDesktop() uint32 {
	return uint32(C.gdk_x11_screen_get_current_desktop(v.native()))
}
