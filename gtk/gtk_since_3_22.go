// +build !gtk_3_6,!gtk_3_8,!gtk_3_10,!gtk_3_12,!gtk_3_14,!gtk_3_16,!gtk_3_18,!gtk_3_20

// See: https://developer.gnome.org/gtk3/3.22/api-index-3-22.html

package gtk

// #include <gtk/gtk.h>
// #include "gtk.go.h"
import "C"
import (
	"unsafe"
)

// Popup is a wrapper around gtk_popover_popup().
func (v *Popover) Popup() {
	C.gtk_popover_popup(v.native())
}

// Popdown is a wrapper around gtk_popover_popdown().
func (v *Popover) Popdown() {
	C.gtk_popover_popdown(v.native())
}

/*
 * GtkFileChooser
 */

// AddChoice is a wrapper around gtk_file_chooser_add_choice().
func (v *FileChooser) AddChoice(id, label string, options, optionLabels []string) {
	cId := C.CString(id)
	defer C.free(unsafe.Pointer(cId))

	cLabel := C.CString(label)
	defer C.free(unsafe.Pointer(cLabel))

	if options == nil || optionLabels == nil {
		C.gtk_file_chooser_add_choice(v.native(), (*C.gchar)(cId), (*C.gchar)(cLabel), nil, nil)
		return
	}

	cOptions := C.make_strings(C.int(len(options) + 1))
	for i, option := range options {
		cstr := C.CString(option)
		defer C.free(unsafe.Pointer(cstr))
		C.set_string(cOptions, C.int(i), (*C.gchar)(cstr))
	}
	C.set_string(cOptions, C.int(len(options)), nil)

	cOptionLabels := C.make_strings(C.int(len(optionLabels) + 1))
	for i, optionLabel := range optionLabels {
		cstr := C.CString(optionLabel)
		defer C.free(unsafe.Pointer(cstr))
		C.set_string(cOptionLabels, C.int(i), (*C.gchar)(cstr))
	}
	C.set_string(cOptionLabels, C.int(len(optionLabels)), nil)

	C.gtk_file_chooser_add_choice(v.native(), (*C.gchar)(cId), (*C.gchar)(cLabel), cOptions, cOptionLabels)
}

// RemoveChoice is a wrapper around gtk_file_chooser_remove_choice().
func (v *FileChooser) RemoveChoice(id string) {
	cId := C.CString(id)
	defer C.free(unsafe.Pointer(cId))
	C.gtk_file_chooser_remove_choice(v.native(), (*C.gchar)(cId))
}

// SetChoice is a wrapper around gtk_file_chooser_set_choice().
func (v *FileChooser) SetChoice(id, option string) {
	cId := C.CString(id)
	defer C.free(unsafe.Pointer(cId))
	cOption := C.CString(option)
	defer C.free(unsafe.Pointer(cOption))
	C.gtk_file_chooser_set_choice(v.native(), (*C.gchar)(cId), (*C.gchar)(cOption))
}

// GetChoice is a wrapper around gtk_file_chooser_get_choice().
func (v *FileChooser) GetChoice(id string) string {
	cId := C.CString(id)
	defer C.free(unsafe.Pointer(cId))
	c := C.gtk_file_chooser_get_choice(v.native(), (*C.gchar)(cId))
	return C.GoString(c)
}

/*
 * GtkScrolledWindow
 */

// GetMaxContentWidth is a wrapper around gtk_scrolled_window_get_max_content_width().
func (v *ScrolledWindow) GetMaxContentWidth() int {
	c := C.gtk_scrolled_window_get_max_content_width(v.native())
	return int(c)
}

// SetMaxContentWidth is a wrapper around gtk_scrolled_window_set_max_content_width().
func (v *ScrolledWindow) SetMaxContentWidth(width int) {
	C.gtk_scrolled_window_set_max_content_width(v.native(), C.gint(width))
}

// GetMaxContentHeight is a wrapper around gtk_scrolled_window_get_max_content_height().
func (v *ScrolledWindow) GetMaxContentHeight() int {
	c := C.gtk_scrolled_window_get_max_content_height(v.native())
	return int(c)
}

// SetMaxContentHeight is a wrapper around gtk_scrolled_window_set_max_content_height().
func (v *ScrolledWindow) SetMaxContentHeight(width int) {
	C.gtk_scrolled_window_set_max_content_height(v.native(), C.gint(width))
}

// GetPropagateNaturalWidth is a wrapper around gtk_scrolled_window_get_propagate_natural_width().
func (v *ScrolledWindow) GetPropagateNaturalWidth() bool {
	c := C.gtk_scrolled_window_get_propagate_natural_width(v.native())
	return gobool(c)
}

// SetPropagateNaturalWidth is a wrapper around gtk_scrolled_window_set_propagate_natural_width().
func (v *ScrolledWindow) SetPropagateNaturalWidth(propagate bool) {
	C.gtk_scrolled_window_set_propagate_natural_width(v.native(), gbool(propagate))
}

// GetPropagateNaturalHeight is a wrapper around gtk_scrolled_window_get_propagate_natural_height().
func (v *ScrolledWindow) GetPropagateNaturalHeight() bool {
	c := C.gtk_scrolled_window_get_propagate_natural_height(v.native())
	return gobool(c)
}

// SetPropagateNaturalHeight is a wrapper around gtk_scrolled_window_set_propagate_natural_height().
func (v *ScrolledWindow) SetPropagateNaturalHeight(propagate bool) {
	C.gtk_scrolled_window_set_propagate_natural_height(v.native(), gbool(propagate))
}
