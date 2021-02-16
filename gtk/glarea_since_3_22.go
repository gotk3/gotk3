// +build !gtk_3_6,!gtk_3_8,!gtk_3_10,!gtk_3_12,!gtk_3_14,!gtk_3_16,!gtk_3_18,!gtk_3_20
// Supports building with gtk 3.22+

package gtk

// #include <gtk/gtk.h>
import "C"

// GetUseES is a wrapper around gtk_gl_area_get_use_es().
func (v *GLArea) GetUseES() bool {
	return gobool(C.gtk_gl_area_get_use_es(v.native()))
}

// SetUseES is a wrapper around gtk_gl_area_set_use_es().
func (v *GLArea) SetUseES(es bool) {
	C.gtk_gl_area_set_use_es(v.native(), gbool(es))
}
