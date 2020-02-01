// Same copyright and license as the rest of the files in this project

// +build !gtk_3_6,!gtk_3_8,!gtk_3_10,!gtk_3_12,!gtk_3_14

package gtk

// #include <gtk/gtk.h>
import "C"
import (
	"unsafe"

	"github.com/gotk3/gotk3/glib"
)

// GetTitlebar is a wrapper around gtk_window_get_titlebar().
// TODO: Use IWidget here
func (v *Window) GetTitlebar() *Widget {
	c := C.gtk_window_get_titlebar(v.native())
	if c == nil {
		return nil
	}
	obj := glib.Take(unsafe.Pointer(c))
	return wrapWidget(obj)
}
