// Same copyright and license as the rest of the files in this project

// +build !gtk_3_6,!gtk_3_8,!gtk_3_10,!gtk_3_12,!gtk_3_14

package gtk

// #include <gtk/gtk.h>
import "C"

// GetTitlebar is a wrapper around gtk_window_get_titlebar().
func (v *Window) GetTitlebar() (IWidget, error) {
	c := C.gtk_window_get_titlebar(v.native())
	if c == nil {
		return nil, nil
	}
	return castWidget(c)
}
