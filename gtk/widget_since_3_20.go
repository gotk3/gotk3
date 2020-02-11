// +build !gtk_3_6,!gtk_3_8,!gtk_3_10,!gtk_3_12,!gtk_3_14,!gtk_3_16,!gtk_3_18

package gtk

// #include <gtk/gtk.h>
import "C"

// GetFocusOnClick is a wrapper around gtk_widget_get_focus_on_click().
func (v *Widget) GetFocusOnClick() bool {
	c := C.gtk_widget_get_focus_on_click(v.native())
	return gobool(c)
}

// SetFocusOnClick is a wrapper around gtk_widget_set_focus_on_click().
func (v *Widget) SetFocusOnClick(focusOnClick bool) {
	C.gtk_widget_set_focus_on_click(v.native(), gbool(focusOnClick))
}

// TODO:
// gtk_widget_class_get_css_name().
// gtk_widget_class_set_css_name().
