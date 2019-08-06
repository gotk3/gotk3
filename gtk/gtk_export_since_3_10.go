// +build !gtk_3_6,!gtk_3_8

package gtk

// #include <gtk/gtk.h>
import "C"
import (
	"unsafe"

	"github.com/gotk3/gotk3/glib"
)

// TODO: figure out a way to determine when we can clean up

//export goListBoxFilterFuncs
func goListBoxFilterFuncs(row *C.GtkListBoxRow, userData C.gpointer) C.gboolean {
	id := int(uintptr(userData))

	listBoxFilterFuncRegistry.Lock()
	r := listBoxFilterFuncRegistry.m[id]
	//delete(listBoxFilterFuncRegistry.m, id)
	listBoxFilterFuncRegistry.Unlock()

	return gbool(r.fn(wrapListBoxRow(glib.Take(unsafe.Pointer(row))), r.userData))
}

//export goListBoxHeaderFuncs
func goListBoxHeaderFuncs(row *C.GtkListBoxRow, before *C.GtkListBoxRow, userData C.gpointer) {
	id := int(uintptr(userData))

	listBoxHeaderFuncRegistry.Lock()
	r := listBoxHeaderFuncRegistry.m[id]
	//delete(listBoxHeaderFuncRegistry.m, id)
	listBoxHeaderFuncRegistry.Unlock()

	r.fn(wrapListBoxRow(glib.Take(unsafe.Pointer(row))), wrapListBoxRow(glib.Take(unsafe.Pointer(before))), r.userData)
}

//export goListBoxSortFuncs
func goListBoxSortFuncs(row1 *C.GtkListBoxRow, row2 *C.GtkListBoxRow, userData C.gpointer) C.gint {
	id := int(uintptr(userData))

	listBoxSortFuncRegistry.Lock()
	r := listBoxSortFuncRegistry.m[id]
	//delete(listBoxSortFuncRegistry.m, id)
	listBoxSortFuncRegistry.Unlock()

	return C.gint(r.fn(wrapListBoxRow(glib.Take(unsafe.Pointer(row1))), wrapListBoxRow(glib.Take(unsafe.Pointer(row2))), r.userData))
}
