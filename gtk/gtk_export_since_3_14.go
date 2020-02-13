// Same copyright and license as the rest of the files in this project
// +build !gtk_3_6,!gtk_3_8,!gtk_3_10,!gtk_3_12

package gtk

// #include <gtk/gtk.h>
import "C"
import (
	"unsafe"

	"github.com/gotk3/gotk3/glib"
)

//export goListBoxForEachFuncs
func goListBoxForEachFuncs(box *C.GtkListBox, row *C.GtkListBoxRow, userData C.gpointer) {
	id := int(uintptr(userData))

	listBoxForeachFuncRegistry.RLock()
	r := listBoxForeachFuncRegistry.m[id]
	listBoxForeachFuncRegistry.RUnlock()

	r.fn(wrapListBox(glib.Take(unsafe.Pointer(box))), wrapListBoxRow(glib.Take(unsafe.Pointer(row))), r.userData)
}
