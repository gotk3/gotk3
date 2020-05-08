// Same copyright and license as the rest of the files in this project
// See: https://developer.gnome.org/gtk3/3.14/api-index-3-14.html

// +build !gtk_3_6,!gtk_3_8,!gtk_3_10,!gtk_3_12

package gtk

// #include <gtk/gtk.h>
// #include "gtk.go.h"
import "C"
import "unsafe"

/*
 * GtkCellArea
 */

// AttributeGetColumn is a wrapper around gtk_cell_area_attribute_get_column().
func (v *CellArea) AttributeGetColumn(renderer ICellRenderer, attribute string) int {
	cstr := C.CString(attribute)
	defer C.free(unsafe.Pointer(cstr))
	column := C.gtk_cell_area_attribute_get_column(v.native(), renderer.toCellRenderer(), (*C.gchar)(cstr))
	return int(column)
}
