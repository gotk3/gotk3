// Same copyright and license as the rest of the files in this project

// +build !gtk_3_6,!gtk_3_8,!gtk_3_10,!gtk_3_12,!gtk_3_14,!gtk_3_16
// Supports building with gtk 3.18+

package gdk

// #include <gdk/gdk.h>
import "C"

/*
 * GdkKeymap
 */

// GetScrollLockState is a wrapper around gdk_keymap_get_scroll_lock_state().
func (v *Keymap) GetScrollLockState() bool {
	return gobool(C.gdk_keymap_get_scroll_lock_state(v.native()))
}

/*
 * GdkWindow
 */

// SetPassThrough is a wrapper around gdk_window_set_pass_through().
func (v *Window) SetPassThrough(passThrough bool) {
	C.gdk_window_set_pass_through(v.native(), gbool(passThrough))
}

// GetPassThrough is a wrapper around gdk_window_get_pass_through().
func (v *Window) GetPassThrough() bool {
	return gobool(C.gdk_window_get_pass_through(v.native()))
}
