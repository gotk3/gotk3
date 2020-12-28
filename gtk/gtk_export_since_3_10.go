// +build !gtk_3_6,!gtk_3_8

package gtk

// #include <gtk/gtk.h>
import "C"
import (
	"unsafe"

	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/internal/callback"
)

//export goListBoxFilterFuncs
func goListBoxFilterFuncs(row *C.GtkListBoxRow, userData C.gpointer) C.gboolean {
	fn := callback.Get(uintptr(userData)).(ListBoxFilterFunc)
	return gbool(fn(wrapListBoxRow(glib.Take(unsafe.Pointer(row)))))
}

//export goListBoxHeaderFuncs
func goListBoxHeaderFuncs(row *C.GtkListBoxRow, before *C.GtkListBoxRow, userData C.gpointer) {
	fn := callback.Get(uintptr(userData)).(ListBoxHeaderFunc)
	fn(
		wrapListBoxRow(glib.Take(unsafe.Pointer(row))),
		wrapListBoxRow(glib.Take(unsafe.Pointer(before))),
	)
}

//export goListBoxSortFuncs
func goListBoxSortFuncs(row1 *C.GtkListBoxRow, row2 *C.GtkListBoxRow, userData C.gpointer) C.gint {
	fn := callback.Get(uintptr(userData)).(ListBoxSortFunc)
	return C.gint(fn(
		wrapListBoxRow(glib.Take(unsafe.Pointer(row1))),
		wrapListBoxRow(glib.Take(unsafe.Pointer(row2))),
	))
}
