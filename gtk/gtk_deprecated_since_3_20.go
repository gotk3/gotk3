//+build gtk_3_6 gtk_3_8 gtk_3_10 gtk_3_12 gtk_3_14 gtk_3_16 gtk_3_18 gtk_deprecated

package gtk

// #include <gtk/gtk.h>
// #include <stdlib.h>
import "C"
import (
	"unsafe"
)

/*
 * GtkRange
 */

// TODO:
// gtk_range_get_min_slider_size().
// gtk_range_set_min_slider_size().

/*
 * GtkToolbar
 */

// TODO:
// GtkToolbarSpaceStyle

/*
 * GtkFileChooserButton
 */

// GetFocusOnClick is a wrapper around gtk_file_chooser_button_get_focus_on_click().
func (v *FileChooserButton) GetFocusOnClick() bool {
	return gobool(C.gtk_file_chooser_button_get_focus_on_click(v.native()))
}

// SetFocusOnClick is a wrapper around gtk_file_chooser_button_set_focus_on_click().
func (v *FileChooserButton) SetFocusOnClick(grabFocus bool) {
	C.gtk_file_chooser_button_set_focus_on_click(v.native(), gbool(grabFocus))
}

/*
 * GtkButton
 */

// GetFocusOnClick is a wrapper around gtk_button_get_focus_on_click().
func (v *Button) GetFocusOnClick() bool {
	c := C.gtk_button_get_focus_on_click(v.native())
	return gobool(c)
}

// SetFocusOnClick is a wrapper around gtk_button_set_focus_on_click().
func (v *Button) SetFocusOnClick(focusOnClick bool) {
	C.gtk_button_set_focus_on_click(v.native(), gbool(focusOnClick))
}

/*
 * GtkTextIter
 */

// BeginsTag is a wrapper around gtk_text_iter_begins_tag().
func (v *TextIter) BeginsTag(v1 *TextTag) bool {
	return gobool(C.gtk_text_iter_begins_tag(v.native(), v1.native()))
}

/*
 * GtkWindow
 */

// ParseGeometry is a wrapper around gtk_window_parse_geometry().
func (v *Window) ParseGeometry(geometry string) bool {
	cstr := C.CString(geometry)
	defer C.free(unsafe.Pointer(cstr))
	c := C.gtk_window_parse_geometry(v.native(), (*C.gchar)(cstr))
	return gobool(c)
}

// ResizeToGeometry is a wrapper around gtk_window_resize_to_geometry().
func (v *Window) ResizeToGeometry(width, height int) {
	C.gtk_window_resize_to_geometry(v.native(), C.gint(width), C.gint(height))
}

// SetDefaultGeometry is a wrapper around gtk_window_set_default_geometry().
func (v *Window) SetDefaultGeometry(width, height int) {
	C.gtk_window_set_default_geometry(v.native(), C.gint(width),
		C.gint(height))
}
