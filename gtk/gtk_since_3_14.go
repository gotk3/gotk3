// +build !gtk_3_6,!gtk_3_8,!gtk_3_10,!gtk_3_12

package gtk

// #include <stdlib.h>
// #include <gtk/gtk.h>
import "C"
import (
	"unsafe"
	"github.com/gotk3/gotk3/glib"
)

/*
 * GtkListBox
 */

// UnselectRow is a wrapper around gtk_list_box_unselect_row().
func (v *ListBox) UnselectRow(row *ListBoxRow) {
	C.gtk_list_box_unselect_row(v.native(), row.native())
}

// SelectAll is a wrapper around gtk_list_box_select_all().
func (v *ListBox) SelectAll() {
	C.gtk_list_box_select_all(v.native())
}

// UnselectAll is a wrapper around gtk_list_box_unselect_all().
func (v *ListBox) UnselectAll() {
	C.gtk_list_box_unselect_all(v.native())
}

// TODO: gtk_list_box_selected_foreach()

// GetSelectedRows is a wrapper around gtk_list_box_get_selected_rows().
func (v *ListBox) GetSelectedRows() *glib.List {
	clist := C.gtk_list_box_get_selected_rows(v.native())
	if clist == nil {
		return nil
	}

	glist := glib.WrapList(uintptr(unsafe.Pointer(clist)))
	glist.DataWrapper(func(ptr unsafe.Pointer) interface{} {
		return wrapListBoxRow(glib.Take(ptr))
	})

	return glist
}

/*
 * GtkListBoxRow
 */

// IsSelected is a wrapper around gtk_list_box_row_is_selected().
func (v *ListBoxRow) IsSelected() bool {
	c := C.gtk_list_box_row_is_selected(v.native())
	return gobool(c)
}

// SetActivatable is a wrapper around gtk_list_box_row_set_activatable().
func (v *ListBoxRow) SetActivatable(activatable bool) {
	C.gtk_list_box_row_set_activatable(v.native(), gbool(activatable))
}

// GetActivatable is a wrapper around gtk_list_box_row_get_activatable().
func (v *ListBoxRow) GetActivatable() bool {
	c := C.gtk_list_box_row_get_activatable(v.native())
	return gobool(c)
}

// SetSelectable is a wrapper around gtk_list_box_row_set_selectable().
func (v *ListBoxRow) SetSelectable(selectable bool) {
	C.gtk_list_box_row_set_selectable(v.native(), gbool(selectable))
}

// GetSelectable is a wrapper around gtk_list_box_row_get_selectable().
func (v *ListBoxRow) GetSelectable() bool {
	c := C.gtk_list_box_row_get_selectable(v.native())
	return gobool(c)
}
