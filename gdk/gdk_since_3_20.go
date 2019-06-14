// +build !gtk_3_6,!gtk_3_8,!gtk_3_10,!gtk_3_12,!gtk_3_14,!gtk_3_16,!gtk_3_18
// Supports building with gtk 3.20+

package gdk

// #include <gdk/gdk.h>
import "C"

// IsLegacy is a wrapper around gdk_gl_context_is_legacy().
func (v *GLContext) IsLegacy() bool {
	return gobool(C.gdk_gl_context_is_legacy(v.native()))
}
